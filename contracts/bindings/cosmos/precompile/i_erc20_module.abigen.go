// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package precompile

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IERC20ModuleCoin is an auto generated low-level Go binding around an user-defined struct.
type IERC20ModuleCoin struct {
	Amount *big.Int
	Denom  string
}

// ERC20ModuleMetaData contains all meta data concerning the ERC20Module contract.
var ERC20ModuleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIERC20Module.Coin[]\",\"name\":\"amount\",\"type\":\"tuple[]\"}],\"name\":\"ConvertCoinToErc20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIERC20Module.Coin[]\",\"name\":\"amount\",\"type\":\"tuple[]\"}],\"name\":\"ConvertErc20ToCoin\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"coinDenomForERC20Address\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"token\",\"type\":\"string\"}],\"name\":\"coinDenomForERC20Address\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"convertCoinToERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"convertCoinToERC20From\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"owner\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"convertCoinToERC20From\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"convertCoinToERC20To\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"convertCoinToERC20To\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"convertERC20ToCoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"convertERC20ToCoinFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"owner\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"convertERC20ToCoinFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"convertERC20ToCoinTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"convertERC20ToCoinTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"erc20AddressForCoinDenom\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ERC20ModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20ModuleMetaData.ABI instead.
var ERC20ModuleABI = ERC20ModuleMetaData.ABI

// ERC20Module is an auto generated Go binding around an Ethereum contract.
type ERC20Module struct {
	ERC20ModuleCaller     // Read-only binding to the contract
	ERC20ModuleTransactor // Write-only binding to the contract
	ERC20ModuleFilterer   // Log filterer for contract events
}

