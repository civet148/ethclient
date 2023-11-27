package ethclient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"testing"
)

const (
	txHash = "0xea665dfdc74dec98fbafc58f066b170411b7e28884255a0dcce312a74249bff0"
)

func TestEthereumClient(t *testing.T) {
	cli := NewEthereumClient("http://127.0.0.1:8545")
	height, err := cli.BlockNumber(context.Background())
	if err != nil {
		fmt.Printf("get block number error %s\n", err)
		return
	}
	fmt.Printf("blockchain height %v\n", height)
	var tx *types.Transaction
	tx, _, err = cli.TransactionByHash(context.Background(), txHash)
	if err != nil {
		fmt.Printf("get tx error %s\n", err)
		return
	}
	printJson("tx", tx)
	var method *CallMethod
	method, err = cli.GetTxCallMethod(context.Background(), txHash, "NFT.abi")
	if err != nil {
		fmt.Printf("get receipt error %s\n", err)
		return
	}
	printJson("method", method)

	var events []*CallMethod
	events, err = cli.GetTxEvents(context.Background(), txHash, "NFT.abi")
	if err != nil {
		fmt.Printf("get tx events error %s\n", err)
		return
	}
	_ = events
	printJson("events", events)
}

func printJson(title string, v interface{}) {
	fmt.Printf("------------------------------------- %s -------------------------------------\n", title)
	data, _ := json.MarshalIndent(v, "", "\t")
	fmt.Printf("%s\n", data)
	fmt.Printf("------------------------------------------------------------------------------\n")
}
