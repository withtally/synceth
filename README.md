# ethgen - A smart contract development toolchain for Go

A simple yet powerful toolchain for Go based smart contract development

- Compile solidity contracts and generate `.abi` and `.bin` artifacts.
- Automatically resolve solidity compiler versions.
- Generate golang bindings from contract artifacts.
- Generate contract fakes for smart contract testing.
- Generate event handlers for smart contract indexing.

## Compiling smart contracts

Create a `generate.go` file and add the following:

```
//go:generate go run -mod=mod github.com/withtally/ethgen compile --outdir ./artifacts ./solidity
```

## Generating bindings

Create a `generate.go` file and add the following:

```
//go:generate go run -mod=mod github.com/withtally/ethgen bind --handlers --fakes --outdir ./bindings ./artifacts
```

## Using Fakes

When `--fakes` is specified for binding generation, smart contract fakes will be generated based on the abi. A smart contract fake implements the full abi interface and adds additional methods for easily setting function return values and emitting events.

For example, a contract like:

```solidity
contract Example {
    event ExampleEvent(string value);
    function getValue() public returns (string value);
}
```

Will generate a fake with interface:

```solidity
contract FakeExample {
    event ExampleEvent(string value);
    function exampleValue() public returns (string value);

    function fakeSetExampleValue(string value);
    function fakeEmitExampleEvent(string value);
}
```

The fake methods can be used to set the return value of the `exampleValue` function and to easily emit the `ExampleEvent`.

Using the fake:

```go
address, tx, c, err := bindings.DeployFakeExample(auth, eth)
if err != nil {
    t.Fatalf("Deploying contract: %v", err)
}

if _, err := c.FakeSetExampleValue(auth, "tally"); err != nil {
    t.Fatalf("Setting value: %v", err)
}

eth.Commit()

v, err := c.exampleValue(nil)
if err != nil {
    t.Fatalf("Getting value: %v", err)
}

println(v) // tally
```