// ERC20ModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20ModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20ModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20ModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20ModuleSession struct {
	Contract     *ERC20Module      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20ModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20ModuleCallerSession struct {
	Contract *ERC20ModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC20ModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20ModuleTransactorSession struct {
	Contract     *ERC20ModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC20ModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20ModuleRaw struct {
	Contract *ERC20Module // Generic contract binding to access the raw methods on
}

// ERC20ModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20ModuleCallerRaw struct {
	Contract *ERC20ModuleCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20ModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20ModuleTransactorRaw struct {
	Contract *ERC20ModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Module creates a new instance of ERC20Module, bound to a specific deployed contract.
func NewERC20Module(address common.Address, backend bind.ContractBackend) (*ERC20Module, error) {
	contract, err := bindERC20Module(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Module{ERC20ModuleCaller: ERC20ModuleCaller{contract: contract}, ERC20ModuleTransactor: ERC20ModuleTransactor{contract: contract}, ERC20ModuleFilterer: ERC20ModuleFilterer{contract: contract}}, nil
}

// NewERC20ModuleCaller creates a new read-only instance of ERC20Module, bound to a specific deployed contract.
func NewERC20ModuleCaller(address common.Address, caller bind.ContractCaller) (*ERC20ModuleCaller, error) {
	contract, err := bindERC20Module(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20ModuleCaller{contract: contract}, nil
}

// NewERC20ModuleTransactor creates a new write-only instance of ERC20Module, bound to a specific deployed contract.
func NewERC20ModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20ModuleTransactor, error) {
	contract, err := bindERC20Module(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20ModuleTransactor{contract: contract}, nil
}

// NewERC20ModuleFilterer creates a new log filterer instance of ERC20Module, bound to a specific deployed contract.
func NewERC20ModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20ModuleFilterer, error) {
	contract, err := bindERC20Module(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20ModuleFilterer{contract: contract}, nil
}

// bindERC20Module binds a generic wrapper to an already deployed contract.
func bindERC20Module(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20ModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Module *ERC20ModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Module.Contract.ERC20ModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Module *ERC20ModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Module.Contract.ERC20ModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Module *ERC20ModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Module.Contract.ERC20ModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Module *ERC20ModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Module.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Module *ERC20ModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Module.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Module *ERC20ModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Module.Contract.contract.Transact(opts, method, params...)
}

// CoinDenomForERC20Address is a free data retrieval call binding the contract method 0xcd22a018.
//
// Solidity: function coinDenomForERC20Address(address token) view returns(string)
func (_ERC20Module *ERC20ModuleCaller) CoinDenomForERC20Address(opts *bind.CallOpts, token common.Address) (string, error) {
	var out []interface{}
	err := _ERC20Module.contract.Call(opts, &out, "coinDenomForERC20Address", token)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CoinDenomForERC20Address is a free data retrieval call binding the contract method 0xcd22a018.
//
// Solidity: function coinDenomForERC20Address(address token) view returns(string)
func (_ERC20Module *ERC20ModuleSession) CoinDenomForERC20Address(token common.Address) (string, error) {
	return _ERC20Module.Contract.CoinDenomForERC20Address(&_ERC20Module.CallOpts, token)
}

// CoinDenomForERC20Address is a free data retrieval call binding the contract method 0xcd22a018.
//
// Solidity: function coinDenomForERC20Address(address token) view returns(string)
func (_ERC20Module *ERC20ModuleCallerSession) CoinDenomForERC20Address(token common.Address) (string, error) {
	return _ERC20Module.Contract.CoinDenomForERC20Address(&_ERC20Module.CallOpts, token)
}

// CoinDenomForERC20Address0 is a free data retrieval call binding the contract method 0xe2bea1fe.
//
// Solidity: function coinDenomForERC20Address(string token) view returns(string)
func (_ERC20Module *ERC20ModuleCaller) CoinDenomForERC20Address0(opts *bind.CallOpts, token string) (string, error) {
	var out []interface{}
	err := _ERC20Module.contract.Call(opts, &out, "coinDenomForERC20Address0", token)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CoinDenomForERC20Address0 is a free data retrieval call binding the contract method 0xe2bea1fe.
//
// Solidity: function coinDenomForERC20Address(string token) view returns(string)
func (_ERC20Module *ERC20ModuleSession) CoinDenomForERC20Address0(token string) (string, error) {
	return _ERC20Module.Contract.CoinDenomForERC20Address0(&_ERC20Module.CallOpts, token)
}

// CoinDenomForERC20Address0 is a free data retrieval call binding the contract method 0xe2bea1fe.
//
// Solidity: function coinDenomForERC20Address(string token) view returns(string)
func (_ERC20Module *ERC20ModuleCallerSession) CoinDenomForERC20Address0(token string) (string, error) {
	return _ERC20Module.Contract.CoinDenomForERC20Address0(&_ERC20Module.CallOpts, token)
}

// Erc20AddressForCoinDenom is a free data retrieval call binding the contract method 0xa333e57c.
//
// Solidity: function erc20AddressForCoinDenom(string denom) view returns(address)
func (_ERC20Module *ERC20ModuleCaller) Erc20AddressForCoinDenom(opts *bind.CallOpts, denom string) (common.Address, error) {
	var out []interface{}
	err := _ERC20Module.contract.Call(opts, &out, "erc20AddressForCoinDenom", denom)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20AddressForCoinDenom is a free data retrieval call binding the contract method 0xa333e57c.
//
// Solidity: function erc20AddressForCoinDenom(string denom) view returns(address)
func (_ERC20Module *ERC20ModuleSession) Erc20AddressForCoinDenom(denom string) (common.Address, error) {
	return _ERC20Module.Contract.Erc20AddressForCoinDenom(&_ERC20Module.CallOpts, denom)
}

// Erc20AddressForCoinDenom is a free data retrieval call binding the contract method 0xa333e57c.
//
// Solidity: function erc20AddressForCoinDenom(string denom) view returns(address)
func (_ERC20Module *ERC20ModuleCallerSession) Erc20AddressForCoinDenom(denom string) (common.Address, error) {
	return _ERC20Module.Contract.Erc20AddressForCoinDenom(&_ERC20Module.CallOpts, denom)
}

// ConvertCoinToERC20 is a paid mutator transaction binding the contract method 0xdbeeeb5c.
//
// Solidity: function convertCoinToERC20(string denom, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactor) ConvertCoinToERC20(opts *bind.TransactOpts, denom string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.contract.Transact(opts, "convertCoinToERC20", denom, amount)
}

// ConvertCoinToERC20 is a paid mutator transaction binding the contract method 0xdbeeeb5c.
//
// Solidity: function convertCoinToERC20(string denom, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleSession) ConvertCoinToERC20(denom string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.ConvertCoinToERC20(&_ERC20Module.TransactOpts, denom, amount)
}

// ConvertCoinToERC20 is a paid mutator transaction binding the contract method 0xdbeeeb5c.
//
// Solidity: function convertCoinToERC20(string denom, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactorSession) ConvertCoinToERC20(denom string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.ConvertCoinToERC20(&_ERC20Module.TransactOpts, denom, amount)
}

// ConvertCoinToERC20From is a paid mutator transaction binding the contract method 0x10e0f725.
//
// Solidity: function convertCoinToERC20From(string denom, address owner, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactor) ConvertCoinToERC20From(opts *bind.TransactOpts, denom string, owner common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.contract.Transact(opts, "convertCoinToERC20From", denom, owner, recipient, amount)
}

// ConvertCoinToERC20From is a paid mutator transaction binding the contract method 0x10e0f725.
//
// Solidity: function convertCoinToERC20From(string denom, address owner, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleSession) ConvertCoinToERC20From(denom string, owner common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.ConvertCoinToERC20From(&_ERC20Module.TransactOpts, denom, owner, recipient, amount)
}

// ConvertCoinToERC20From is a paid mutator transaction binding the contract method 0x10e0f725.
//
// Solidity: function convertCoinToERC20From(string denom, address owner, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactorSession) ConvertCoinToERC20From(denom string, owner common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.ConvertCoinToERC20From(&_ERC20Module.TransactOpts, denom, owner, recipient, amount)
}

// ConvertCoinToERC20From0 is a paid mutator transaction binding the contract method 0x7e943c78.
//
// Solidity: function convertCoinToERC20From(string denom, string owner, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactor) ConvertCoinToERC20From0(opts *bind.TransactOpts, denom string, owner string, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.contract.Transact(opts, "convertCoinToERC20From0", denom, owner, recipient, amount)
}

// ConvertCoinToERC20From0 is a paid mutator transaction binding the contract method 0x7e943c78.
//
// Solidity: function convertCoinToERC20From(string denom, string owner, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleSession) ConvertCoinToERC20From0(denom string, owner string, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.ConvertCoinToERC20From0(&_ERC20Module.TransactOpts, denom, owner, recipient, amount)
}

