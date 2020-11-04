// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dataContract

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

// AccessControlContractABI is the input ABI used to generate the binding from.
const AccessControlContractABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedAccounts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AccessControlContractFuncSigs maps the 4-byte function signature to its string representation.
var AccessControlContractFuncSigs = map[string]string{
	"e04610ed": "allowedAccounts(address)",
}

// AccessControlContractBin is the compiled bytecode used for deploying new contracts.
var AccessControlContractBin = "0x6080604052348015600f57600080fd5b5060ae8061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063e04610ed14602d575b600080fd5b605060048036036020811015604157600080fd5b50356001600160a01b03166064565b604080519115158252519081900360200190f35b60006020819052908152604090205460ff168156fea265627a7a723158205a4c0306d1fec8e6ab95d9dbad75a2db169df528a59a6221389c45ba20b8194a64736f6c63430005100032"

// DeployAccessControlContract deploys a new Ethereum contract, binding an instance of AccessControlContract to it.
func DeployAccessControlContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccessControlContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccessControlContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccessControlContract{AccessControlContractCaller: AccessControlContractCaller{contract: contract}, AccessControlContractTransactor: AccessControlContractTransactor{contract: contract}, AccessControlContractFilterer: AccessControlContractFilterer{contract: contract}}, nil
}

// AccessControlContract is an auto generated Go binding around an Ethereum contract.
type AccessControlContract struct {
	AccessControlContractCaller     // Read-only binding to the contract
	AccessControlContractTransactor // Write-only binding to the contract
	AccessControlContractFilterer   // Log filterer for contract events
}

// AccessControlContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlContractSession struct {
	Contract     *AccessControlContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AccessControlContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlContractCallerSession struct {
	Contract *AccessControlContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// AccessControlContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlContractTransactorSession struct {
	Contract     *AccessControlContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// AccessControlContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlContractRaw struct {
	Contract *AccessControlContract // Generic contract binding to access the raw methods on
}

// AccessControlContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlContractCallerRaw struct {
	Contract *AccessControlContractCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlContractTransactorRaw struct {
	Contract *AccessControlContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControlContract creates a new instance of AccessControlContract, bound to a specific deployed contract.
func NewAccessControlContract(address common.Address, backend bind.ContractBackend) (*AccessControlContract, error) {
	contract, err := bindAccessControlContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControlContract{AccessControlContractCaller: AccessControlContractCaller{contract: contract}, AccessControlContractTransactor: AccessControlContractTransactor{contract: contract}, AccessControlContractFilterer: AccessControlContractFilterer{contract: contract}}, nil
}

// NewAccessControlContractCaller creates a new read-only instance of AccessControlContract, bound to a specific deployed contract.
func NewAccessControlContractCaller(address common.Address, caller bind.ContractCaller) (*AccessControlContractCaller, error) {
	contract, err := bindAccessControlContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlContractCaller{contract: contract}, nil
}

// NewAccessControlContractTransactor creates a new write-only instance of AccessControlContract, bound to a specific deployed contract.
func NewAccessControlContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlContractTransactor, error) {
	contract, err := bindAccessControlContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlContractTransactor{contract: contract}, nil
}

// NewAccessControlContractFilterer creates a new log filterer instance of AccessControlContract, bound to a specific deployed contract.
func NewAccessControlContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlContractFilterer, error) {
	contract, err := bindAccessControlContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlContractFilterer{contract: contract}, nil
}

// bindAccessControlContract binds a generic wrapper to an already deployed contract.
func bindAccessControlContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlContract *AccessControlContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlContract.Contract.AccessControlContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlContract *AccessControlContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlContract.Contract.AccessControlContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlContract *AccessControlContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlContract.Contract.AccessControlContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlContract *AccessControlContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlContract *AccessControlContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlContract *AccessControlContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlContract.Contract.contract.Transact(opts, method, params...)
}

