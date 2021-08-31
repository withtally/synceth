package test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
)

var (
	TestKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
)

func NewAuth(ctx context.Context) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(TestKey, big.NewInt(1337))
	if err != nil {
		return nil, err
	}

	auth.Context = ctx
	return auth, nil
}

func NewTestChain(t testing.TB, auth *bind.TransactOpts) *ethclient.Client {
	t.Helper()

	address := auth.From
	db := rawdb.NewMemoryDatabase()
	genesis := &core.Genesis{
		Config:    params.AllEthashProtocolChanges,
		Alloc:     core.GenesisAlloc{address: {Balance: big.NewInt(10000000000000000)}},
		ExtraData: []byte("test genesis"),
		Timestamp: 9000,
		BaseFee:   big.NewInt(params.InitialBaseFee),
	}
	generate := func(i int, g *core.BlockGen) {
		g.OffsetTime(5)
		g.SetExtra([]byte("test"))
	}
	gblock := genesis.ToBlock(db)
	engine := ethash.NewFaker()
	blocks, _ := core.GenerateChain(params.AllEthashProtocolChanges, gblock, engine, db, 1, generate)
	blocks = append([]*types.Block{gblock}, blocks...)

	// Create node
	n, err := node.New(&node.Config{})
	if err != nil {
		t.Fatalf("can't create new node: %v", err)
	}
	// Create Ethereum Service
	config := &ethconfig.Config{Genesis: genesis}
	config.Ethash.PowMode = ethash.ModeFake
	ethservice, err := eth.New(n, config)
	if err != nil {
		t.Fatalf("can't create new ethereum service: %v", err)
	}
	// Import the test chain.
	if err := n.Start(); err != nil {
		t.Fatalf("can't start test node: %v", err)
	}
	if _, err := ethservice.BlockChain().InsertChain(blocks[1:]); err != nil {
		t.Fatalf("can't import test blocks: %v", err)
	}

	rpc, err := n.Attach()
	if err != nil {
		t.Fatalf("creating rpc: %v", err)
	}
	m := ethservice.Miner()
	go m.Start(auth.From)

	client := ethclient.NewClient(rpc)

	t.Cleanup(func() {
		client.Close()
		m.Stop()
	})

	return client
}

// GETH's simulated backend is missing a few methods from the ethclient that we use.
// This backend augments the simulated backend with those methods to simplify testing.
type SimulatedBackend struct {
	*backends.SimulatedBackend
}

func NewSimulatedBackend(t testing.TB, auth *bind.TransactOpts) *SimulatedBackend {
	t.Helper()
	address := auth.From
	genesisAlloc := map[common.Address]core.GenesisAccount{
		address: {
			Balance: big.NewInt(10000000000000000),
		},
	}

	be := backends.NewSimulatedBackend(genesisAlloc, 10000000)

	// Commit an empty block to avoid negative indexer start blocks.
	be.Commit()
	return &SimulatedBackend{
		be,
	}
}

func (b *SimulatedBackend) NetworkID(ctx context.Context) (*big.Int, error) {
	return b.SimulatedBackend.Blockchain().Config().ChainID, nil
}

func (b *SimulatedBackend) BlockNumber(ctx context.Context) (uint64, error) {
	return b.Blockchain().CurrentBlock().Number().Uint64(), nil
}
