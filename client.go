package ethclient

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
	"strings"
)

const (
	hexPrefix   = "0x"
	NullAddress = "0x0000000000000000000000000000000000000000"
)

type CallInput struct {
	Argument abi.Argument
	Value    interface{}
}

type CallMethod struct {
	Name   string
	Sig    string
	ID     string
	Inputs []*CallInput
}

type EthereumClient struct {
	ethcli *ethclient.Client
}

func NewEthereumClient(strNodeUrl string) *EthereumClient {
	ethcli, err := ethclient.Dial(strNodeUrl)
	if err != nil {
		panic(fmt.Sprintf("dial to ethereum node [%s] error [%s]", strNodeUrl, err.Error()))
	}
	return &EthereumClient{
		ethcli: ethcli,
	}
}

func (m *EthereumClient) Client() *ethclient.Client {
	return m.ethcli
}

func (m *EthereumClient) ChainID(ctx context.Context) (chainId int64, err error) {
	var id *big.Int
	id, err = m.ethcli.ChainID(ctx)
	if err != nil {
		return 0, err
	}
	return id.Int64(), nil
}

func (m *EthereumClient) BlockNumber(ctx context.Context) (uint64, error) {
	return m.ethcli.BlockNumber(ctx)
}

func (m *EthereumClient) BlockByHash(ctx context.Context, hash string) (*types.Block, error) {
	return m.ethcli.BlockByHash(ctx, Hex2Hash(hash))
}

func (m *EthereumClient) BlockByNumber(ctx context.Context, number uint64) (*types.Block, error) {
	return m.ethcli.BlockByNumber(ctx, big.NewInt(int64(number)))
}

func (m *EthereumClient) TransactionByHash(ctx context.Context, hash string) (tx *types.Transaction, pending bool, err error) {
	return m.ethcli.TransactionByHash(ctx, Hex2Hash(hash))
}

func (m *EthereumClient) TransactionReceipt(ctx context.Context, hash string) (tx *types.Receipt, err error) {
	return m.ethcli.TransactionReceipt(ctx, Hex2Hash(hash))
}

func (m *EthereumClient) PeerCount(ctx context.Context) (uint64, error) {
	return m.ethcli.PeerCount(ctx)
}

//func (m *EthereumClient) BlockReceipts(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) ([]*types.Receipt, error) {
//	return m.ethcli.BlockReceipts(ctx, blockNrOrHash)
//}

func (m *EthereumClient) HeaderByHash(ctx context.Context, hash string) (*types.Header, error) {
	return m.ethcli.HeaderByHash(ctx, Hex2Hash(hash))
}

func (m *EthereumClient) HeaderByNumber(ctx context.Context, number uint64) (*types.Header, error) {
	return m.ethcli.HeaderByNumber(ctx, big.NewInt(int64(number)))
}

func (m *EthereumClient) TransactionSender(ctx context.Context, tx *types.Transaction, strBlockHash string, index uint) (common.Address, error) {
	return m.ethcli.TransactionSender(ctx, tx, Hex2Hash(strBlockHash), index)
}

func (m *EthereumClient) TransactionCount(ctx context.Context, strBlockHash string) (uint, error) {
	return m.ethcli.TransactionCount(ctx, Hex2Hash(strBlockHash))
}

func (m *EthereumClient) TransactionInBlock(ctx context.Context, strBlockHash string, index uint) (*types.Transaction, error) {
	return m.ethcli.TransactionInBlock(ctx, Hex2Hash(strBlockHash), index)
}

func (m *EthereumClient) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	return m.ethcli.SyncProgress(ctx)
}

func (m *EthereumClient) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	return m.ethcli.SubscribeNewHead(ctx, ch)
}

func (m *EthereumClient) NetworkID(ctx context.Context) (int64, error) {
	id, err := m.ethcli.NetworkID(ctx)
	if err != nil {
		return 0, err
	}
	return id.Int64(), nil
}

func (m *EthereumClient) BalanceAt(ctx context.Context, strAddress string, number uint64) (*big.Int, error) {
	return m.ethcli.BalanceAt(ctx, Hex2Address(strAddress), Int642Big(int64(number)))
}

//func (m *EthereumClient) BalanceAtHash(ctx context.Context, strAddress string, strBlockHash string) (*big.Int, error) {
//	return m.ethcli.BalanceAtHash(ctx, Hex2Address(strAddress), Hex2Hash(strBlockHash))
//}

func (m *EthereumClient) StorageAt(ctx context.Context, strAddress, strKey string, number uint64) ([]byte, error) {
	return m.ethcli.StorageAt(ctx, Hex2Address(strAddress), Hex2Hash(strKey), Int642Big(int64(number)))
}

//func (m *EthereumClient) StorageAtHash(ctx context.Context, strAddress, strKey, strBlockHash string) ([]byte, error) {
//	return m.ethcli.StorageAtHash(ctx, Hex2Address(strAddress), Hex2Hash(strKey), Hex2Hash(strBlockHash))
//}