// AllowedAccounts is a free data retrieval call binding the contract method 0xe04610ed.
//
// Solidity: function allowedAccounts(address ) view returns(bool)
func (_AccessControlContract *AccessControlContractCaller) AllowedAccounts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AccessControlContract.contract.Call(opts, &out, "allowedAccounts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedAccounts is a free data retrieval call binding the contract method 0xe04610ed.
//
// Solidity: function allowedAccounts(address ) view returns(bool)
func (_AccessControlContract *AccessControlContractSession) AllowedAccounts(arg0 common.Address) (bool, error) {
	return _AccessControlContract.Contract.AllowedAccounts(&_AccessControlContract.CallOpts, arg0)
}

// AllowedAccounts is a free data retrieval call binding the contract method 0xe04610ed.
//
// Solidity: function allowedAccounts(address ) view returns(bool)
func (_AccessControlContract *AccessControlContractCallerSession) AllowedAccounts(arg0 common.Address) (bool, error) {
	return _AccessControlContract.Contract.AllowedAccounts(&_AccessControlContract.CallOpts, arg0)
}

// DataLedgerContractABI is the input ABI used to generate the binding from.
const DataLedgerContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"deleteInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"evtStoreInfo\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"deleteMeasurement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"retrieveInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"storeInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DataLedgerContractFuncSigs maps the 4-byte function signature to its string representation.
var DataLedgerContractFuncSigs = map[string]string{
	"77ad95ca": "deleteMeasurement(bytes32)",
	"c9776a6d": "retrieveInfo(bytes32)",
	"e30081a0": "setAddress(address)",
	"b7e2a1b8": "storeInfo(bytes32,string,string)",
}

// DataLedgerContractBin is the compiled bytecode used for deploying new contracts.
var DataLedgerContractBin = "0x6080604052600180546001600160a01b03191673647f089f75db1874e574419d20c34b078797c4c517905534801561003657600080fd5b506108cb806100466000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806377ad95ca14610051578063b7e2a1b814610070578063c9776a6d146101a4578063e30081a01461029f575b600080fd5b61006e6004803603602081101561006757600080fd5b50356102c5565b005b61006e6004803603606081101561008657600080fd5b813591908101906040810160208201356401000000008111156100a857600080fd5b8201836020820111156100ba57600080fd5b803590602001918460018302840111640100000000831117156100dc57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561012f57600080fd5b82018360208201111561014157600080fd5b8035906020019184600183028401116401000000008311171561016357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610375945050505050565b6101c1600480360360208110156101ba57600080fd5b50356104d9565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b838110156102025781810151838201526020016101ea565b50505050905090810190601f16801561022f5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561026257818101518382015260200161024a565b50505050905090810190601f16801561028f5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b61006e600480360360208110156102b557600080fd5b50356001600160a01b0316610615565b6001546001600160a01b0316331461030e5760405162461bcd60e51b815260040180806020018281038252603381526020018061080d6033913960400191505060405180910390fd5b6000818152600260205260408120906103278282610700565b610335600183016000610700565b5060020180546001600160a01b031916905560405181907f072007d551e16de6c1b8938fdd0559f70033d87037e5dffa28631256df69f9fe90600090a250565b61037e33610680565b15156001146103be5760405162461bcd60e51b815260040180806020018281038252602b81526020018061086c602b913960400191505060405180910390fd5b6103c6610747565b828152602080820183905233604080840191909152600086815260028352208251805184936103f9928492910190610771565b5060208281015180516104129260018501920190610771565b5060409182015160029190910180546001600160a01b0319166001600160a01b0390921691909117905580516020808252855181830152855187937f06c4f2b8beb126621a4c74187a1573e7f48017b16d171df264507e2131b20f429388939092839283019185019080838360005b83811015610499578181015183820152602001610481565b50505050905090810190601f1680156104c65780820380516001836020036101000a031916815260200191505b509250505060405180910390a250505050565b60008181526002602081815260409283902080548451600180831615610100026000190190921694909404601f81018490048402850184019095528484526060948594929391840192918491908301828280156105775780601f1061054c57610100808354040283529160200191610577565b820191906000526020600020905b81548152906001019060200180831161055a57829003601f168201915b5050845460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959750869450925084019050828280156106055780601f106105da57610100808354040283529160200191610605565b820191906000526020600020905b8154815290600101906020018083116105e857829003601f168201915b5050505050905091509150915091565b6001546001600160a01b0316331461065e5760405162461bcd60e51b815260040180806020018281038252602c815260200180610840602c913960400191505060405180910390fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b600080546040805163e04610ed60e01b81526001600160a01b0385811660048301529151919092169163e04610ed916024808301926020929190829003018186803b1580156106ce57600080fd5b505afa1580156106e2573d6000803e3d6000fd5b505050506040513d60208110156106f857600080fd5b505192915050565b50805460018160011615610100020316600290046000825580601f106107265750610744565b601f01602090049060005260206000209081019061074491906107ef565b50565b6040518060600160405280606081526020016060815260200160006001600160a01b031681525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106107b257805160ff19168380011785556107df565b828001600101855582156107df579182015b828111156107df5782518255916020019190600101906107c4565b506107eb9291506107ef565b5090565b61080991905b808211156107eb57600081556001016107f5565b9056fe596f7520646f206e6f74206861766520656e6f7567682070726976696c6567657320746f20646f207468697320616374696f6e596f7520646f206e6f7420686176652070726976696c6567657320746f20646f207468697320616374696f6e546865204944207468617420796f7520617265207573696e67206973206e6f742072656769737465726564a265627a7a723158202d81c1c218a44114b4ed4851bd355e6c04e4fbba0e091fe4c6d1b23c18e2f11564736f6c63430005100032"

// DeployDataLedgerContract deploys a new Ethereum contract, binding an instance of DataLedgerContract to it.
func DeployDataLedgerContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DataLedgerContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataLedgerContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DataLedgerContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DataLedgerContract{DataLedgerContractCaller: DataLedgerContractCaller{contract: contract}, DataLedgerContractTransactor: DataLedgerContractTransactor{contract: contract}, DataLedgerContractFilterer: DataLedgerContractFilterer{contract: contract}}, nil
}

