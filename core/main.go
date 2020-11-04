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

	accessControlContract "./contracts/accessContract"
	balanceContract "./contracts/balanceContract"
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

// EventListener listens to new purchases in the marketplace and process them
func (localEthClient *localClient) EventListener(w http.ResponseWriter, req *http.Request) {}

func main() {
	fmt.Printf("----------- Initializing Storage module -----------\n\n")
	localEthClient := initialize()

	fmt.Printf("Starting server on port %s\n\n", localEthClient.GeneralConfig["HTTPport"].(string))
	http.HandleFunc("/marketplace", localEthClient.EventListener)

	err := http.ListenAndServe(":"+localEthClient.GeneralConfig["HTTPport"].(string), nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}