func (m *EthereumClient) CodeAt(ctx context.Context, strAddress string, number uint64) ([]byte, error) {
	return m.ethcli.CodeAt(ctx, Hex2Address(strAddress), Int642Big(int64(number)))
}

//func (m *EthereumClient) CodeAtHash(ctx context.Context, strAddress, strBlockHash string) ([]byte, error) {
//	return m.ethcli.CodeAtHash(ctx, Hex2Address(strAddress), Hex2Hash(strBlockHash))
//}

func (m *EthereumClient) NonceAt(ctx context.Context, strAddress string, number uint64) (uint64, error) {
	return m.ethcli.NonceAt(ctx, Hex2Address(strAddress), Int642Big(int64(number)))
}

//func (m *EthereumClient) NonceAtHash(ctx context.Context, strAddress string, strBlockHash string) (uint64, error) {
//	return m.ethcli.NonceAtHash(ctx, Hex2Address(strAddress), Hex2Hash(strBlockHash))
//}

func (m *EthereumClient) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return m.ethcli.FilterLogs(ctx, q)
}

func (m *EthereumClient) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return m.ethcli.SubscribeFilterLogs(ctx, q, ch)
}

func (m *EthereumClient) PendingBalanceAt(ctx context.Context, strAddress string) (*big.Int, error) {
	return m.ethcli.PendingBalanceAt(ctx, Hex2Address(strAddress))
}

func (m *EthereumClient) PendingStorageAt(ctx context.Context, strAddress string, strKey string) ([]byte, error) {
	return m.ethcli.PendingStorageAt(ctx, Hex2Address(strAddress), Hex2Hash(strKey))
}

func (m *EthereumClient) PendingCodeAt(ctx context.Context, strAddress string) ([]byte, error) {
	return m.ethcli.PendingCodeAt(ctx, Hex2Address(strAddress))
}

func (m *EthereumClient) PendingNonceAt(ctx context.Context, strAddress string) (uint64, error) {
	return m.ethcli.PendingNonceAt(ctx, Hex2Address(strAddress))
}

func (m *EthereumClient) PendingTransactionCount(ctx context.Context) (uint, error) {
	return m.ethcli.PendingTransactionCount(ctx)
}

func (m *EthereumClient) CallContract(ctx context.Context, msg ethereum.CallMsg, number uint64) ([]byte, error) {
	return m.ethcli.CallContract(ctx, msg, Uint642Big(number))
}

func (m *EthereumClient) CallContractAtHash(ctx context.Context, msg ethereum.CallMsg, strBlockHash string) ([]byte, error) {
	return m.ethcli.CallContractAtHash(ctx, msg, Hex2Hash(strBlockHash))
}

func (m *EthereumClient) PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
	return m.ethcli.PendingCallContract(ctx, msg)
}

func (m *EthereumClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return m.ethcli.SuggestGasPrice(ctx)
}

func (m *EthereumClient) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return m.ethcli.SuggestGasTipCap(ctx)
}

func (m *EthereumClient) FeeHistory(ctx context.Context, blockCount uint64, blockNumber uint64, rewardPercentiles []float64) (*ethereum.FeeHistory, error) {
	return m.ethcli.FeeHistory(ctx, blockCount, Uint642Big(blockNumber), rewardPercentiles)
}

func (m *EthereumClient) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return m.ethcli.EstimateGas(ctx, msg)
}

func (m *EthereumClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return m.ethcli.SendTransaction(ctx, tx)
}

func (m *EthereumClient) ContractTransactor() bind.ContractTransactor {
	return m.ethcli
}

func (m *EthereumClient) ContractCaller() bind.ContractCaller {
	return m.ethcli
}

func (m *EthereumClient) ContractBackend() bind.ContractBackend {
	return m.ethcli
}

func (m *EthereumClient) ContractFilter() bind.ContractFilterer {
	return m.ethcli
}

func (m *EthereumClient) PendingContractCaller() bind.PendingContractCaller {
	return m.ethcli
}

func (m *EthereumClient) GetContractAddrByTxHash(ctx context.Context, hash string) (strContractAddr string, err error) {
	var tx *types.Transaction
	var receipt *types.Receipt
	var contractAddress common.Address
	tx, _, err = m.TransactionByHash(ctx, hash)
	if err != nil {
		return "", fmt.Errorf("get receipt by hash [%s] error [%s]\n", hash, err)
	}

	if tx.To() != nil {
		contractAddress = *tx.To()
	}
	receipt, err = m.TransactionReceipt(ctx, hash)
	if err != nil {
		return "", fmt.Errorf("get receipt by hash [%s] error [%s]\n", hash, err)
	}
	if contractAddress.String() == NullAddress && receipt.ContractAddress.String() != NullAddress {
		contractAddress = receipt.ContractAddress
	}
	if contractAddress.String() == NullAddress {
		return "", fmt.Errorf("get contract address by tx hash [%s] error: not found", hash)
	}
	return contractAddress.String(), nil
}