// ConvertCoinToERC20From0 is a paid mutator transaction binding the contract method 0x7e943c78.
//
// Solidity: function convertCoinToERC20From(string denom, string owner, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactorSession) ConvertCoinToERC20From0(denom string, owner string, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.ConvertCoinToERC20From0(&_ERC20Module.TransactOpts, denom, owner, recipient, amount)
}

// ConvertCoinToERC20To is a paid mutator transaction binding the contract method 0x501c3515.
//
// Solidity: function convertCoinToERC20To(string denom, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactor) ConvertCoinToERC20To(opts *bind.TransactOpts, denom string, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.contract.Transact(opts, "convertCoinToERC20To", denom, recipient, amount)
}

// ConvertCoinToERC20To is a paid mutator transaction binding the contract method 0x501c3515.
//
// Solidity: function convertCoinToERC20To(string denom, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleSession) ConvertCoinToERC20To(denom string, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.ConvertCoinToERC20To(&_ERC20Module.TransactOpts, denom, recipient, amount)
}

// ConvertCoinToERC20To is a paid mutator transaction binding the contract method 0x501c3515.
//
// Solidity: function convertCoinToERC20To(string denom, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactorSession) ConvertCoinToERC20To(denom string, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.ConvertCoinToERC20To(&_ERC20Module.TransactOpts, denom, recipient, amount)
}

// ConvertCoinToERC20To0 is a paid mutator transaction binding the contract method 0xacf58886.
//
// Solidity: function convertCoinToERC20To(string denom, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactor) ConvertCoinToERC20To0(opts *bind.TransactOpts, denom string, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.contract.Transact(opts, "convertCoinToERC20To0", denom, recipient, amount)
}

