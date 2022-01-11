package engine

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type Client interface {
	ethereum.ChainReader
	ethereum.ChainStateReader
	ethereum.TransactionReader
	bind.ContractBackend
	ChainID(ctx context.Context) (*big.Int, error)
	BlockNumber(ctx context.Context) (uint64, error)
}
