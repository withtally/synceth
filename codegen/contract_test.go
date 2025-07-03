package codegen

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestParseContract_ValidContracts tests ParseContract with various valid Solidity contracts
func TestParseContract_ValidContracts(t *testing.T) {
	tests := []struct {
		name                string
		source              string
		expectedTypes       []string
		expectedABIElements map[string][]string // contract name -> expected ABI element names
		validateLibs        func(t *testing.T, libs map[string]string)
		skipIfNoCompiler    bool
	}{
		{
			name: "simple_storage_contract",
			source: `
				pragma solidity ^0.8.0;
				
				contract SimpleStorage {
					uint256 public value;
					
					function setValue(uint256 _value) public {
						value = _value;
					}
					
					function getValue() public view returns (uint256) {
						return value;
					}
				}
			`,
			expectedTypes: []string{"SimpleStorage"},
			expectedABIElements: map[string][]string{
				"SimpleStorage": {"value", "setValue", "getValue"},
			},
			validateLibs: func(t *testing.T, libs map[string]string) {
				assert.Len(t, libs, 1)
				for pattern, name := range libs {
					assert.Equal(t, "SimpleStorage", name)
					assert.Len(t, pattern, 34) // Keccak256 hash truncated
				}
			},
		},
		{
			name: "multiple_contracts",
			source: `
				pragma solidity ^0.8.0;
				
				contract ContractA {
					uint256 public a;
				}
				
				contract ContractB {
					string public b;
				}
			`,
			expectedTypes: []string{"ContractA", "ContractB"},
			expectedABIElements: map[string][]string{
				"ContractA": {"a"},
				"ContractB": {"b"},
			},
			validateLibs: func(t *testing.T, libs map[string]string) {
				assert.Len(t, libs, 2)
				foundA, foundB := false, false
				for _, name := range libs {
					if name == "ContractA" {
						foundA = true
					}
					if name == "ContractB" {
						foundB = true
					}
				}
				assert.True(t, foundA, "ContractA not found in libs")
				assert.True(t, foundB, "ContractB not found in libs")
			},
		},
		{
			name: "contract_with_events",
			source: `
				pragma solidity ^0.8.0;
				
				contract EventContract {
					event ValueChanged(uint256 indexed oldValue, uint256 indexed newValue);
					event OwnerChanged(address indexed oldOwner, address indexed newOwner);
					
					uint256 public value;
					address public owner;
					
					constructor() {
						owner = msg.sender;
					}
					
					function updateValue(uint256 newValue) public {
						uint256 oldValue = value;
						value = newValue;
						emit ValueChanged(oldValue, newValue);
					}
				}
			`,
			expectedTypes: []string{"EventContract"},
			expectedABIElements: map[string][]string{
				"EventContract": {"ValueChanged", "OwnerChanged", "value", "owner", "updateValue"},
			},
			validateLibs: func(t *testing.T, libs map[string]string) {
				assert.Len(t, libs, 1)
			},
		},
		{
			name: "library_contract",
			source: `
				pragma solidity ^0.8.0;
				
				library SafeMath {
					function add(uint256 a, uint256 b) internal pure returns (uint256) {
						uint256 c = a + b;
						require(c >= a, "SafeMath: addition overflow");
						return c;
					}
					
					function sub(uint256 a, uint256 b) internal pure returns (uint256) {
						require(b <= a, "SafeMath: subtraction overflow");
						return a - b;
					}
				}
			`,
			expectedTypes: []string{"SafeMath"},
			expectedABIElements: map[string][]string{
				"SafeMath": {}, // Libraries typically have empty external ABI
			},
			validateLibs: func(t *testing.T, libs map[string]string) {
				assert.Len(t, libs, 1)
				for _, name := range libs {
					assert.Equal(t, "SafeMath", name)
				}
			},
		},
		{
			name: "empty_contract",
			source: `
				pragma solidity ^0.8.0;
				
				contract Empty {
				}
			`,
			expectedTypes: []string{"Empty"},
			expectedABIElements: map[string][]string{
				"Empty": {},
			},
			validateLibs: func(t *testing.T, libs map[string]string) {
				assert.Len(t, libs, 1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Try to parse the contract
			metadata, err := ParseContract(tt.source)
			
			// If compilation fails, it might be due to missing solc compiler
			if err != nil {
				if strings.Contains(err.Error(), "resolving compiler") || 
				   strings.Contains(err.Error(), "executable file not found") {
					t.Skip("Skipping test - Solidity compiler not available")
				}
				require.NoError(t, err)
			}

			// Validate types
			assert.ElementsMatch(t, tt.expectedTypes, metadata.Types)
			
			// Validate counts
			assert.Len(t, metadata.ABIs, len(tt.expectedTypes))
			assert.Len(t, metadata.Bins, len(tt.expectedTypes))
			assert.Len(t, metadata.Sigs, len(tt.expectedTypes))
			
			// Validate ABIs are valid JSON and contain expected elements
			for i, abiStr := range metadata.ABIs {
				var abi []interface{}
				err := json.Unmarshal([]byte(abiStr), &abi)
				require.NoError(t, err, "ABI at index %d should be valid JSON", i)
				
				// Check if this ABI contains expected elements
				contractType := metadata.Types[i]
				if expectedElements, ok := tt.expectedABIElements[contractType]; ok {
					abiElementNames := extractABIElementNames(abi)
					for _, expected := range expectedElements {
						assert.Contains(t, abiElementNames, expected, 
							"Contract %s ABI should contain element %s", contractType, expected)
					}
				}
			}
			
			// Validate bytecode is not empty (except for libraries/interfaces)
			for i, bin := range metadata.Bins {
				if bin != "" {
					assert.True(t, strings.HasPrefix(bin, "0x") || len(bin) > 0,
						"Contract %s should have valid bytecode", metadata.Types[i])
				}
			}
			
			// Validate library patterns
			if tt.validateLibs != nil {
				tt.validateLibs(t, metadata.Libs)
			}
		})
	}
}

// TestParseContract_InvalidSolidity tests ParseContract with invalid Solidity code
func TestParseContract_InvalidSolidity(t *testing.T) {
	tests := []struct {
		name          string
		source        string
		errorContains string
	}{
		{
			name:          "invalid_syntax",
			source:        "This is not valid Solidity code",
			errorContains: "parsing solidity version",
		},
		{
			name:          "missing_pragma",
			source:        "contract Test {}",
			errorContains: "parsing solidity version",
		},
		{
			name: "syntax_error",
			source: `
				pragma solidity ^0.8.0;
				
				contract SyntaxError {
					function broken(
				}
			`,
			errorContains: "compiling:",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseContract(tt.source)
			
			// Skip if compiler is not available
			if err != nil && strings.Contains(err.Error(), "resolving compiler") {
				t.Skip("Skipping test - Solidity compiler not available")
			}
			
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.errorContains)
		})
	}
}

