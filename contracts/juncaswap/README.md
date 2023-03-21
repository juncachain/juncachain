## build JuncaswapFactory
```shell
solc0.5.16 --abi --bin contract/JuncaswapFactory.sol -o contract/build --overwrite --optimize
go run ../../cmd/abigen/ --abi contract/build/JuncaswapV2Factory.abi --bin contract/build/JuncaswapV2Factory.bin --pkg contract --type JuncaswapFactory --out contract/JuncaswapFactory.go
```

## build JuncaswapRouter
```shell
solc0.6.6 --abi --bin contract/JuncaswapRouter.sol -o contract/build/ --overwrite --optimize
go run ../../cmd/abigen/ --abi contract/build/JuncaswapV2Router02.abi --bin contract/build/JuncaswapV2Router02.bin --pkg contract --type JuncaswapRouter --out contract/JuncaswapRouter.go
```