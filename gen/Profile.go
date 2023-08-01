// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package profile

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

// ProfileProfileStruct is an auto generated low-level Go binding around an user-defined struct.
type ProfileProfileStruct struct {
	ProfileAddress common.Address
	Username       string
	ProfileQRCode  string
}

// ProfileMetaData contains all meta data concerning the Profile contract.
var ProfileMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_qrCode\",\"type\":\"string\"}],\"name\":\"createProfile\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"profileAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"profileQRCode\",\"type\":\"string\"}],\"internalType\":\"structProfile.ProfileStruct\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"profiles\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"profileAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"profileQRCode\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610d108061005c5f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c80638da5cb5b14610043578063ba1feb8a14610061578063bbe1562714610091575b5f80fd5b61004b6100c3565b60405161005891906105e7565b60405180910390f35b61007b6004803603810190610076919061074d565b6100e6565b60405161008891906108a0565b60405180910390f35b6100ab60048036038101906100a691906108ea565b610421565b6040516100ba9392919061095d565b60405180910390f35b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6100ee610572565b5f73ffffffffffffffffffffffffffffffffffffffff1660015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146101ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101b1906109ea565b60405180910390fd5b60405180606001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018381525060015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010190816102819190610c0b565b5060408201518160020190816102979190610c0b565b5090506040518060600160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461030990610a35565b80601f016020809104026020016040519081016040528092919081815260200182805461033590610a35565b80156103805780601f1061035757610100808354040283529160200191610380565b820191905f5260205f20905b81548152906001019060200180831161036357829003601f168201915b5050505050815260200160028201805461039990610a35565b80601f01602080910402602001604051908101604052809291908181526020018280546103c590610a35565b80156104105780601f106103e757610100808354040283529160200191610410565b820191905f5260205f20905b8154815290600101906020018083116103f357829003601f168201915b505050505081525050905092915050565b6001602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600101805461046590610a35565b80601f016020809104026020016040519081016040528092919081815260200182805461049190610a35565b80156104dc5780601f106104b3576101008083540402835291602001916104dc565b820191905f5260205f20905b8154815290600101906020018083116104bf57829003601f168201915b5050505050908060020180546104f190610a35565b80601f016020809104026020016040519081016040528092919081815260200182805461051d90610a35565b80156105685780601f1061053f57610100808354040283529160200191610568565b820191905f5260205f20905b81548152906001019060200180831161054b57829003601f168201915b5050505050905083565b60405180606001604052805f73ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001606081525090565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6105d1826105a8565b9050919050565b6105e1816105c7565b82525050565b5f6020820190506105fa5f8301846105d8565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61065f82610619565b810181811067ffffffffffffffff8211171561067e5761067d610629565b5b80604052505050565b5f610690610600565b905061069c8282610656565b919050565b5f67ffffffffffffffff8211156106bb576106ba610629565b5b6106c482610619565b9050602081019050919050565b828183375f83830152505050565b5f6106f16106ec846106a1565b610687565b90508281526020810184848401111561070d5761070c610615565b5b6107188482856106d1565b509392505050565b5f82601f83011261073457610733610611565b5b81356107448482602086016106df565b91505092915050565b5f806040838503121561076357610762610609565b5b5f83013567ffffffffffffffff8111156107805761077f61060d565b5b61078c85828601610720565b925050602083013567ffffffffffffffff8111156107ad576107ac61060d565b5b6107b985828601610720565b9150509250929050565b6107cc816105c7565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156108095780820151818401526020810190506107ee565b5f8484015250505050565b5f61081e826107d2565b61082881856107dc565b93506108388185602086016107ec565b61084181610619565b840191505092915050565b5f606083015f8301516108615f8601826107c3565b50602083015184820360208601526108798282610814565b915050604083015184820360408601526108938282610814565b9150508091505092915050565b5f6020820190508181035f8301526108b8818461084c565b905092915050565b6108c9816105c7565b81146108d3575f80fd5b50565b5f813590506108e4816108c0565b92915050565b5f602082840312156108ff576108fe610609565b5b5f61090c848285016108d6565b91505092915050565b5f82825260208201905092915050565b5f61092f826107d2565b6109398185610915565b93506109498185602086016107ec565b61095281610619565b840191505092915050565b5f6060820190506109705f8301866105d8565b81810360208301526109828185610925565b905081810360408301526109968184610925565b9050949350505050565b7f50726f66696c653a2070726f66696c6520616c726561647920657869737473005f82015250565b5f6109d4601f83610915565b91506109df826109a0565b602082019050919050565b5f6020820190508181035f830152610a01816109c8565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610a4c57607f821691505b602082108103610a5f57610a5e610a08565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610ac17fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610a86565b610acb8683610a86565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f610b0f610b0a610b0584610ae3565b610aec565b610ae3565b9050919050565b5f819050919050565b610b2883610af5565b610b3c610b3482610b16565b848454610a92565b825550505050565b5f90565b610b50610b44565b610b5b818484610b1f565b505050565b5b81811015610b7e57610b735f82610b48565b600181019050610b61565b5050565b601f821115610bc357610b9481610a65565b610b9d84610a77565b81016020851015610bac578190505b610bc0610bb885610a77565b830182610b60565b50505b505050565b5f82821c905092915050565b5f610be35f1984600802610bc8565b1980831691505092915050565b5f610bfb8383610bd4565b9150826002028217905092915050565b610c14826107d2565b67ffffffffffffffff811115610c2d57610c2c610629565b5b610c378254610a35565b610c42828285610b82565b5f60209050601f831160018114610c73575f8415610c61578287015190505b610c6b8582610bf0565b865550610cd2565b601f198416610c8186610a65565b5f5b82811015610ca857848901518255600182019150602085019450602081019050610c83565b86831015610cc55784890151610cc1601f891682610bd4565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220d77086820e55fe193202ac722bea97d44acb44e03b192e97ec125ca6c15a9ffe64736f6c63430008150033",
}

