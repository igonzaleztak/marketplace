package main

import (
	"crypto/elliptic"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	accessControlContract "marketplace/core/contracts/accessContract"
	balanceContract "marketplace/core/contracts/balanceContract"
	dataContract "marketplace/core/contracts/dataContract"
	libs "marketplace/core/libs"
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

	// Set the access contract address in the data contract so the data
	// contract can interact with the  access contract
	auth := bind.NewKeyedTransactor(privKey)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(400000)
	auth.GasPrice = big.NewInt(0)

	_, err = dataContract.SetAddress(auth, common.HexToAddress(config["accessContractAddr"].(string)))
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

	// Set admin public address in the access smart contract
	publicKeyECDSA := privKey.PublicKey
	publicKeyBytes := elliptic.Marshal(publicKeyECDSA.Curve, publicKeyECDSA.X, publicKeyECDSA.Y)
	publicKeyString := fmt.Sprintf("%x", publicKeyBytes)
	_, err = accessContract.AddPubKey(auth, publicKeyString)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Initialize the balanceContract
	balanceContract, err := balanceContract.NewBalanceContract(common.HexToAddress(config["balanceContractAddr"].(string)), client)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Set the data contract address in the balance contract
	_, err = balanceContract.SetAddress(auth, common.HexToAddress(config["dataContractAddr"].(string)))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Set the total amount of tokens in the marketplace
	_, err = balanceContract.SetTotalSupply(auth, big.NewInt(int64(config["totalSupply"].(float64))))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	localEthClient := localClient{
		client,
		common.HexToAddress(config["ethAddress"].(string)),
		privKey,
		publicKeyECDSA,
		dataContract,
		accessContract,
		balanceContract,
		config,
	}

	return localEthClient

}

// AdminEP is the end point for the admin user. Here, the admin can send
// new tokens to clients, add new producers to the access list, ...
func (localClient *localClient) AdminEP(w http.ResponseWriter, req *http.Request) {
	// Check whether the URL is correct or not
	if req.URL.Path != "/admin" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Create a map with body of the message
	//var bodyArray bodyArray
	bodyMap := make(map[string]interface{})

	// Read the body of the message
	err := json.NewDecoder(req.Body).Decode(&bodyMap)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "401 Could not introduce the event in the blockchain because of: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Convert the localClient to libs.ComponentConfig
	ethClient := libs.ComponentConfig{
		localClient.EthereumClient,
		localClient.Address,
		localClient.PrivateKey,
		localClient.PublicKey,
		localClient.DataCon,
		localClient.AccessCon,
		localClient.BalanceCon,
		localClient.GeneralConfig,
	}

	// Check if the method of the request is correct
	switch req.Method {
	case "POST":
		err := libs.AdminListener(ethClient, bodyMap)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "401 Could not introduce the event in the blockchain because of: "+err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "200 ok", http.StatusOK)

	default:
		http.Error(w, "401 Only POST methods are supported", http.StatusBadRequest)
		fmt.Fprintf(w, "Only Post methods are supported")
	}
}


func main() {
	fmt.Printf("----------- Initializing Core module -----------\n\n")
	localClient := initialize()

	fmt.Printf("Starting server on port %s\n\n", localClient.GeneralConfig["HTTPport"].(string))
	http.HandleFunc("/admin", localClient.AdminEP)

	// Convert the localClient to libs.ComponentConfig
	ethClient := libs.ComponentConfig{
		localClient.EthereumClient,
		localClient.Address,
		localClient.PrivateKey,
		localClient.PublicKey,
		localClient.DataCon,
		localClient.AccessCon,
		localClient.BalanceCon,
		localClient.GeneralConfig,
	}

	// Run the purchases manager thread
	go libs.ManagePurchases(ethClient)

	err := http.ListenAndServe(":"+localClient.GeneralConfig["HTTPport"].(string), nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}