func (m *EthereumClient) GetTxCallMethod(ctx context.Context, hash string, strABI string) (calldata *CallMethod, err error) {
	var tx *types.Transaction
	var receipt *types.Receipt
	var contractAddress common.Address
	tx, _, err = m.TransactionByHash(ctx, hash)
	if err != nil {
		return nil, fmt.Errorf("get receipt by hash [%s] error [%s]\n", hash, err)
	}
	if tx.To() != nil {
		contractAddress = *tx.To()
	}
	receipt, err = m.TransactionReceipt(ctx, hash)
	if err != nil {
		return nil, fmt.Errorf("get receipt by hash [%s] error [%s]\n", hash, err)
	}
	if contractAddress.String() == NullAddress && receipt.ContractAddress.String() != NullAddress {
		contractAddress = receipt.ContractAddress
	}
	if contractAddress.String() == NullAddress {
		return nil, fmt.Errorf("get contract address by tx hash [%s] error: not found", hash)
	}
	var contractABI abi.ABI
	contractABI, err = m.LoadABI(strABI)
	if err != nil {
		return nil, err
	}
	data := tx.Data()
	id := data[0:4]
	var method *abi.Method
	method, err = contractABI.MethodById(id)
	if method == nil {
		if err != nil {
			return nil, fmt.Errorf("search by method id 0x%x not found", id)
		}
	}

	var inputValues []interface{}
	inputValues, err = method.Inputs.UnpackValues(data[4:])
	if err != nil {
		return nil, fmt.Errorf("unpack values error [%s]", err)
	}
	var inputs []*CallInput
	for i, v := range inputValues {
		inputs = append(inputs, &CallInput{
			Argument: method.Inputs[i],
			Value:    v,
		})
	}
	return &CallMethod{
		Name:   method.Name,
		Sig:    method.Sig,
		ID:     fmt.Sprintf("0x%x", method.ID),
		Inputs: inputs,
	}, nil
}

func (m *EthereumClient) GetTxEvents(ctx context.Context, hash string, strABI string) (events []*CallMethod, err error) {
	var tx *types.Transaction
	var receipt *types.Receipt
	var contractAddress common.Address
	tx, _, err = m.TransactionByHash(ctx, hash)
	if err != nil {
		return nil, fmt.Errorf("get receipt by hash [%s] error [%s]\n", hash, err)
	}

	if tx.To() != nil {
		contractAddress = *tx.To()
	}
	receipt, err = m.TransactionReceipt(ctx, hash)
	if err != nil {
		return nil, fmt.Errorf("get receipt by hash [%s] error [%s]\n", hash, err)
	}
	if contractAddress.String() == NullAddress && receipt.ContractAddress.String() != NullAddress {
		contractAddress = receipt.ContractAddress
	}
	if contractAddress.String() == NullAddress {
		return nil, fmt.Errorf("get contract address by tx hash [%s] error: not found", hash)
	}
	var contractABI abi.ABI
	contractABI, err = m.LoadABI(strABI)
	if err != nil {
		return nil, err
	}

	for _, lo := range receipt.Logs {
		var evt *abi.Event
		evt, err = contractABI.EventByID(lo.Topics[0])
		if err != nil {
			continue //event id not found
		}
		if evt == nil {
			continue
		}
		inputValues := lo.Topics[1:]
		var inputs []*CallInput
		for i, v := range inputValues {
			inputs = append(inputs, &CallInput{
				Argument: evt.Inputs[i],
				Value:    v.String(),
			})
		}
		events = append(events, &CallMethod{
			Name:   evt.Name,
			Sig:    evt.Sig,
			ID:     evt.ID.String(),
			Inputs: inputs,
		})
	}
	return
}

// LoadABI load ABI from file or json string
func (m *EthereumClient) LoadABI(strABI string) (contractABI abi.ABI, err error) {
	if strings.HasSuffix(strABI, ".abi") {
		return m.loadABIFromFile(strABI)
	}
	return m.loadABIFromString(strABI)
}

func (m *EthereumClient) loadABIFromFile(strAbiFile string) (contractABI abi.ABI, err error) {
	var file *os.File
	file, err = os.Open(strAbiFile)
	if err != nil {
		return abi.ABI{}, fmt.Errorf("read abi file %s error: %s", strAbiFile, err.Error())
	}

	contractABI, err = abi.JSON(file)
	if err != nil {
		return abi.ABI{}, fmt.Errorf("unmarshal abi json error: %s", err.Error())
	}
	return contractABI, nil
}

func (m *EthereumClient) loadABIFromString(strAbi string) (contractABI abi.ABI, err error) {
	reader := bytes.NewBufferString(strAbi)
	contractABI, err = abi.JSON(reader)
	if err != nil {
		return abi.ABI{}, fmt.Errorf("unmarshal abi json error: %s", err.Error())
	}
	return contractABI, nil
}