// DataLedgerContract is an auto generated Go binding around an Ethereum contract.
type DataLedgerContract struct {
	DataLedgerContractCaller     // Read-only binding to the contract
	DataLedgerContractTransactor // Write-only binding to the contract
	DataLedgerContractFilterer   // Log filterer for contract events
}

// DataLedgerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataLedgerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataLedgerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataLedgerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataLedgerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataLedgerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataLedgerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataLedgerContractSession struct {
	Contract     *DataLedgerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DataLedgerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataLedgerContractCallerSession struct {
	Contract *DataLedgerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// DataLedgerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataLedgerContractTransactorSession struct {
	Contract     *DataLedgerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// DataLedgerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataLedgerContractRaw struct {
	Contract *DataLedgerContract // Generic contract binding to access the raw methods on
}

// DataLedgerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataLedgerContractCallerRaw struct {
	Contract *DataLedgerContractCaller // Generic read-only contract binding to access the raw methods on
}

// DataLedgerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataLedgerContractTransactorRaw struct {
	Contract *DataLedgerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataLedgerContract creates a new instance of DataLedgerContract, bound to a specific deployed contract.
func NewDataLedgerContract(address common.Address, backend bind.ContractBackend) (*DataLedgerContract, error) {
	contract, err := bindDataLedgerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContract{DataLedgerContractCaller: DataLedgerContractCaller{contract: contract}, DataLedgerContractTransactor: DataLedgerContractTransactor{contract: contract}, DataLedgerContractFilterer: DataLedgerContractFilterer{contract: contract}}, nil
}

// NewDataLedgerContractCaller creates a new read-only instance of DataLedgerContract, bound to a specific deployed contract.
func NewDataLedgerContractCaller(address common.Address, caller bind.ContractCaller) (*DataLedgerContractCaller, error) {
	contract, err := bindDataLedgerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContractCaller{contract: contract}, nil
}

// NewDataLedgerContractTransactor creates a new write-only instance of DataLedgerContract, bound to a specific deployed contract.
func NewDataLedgerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*DataLedgerContractTransactor, error) {
	contract, err := bindDataLedgerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContractTransactor{contract: contract}, nil
}

// NewDataLedgerContractFilterer creates a new log filterer instance of DataLedgerContract, bound to a specific deployed contract.
func NewDataLedgerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*DataLedgerContractFilterer, error) {
	contract, err := bindDataLedgerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContractFilterer{contract: contract}, nil
}

// bindDataLedgerContract binds a generic wrapper to an already deployed contract.
func bindDataLedgerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataLedgerContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataLedgerContract *DataLedgerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataLedgerContract.Contract.DataLedgerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataLedgerContract *DataLedgerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.DataLedgerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataLedgerContract *DataLedgerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.DataLedgerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataLedgerContract *DataLedgerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataLedgerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataLedgerContract *DataLedgerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataLedgerContract *DataLedgerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.contract.Transact(opts, method, params...)
}

