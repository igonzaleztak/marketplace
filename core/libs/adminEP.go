package libs

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// AdminListener listens to the request made by the admin user
func AdminListener(ethClient ComponentConfig, body map[string]interface{}) error {
	action := body["action"].(string)
	addr := body["addr"].(string)
	tokens := int64(body["tokens"].(float64))

	fmt.Println(fmt.Sprintf("+ Sending %d to %s", tokens, addr))

	switch action {
	case "sendTokens":
		// Send the tokens to the client
		auth := bind.NewKeyedTransactor(ethClient.PrivateKey)
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(400000)
		auth.GasPrice = big.NewInt(0)
		_, err := ethClient.BalanceCon.SendTokenToClient(auth, common.HexToAddress(addr), big.NewInt(tokens))
		if err != nil {
			return err
		}
	default:
		return errors.New("Unelegible action")
	}

	return nil
}