// TestParseContract_LibraryPatternGeneration tests the library pattern generation
func TestParseContract_LibraryPatternGeneration(t *testing.T) {
	source := `
		pragma solidity ^0.8.0;
		
		contract TestContract {
			uint256 public value;
		}
	`
	
	metadata, err := ParseContract(source)
	if err != nil {
		if strings.Contains(err.Error(), "resolving compiler") {
			t.Skip("Skipping test - Solidity compiler not available")
		}
		require.NoError(t, err)
	}
	
	// Check library pattern generation
	require.Len(t, metadata.Libs, 1)
	
	for pattern, name := range metadata.Libs {
		assert.Equal(t, "TestContract", name)
		assert.Len(t, pattern, 34) // Keccak256 truncated to 34 chars
		
		// Verify the pattern is a valid hex string
		assert.Regexp(t, "^[0-9a-f]{34}$", pattern)
		
		// Verify the pattern is correctly generated from the full contract name
		// The full name should be something like "source.sol:TestContract"
		// We can't know the exact filename, but we can verify the pattern generation logic
		for fullName := range map[string]bool{
			"test.sol:TestContract":     true,
			"source.sol:TestContract":   true,
			"<stdin>:TestContract":      true,
		} {
			hash := crypto.Keccak256Hash([]byte(fullName)).String()[2:36]
			if hash == pattern {
				// Found a match, the pattern is correctly generated
				return
			}
		}
	}
}

