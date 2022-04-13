# Examples for implementing applications with chain interactions for the EV3

* Print EV3 key presses tot he console: [examples/key-press/waitforkeys.go]()
* Run a motor on pressing a button: [examples/touchandgo/touch-and-go.go]()
* Control device state and collect metrics via smart contract: [examples/devicecontrol/controlled-drill.go]()

## Compiling Go for ARM
Compile for EV3 or BrickPi3

```GOOS=linux GOARCH=arm GOARM=5 go build```

Copy the compiled binaries to the device, i.e., you don't have to compile on the device.

Execute a program

```./{binaryName}```

Alternatively, if you are using the EV3, you may start the program from the File Browser.


## Ethereum smart contract interaction

Compile go-ethereum devtools
```
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools
```

Compile contracts

```solc --abi --bin --overwrite -o bin solidity/*```

Generate Go contract file

```abigen --bin=bin/{ContractName}.bin --abi=bin/{ContractName}.abi --pkg=contract --type={ContractName} --out=chain/contract/{contractFile}.go```

All in one

```abigen --sol=solidity/{contractFile}.sol --pkg=contract --type={ContractName} --out=chain/contract/{contractFile}.go```


## Resources:
* https://github.com/ev3go/ev3dev
* https://goethereumbook.org/smart-contract-compile/