# How to generate golang code for contract deploy and call
An example is blocksigner

step 1: generate contract abi and bin with [solc](https://github.com/ethereum/solidity/releases)
```
- cd contracts/blocksigner
- solc --abi --bin contract/BlockSigner.sol -o contract/build/ --overwrite --optimize
```

step 2:generate golang code
```
- go run ../../cmd/abigen/ --abi contract/build/BlockSigner.abi --bin contract/build/BlockSigner.bin --pkg contract --type BlockSigner --out contract/BlockSigner.go
```
