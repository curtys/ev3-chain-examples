// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DeviceControllerABI is the input ABI used to generate the binding from.
const DeviceControllerABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"devices\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ready\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"metric\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"deviceType\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_deviceId\",\"type\":\"string\"}],\"name\":\"getReadyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"readyState\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_deviceId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_type\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_deviceId\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"_delta\",\"type\":\"uint32\"}],\"name\":\"updateMetric\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"metric\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_deviceId\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_ready\",\"type\":\"bool\"}],\"name\":\"updateReadyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DeviceControllerFuncSigs maps the 4-byte function signature to its string representation.
var DeviceControllerFuncSigs = map[string]string{
	"22233396": "devices(string)",
	"924211f1": "getReadyState(string)",
	"3ffbd47f": "register(string,string)",
	"b0185690": "updateMetric(string,uint32)",
	"8c239bca": "updateReadyState(string,bool)",
}

// DeviceControllerBin is the compiled bytecode used for deploying new contracts.
var DeviceControllerBin = "0x608060405234801561001057600080fd5b506108b9806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063222333961461005c5780633ffbd47f146100875780638c239bca1461009c578063924211f1146100af578063b0185690146100d2575b600080fd5b61006f61006a36600461053d565b6100fa565b60405161007e939291906105ee565b60405180910390f35b61009a6100953660046106a3565b6101bb565b005b61009a6100aa36600461070f565b6102c5565b6100c26100bd36600461076b565b610338565b604051901515815260200161007e565b6100e56100e03660046107ad565b6103a2565b60405163ffffffff909116815260200161007e565b80516020818301810180516000825292820191909301209152805460018201805460ff83169361010090930463ffffffff1692919061013890610802565b80601f016020809104026020016040519081016040528092919081815260200182805461016490610802565b80156101b15780601f10610186576101008083540402835291602001916101b1565b820191906000526020600020905b81548152906001019060200180831161019457829003601f168201915b5050505050905083565b600084846040516101cd92919061083d565b908152602001604051809103902060010180546101e990610802565b1590506101f557600080fd5b6040518060600160405280600115158152602001600063ffffffff16815260200183838080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250939094525050604051909150610261908790879061083d565b90815260408051918290036020908101909220835181548585015163ffffffff166101000264ffffffff00199215159290921664ffffffffff199091161717815590830151805191926102bc9260018501929091019061048e565b50505050505050565b60008084846040516102d892919061083d565b908152602001604051809103902060010180546102f490610802565b90501161030057600080fd5b806000848460405161031392919061083d565b908152604051908190036020019020805491151560ff19909216919091179055505050565b6000806000848460405161034d92919061083d565b9081526020016040518091039020600101805461036990610802565b90501161037557600080fd5b6000838360405161038792919061083d565b9081526040519081900360200190205460ff16905092915050565b600080600085856040516103b792919061083d565b908152602001604051809103902060010180546103d390610802565b9050116103df57600080fd5b81600085856040516103f292919061083d565b90815260405190819003602001902054610418919063ffffffff6101009091041661084d565b6000858560405161042a92919061083d565b908152604051908190036020018120805463ffffffff939093166101000264ffffffff001990931692909217909155600090610469908690869061083d565b9081526040519081900360200190205463ffffffff6101009091041690509392505050565b82805461049a90610802565b90600052602060002090601f0160209004810192826104bc5760008555610502565b82601f106104d557805160ff1916838001178555610502565b82800160010185558215610502579182015b828111156105025782518255916020019190600101906104e7565b5061050e929150610512565b5090565b5b8082111561050e5760008155600101610513565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561054f57600080fd5b813567ffffffffffffffff8082111561056757600080fd5b818401915084601f83011261057b57600080fd5b81358181111561058d5761058d610527565b604051601f8201601f19908116603f011681019083821181831017156105b5576105b5610527565b816040528281528760208487010111156105ce57600080fd5b826020860160208301376000928101602001929092525095945050505050565b83151581526000602063ffffffff85168184015260606040840152835180606085015260005b8181101561063057858101830151858201608001528201610614565b81811115610642576000608083870101525b50601f01601f19169290920160800195945050505050565b60008083601f84011261066c57600080fd5b50813567ffffffffffffffff81111561068457600080fd5b60208301915083602082850101111561069c57600080fd5b9250929050565b600080600080604085870312156106b957600080fd5b843567ffffffffffffffff808211156106d157600080fd5b6106dd8883890161065a565b909650945060208701359150808211156106f657600080fd5b506107038782880161065a565b95989497509550505050565b60008060006040848603121561072457600080fd5b833567ffffffffffffffff81111561073b57600080fd5b6107478682870161065a565b9094509250506020840135801515811461076057600080fd5b809150509250925092565b6000806020838503121561077e57600080fd5b823567ffffffffffffffff81111561079557600080fd5b6107a18582860161065a565b90969095509350505050565b6000806000604084860312156107c257600080fd5b833567ffffffffffffffff8111156107d957600080fd5b6107e58682870161065a565b909450925050602084013563ffffffff8116811461076057600080fd5b600181811c9082168061081657607f821691505b6020821081141561083757634e487b7160e01b600052602260045260246000fd5b50919050565b8183823760009101908152919050565b600063ffffffff80831681851680830382111561087a57634e487b7160e01b600052601160045260246000fd5b0194935050505056fea2646970667358221220439776788319ccedbaf190576045dde050da67b8c15d0893e2938becc160c50664736f6c63430008090033"

