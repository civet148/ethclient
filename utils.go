package ethclient

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"reflect"
	"strings"
)

// NewTransactOpts new transact options by private key string or *ecdsa.PrivateKey object and chain id
func NewTransactOpts(privateKey interface{}, chainId int64) (txOpts *bind.TransactOpts, err error) {
	var pk *ecdsa.PrivateKey
	pk, err = NewPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	return bind.NewKeyedTransactorWithChainID(pk, big.NewInt(chainId))
}

// NewTransactOptsWithValue new transact options by private key string or *ecdsa.PrivateKey object and chain id and value
func NewTransactOptsWithValue(privateKey interface{}, chainId int64, value interface{}) (txOpts *bind.TransactOpts, err error) {
	var pk *ecdsa.PrivateKey
	pk, err = NewPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	if value == nil {
		return nil, fmt.Errorf("value must be non-nil")
	}
	return newKeyedTransactorWithValue(pk, big.NewInt(chainId), value)
}

// NewPrivateKey new ecdsa private key from hex string, bytes or *ecdsa.PrivateKey
func NewPrivateKey(privateKey interface{}) (pk *ecdsa.PrivateKey, err error) {
	switch privateKey.(type) {
	case string:
		var pkBytes []byte
		strPriKey := privateKey.(string)
		strPriKey = TrimHexPrefix(strPriKey)
		pkBytes, err = hex.DecodeString(strPriKey)
		pk, err = crypto.ToECDSA(pkBytes)
		if err != nil {
			return nil, err
		}
	case []byte:
		pk, err = crypto.ToECDSA(privateKey.([]byte))
		if err != nil {
			return nil, err
		}
	case *ecdsa.PrivateKey:
		pk = privateKey.(*ecdsa.PrivateKey)
	default:
		return nil, fmt.Errorf("unsupported private key type: %v", reflect.TypeOf(privateKey).String())
	}
	return pk, nil
}

func NewCallOpts(strFromAddr string) *bind.CallOpts {
	address := Hex2Address(strFromAddr)
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

// newKeyedTransactorWithValue is a utility method to easily create a transaction signer
// from a single private key and value.
func newKeyedTransactorWithValue(key *ecdsa.PrivateKey, chainID *big.Int, value interface{}) (*bind.TransactOpts, error) {
	var ok bool
	keyAddr := crypto.PubkeyToAddress(key.PublicKey)
	if chainID == nil {
		return nil, bind.ErrNoChainID
	}
	var bigValue = big.NewInt(0)
	switch value.(type) {
	case *big.Int:
		bigValue = value.(*big.Int)
	case string:
		bigValue, ok = bigValue.SetString(value.(string), 10)
		if !ok {
			return nil, fmt.Errorf("value '%v' invalid", value.(string))
		}
	default:
		strValue := fmt.Sprintf("%v", value)
		bigValue, ok = bigValue.SetString(strValue, 10)
		if !ok {
			return nil, fmt.Errorf("value '%v' invalid", value)
		}
	}
	signer := types.LatestSignerForChainID(chainID)
	return &bind.TransactOpts{
		From: keyAddr,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != keyAddr {
				return nil, bind.ErrNotAuthorized
			}
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), key)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
		Context: context.Background(),
		Value:   bigValue,
	}, nil
}