// ProfileABI is the input ABI used to generate the binding from.
// Deprecated: Use ProfileMetaData.ABI instead.
var ProfileABI = ProfileMetaData.ABI

// ProfileBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProfileMetaData.Bin instead.
var ProfileBin = ProfileMetaData.Bin

// DeployProfile deploys a new Ethereum contract, binding an instance of Profile to it.
func DeployProfile(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Profile, error) {
	parsed, err := ProfileMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProfileBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Profile{ProfileCaller: ProfileCaller{contract: contract}, ProfileTransactor: ProfileTransactor{contract: contract}, ProfileFilterer: ProfileFilterer{contract: contract}}, nil
}

// Profile is an auto generated Go binding around an Ethereum contract.
type Profile struct {
	ProfileCaller     // Read-only binding to the contract
	ProfileTransactor // Write-only binding to the contract
	ProfileFilterer   // Log filterer for contract events
}

// ProfileCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProfileCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProfileTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProfileFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProfileSession struct {
	Contract     *Profile          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProfileCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProfileCallerSession struct {
	Contract *ProfileCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ProfileTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProfileTransactorSession struct {
	Contract     *ProfileTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ProfileRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProfileRaw struct {
	Contract *Profile // Generic contract binding to access the raw methods on
}

// ProfileCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProfileCallerRaw struct {
	Contract *ProfileCaller // Generic read-only contract binding to access the raw methods on
}

// ProfileTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProfileTransactorRaw struct {
	Contract *ProfileTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProfile creates a new instance of Profile, bound to a specific deployed contract.
func NewProfile(address common.Address, backend bind.ContractBackend) (*Profile, error) {
	contract, err := bindProfile(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Profile{ProfileCaller: ProfileCaller{contract: contract}, ProfileTransactor: ProfileTransactor{contract: contract}, ProfileFilterer: ProfileFilterer{contract: contract}}, nil
}

// NewProfileCaller creates a new read-only instance of Profile, bound to a specific deployed contract.
func NewProfileCaller(address common.Address, caller bind.ContractCaller) (*ProfileCaller, error) {
	contract, err := bindProfile(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProfileCaller{contract: contract}, nil
}

// NewProfileTransactor creates a new write-only instance of Profile, bound to a specific deployed contract.
func NewProfileTransactor(address common.Address, transactor bind.ContractTransactor) (*ProfileTransactor, error) {
	contract, err := bindProfile(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProfileTransactor{contract: contract}, nil
}

// NewProfileFilterer creates a new log filterer instance of Profile, bound to a specific deployed contract.
func NewProfileFilterer(address common.Address, filterer bind.ContractFilterer) (*ProfileFilterer, error) {
	contract, err := bindProfile(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProfileFilterer{contract: contract}, nil
}

// bindProfile binds a generic wrapper to an already deployed contract.
func bindProfile(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProfileMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Profile *ProfileRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Profile.Contract.ProfileCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Profile *ProfileRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Profile.Contract.ProfileTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Profile *ProfileRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Profile.Contract.ProfileTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Profile *ProfileCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Profile.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Profile *ProfileTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Profile.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Profile *ProfileTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Profile.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Profile *ProfileCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Profile *ProfileSession) Owner() (common.Address, error) {
	return _Profile.Contract.Owner(&_Profile.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Profile *ProfileCallerSession) Owner() (common.Address, error) {
	return _Profile.Contract.Owner(&_Profile.CallOpts)
}

// Profiles is a free data retrieval call binding the contract method 0xbbe15627.
//
// Solidity: function profiles(address ) view returns(address profileAddress, string username, string profileQRCode)
func (_Profile *ProfileCaller) Profiles(opts *bind.CallOpts, arg0 common.Address) (struct {
	ProfileAddress common.Address
	Username       string
	ProfileQRCode  string
}, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "profiles", arg0)

	outstruct := new(struct {
		ProfileAddress common.Address
		Username       string
		ProfileQRCode  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProfileAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Username = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.ProfileQRCode = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// Profiles is a free data retrieval call binding the contract method 0xbbe15627.
//
// Solidity: function profiles(address ) view returns(address profileAddress, string username, string profileQRCode)
func (_Profile *ProfileSession) Profiles(arg0 common.Address) (struct {
	ProfileAddress common.Address
	Username       string
	ProfileQRCode  string
}, error) {
	return _Profile.Contract.Profiles(&_Profile.CallOpts, arg0)
}

// Profiles is a free data retrieval call binding the contract method 0xbbe15627.
//
// Solidity: function profiles(address ) view returns(address profileAddress, string username, string profileQRCode)
func (_Profile *ProfileCallerSession) Profiles(arg0 common.Address) (struct {
	ProfileAddress common.Address
	Username       string
	ProfileQRCode  string
}, error) {
	return _Profile.Contract.Profiles(&_Profile.CallOpts, arg0)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xba1feb8a.
//
// Solidity: function createProfile(string _username, string _qrCode) returns((address,string,string))
func (_Profile *ProfileTransactor) CreateProfile(opts *bind.TransactOpts, _username string, _qrCode string) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "createProfile", _username, _qrCode)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xba1feb8a.
//
// Solidity: function createProfile(string _username, string _qrCode) returns((address,string,string))
func (_Profile *ProfileSession) CreateProfile(_username string, _qrCode string) (*types.Transaction, error) {
	return _Profile.Contract.CreateProfile(&_Profile.TransactOpts, _username, _qrCode)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xba1feb8a.
//
// Solidity: function createProfile(string _username, string _qrCode) returns((address,string,string))
func (_Profile *ProfileTransactorSession) CreateProfile(_username string, _qrCode string) (*types.Transaction, error) {
	return _Profile.Contract.CreateProfile(&_Profile.TransactOpts, _username, _qrCode)
}