// DeployDeviceController deploys a new Ethereum contract, binding an instance of DeviceController to it.
func DeployDeviceController(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DeviceController, error) {
	parsed, err := abi.JSON(strings.NewReader(DeviceControllerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DeviceControllerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DeviceController{DeviceControllerCaller: DeviceControllerCaller{contract: contract}, DeviceControllerTransactor: DeviceControllerTransactor{contract: contract}, DeviceControllerFilterer: DeviceControllerFilterer{contract: contract}}, nil
}

// DeviceController is an auto generated Go binding around an Ethereum contract.
type DeviceController struct {
	DeviceControllerCaller     // Read-only binding to the contract
	DeviceControllerTransactor // Write-only binding to the contract
	DeviceControllerFilterer   // Log filterer for contract events
}

// DeviceControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeviceControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeviceControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeviceControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeviceControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeviceControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeviceControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeviceControllerSession struct {
	Contract     *DeviceController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeviceControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeviceControllerCallerSession struct {
	Contract *DeviceControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DeviceControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeviceControllerTransactorSession struct {
	Contract     *DeviceControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DeviceControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeviceControllerRaw struct {
	Contract *DeviceController // Generic contract binding to access the raw methods on
}

// DeviceControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeviceControllerCallerRaw struct {
	Contract *DeviceControllerCaller // Generic read-only contract binding to access the raw methods on
}

// DeviceControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeviceControllerTransactorRaw struct {
	Contract *DeviceControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeviceController creates a new instance of DeviceController, bound to a specific deployed contract.
func NewDeviceController(address common.Address, backend bind.ContractBackend) (*DeviceController, error) {
	contract, err := bindDeviceController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DeviceController{DeviceControllerCaller: DeviceControllerCaller{contract: contract}, DeviceControllerTransactor: DeviceControllerTransactor{contract: contract}, DeviceControllerFilterer: DeviceControllerFilterer{contract: contract}}, nil
}

// NewDeviceControllerCaller creates a new read-only instance of DeviceController, bound to a specific deployed contract.
func NewDeviceControllerCaller(address common.Address, caller bind.ContractCaller) (*DeviceControllerCaller, error) {
	contract, err := bindDeviceController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeviceControllerCaller{contract: contract}, nil
}

// NewDeviceControllerTransactor creates a new write-only instance of DeviceController, bound to a specific deployed contract.
func NewDeviceControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*DeviceControllerTransactor, error) {
	contract, err := bindDeviceController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeviceControllerTransactor{contract: contract}, nil
}

// NewDeviceControllerFilterer creates a new log filterer instance of DeviceController, bound to a specific deployed contract.
func NewDeviceControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*DeviceControllerFilterer, error) {
	contract, err := bindDeviceController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeviceControllerFilterer{contract: contract}, nil
}

// bindDeviceController binds a generic wrapper to an already deployed contract.
func bindDeviceController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DeviceControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeviceController *DeviceControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeviceController.Contract.DeviceControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeviceController *DeviceControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeviceController.Contract.DeviceControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeviceController *DeviceControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeviceController.Contract.DeviceControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeviceController *DeviceControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeviceController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeviceController *DeviceControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeviceController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeviceController *DeviceControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeviceController.Contract.contract.Transact(opts, method, params...)
}