// RetrieveInfo is a free data retrieval call binding the contract method 0xc9776a6d.
//
// Solidity: function retrieveInfo(bytes32 hash) view returns(string, string)
func (_DataLedgerContract *DataLedgerContractCaller) RetrieveInfo(opts *bind.CallOpts, hash [32]byte) (string, string, error) {
	var out []interface{}
	err := _DataLedgerContract.contract.Call(opts, &out, "retrieveInfo", hash)

	if err != nil {
		return *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// RetrieveInfo is a free data retrieval call binding the contract method 0xc9776a6d.
//
// Solidity: function retrieveInfo(bytes32 hash) view returns(string, string)
func (_DataLedgerContract *DataLedgerContractSession) RetrieveInfo(hash [32]byte) (string, string, error) {
	return _DataLedgerContract.Contract.RetrieveInfo(&_DataLedgerContract.CallOpts, hash)
}

// RetrieveInfo is a free data retrieval call binding the contract method 0xc9776a6d.
//
// Solidity: function retrieveInfo(bytes32 hash) view returns(string, string)
func (_DataLedgerContract *DataLedgerContractCallerSession) RetrieveInfo(hash [32]byte) (string, string, error) {
	return _DataLedgerContract.Contract.RetrieveInfo(&_DataLedgerContract.CallOpts, hash)
}

// DeleteMeasurement is a paid mutator transaction binding the contract method 0x77ad95ca.
//
// Solidity: function deleteMeasurement(bytes32 hash) returns()
func (_DataLedgerContract *DataLedgerContractTransactor) DeleteMeasurement(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _DataLedgerContract.contract.Transact(opts, "deleteMeasurement", hash)
}

// DeleteMeasurement is a paid mutator transaction binding the contract method 0x77ad95ca.
//
// Solidity: function deleteMeasurement(bytes32 hash) returns()
func (_DataLedgerContract *DataLedgerContractSession) DeleteMeasurement(hash [32]byte) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.DeleteMeasurement(&_DataLedgerContract.TransactOpts, hash)
}

// DeleteMeasurement is a paid mutator transaction binding the contract method 0x77ad95ca.
//
// Solidity: function deleteMeasurement(bytes32 hash) returns()
func (_DataLedgerContract *DataLedgerContractTransactorSession) DeleteMeasurement(hash [32]byte) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.DeleteMeasurement(&_DataLedgerContract.TransactOpts, hash)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _address) returns()
func (_DataLedgerContract *DataLedgerContractTransactor) SetAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _DataLedgerContract.contract.Transact(opts, "setAddress", _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _address) returns()
func (_DataLedgerContract *DataLedgerContractSession) SetAddress(_address common.Address) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.SetAddress(&_DataLedgerContract.TransactOpts, _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _address) returns()
func (_DataLedgerContract *DataLedgerContractTransactorSession) SetAddress(_address common.Address) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.SetAddress(&_DataLedgerContract.TransactOpts, _address)
}

// StoreInfo is a paid mutator transaction binding the contract method 0xb7e2a1b8.
//
// Solidity: function storeInfo(bytes32 hash, string uri, string description) returns()
func (_DataLedgerContract *DataLedgerContractTransactor) StoreInfo(opts *bind.TransactOpts, hash [32]byte, uri string, description string) (*types.Transaction, error) {
	return _DataLedgerContract.contract.Transact(opts, "storeInfo", hash, uri, description)
}

// StoreInfo is a paid mutator transaction binding the contract method 0xb7e2a1b8.
//
// Solidity: function storeInfo(bytes32 hash, string uri, string description) returns()
func (_DataLedgerContract *DataLedgerContractSession) StoreInfo(hash [32]byte, uri string, description string) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.StoreInfo(&_DataLedgerContract.TransactOpts, hash, uri, description)
}

// StoreInfo is a paid mutator transaction binding the contract method 0xb7e2a1b8.
//
// Solidity: function storeInfo(bytes32 hash, string uri, string description) returns()
func (_DataLedgerContract *DataLedgerContractTransactorSession) StoreInfo(hash [32]byte, uri string, description string) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.StoreInfo(&_DataLedgerContract.TransactOpts, hash, uri, description)
}

