package libs

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

// Checks if there is a CompletePurchase event indexed by the clientAddr and the hash
func checkEventLog(ethClient ComponentConfig, clientAddr, hash string) (bool, error) {
	topic := []byte("CompletePurchase(bytes32,address,address,bytes32)")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		Addresses: []common.Address{
			common.HexToAddress(ethClient.GeneralConfig["balanceContractAddr"].(string))},
		Topics: [][]common.Hash{{crypto.Keccak256Hash(topic)}},
	}

	logs, err := ethClient.EthereumClient.FilterLogs(context.Background(), query)
	if err != nil {
		return false, err
	}

	for _, vlog := range logs {
		hashLog := vlog.Topics[1].Hex()
		fromLog := common.HexToAddress(vlog.Topics[2].Hex())

		if hashLog == hash && fromLog == common.HexToAddress(clientAddr) {
			return true, nil
		}
	}

	return false, errors.New("Did not find any event that matches the requirements")
}

// Queries the database to retrieve the value of a measurement
func queryDB(configParams map[string]interface{}, hash, tableName string) ([]byte, error) {
	type dbResponseStruct struct {
		Hash        string `json:"hash"`
		Measurement string `json:"measurement"`
	}

	var dbResponse dbResponseStruct

	// The format of the URL to the database is:
	// 	user:password@tcp(dbHost)/dbName
	connParams := configParams["dbUsername"].(string) + ":" +
		configParams["dbPassword"].(string) + "@tcp(" +
		configParams["dbHost"].(string) + ")/" +
		configParams["dbName"].(string)

	// Connect to the MySQL database
	db, err := sql.Open("mysql", connParams)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	// Create query
	query := fmt.Sprintf("SELECT * FROM %s WHERE Hash='%s';", tableName, hash)
	err = db.QueryRow(query).Scan(&dbResponse.Hash, &dbResponse.Measurement)
	if err != nil {
		return nil, err
	}

	// Convert the struct to JSON
	response, err := json.Marshal(dbResponse)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// RetrieveMeasurement sends the purchases to customers
// 	1) The body of the request must be a json string containing the following fields:
//		+ clientAddr -> address of the customer who purchased the measurement
//		+ hash -> hash of the measurement
//		+ timestamp -> timestamp of the request
//		+ signature -> signature of the previous fields
//	2) Gets the public key of the customer from the Access SC
// 	3) Verifies the identity of the customer by checking the signature of the message and the
//	   the public key obtained in the previous point.
//	4) Verifies that there is an event in the Blockchain matching the purchase
//	5) Retrieves the measurement from the DB and sends it to the customer
func RetrieveMeasurement(ethClient ComponentConfig, body map[string]interface{}, req *http.Request) ([]byte, error) {

	// Extract the parameters from the URL
	tableName := mux.Vars(req)["table"]
	hashURL := mux.Vars(req)["hash"]

	clientAddr := body["clientAddr"].(string)
	hash := body["hash"].(string)

	// The hash indicated in the path of the URL must match the one in the body of the message
	if hashURL != hash {
		log.Println("Wrong URL")
		return nil, errors.New("Wrong URL")
	}

	log.Printf("%s: Resquest measurement %s\n", clientAddr, hash)

	// Get the public key of the customer
	pubKey, err := ethClient.AccessCon.PubKeysKeystore(nil, common.HexToAddress(clientAddr))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Verify the signature of the message
	log.Printf("%s: Verifying customer signature\n", clientAddr)
	msg := clientAddr + hash + fmt.Sprintf("%d", uint64(body["timestamp"].(float64)))
	msgHash := sha256.Sum256([]byte(msg))
	isSignatureOK := secp256k1.VerifySignature(common.Hex2Bytes(pubKey), msgHash[:], common.Hex2Bytes(body["signature"].(string)))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if !isSignatureOK {
		log.Printf("%s: Wrong signature\n", clientAddr)
		return nil, errors.New("Wrong signature")
	}

	log.Printf("%s: Signature ok\n", clientAddr)

	// Check that there is an event that matches the purchase
	_, err = checkEventLog(ethClient, "0x"+clientAddr, "0x"+hash)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Query the database
	response, err := queryDB(ethClient.GeneralConfig, hashURL, tableName)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return response, nil
}
