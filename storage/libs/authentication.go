package libs

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/accounts/keystore"

	cipher "../cipherLibs"
)

// Gets the file that contains the private key of the admin
func getUTCFile(address, folderPath string) (string, error) {
	// Compile the regex expresion
	libRegEx, err := regexp.Compile("(?i).*" + address)
	if err != nil {
		log.Fatal(err)
	}

	// Get all the files of the folder
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		log.Fatal(err)
	}

	// Get the name of the file that matches the expression
	for _, f := range files {
		if libRegEx.MatchString(f.Name()) {
			return folderPath + f.Name(), nil
		}
	}

	return "", errors.New("UTC File not found")
}

// GetPrivateKey gets the Ethereum's private key of the Market component
func GetPrivateKey(address, password, folderPath string) (*ecdsa.PrivateKey, error) {

	// Get the file that contains the private key
	file, err := getUTCFile(address[2:], folderPath)
	if err != nil {
		return nil, err
	}

	// Read the file
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// Get the private key
	keyWrapper, err := keystore.DecryptKey(jsonBytes, password)
	if err != nil {
		return nil, err
	}

	return keyWrapper.PrivateKey, nil
}

// CheckAccess checks whether an IoT producer can insert a measurement
// in the Blockchain. To do so, it checks if the signature has been produced
// by an authorized IoT producer.
func CheckAccess(ethClient ComponentConfig, body []byte) (string, []byte, error) {
	// Separate the address and encrypted body from the signature to check the signature.
	// The last 65 bytes of the body is the signature [R || S || V].
	msg := body[:len(body)-65]
	signature := body[len(body)-65:]

	// Get the hash of the measurement
	hash := cipher.HashData(msg)

	// Recover the public key and the Ethereum address that signed the measurement
	iotAddr, _, err := cipher.RecoverDataFromSignature(hash, signature)
	if err != nil {
		return "", nil, err
	}

	// Check if the address has access
	hasAccess, err := ethClient.AccessCon.AllowedAccounts(nil, iotAddr)
	if err != nil {
		return "", nil, err
	}
	if !hasAccess {
		return "", nil, errors.New("This address is not authorized to introduce content in the marketplace")
	}

	fmt.Println("The measurement is properly signed")

	// Get the name of the IoT producer from its address
	iotName, err := ethClient.AccessCon.ProducersNameMap(nil, iotAddr)
	if err != nil {
		return "", nil, err
	}
	if iotName == "" {
		return "", nil, errors.New("The IoT producer's name is not registered in the Blockchain")
	}

	fmt.Println("Measurement received from: " + fmt.Sprintf("%x", iotAddr) + " (" + iotName + ")")

	return iotName, hash, nil
}