// DataLedgerContractDeleteInfoIterator is returned from FilterDeleteInfo and is used to iterate over the raw logs and unpacked data for DeleteInfo events raised by the DataLedgerContract contract.
type DataLedgerContractDeleteInfoIterator struct {
	Event *DataLedgerContractDeleteInfo // Event containing the contract specifics and raw log

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
func (it *DataLedgerContractDeleteInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLedgerContractDeleteInfo)
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
		it.Event = new(DataLedgerContractDeleteInfo)
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
func (it *DataLedgerContractDeleteInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLedgerContractDeleteInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLedgerContractDeleteInfo represents a DeleteInfo event raised by the DataLedgerContract contract.
type DataLedgerContractDeleteInfo struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDeleteInfo is a free log retrieval operation binding the contract event 0x072007d551e16de6c1b8938fdd0559f70033d87037e5dffa28631256df69f9fe.
//
// Solidity: event deleteInfo(bytes32 indexed _hash)
func (_DataLedgerContract *DataLedgerContractFilterer) FilterDeleteInfo(opts *bind.FilterOpts, _hash [][32]byte) (*DataLedgerContractDeleteInfoIterator, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _DataLedgerContract.contract.FilterLogs(opts, "deleteInfo", _hashRule)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContractDeleteInfoIterator{contract: _DataLedgerContract.contract, event: "deleteInfo", logs: logs, sub: sub}, nil
}

// WatchDeleteInfo is a free log subscription operation binding the contract event 0x072007d551e16de6c1b8938fdd0559f70033d87037e5dffa28631256df69f9fe.
//
// Solidity: event deleteInfo(bytes32 indexed _hash)
func (_DataLedgerContract *DataLedgerContractFilterer) WatchDeleteInfo(opts *bind.WatchOpts, sink chan<- *DataLedgerContractDeleteInfo, _hash [][32]byte) (event.Subscription, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _DataLedgerContract.contract.WatchLogs(opts, "deleteInfo", _hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLedgerContractDeleteInfo)
				if err := _DataLedgerContract.contract.UnpackLog(event, "deleteInfo", log); err != nil {
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

// ParseDeleteInfo is a log parse operation binding the contract event 0x072007d551e16de6c1b8938fdd0559f70033d87037e5dffa28631256df69f9fe.
//
// Solidity: event deleteInfo(bytes32 indexed _hash)
func (_DataLedgerContract *DataLedgerContractFilterer) ParseDeleteInfo(log types.Log) (*DataLedgerContractDeleteInfo, error) {
	event := new(DataLedgerContractDeleteInfo)
	if err := _DataLedgerContract.contract.UnpackLog(event, "deleteInfo", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DataLedgerContractEvtStoreInfoIterator is returned from FilterEvtStoreInfo and is used to iterate over the raw logs and unpacked data for EvtStoreInfo events raised by the DataLedgerContract contract.
type DataLedgerContractEvtStoreInfoIterator struct {
	Event *DataLedgerContractEvtStoreInfo // Event containing the contract specifics and raw log

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
func (it *DataLedgerContractEvtStoreInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLedgerContractEvtStoreInfo)
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
		it.Event = new(DataLedgerContractEvtStoreInfo)
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
func (it *DataLedgerContractEvtStoreInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLedgerContractEvtStoreInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLedgerContractEvtStoreInfo represents a EvtStoreInfo event raised by the DataLedgerContract contract.
type DataLedgerContractEvtStoreInfo struct {
	Hash [32]byte
	Uri  string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEvtStoreInfo is a free log retrieval operation binding the contract event 0x06c4f2b8beb126621a4c74187a1573e7f48017b16d171df264507e2131b20f42.
//
// Solidity: event evtStoreInfo(bytes32 indexed _hash, string _uri)
func (_DataLedgerContract *DataLedgerContractFilterer) FilterEvtStoreInfo(opts *bind.FilterOpts, _hash [][32]byte) (*DataLedgerContractEvtStoreInfoIterator, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _DataLedgerContract.contract.FilterLogs(opts, "evtStoreInfo", _hashRule)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContractEvtStoreInfoIterator{contract: _DataLedgerContract.contract, event: "evtStoreInfo", logs: logs, sub: sub}, nil
}

// WatchEvtStoreInfo is a free log subscription operation binding the contract event 0x06c4f2b8beb126621a4c74187a1573e7f48017b16d171df264507e2131b20f42.
//
// Solidity: event evtStoreInfo(bytes32 indexed _hash, string _uri)
func (_DataLedgerContract *DataLedgerContractFilterer) WatchEvtStoreInfo(opts *bind.WatchOpts, sink chan<- *DataLedgerContractEvtStoreInfo, _hash [][32]byte) (event.Subscription, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _DataLedgerContract.contract.WatchLogs(opts, "evtStoreInfo", _hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLedgerContractEvtStoreInfo)
				if err := _DataLedgerContract.contract.UnpackLog(event, "evtStoreInfo", log); err != nil {
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

// ParseEvtStoreInfo is a log parse operation binding the contract event 0x06c4f2b8beb126621a4c74187a1573e7f48017b16d171df264507e2131b20f42.
//
// Solidity: event evtStoreInfo(bytes32 indexed _hash, string _uri)
func (_DataLedgerContract *DataLedgerContractFilterer) ParseEvtStoreInfo(log types.Log) (*DataLedgerContractEvtStoreInfo, error) {
	event := new(DataLedgerContractEvtStoreInfo)
	if err := _DataLedgerContract.contract.UnpackLog(event, "evtStoreInfo", log); err != nil {
		return nil, err
	}
	return event, nil
}
