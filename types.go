package ethclient

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	hexPrefix   = "0x"
	NullAddress = "0x0000000000000000000000000000000000000000"
)

type CallEvent struct {
	Event *abi.Event
	ABI   abi.ABI
	Log   types.Log
}

func (e *CallEvent) Unpack(v interface{}) error {
	return UnpackLog(e.ABI, v, e.Event.Name, e.Log)
}
func (m *CallEvent) Prototype() string {
	return m.Event.Sig
}

func (m *CallEvent) Sig() string {
	return m.Event.Sig
}

func (m *CallEvent) Name() string {
	return m.Event.Name
}

func (m *CallEvent) ID() string {
	return hex.EncodeToString(m.Event.ID.Bytes())
}

type CallMethod struct {
	Method *abi.Method
	ABI    abi.ABI
	Data   []byte
}

func (m *CallMethod) Unpack(v interface{}) error {
	values, err := m.Method.Inputs.UnpackValues(m.Data)
	if err != nil {
		return err
	}
	return m.Method.Inputs.Copy(v, values)
}

func (m *CallMethod) Prototype() string {
	return m.Method.Sig
}

func (m *CallMethod) Sig() string {
	return m.Method.Sig
}

func (m *CallMethod) Name() string {
	return m.Method.Name
}

func (m *CallMethod) ID() string {
	return hex.EncodeToString(m.Method.ID)
}

func (m *CallMethod) InputValues() []interface{} {
	values, err := m.Method.Inputs.UnpackValues(m.Data)
	if err != nil {
		return nil
	}
	return values
}