// Devices is a free data retrieval call binding the contract method 0x22233396.
//
// Solidity: function devices(string ) view returns(bool ready, uint32 metric, string deviceType)
func (_DeviceController *DeviceControllerCaller) Devices(opts *bind.CallOpts, arg0 string) (struct {
	Ready      bool
	Metric     uint32
	DeviceType string
}, error) {
	var out []interface{}
	err := _DeviceController.contract.Call(opts, &out, "devices", arg0)

	outstruct := new(struct {
		Ready      bool
		Metric     uint32
		DeviceType string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ready = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Metric = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.DeviceType = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// Devices is a free data retrieval call binding the contract method 0x22233396.
//
// Solidity: function devices(string ) view returns(bool ready, uint32 metric, string deviceType)
func (_DeviceController *DeviceControllerSession) Devices(arg0 string) (struct {
	Ready      bool
	Metric     uint32
	DeviceType string
}, error) {
	return _DeviceController.Contract.Devices(&_DeviceController.CallOpts, arg0)
}

// Devices is a free data retrieval call binding the contract method 0x22233396.
//
// Solidity: function devices(string ) view returns(bool ready, uint32 metric, string deviceType)
func (_DeviceController *DeviceControllerCallerSession) Devices(arg0 string) (struct {
	Ready      bool
	Metric     uint32
	DeviceType string
}, error) {
	return _DeviceController.Contract.Devices(&_DeviceController.CallOpts, arg0)
}

// GetReadyState is a free data retrieval call binding the contract method 0x924211f1.
//
// Solidity: function getReadyState(string _deviceId) view returns(bool readyState)
func (_DeviceController *DeviceControllerCaller) GetReadyState(opts *bind.CallOpts, _deviceId string) (bool, error) {
	var out []interface{}
	err := _DeviceController.contract.Call(opts, &out, "getReadyState", _deviceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetReadyState is a free data retrieval call binding the contract method 0x924211f1.
//
// Solidity: function getReadyState(string _deviceId) view returns(bool readyState)
func (_DeviceController *DeviceControllerSession) GetReadyState(_deviceId string) (bool, error) {
	return _DeviceController.Contract.GetReadyState(&_DeviceController.CallOpts, _deviceId)
}

// GetReadyState is a free data retrieval call binding the contract method 0x924211f1.
//
// Solidity: function getReadyState(string _deviceId) view returns(bool readyState)
func (_DeviceController *DeviceControllerCallerSession) GetReadyState(_deviceId string) (bool, error) {
	return _DeviceController.Contract.GetReadyState(&_DeviceController.CallOpts, _deviceId)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string _deviceId, string _type) returns()
func (_DeviceController *DeviceControllerTransactor) Register(opts *bind.TransactOpts, _deviceId string, _type string) (*types.Transaction, error) {
	return _DeviceController.contract.Transact(opts, "register", _deviceId, _type)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string _deviceId, string _type) returns()
func (_DeviceController *DeviceControllerSession) Register(_deviceId string, _type string) (*types.Transaction, error) {
	return _DeviceController.Contract.Register(&_DeviceController.TransactOpts, _deviceId, _type)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string _deviceId, string _type) returns()
func (_DeviceController *DeviceControllerTransactorSession) Register(_deviceId string, _type string) (*types.Transaction, error) {
	return _DeviceController.Contract.Register(&_DeviceController.TransactOpts, _deviceId, _type)
}

// UpdateMetric is a paid mutator transaction binding the contract method 0xb0185690.
//
// Solidity: function updateMetric(string _deviceId, uint32 _delta) returns(uint32 metric)
func (_DeviceController *DeviceControllerTransactor) UpdateMetric(opts *bind.TransactOpts, _deviceId string, _delta uint32) (*types.Transaction, error) {
	return _DeviceController.contract.Transact(opts, "updateMetric", _deviceId, _delta)
}

// UpdateMetric is a paid mutator transaction binding the contract method 0xb0185690.
//
// Solidity: function updateMetric(string _deviceId, uint32 _delta) returns(uint32 metric)
func (_DeviceController *DeviceControllerSession) UpdateMetric(_deviceId string, _delta uint32) (*types.Transaction, error) {
	return _DeviceController.Contract.UpdateMetric(&_DeviceController.TransactOpts, _deviceId, _delta)
}

// UpdateMetric is a paid mutator transaction binding the contract method 0xb0185690.
//
// Solidity: function updateMetric(string _deviceId, uint32 _delta) returns(uint32 metric)
func (_DeviceController *DeviceControllerTransactorSession) UpdateMetric(_deviceId string, _delta uint32) (*types.Transaction, error) {
	return _DeviceController.Contract.UpdateMetric(&_DeviceController.TransactOpts, _deviceId, _delta)
}

// UpdateReadyState is a paid mutator transaction binding the contract method 0x8c239bca.
//
// Solidity: function updateReadyState(string _deviceId, bool _ready) returns()
func (_DeviceController *DeviceControllerTransactor) UpdateReadyState(opts *bind.TransactOpts, _deviceId string, _ready bool) (*types.Transaction, error) {
	return _DeviceController.contract.Transact(opts, "updateReadyState", _deviceId, _ready)
}

// UpdateReadyState is a paid mutator transaction binding the contract method 0x8c239bca.
//
// Solidity: function updateReadyState(string _deviceId, bool _ready) returns()
func (_DeviceController *DeviceControllerSession) UpdateReadyState(_deviceId string, _ready bool) (*types.Transaction, error) {
	return _DeviceController.Contract.UpdateReadyState(&_DeviceController.TransactOpts, _deviceId, _ready)
}

// UpdateReadyState is a paid mutator transaction binding the contract method 0x8c239bca.
//
// Solidity: function updateReadyState(string _deviceId, bool _ready) returns()
func (_DeviceController *DeviceControllerTransactorSession) UpdateReadyState(_deviceId string, _ready bool) (*types.Transaction, error) {
	return _DeviceController.Contract.UpdateReadyState(&_DeviceController.TransactOpts, _deviceId, _ready)
}
