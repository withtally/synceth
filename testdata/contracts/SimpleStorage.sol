// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SimpleStorage {
    uint256 private storedValue;
    
    event ValueChanged(uint256 newValue);
    
    function store(uint256 value) public {
        storedValue = value;
        emit ValueChanged(value);
    }
    
    function retrieve() public view returns (uint256) {
        return storedValue;
    }
}