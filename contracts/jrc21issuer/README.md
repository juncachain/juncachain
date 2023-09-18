## how generate code
solc --abi --bin contract/JRC21Issuer.sol -o contract/build/ --overwrite --optimize

go run ../../cmd/abigen/ --abi contract/build/JRC21Issuer.abi --bin contract/build/JRC21Issuer.bin --pkg contract --type JRC21Issuer --out contract/JRC21Issuer.go
go run ../../cmd/abigen/ --abi contract/build/JRC21PresetFixed.abi --bin contract/build/JRC21PresetFixed.bin --pkg contract --type JRC21PresetFixed --out contract/JRC21PresetFixed.go
go run ../../cmd/abigen/ --abi contract/build/JRC21PresetMinter.abi --bin contract/build/JRC21PresetMinter.bin --pkg contract --type JRC21PresetMinter --out contract/JRC21PresetMinter.go 

