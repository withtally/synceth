# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

synceth (also known as ethgen) is a smart contract development toolchain for Go that provides compilation, binding generation, and testing utilities for Ethereum smart contracts.

## Essential Commands

### Development
```bash
# Run all tests (builds ethgen binary and runs example tests)
task test

# Run tests for a specific package
go test ./example/test/...

# Run all tests
go test ./...

# Format code
go fmt ./...

# Vet code
go vet ./...
```

### Code Generation Workflow
```bash
# 1. Compile Solidity contracts to artifacts (.abi.json, .bin)
go run . compile --outdir ./artifacts ./path/to/contracts

# 2. Generate Go bindings from artifacts
go run . bind --handlers --fakes --outdir ./bindings ./artifacts
```

## Architecture Overview

### Core Components

1. **Parser (parser/)**: ANTLR-based Solidity parser that analyzes contract structure
   - Entry: `parser/parse.go` - Parses Solidity files into AST
   - Grammar: `parser/Solidity.g4` - Defines Solidity syntax

2. **Compiler Integration (solc/)**: Manages Solidity compiler versions
   - `solc/resolver.go` - Resolves and downloads compiler versions
   - `solc/compile.go` - Handles compilation process

3. **Code Generation (codegen/)**: Generates Go code from contract artifacts
   - `codegen/contract.go` - Main contract binding generation
   - `codegen/fake.go` - Test fake generation with setter/emitter methods
   - `codegen/processor.go` - Processes artifacts and coordinates generation

4. **CLI (cmd/)**: Command-line interface implementations
   - `cmd/compile.go` - Compile command implementation
   - `cmd/bind.go` - Bind command with --handlers and --fakes flags
   - `cmd/root.go` - CLI configuration and command registration

### Key Design Patterns

1. **Fake Contracts**: For testing, generates contracts with additional methods:
   - `fakeSet[MethodName]()` - Set return values for view/pure functions
   - `fakeEmit[EventName]()` - Programmatically emit events
   - Supports struct types as of recent updates

2. **Event Handlers**: Generated with `--handlers` flag, creates typed event handler interfaces for indexing

3. **Compiler Resolution**: Automatically downloads and manages Solidity compiler versions based on pragma statements

### Testing Approach

Tests use simulated blockchain backend (`test.NewSimulatedBackend`) with deployed contracts. See `example/test/example_test.go` for patterns:
- Deploy contracts using generated bindings
- Use `eth.Commit()` to mine blocks
- Test fakes by setting return values before calling methods