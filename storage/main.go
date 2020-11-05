package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	accessControlContract "./contracts/accessContract"
	dataContract "./contracts/dataContract"
	libs "./libs"
)

// Local definition of the struct libs.ComponentConfig
type localClient libs.ComponentConfig

func readConfigFile() map[string]interface{} {
	config := make(map[string]interface{})

	// Open the configuration file
	jsonFile, err := os.Open("./config/config.json")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer jsonFile.Close()

	// Parse to bytes
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Load the object in the map[string]interface{} variable
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return config

}

func initialize() localClient {
	// Read the configuration file
	config := readConfigFile()

	// Connect to the IPC endpoint of the Ethereum node
	client, err := ethclient.Dial(config["nodePath"].(string) + "geth.ipc")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Get the private key of the admin
	privKey, err := libs.GetPrivateKey(config["ethAddress"].(string),
		config["ethPassword"].(string),
		config["nodePath"].(string)+"keystore/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Initialize the data contract
	dataContract, err := dataContract.NewDataLedgerContract(common.HexToAddress(config["dataContractAddr"].(string)), client)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Initialize the accessControlContract
	accessContract, err := accessControlContract.NewAccessControlContract(common.HexToAddress(config["accessContractAddr"].(string)), client)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	localEthClient := localClient{
		client,
		common.HexToAddress(config["ethAddress"].(string)),
		privKey,
		dataContract,
		accessContract,
		config,
	}

	return localEthClient

}

// Listens the measurements provided by the IoT producers and
// responds them with the URL where the measurement has been stored
func (localEthClient localClient) measurementListener(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/store" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Convert the local struct to global struct
	userEthClien := libs.ComponentConfig{
		localEthClient.EthereumClient,
		localEthClient.Address,
		localEthClient.PrivateKey,
		localEthClient.DataCon,
		localEthClient.AccessCon,
		localEthClient.GeneralConfig,
	}

	switch req.Method {
	case "POST":
		fmt.Printf("Request received from %s\n", req.Host)
		bodyBytes := libs.StreamToByte(req.Body)

		// Check if the producer can introduce information in the Blockchain
		iotName, hash, err := libs.CheckAccess(userEthClien, bodyBytes)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "401 Could not introduce the event in the blockchain due to: "+err.Error(), http.StatusBadRequest)
		}

		// Store the information in the database
		url, err := libs.StoreInDB(userEthClien.GeneralConfig, bodyBytes[:len(bodyBytes)-65], fmt.Sprintf("%x", hash), iotName)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "500 Could not introduce the event in the blockchain due to: "+err.Error(), http.StatusBadRequest)
		}

		// Create JSON anwser
		jsonResponse := make(map[string]interface{})
		jsonResponse["url"] = url
		err = json.NewEncoder(w).Encode(jsonResponse)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "500 Could not introduce the event in the blockchain due to: "+err.Error(), http.StatusBadRequest)
		}

	default:
		http.Error(w, "401 Only POST methods are supported", http.StatusBadRequest)
		fmt.Fprintf(w, "Only Post methods are supported")
	}
}

func main() {
	fmt.Printf("----------- Initializing Storage module -----------\n\n")
	localEthClient := initialize()

	fmt.Printf("Starting server on port %s\n\n", localEthClient.GeneralConfig["HTTPport"].(string))

	// Route to process the measurements of the IoT producers
	http.HandleFunc("/store", localEthClient.measurementListener)

	err := http.ListenAndServe(":"+localEthClient.GeneralConfig["HTTPport"].(string), nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}
