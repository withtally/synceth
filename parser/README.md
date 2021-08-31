# Parser

Golang solidity parser. antlr spec taken from https://github.com/solidity-parser/antlr/blob/master/Solidity.g4

## Generating

```
antlr -Dlanguage=Go parser/Solidity.g4
```