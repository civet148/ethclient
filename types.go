package ethclient

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	hexPrefix   = "0x"
	NullAddress = "0x0000000000000000000000000000000000000000"
)

type CallMethod struct {
	Method *abi.Method
	ABI    abi.ABI
	Data   []byte
}

type CallEvent struct {
	Event *abi.Event
	ABI   abi.ABI
	Log   types.Log
}

func (e *CallEvent) Unpack(v interface{}) error {
	return UnpackLog(e.ABI, v, e.Event.Name, e.Log)
}

func (m *CallMethod) Unpack(v interface{}) error {
	values, err := m.Method.Inputs.UnpackValues(m.Data)
	if err != nil {
		return err
	}
	return m.Method.Inputs.Copy(v, values)
}
