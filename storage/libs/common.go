package libs

import (
	"bytes"
	"crypto/ecdsa"
	"io"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	accessControlContract "../contracts/accessContract"
	dataContract "../contracts/dataContract"
)

// ComponentConfig is a struct that stores the parameters of the node
type ComponentConfig struct {
	EthereumClient *ethclient.Client
	Address        common.Address
	PrivateKey     *ecdsa.PrivateKey
	DataCon        *dataContract.DataLedgerContract
	AccessCon      *accessControlContract.AccessControlContract
	GeneralConfig  map[string]interface{}
}

// StreamToByte converts io.Reader stream to string or byte slice
func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

// ByteToByte32 converts []byte to [32]byte
func ByteToByte32(bytes []byte) [32]byte {
	var b [32]byte
	copy(b[:], bytes)
	return b
}
