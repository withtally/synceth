package example

import (
	"context"
	"testing"

	"github.com/withtally/ethgen/example/bindings"
	"github.com/withtally/ethgen/test"
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

	// Set return value for `exampleValue`
	if _, err := c.FakeSetExampleValue(auth, "notethgen"); err != nil {
		t.Fatalf("Failed to set example value: %v", err)
	}

	// Commit block to the blockchain.
	eth.Commit()

	want := "notethgen"
	got, err := c.ExampleValue(nil)
	if err != nil {
		t.Fatalf("Calling ExampleValue: %v", err)
	}

	if got != want {
		t.Errorf("ExampleValue got: %s, want: %s", got, want)
	}
}
