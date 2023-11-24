package ethclient

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

const (
	hexPrefix = "0x"
)

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

func CallOpts(strAddress string) *bind.CallOpts {
	address := Hex2Address(strAddress)
	return &bind.CallOpts{
		From: address,
	}
}

func TrimHexPrefix(str string) string {
	if strings.HasPrefix(str, hexPrefix) {
		str = strings.TrimPrefix(str, hexPrefix)
	}
	return str
}

func Hex2Hash(hash string) common.Hash {
	hash = TrimHexPrefix(hash)
	return common.HexToHash(hash)
}

func Big2Int64(n *big.Int) int64 {
	return n.Int64()
}

func Int642Big(n int64) *big.Int {
	return big.NewInt(n)
}

func Uint642Big(n uint64) *big.Int {
	return big.NewInt(int64(n))
}

func Hex2Address(addr string) common.Address {
	mixedAddr, err := common.NewMixedcaseAddressFromString(addr)
	if err != nil {
		return common.Address{}
	}
	return mixedAddr.Address()
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