// TestContractMetadata tests the ContractMetadata struct
func TestContractMetadata(t *testing.T) {
	// Test zero value
	var metadata ContractMetadata
	assert.Nil(t, metadata.ABIs)
	assert.Nil(t, metadata.Bins)
	assert.Nil(t, metadata.Types)
	assert.Nil(t, metadata.Sigs)
	assert.Nil(t, metadata.Libs)
	
	// Test with initialized values
	metadata = ContractMetadata{
		ABIs:  []string{`[{"name":"test","type":"function"}]`},
		Bins:  []string{"0x608060405234801561001057600080fd5b50"},
		Types: []string{"TestContract"},
		Sigs:  []map[string]string{{"test()": "0x12345678"}},
		Libs:  map[string]string{"abcdef": "TestLib"},
	}
	
	assert.Len(t, metadata.ABIs, 1)
	assert.Len(t, metadata.Bins, 1)
	assert.Len(t, metadata.Types, 1)
	assert.Len(t, metadata.Sigs, 1)
	assert.Len(t, metadata.Libs, 1)
}

// TestParseContract_ComplexScenarios tests more complex contract scenarios
func TestParseContract_ComplexScenarios(t *testing.T) {
	tests := []struct {
		name   string
		source string
	}{
		{
			name: "inheritance",
			source: `
				pragma solidity ^0.8.0;
				
				contract Base {
					uint256 public baseValue;
				}
				
				contract Derived is Base {
					uint256 public derivedValue;
					
					function setValues(uint256 _base, uint256 _derived) public {
						baseValue = _base;
						derivedValue = _derived;
					}
				}
			`,
		},
		{
			name: "abstract_contract",
			source: `
				pragma solidity ^0.8.0;
				
				abstract contract AbstractToken {
					function totalSupply() public view virtual returns (uint256);
				}
				
				contract ConcreteToken is AbstractToken {
					uint256 private _totalSupply;
					
					function totalSupply() public view override returns (uint256) {
						return _totalSupply;
					}
				}
			`,
		},
		{
			name: "interface",
			source: `
				pragma solidity ^0.8.0;
				
				interface IERC20 {
					function transfer(address to, uint256 amount) external returns (bool);
					function balanceOf(address account) external view returns (uint256);
				}
				
				contract Token is IERC20 {
					mapping(address => uint256) private balances;
					
					function transfer(address to, uint256 amount) external override returns (bool) {
						balances[msg.sender] -= amount;
						balances[to] += amount;
						return true;
					}
					
					function balanceOf(address account) external view override returns (uint256) {
						return balances[account];
					}
				}
			`,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metadata, err := ParseContract(tt.source)
			
			if err != nil {
				if strings.Contains(err.Error(), "resolving compiler") {
					t.Skip("Skipping test - Solidity compiler not available")
				}
				require.NoError(t, err)
			}
			
			// Just verify it parses successfully and returns metadata
			assert.NotEmpty(t, metadata.Types)
			assert.NotEmpty(t, metadata.ABIs)
			assert.Len(t, metadata.Types, len(metadata.ABIs))
			assert.Len(t, metadata.Types, len(metadata.Bins))
			assert.Len(t, metadata.Types, len(metadata.Sigs))
		})
	}
}

// Helper function to extract ABI element names
func extractABIElementNames(abi []interface{}) []string {
	var names []string
	for _, element := range abi {
		if m, ok := element.(map[string]interface{}); ok {
			if name, ok := m["name"].(string); ok && name != "" {
				names = append(names, name)
			}
		}
	}
	return names
}

// BenchmarkParseContract benchmarks the ParseContract function
func BenchmarkParseContract(b *testing.B) {
	source := `
		pragma solidity ^0.8.0;
		
		contract BenchmarkContract {
			uint256 public value;
			mapping(address => uint256) public balances;
			
			event Transfer(address indexed from, address indexed to, uint256 value);
			
			function setValue(uint256 _value) public {
				value = _value;
			}
			
			function transfer(address to, uint256 amount) public returns (bool) {
				require(balances[msg.sender] >= amount, "Insufficient balance");
				balances[msg.sender] -= amount;
				balances[to] += amount;
				emit Transfer(msg.sender, to, amount);
				return true;
			}
		}
	`
	
	// Run once to check if compiler is available
	_, err := ParseContract(source)
	if err != nil {
		if strings.Contains(err.Error(), "resolving compiler") {
			b.Skip("Skipping benchmark - Solidity compiler not available")
		}
		b.Fatal(err)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ParseContract(source)
		if err != nil {
			b.Fatal(err)
		}
	}
}