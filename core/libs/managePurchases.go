package libs

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ManagePurchases listens to the purchases events in
// the Blockchain and processes them
func ManagePurchases(ethClient ComponentConfig) {

	// Prepare the filter to filter the Balance contract events
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			common.HexToAddress(ethClient.GeneralConfig["balanceContractAddr"].(string))},
	}

	// Make channel to receive the events
	logs := make(chan types.Log)
	sub, err := ethClient.EthereumClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	// Process the events
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
