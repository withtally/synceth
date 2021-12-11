package codegen

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/withtally/synceth/solc"
)

const (
	SolidityBase  = "https://github.com/ethereum/solidity/releases/download/"
	SolidityLinux = "solc-static-linux"
	SolidityMacos = "solc-macos"
)

type ContractMetadata struct {
	ABIs  []string
	Bins  []string
	Types []string
	Sigs  []map[string]string
	Libs  map[string]string
}

func ParseContract(src string) (ContractMetadata, error) {
	var (
		abis  []string
		bins  []string
		types []string
		sigs  []map[string]string
		libs  = make(map[string]string)
	)

	contracts, err := solc.CompileSolidityString(src)
	if err != nil {
		return ContractMetadata{}, fmt.Errorf("compiling: %w", err)
	}

	for n, c := range contracts {
		abi, err := json.Marshal(c.Info.AbiDefinition)
		if err != nil {
			return ContractMetadata{}, fmt.Errorf("marshalling abi: %w", err)
		}

		abis = append(abis, string(abi))
		bins = append(bins, c.Code)
		sigs = append(sigs, c.Hashes)
		nameParts := strings.Split(n, ":")
		types = append(types, nameParts[len(nameParts)-1])
		libPattern := crypto.Keccak256Hash([]byte(n)).String()[2:36]
		libs[libPattern] = nameParts[len(nameParts)-1]
	}

	return ContractMetadata{
		ABIs:  abis,
		Bins:  bins,
		Types: types,
		Sigs:  sigs,
		Libs:  libs,
	}, nil
}
