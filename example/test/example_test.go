package test

import (
	"context"
	"testing"

	"github.com/withtally/synceth/example/bindings"
	"github.com/withtally/synceth/test"
)

func TestExample(t *testing.T) {
	ctx := context.Background()

	// Create eth signer.
	auth, err := test.NewAuth(ctx)
	if err != nil {
		t.Fatalf("Creating auth: %v", err)
	}

	eth := test.NewSimulatedBackend(t, auth)

	_, _, c, err := bindings.DeployExample(auth, eth)
	if err != nil {
		t.Fatalf("Failed to deploy contract: %v", err)
	}

	// Commit block to the blockchain.
	eth.Commit()

	want := "ethgen"
	got, err := c.ExampleValue(nil)
	if err != nil {
		t.Fatalf("Calling ExampleValue: %v", err)
	}

	if got != want {
		t.Errorf("ExampleValue got: %s, want: %s", got, want)
	}
}

func TestFakeExample(t *testing.T) {
	ctx := context.Background()

	// Create eth signer.
	auth, err := test.NewAuth(ctx)
	if err != nil {
		t.Fatalf("Creating auth: %v", err)
	}

	eth := test.NewSimulatedBackend(t, auth)

	_, _, c, err := bindings.DeployFakeExample(auth, eth)
	if err != nil {
		t.Fatalf("Failed to deploy contract: %v", err)
	}

	// Commit the deployment
	eth.Commit()

	// Set return value for `exampleValue`
	tx, err := c.FakeSetExampleValue(auth, "notethgen")
	if err != nil {
		t.Fatalf("Failed to set example value: %v", err)
	}

	// Commit block to include the transaction
	eth.Commit()

	// Wait for transaction receipt to ensure it's mined
	receipt, err := eth.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		t.Fatalf("Failed to get transaction receipt: %v", err)
	}
	if receipt.Status != 1 {
		t.Fatalf("Transaction failed with status: %d", receipt.Status)
	}

	want := "notethgen"
	got, err := c.ExampleValue(nil)
	if err != nil {
		t.Fatalf("Calling ExampleValue: %v", err)
	}

	if got != want {
		t.Errorf("ExampleValue got: %s, want: %s", got, want)
	}
}