// ConvertCoinToERC20To0 is a paid mutator transaction binding the contract method 0xacf58886.
//
// Solidity: function convertCoinToERC20To(string denom, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleSession) ConvertCoinToERC20To0(denom string, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.ConvertCoinToERC20To0(&_ERC20Module.TransactOpts, denom, recipient, amount)
}

// ConvertCoinToERC20To0 is a paid mutator transaction binding the contract method 0xacf58886.
//
// Solidity: function convertCoinToERC20To(string denom, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactorSession) ConvertCoinToERC20To0(denom string, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.ConvertCoinToERC20To0(&_ERC20Module.TransactOpts, denom, recipient, amount)
}

// ConvertERC20ToCoin is a paid mutator transaction binding the contract method 0x3acdb33b.
//
// Solidity: function convertERC20ToCoin(address token, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactor) ConvertERC20ToCoin(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.contract.Transact(opts, "convertERC20ToCoin", token, amount)
}

// ConvertERC20ToCoin is a paid mutator transaction binding the contract method 0x3acdb33b.
//
// Solidity: function convertERC20ToCoin(address token, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleSession) ConvertERC20ToCoin(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.ConvertERC20ToCoin(&_ERC20Module.TransactOpts, token, amount)
}

// ConvertERC20ToCoin is a paid mutator transaction binding the contract method 0x3acdb33b.
//
// Solidity: function convertERC20ToCoin(address token, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactorSession) ConvertERC20ToCoin(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.ConvertERC20ToCoin(&_ERC20Module.TransactOpts, token, amount)
}

// ConvertERC20ToCoinFrom is a paid mutator transaction binding the contract method 0x24b8f0fe.
//
// Solidity: function convertERC20ToCoinFrom(address token, address owner, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactor) ConvertERC20ToCoinFrom(opts *bind.TransactOpts, token common.Address, owner common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.contract.Transact(opts, "convertERC20ToCoinFrom", token, owner, recipient, amount)
}

// ConvertERC20ToCoinFrom is a paid mutator transaction binding the contract method 0x24b8f0fe.
//
// Solidity: function convertERC20ToCoinFrom(address token, address owner, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleSession) ConvertERC20ToCoinFrom(token common.Address, owner common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.ConvertERC20ToCoinFrom(&_ERC20Module.TransactOpts, token, owner, recipient, amount)
}

// ConvertERC20ToCoinFrom is a paid mutator transaction binding the contract method 0x24b8f0fe.
//
// Solidity: function convertERC20ToCoinFrom(address token, address owner, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactorSession) ConvertERC20ToCoinFrom(token common.Address, owner common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.ConvertERC20ToCoinFrom(&_ERC20Module.TransactOpts, token, owner, recipient, amount)
}

// ConvertERC20ToCoinFrom0 is a paid mutator transaction binding the contract method 0xcab01e7a.
//
// Solidity: function convertERC20ToCoinFrom(address token, string owner, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactor) ConvertERC20ToCoinFrom0(opts *bind.TransactOpts, token common.Address, owner string, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.contract.Transact(opts, "convertERC20ToCoinFrom0", token, owner, recipient, amount)
}

// ConvertERC20ToCoinFrom0 is a paid mutator transaction binding the contract method 0xcab01e7a.
//
// Solidity: function convertERC20ToCoinFrom(address token, string owner, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleSession) ConvertERC20ToCoinFrom0(token common.Address, owner string, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.ConvertERC20ToCoinFrom0(&_ERC20Module.TransactOpts, token, owner, recipient, amount)
}

// ConvertERC20ToCoinFrom0 is a paid mutator transaction binding the contract method 0xcab01e7a.
//
// Solidity: function convertERC20ToCoinFrom(address token, string owner, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactorSession) ConvertERC20ToCoinFrom0(token common.Address, owner string, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.ConvertERC20ToCoinFrom0(&_ERC20Module.TransactOpts, token, owner, recipient, amount)
}

// ConvertERC20ToCoinTo is a paid mutator transaction binding the contract method 0xd20551d6.
//
// Solidity: function convertERC20ToCoinTo(address token, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactor) ConvertERC20ToCoinTo(opts *bind.TransactOpts, token common.Address, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.contract.Transact(opts, "convertERC20ToCoinTo", token, recipient, amount)
}

