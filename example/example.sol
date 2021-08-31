// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.3;

contract Example {
    string private _exampleValue = "ethgen";
    event ExampleEvent(string value);

    function exampleValue() public view returns (string memory) {
        return _exampleValue;
    }
}
