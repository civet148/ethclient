package ethclient

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
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

type CallEvent struct {
	Event *abi.Event
	ABI   abi.ABI
	Log   types.Log
}

func (e *CallEvent) Unpack(v interface{}) error {
	return UnpackLog(e.ABI, v, e.Event.Name, e.Log)
}
