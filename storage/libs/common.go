package libs

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	accessControlContract "../contracts/accessContract"
	balanceContract "../contracts/balanceContract"
	dataContract "../contracts/dataContract"
)

// ComponentConfig is a struct that stores the parameters of the node
type ComponentConfig struct {
	EthereumClient *ethclient.Client
	Address        common.Address
	PrivateKey     *ecdsa.PrivateKey
	DataCon        *dataContract.DataLedgerContract
	AccessCon      *accessControlContract.AccessControlContract
	BalanceCon     *balanceContract.BalanceContract
	GeneralConfig  map[string]interface{}
}
