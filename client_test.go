package ethclient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"testing"
)

const (
	nodeUrl = "http://103.39.218.177:8545"
	txHash  = "0xe3e5cf580441963933a0932957cf740cdd63249f2030328216cd15bd63eea7b4"
)

var cli = NewEthereumClient(&Option{
	NodeUrl: nodeUrl,
	ABI:     "./contracts/abis/ERC721.abi",
})

func TestTransactionByHash(t *testing.T) {
	var err error
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
	var receipt *types.Receipt
	receipt, err = cli.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		fmt.Printf("get receipt error %s\n", err)
		return
	}
	printJson("receipt", receipt)
}

func TestGetTxMethod(t *testing.T) {
	var err error
	var method *CallMethod
	method, err = cli.GetTxCallMethod(context.Background(), "0x91cd36e1583ee38b2c724071c638879e5475982b644812087285daf3d61fa6af")
	if err != nil {
		fmt.Printf("get method error %s\n", err)
		return
	}
	type SafeTransferFrom struct {
		From    common.Address //`json:"from"`
		To      common.Address //`json:"to"`
		TokenId *big.Int       //`json:"tokenId"`
	}
	var safeTransfer SafeTransferFrom
	if err = method.Unpack(&safeTransfer); err != nil {
		fmt.Printf("get tx method error %s\n", err)
		return
	}
	printJson("SafeTransferFrom", safeTransfer)
}

func TestGetTxEvents(t *testing.T) {
	var err error
	var events []*CallEvent
	events, err = cli.GetTxEvents(context.Background(), txHash)
	if err != nil {
		fmt.Printf("get tx events error %s\n", err)
		return
	}
	for _, e := range events {
		fmt.Printf("event name %s\n", e.Event.Name)
		//var transfer erc721.Erc721Transfer
		//if err = e.Unpack(&transfer); err != nil {
		//	fmt.Printf("unpack event %s error %s\n", e.Event.Name, err)
		//	return
		//}
		//printJson("[Transfer]", transfer)
	}
}

func printJson(title string, v interface{}) {
	fmt.Printf("------------------------------------- %s -------------------------------------\n", title)
	data, _ := json.MarshalIndent(v, "", "\t")
	fmt.Printf("%s\n", data)
	fmt.Printf("------------------------------------------------------------------------------\n")
}
