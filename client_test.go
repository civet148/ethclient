package ethclient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/civet148/ethclient/contracts/erc721"
	"github.com/ethereum/go-ethereum/core/types"
	"testing"
)

const (
	nodeUrl = "http://103.39.218.177:8545"
	txHash  = "0xe3e5cf580441963933a0932957cf740cdd63249f2030328216cd15bd63eea7b4"
)

func TestTransactionByHash(t *testing.T) {
	var err error
	cli := NewEthereumClient(nodeUrl)
	var tx *types.Transaction
	tx, _, err = cli.TransactionByHash(context.Background(), txHash)
	if err != nil {
		fmt.Printf("get tx error %s\n", err)
		return
	}
	printJson("tx", tx)
}

func TestTransactionReceipt(t *testing.T) {
	var err error
	cli := NewEthereumClient(nodeUrl)
	var receipt *types.Receipt
	receipt, err = cli.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		fmt.Printf("get receipt error %s\n", err)
		return
	}
	printJson("receipt", receipt)
}

func TestGetTxEvents(t *testing.T) {
	var err error
	var events []*CallEvent
	cli := NewEthereumClient(nodeUrl)
	events, err = cli.GetTxEvents(context.Background(), txHash, "./contracts/abis/ERC721.abi")
	if err != nil {
		fmt.Printf("get tx events error %s\n", err)
		return
	}
	for _, e := range events {
		var transfer erc721.Erc721Transfer
		if err = e.Unpack(&transfer); err != nil {
			fmt.Printf("unpack event %s error %s\n", e.Event.Name, err)
			return
		}
		printJson("[Transfer]", transfer)
	}

}

func printJson(title string, v interface{}) {
	fmt.Printf("------------------------------------- %s -------------------------------------\n", title)
	data, _ := json.MarshalIndent(v, "", "\t")
	fmt.Printf("%s\n", data)
	fmt.Printf("------------------------------------------------------------------------------\n")
}