// ConvertERC20ToCoinTo is a paid mutator transaction binding the contract method 0xd20551d6.
//
// Solidity: function convertERC20ToCoinTo(address token, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleSession) ConvertERC20ToCoinTo(token common.Address, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.ConvertERC20ToCoinTo(&_ERC20Module.TransactOpts, token, recipient, amount)
}

// ConvertERC20ToCoinTo is a paid mutator transaction binding the contract method 0xd20551d6.
//
// Solidity: function convertERC20ToCoinTo(address token, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactorSession) ConvertERC20ToCoinTo(token common.Address, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.ConvertERC20ToCoinTo(&_ERC20Module.TransactOpts, token, recipient, amount)
}

// ConvertERC20ToCoinTo0 is a paid mutator transaction binding the contract method 0xf7ea84bc.
//
// Solidity: function convertERC20ToCoinTo(address token, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactor) ConvertERC20ToCoinTo0(opts *bind.TransactOpts, token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.contract.Transact(opts, "convertERC20ToCoinTo0", token, recipient, amount)
}

// ConvertERC20ToCoinTo0 is a paid mutator transaction binding the contract method 0xf7ea84bc.
//
// Solidity: function convertERC20ToCoinTo(address token, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleSession) ConvertERC20ToCoinTo0(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.ConvertERC20ToCoinTo0(&_ERC20Module.TransactOpts, token, recipient, amount)
}

// ConvertERC20ToCoinTo0 is a paid mutator transaction binding the contract method 0xf7ea84bc.
//
// Solidity: function convertERC20ToCoinTo(address token, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactorSession) ConvertERC20ToCoinTo0(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.ConvertERC20ToCoinTo0(&_ERC20Module.TransactOpts, token, recipient, amount)
}

// ERC20ModuleConvertCoinToErc20Iterator is returned from FilterConvertCoinToErc20 and is used to iterate over the raw logs and unpacked data for ConvertCoinToErc20 events raised by the ERC20Module contract.
type ERC20ModuleConvertCoinToErc20Iterator struct {
	Event *ERC20ModuleConvertCoinToErc20 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20ModuleConvertCoinToErc20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20ModuleConvertCoinToErc20)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20ModuleConvertCoinToErc20)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20ModuleConvertCoinToErc20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ModuleConvertCoinToErc20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20ModuleConvertCoinToErc20 represents a ConvertCoinToErc20 event raised by the ERC20Module contract.
