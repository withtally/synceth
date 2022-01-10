// Code generated by github.com/withtally/synceth, DO NOT EDIT.

package bindings

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/withtally/synceth/example"
)

type ExampleProcessor interface {
	Setup(address common.Address, eth interface {
		ethereum.ChainReader
		ethereum.ChainStateReader
		ethereum.TransactionReader
		bind.ContractBackend
	}) error
	Initialize(ctx context.Context, start uint64, tx *example.TestInput) error

	ProcessExampleEvent(ctx context.Context, e *ExampleExampleEvent, cb func(tx *example.TestInput)) error

	mustEmbedBaseExampleProcessor()
}

type BaseExampleProcessor struct {
	Address  common.Address
	ABI      abi.ABI
	Contract *Example
	Eth      interface {
		ethereum.ChainReader
		ethereum.ChainStateReader
		ethereum.TransactionReader
		bind.ContractBackend
	}
}

func (h *BaseExampleProcessor) Setup(address common.Address, eth interface {
	ethereum.ChainReader
	ethereum.ChainStateReader
	ethereum.TransactionReader
	bind.ContractBackend
}) error {
	contract, err := NewExample(address, eth)
	if err != nil {
		return fmt.Errorf("new Example: %w", err)
	}

	abi, err := abi.JSON(strings.NewReader(string(ExampleABI)))
	if err != nil {
		return fmt.Errorf("parsing Example abi: %w", err)
	}

	h.Address = address
	h.ABI = abi
	h.Contract = contract
	h.Eth = eth
	return nil
}

func (h *BaseExampleProcessor) ProcessElement(p interface{}) func(context.Context, types.Log, func(*example.TestInput)) error {
	return func(ctx context.Context, vLog types.Log, cb func(*example.TestInput)) error {
		switch vLog.Topics[0].Hex() {

		case h.ABI.Events["ExampleEvent"].ID.Hex():
			e := new(ExampleExampleEvent)
			if err := h.UnpackLog(e, "ExampleEvent", vLog); err != nil {
				return fmt.Errorf("unpacking ExampleEvent: %w", err)
			}

			e.Raw = vLog
			if err := p.(ExampleProcessor).ProcessExampleEvent(ctx, e, cb); err != nil {
				return fmt.Errorf("processing ExampleEvent: %w", err)
			}

		}
		return nil
	}
}

func (h *BaseExampleProcessor) UnpackLog(out interface{}, event string, log types.Log) error {
	if len(log.Data) > 0 {
		if err := h.ABI.UnpackIntoInterface(out, event, log.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range h.ABI.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return abi.ParseTopics(out, indexed, log.Topics[1:])
}

func (h *BaseExampleProcessor) Initialize(ctx context.Context, start uint64, tx *example.TestInput) error {
	return nil
}

func (h *BaseExampleProcessor) ProcessExampleEvent(ctx context.Context, e *ExampleExampleEvent, cb func(tx *example.TestInput)) error {
	return nil
}

func (h *BaseExampleProcessor) mustEmbedBaseExampleProcessor() {}