type ERC20ModuleConvertCoinToErc20 struct {
	Denom     common.Hash
	Owner     common.Address
	Recipient common.Address
	Amount    []IERC20ModuleCoin
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConvertCoinToErc20 is a free log retrieval operation binding the contract event 0xf4df46dd20cc8c41d752241beb83f671035c02359c85f28648aec90fea873c2b.
//
// Solidity: event ConvertCoinToErc20(string indexed denom, address indexed owner, address indexed recipient, (uint256,string)[] amount)
func (_ERC20Module *ERC20ModuleFilterer) FilterConvertCoinToErc20(opts *bind.FilterOpts, denom []string, owner []common.Address, recipient []common.Address) (*ERC20ModuleConvertCoinToErc20Iterator, error) {

	var denomRule []interface{}
	for _, denomItem := range denom {
		denomRule = append(denomRule, denomItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20Module.contract.FilterLogs(opts, "ConvertCoinToErc20", denomRule, ownerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ModuleConvertCoinToErc20Iterator{contract: _ERC20Module.contract, event: "ConvertCoinToErc20", logs: logs, sub: sub}, nil
}

// WatchConvertCoinToErc20 is a free log subscription operation binding the contract event 0xf4df46dd20cc8c41d752241beb83f671035c02359c85f28648aec90fea873c2b.
//
// Solidity: event ConvertCoinToErc20(string indexed denom, address indexed owner, address indexed recipient, (uint256,string)[] amount)
func (_ERC20Module *ERC20ModuleFilterer) WatchConvertCoinToErc20(opts *bind.WatchOpts, sink chan<- *ERC20ModuleConvertCoinToErc20, denom []string, owner []common.Address, recipient []common.Address) (event.Subscription, error) {

	var denomRule []interface{}
	for _, denomItem := range denom {
		denomRule = append(denomRule, denomItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20Module.contract.WatchLogs(opts, "ConvertCoinToErc20", denomRule, ownerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20ModuleConvertCoinToErc20)
				if err := _ERC20Module.contract.UnpackLog(event, "ConvertCoinToErc20", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConvertCoinToErc20 is a log parse operation binding the contract event 0xf4df46dd20cc8c41d752241beb83f671035c02359c85f28648aec90fea873c2b.
//
// Solidity: event ConvertCoinToErc20(string indexed denom, address indexed owner, address indexed recipient, (uint256,string)[] amount)
func (_ERC20Module *ERC20ModuleFilterer) ParseConvertCoinToErc20(log types.Log) (*ERC20ModuleConvertCoinToErc20, error) {
	event := new(ERC20ModuleConvertCoinToErc20)
	if err := _ERC20Module.contract.UnpackLog(event, "ConvertCoinToErc20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20ModuleConvertErc20ToCoinIterator is returned from FilterConvertErc20ToCoin and is used to iterate over the raw logs and unpacked data for ConvertErc20ToCoin events raised by the ERC20Module contract.
type ERC20ModuleConvertErc20ToCoinIterator struct {
	Event *ERC20ModuleConvertErc20ToCoin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20ModuleConvertErc20ToCoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20ModuleConvertErc20ToCoin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20ModuleConvertErc20ToCoin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20ModuleConvertErc20ToCoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ModuleConvertErc20ToCoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20ModuleConvertErc20ToCoin represents a ConvertErc20ToCoin event raised by the ERC20Module contract.
type ERC20ModuleConvertErc20ToCoin struct {
	Token     common.Address
	Owner     common.Address
	Recipient common.Address
	Amount    []IERC20ModuleCoin
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConvertErc20ToCoin is a free log retrieval operation binding the contract event 0xd6886038dca20553bec918c9dfc0e3e876b5b7e996ee6599100d49356c61903a.
//
// Solidity: event ConvertErc20ToCoin(address indexed token, address indexed owner, address indexed recipient, (uint256,string)[] amount)
func (_ERC20Module *ERC20ModuleFilterer) FilterConvertErc20ToCoin(opts *bind.FilterOpts, token []common.Address, owner []common.Address, recipient []common.Address) (*ERC20ModuleConvertErc20ToCoinIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20Module.contract.FilterLogs(opts, "ConvertErc20ToCoin", tokenRule, ownerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ModuleConvertErc20ToCoinIterator{contract: _ERC20Module.contract, event: "ConvertErc20ToCoin", logs: logs, sub: sub}, nil
}

// WatchConvertErc20ToCoin is a free log subscription operation binding the contract event 0xd6886038dca20553bec918c9dfc0e3e876b5b7e996ee6599100d49356c61903a.
//
// Solidity: event ConvertErc20ToCoin(address indexed token, address indexed owner, address indexed recipient, (uint256,string)[] amount)
func (_ERC20Module *ERC20ModuleFilterer) WatchConvertErc20ToCoin(opts *bind.WatchOpts, sink chan<- *ERC20ModuleConvertErc20ToCoin, token []common.Address, owner []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20Module.contract.WatchLogs(opts, "ConvertErc20ToCoin", tokenRule, ownerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20ModuleConvertErc20ToCoin)
				if err := _ERC20Module.contract.UnpackLog(event, "ConvertErc20ToCoin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConvertErc20ToCoin is a log parse operation binding the contract event 0xd6886038dca20553bec918c9dfc0e3e876b5b7e996ee6599100d49356c61903a.
//
// Solidity: event ConvertErc20ToCoin(address indexed token, address indexed owner, address indexed recipient, (uint256,string)[] amount)
func (_ERC20Module *ERC20ModuleFilterer) ParseConvertErc20ToCoin(log types.Log) (*ERC20ModuleConvertErc20ToCoin, error) {
	event := new(ERC20ModuleConvertErc20ToCoin)
	if err := _ERC20Module.contract.UnpackLog(event, "ConvertErc20ToCoin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
