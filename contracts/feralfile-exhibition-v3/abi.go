// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package feralfilev3

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
)

// FeralfileExhibitionV3Artwork is an auto generated low-level Go binding around an user-defined struct.
type FeralfileExhibitionV3Artwork struct {
	Title       string
	ArtistName  string
	Fingerprint string
	EditionSize *big.Int
	AEAmount    *big.Int
	PPAmount    *big.Int
}

// FeralfileExhibitionV3MintArtworkParam is an auto generated low-level Go binding around an user-defined struct.
type FeralfileExhibitionV3MintArtworkParam struct {
	ArtworkID *big.Int
	Edition   *big.Int
	Artist    common.Address
	Owner     common.Address
	IpfsCID   string
}

// FeralfileExhibitionV3TransferArtworkParam is an auto generated low-level Go binding around an user-defined struct.
type FeralfileExhibitionV3TransferArtworkParam struct {
	From      common.Address
	To        common.Address
	TokenID   *big.Int
	Timestamp *big.Int
	R         [32]byte
	S         [32]byte
	V         uint8
}

// FeralfileExhibitionV3MetaData contains all meta data concerning the FeralfileExhibitionV3 contract.
var FeralfileExhibitionV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"secondarySaleRoyaltyBPS_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"royaltyPayoutAddress_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"contractURI_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenBaseURI_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isBurnable_\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isBridgeable_\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"editionID\",\"type\":\"uint256\"}],\"name\":\"BurnArtworkEdition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"artworkID\",\"type\":\"uint256\"}],\"name\":\"NewArtwork\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"artworkID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"editionID\",\"type\":\"uint256\"}],\"name\":\"NewArtworkEdition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_ROYALITY_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustee\",\"type\":\"address\"}],\"name\":\"addTrustee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"artworkEditions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"editionID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ipfsCID\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"artworks\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"artistName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fingerprint\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"editionSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"AEAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"PPAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"codeVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isBridgeable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isBurnable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustee\",\"type\":\"address\"}],\"name\":\"removeTrustee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyPayoutAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"secondarySaleRoyaltyBPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"trustees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"artistName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fingerprint\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"editionSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"AEAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"PPAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFeralfileExhibitionV3.Artwork[]\",\"name\":\"artworks_\",\"type\":\"tuple[]\"}],\"name\":\"createArtworks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalArtworks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getArtworkByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ipfsCID\",\"type\":\"string\"}],\"name\":\"updateArtworkEditionIPFSCid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"royaltyPayoutAddress_\",\"type\":\"address\"}],\"name\":\"setRoyaltyPayoutAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"artworkID\",\"type\":\"uint256\"}],\"name\":\"totalEditionOfArtwork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"artworkID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getArtworkEditionByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"}],\"name\":\"setTokenBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v_\",\"type\":\"uint8\"}],\"internalType\":\"structFeralfileExhibitionV3.TransferArtworkParam[]\",\"name\":\"transferParams_\",\"type\":\"tuple[]\"}],\"name\":\"authorizedTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"artworkID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"edition\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"artist\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipfsCID\",\"type\":\"string\"}],\"internalType\":\"structFeralfileExhibitionV3.MintArtworkParam[]\",\"name\":\"mintParams_\",\"type\":\"tuple[]\"}],\"name\":\"batchMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"editionIDs_\",\"type\":\"uint256[]\"}],\"name\":\"burnEditions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FeralfileExhibitionV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use FeralfileExhibitionV3MetaData.ABI instead.
var FeralfileExhibitionV3ABI = FeralfileExhibitionV3MetaData.ABI

// FeralfileExhibitionV3 is an auto generated Go binding around an Ethereum contract.
type FeralfileExhibitionV3 struct {
	FeralfileExhibitionV3Caller     // Read-only binding to the contract
	FeralfileExhibitionV3Transactor // Write-only binding to the contract
	FeralfileExhibitionV3Filterer   // Log filterer for contract events
}

// FeralfileExhibitionV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type FeralfileExhibitionV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeralfileExhibitionV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FeralfileExhibitionV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeralfileExhibitionV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeralfileExhibitionV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeralfileExhibitionV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeralfileExhibitionV3Session struct {
	Contract     *FeralfileExhibitionV3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FeralfileExhibitionV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeralfileExhibitionV3CallerSession struct {
	Contract *FeralfileExhibitionV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// FeralfileExhibitionV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeralfileExhibitionV3TransactorSession struct {
	Contract     *FeralfileExhibitionV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// FeralfileExhibitionV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type FeralfileExhibitionV3Raw struct {
	Contract *FeralfileExhibitionV3 // Generic contract binding to access the raw methods on
}

// FeralfileExhibitionV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeralfileExhibitionV3CallerRaw struct {
	Contract *FeralfileExhibitionV3Caller // Generic read-only contract binding to access the raw methods on
}

// FeralfileExhibitionV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeralfileExhibitionV3TransactorRaw struct {
	Contract *FeralfileExhibitionV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFeralfileExhibitionV3 creates a new instance of FeralfileExhibitionV3, bound to a specific deployed contract.
func NewFeralfileExhibitionV3(address common.Address, backend bind.ContractBackend) (*FeralfileExhibitionV3, error) {
	contract, err := bindFeralfileExhibitionV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV3{FeralfileExhibitionV3Caller: FeralfileExhibitionV3Caller{contract: contract}, FeralfileExhibitionV3Transactor: FeralfileExhibitionV3Transactor{contract: contract}, FeralfileExhibitionV3Filterer: FeralfileExhibitionV3Filterer{contract: contract}}, nil
}

// NewFeralfileExhibitionV3Caller creates a new read-only instance of FeralfileExhibitionV3, bound to a specific deployed contract.
func NewFeralfileExhibitionV3Caller(address common.Address, caller bind.ContractCaller) (*FeralfileExhibitionV3Caller, error) {
	contract, err := bindFeralfileExhibitionV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV3Caller{contract: contract}, nil
}

// NewFeralfileExhibitionV3Transactor creates a new write-only instance of FeralfileExhibitionV3, bound to a specific deployed contract.
func NewFeralfileExhibitionV3Transactor(address common.Address, transactor bind.ContractTransactor) (*FeralfileExhibitionV3Transactor, error) {
	contract, err := bindFeralfileExhibitionV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV3Transactor{contract: contract}, nil
}

// NewFeralfileExhibitionV3Filterer creates a new log filterer instance of FeralfileExhibitionV3, bound to a specific deployed contract.
func NewFeralfileExhibitionV3Filterer(address common.Address, filterer bind.ContractFilterer) (*FeralfileExhibitionV3Filterer, error) {
	contract, err := bindFeralfileExhibitionV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV3Filterer{contract: contract}, nil
}

// bindFeralfileExhibitionV3 binds a generic wrapper to an already deployed contract.
func bindFeralfileExhibitionV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeralfileExhibitionV3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeralfileExhibitionV3.Contract.FeralfileExhibitionV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.FeralfileExhibitionV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.FeralfileExhibitionV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeralfileExhibitionV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.contract.Transact(opts, method, params...)
}

// MAXROYALITYBPS is a free data retrieval call binding the contract method 0xec9cbb44.
//
// Solidity: function MAX_ROYALITY_BPS() view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) MAXROYALITYBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "MAX_ROYALITY_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXROYALITYBPS is a free data retrieval call binding the contract method 0xec9cbb44.
//
// Solidity: function MAX_ROYALITY_BPS() view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) MAXROYALITYBPS() (*big.Int, error) {
	return _FeralfileExhibitionV3.Contract.MAXROYALITYBPS(&_FeralfileExhibitionV3.CallOpts)
}

// MAXROYALITYBPS is a free data retrieval call binding the contract method 0xec9cbb44.
//
// Solidity: function MAX_ROYALITY_BPS() view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) MAXROYALITYBPS() (*big.Int, error) {
	return _FeralfileExhibitionV3.Contract.MAXROYALITYBPS(&_FeralfileExhibitionV3.CallOpts)
}

// ArtworkEditions is a free data retrieval call binding the contract method 0x62fe2131.
//
// Solidity: function artworkEditions(uint256 ) view returns(uint256 editionID, string ipfsCID)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) ArtworkEditions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	EditionID *big.Int
	IpfsCID   string
}, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "artworkEditions", arg0)

	outstruct := new(struct {
		EditionID *big.Int
		IpfsCID   string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EditionID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IpfsCID = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// ArtworkEditions is a free data retrieval call binding the contract method 0x62fe2131.
//
// Solidity: function artworkEditions(uint256 ) view returns(uint256 editionID, string ipfsCID)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) ArtworkEditions(arg0 *big.Int) (struct {
	EditionID *big.Int
	IpfsCID   string
}, error) {
	return _FeralfileExhibitionV3.Contract.ArtworkEditions(&_FeralfileExhibitionV3.CallOpts, arg0)
}

// ArtworkEditions is a free data retrieval call binding the contract method 0x62fe2131.
//
// Solidity: function artworkEditions(uint256 ) view returns(uint256 editionID, string ipfsCID)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) ArtworkEditions(arg0 *big.Int) (struct {
	EditionID *big.Int
	IpfsCID   string
}, error) {
	return _FeralfileExhibitionV3.Contract.ArtworkEditions(&_FeralfileExhibitionV3.CallOpts, arg0)
}

// Artworks is a free data retrieval call binding the contract method 0x4b602673.
//
// Solidity: function artworks(uint256 ) view returns(string title, string artistName, string fingerprint, uint256 editionSize, uint256 AEAmount, uint256 PPAmount)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) Artworks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Title       string
	ArtistName  string
	Fingerprint string
	EditionSize *big.Int
	AEAmount    *big.Int
	PPAmount    *big.Int
}, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "artworks", arg0)

	outstruct := new(struct {
		Title       string
		ArtistName  string
		Fingerprint string
		EditionSize *big.Int
		AEAmount    *big.Int
		PPAmount    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Title = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.ArtistName = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Fingerprint = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.EditionSize = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AEAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.PPAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Artworks is a free data retrieval call binding the contract method 0x4b602673.
//
// Solidity: function artworks(uint256 ) view returns(string title, string artistName, string fingerprint, uint256 editionSize, uint256 AEAmount, uint256 PPAmount)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) Artworks(arg0 *big.Int) (struct {
	Title       string
	ArtistName  string
	Fingerprint string
	EditionSize *big.Int
	AEAmount    *big.Int
	PPAmount    *big.Int
}, error) {
	return _FeralfileExhibitionV3.Contract.Artworks(&_FeralfileExhibitionV3.CallOpts, arg0)
}

// Artworks is a free data retrieval call binding the contract method 0x4b602673.
//
// Solidity: function artworks(uint256 ) view returns(string title, string artistName, string fingerprint, uint256 editionSize, uint256 AEAmount, uint256 PPAmount)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) Artworks(arg0 *big.Int) (struct {
	Title       string
	ArtistName  string
	Fingerprint string
	EditionSize *big.Int
	AEAmount    *big.Int
	PPAmount    *big.Int
}, error) {
	return _FeralfileExhibitionV3.Contract.Artworks(&_FeralfileExhibitionV3.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _FeralfileExhibitionV3.Contract.BalanceOf(&_FeralfileExhibitionV3.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _FeralfileExhibitionV3.Contract.BalanceOf(&_FeralfileExhibitionV3.CallOpts, owner)
}

// CodeVersion is a free data retrieval call binding the contract method 0x63e60230.
//
// Solidity: function codeVersion() view returns(string)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) CodeVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "codeVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CodeVersion is a free data retrieval call binding the contract method 0x63e60230.
//
// Solidity: function codeVersion() view returns(string)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) CodeVersion() (string, error) {
	return _FeralfileExhibitionV3.Contract.CodeVersion(&_FeralfileExhibitionV3.CallOpts)
}

// CodeVersion is a free data retrieval call binding the contract method 0x63e60230.
//
// Solidity: function codeVersion() view returns(string)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) CodeVersion() (string, error) {
	return _FeralfileExhibitionV3.Contract.CodeVersion(&_FeralfileExhibitionV3.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) ContractURI() (string, error) {
	return _FeralfileExhibitionV3.Contract.ContractURI(&_FeralfileExhibitionV3.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) ContractURI() (string, error) {
	return _FeralfileExhibitionV3.Contract.ContractURI(&_FeralfileExhibitionV3.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _FeralfileExhibitionV3.Contract.GetApproved(&_FeralfileExhibitionV3.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _FeralfileExhibitionV3.Contract.GetApproved(&_FeralfileExhibitionV3.CallOpts, tokenId)
}

// GetArtworkByIndex is a free data retrieval call binding the contract method 0xb4883703.
//
// Solidity: function getArtworkByIndex(uint256 index) view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) GetArtworkByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "getArtworkByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetArtworkByIndex is a free data retrieval call binding the contract method 0xb4883703.
//
// Solidity: function getArtworkByIndex(uint256 index) view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) GetArtworkByIndex(index *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV3.Contract.GetArtworkByIndex(&_FeralfileExhibitionV3.CallOpts, index)
}

// GetArtworkByIndex is a free data retrieval call binding the contract method 0xb4883703.
//
// Solidity: function getArtworkByIndex(uint256 index) view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) GetArtworkByIndex(index *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV3.Contract.GetArtworkByIndex(&_FeralfileExhibitionV3.CallOpts, index)
}

// GetArtworkEditionByIndex is a free data retrieval call binding the contract method 0x641b18e9.
//
// Solidity: function getArtworkEditionByIndex(uint256 artworkID, uint256 index) view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) GetArtworkEditionByIndex(opts *bind.CallOpts, artworkID *big.Int, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "getArtworkEditionByIndex", artworkID, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetArtworkEditionByIndex is a free data retrieval call binding the contract method 0x641b18e9.
//
// Solidity: function getArtworkEditionByIndex(uint256 artworkID, uint256 index) view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) GetArtworkEditionByIndex(artworkID *big.Int, index *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV3.Contract.GetArtworkEditionByIndex(&_FeralfileExhibitionV3.CallOpts, artworkID, index)
}

// GetArtworkEditionByIndex is a free data retrieval call binding the contract method 0x641b18e9.
//
// Solidity: function getArtworkEditionByIndex(uint256 artworkID, uint256 index) view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) GetArtworkEditionByIndex(artworkID *big.Int, index *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV3.Contract.GetArtworkEditionByIndex(&_FeralfileExhibitionV3.CallOpts, artworkID, index)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _FeralfileExhibitionV3.Contract.IsApprovedForAll(&_FeralfileExhibitionV3.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _FeralfileExhibitionV3.Contract.IsApprovedForAll(&_FeralfileExhibitionV3.CallOpts, owner, operator)
}

// IsBridgeable is a free data retrieval call binding the contract method 0x7ca5ea89.
//
// Solidity: function isBridgeable() view returns(bool)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) IsBridgeable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "isBridgeable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBridgeable is a free data retrieval call binding the contract method 0x7ca5ea89.
//
// Solidity: function isBridgeable() view returns(bool)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) IsBridgeable() (bool, error) {
	return _FeralfileExhibitionV3.Contract.IsBridgeable(&_FeralfileExhibitionV3.CallOpts)
}

// IsBridgeable is a free data retrieval call binding the contract method 0x7ca5ea89.
//
// Solidity: function isBridgeable() view returns(bool)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) IsBridgeable() (bool, error) {
	return _FeralfileExhibitionV3.Contract.IsBridgeable(&_FeralfileExhibitionV3.CallOpts)
}

// IsBurnable is a free data retrieval call binding the contract method 0x883356d9.
//
// Solidity: function isBurnable() view returns(bool)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) IsBurnable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "isBurnable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBurnable is a free data retrieval call binding the contract method 0x883356d9.
//
// Solidity: function isBurnable() view returns(bool)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) IsBurnable() (bool, error) {
	return _FeralfileExhibitionV3.Contract.IsBurnable(&_FeralfileExhibitionV3.CallOpts)
}

// IsBurnable is a free data retrieval call binding the contract method 0x883356d9.
//
// Solidity: function isBurnable() view returns(bool)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) IsBurnable() (bool, error) {
	return _FeralfileExhibitionV3.Contract.IsBurnable(&_FeralfileExhibitionV3.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) Name() (string, error) {
	return _FeralfileExhibitionV3.Contract.Name(&_FeralfileExhibitionV3.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) Name() (string, error) {
	return _FeralfileExhibitionV3.Contract.Name(&_FeralfileExhibitionV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) Owner() (common.Address, error) {
	return _FeralfileExhibitionV3.Contract.Owner(&_FeralfileExhibitionV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) Owner() (common.Address, error) {
	return _FeralfileExhibitionV3.Contract.Owner(&_FeralfileExhibitionV3.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _FeralfileExhibitionV3.Contract.OwnerOf(&_FeralfileExhibitionV3.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _FeralfileExhibitionV3.Contract.OwnerOf(&_FeralfileExhibitionV3.CallOpts, tokenId)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _FeralfileExhibitionV3.Contract.RoyaltyInfo(&_FeralfileExhibitionV3.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _FeralfileExhibitionV3.Contract.RoyaltyInfo(&_FeralfileExhibitionV3.CallOpts, tokenId, salePrice)
}

// RoyaltyPayoutAddress is a free data retrieval call binding the contract method 0x3f6805ba.
//
// Solidity: function royaltyPayoutAddress() view returns(address)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) RoyaltyPayoutAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "royaltyPayoutAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoyaltyPayoutAddress is a free data retrieval call binding the contract method 0x3f6805ba.
//
// Solidity: function royaltyPayoutAddress() view returns(address)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) RoyaltyPayoutAddress() (common.Address, error) {
	return _FeralfileExhibitionV3.Contract.RoyaltyPayoutAddress(&_FeralfileExhibitionV3.CallOpts)
}

// RoyaltyPayoutAddress is a free data retrieval call binding the contract method 0x3f6805ba.
//
// Solidity: function royaltyPayoutAddress() view returns(address)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) RoyaltyPayoutAddress() (common.Address, error) {
	return _FeralfileExhibitionV3.Contract.RoyaltyPayoutAddress(&_FeralfileExhibitionV3.CallOpts)
}

// SecondarySaleRoyaltyBPS is a free data retrieval call binding the contract method 0xea211d7c.
//
// Solidity: function secondarySaleRoyaltyBPS() view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) SecondarySaleRoyaltyBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "secondarySaleRoyaltyBPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecondarySaleRoyaltyBPS is a free data retrieval call binding the contract method 0xea211d7c.
//
// Solidity: function secondarySaleRoyaltyBPS() view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) SecondarySaleRoyaltyBPS() (*big.Int, error) {
	return _FeralfileExhibitionV3.Contract.SecondarySaleRoyaltyBPS(&_FeralfileExhibitionV3.CallOpts)
}

// SecondarySaleRoyaltyBPS is a free data retrieval call binding the contract method 0xea211d7c.
//
// Solidity: function secondarySaleRoyaltyBPS() view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) SecondarySaleRoyaltyBPS() (*big.Int, error) {
	return _FeralfileExhibitionV3.Contract.SecondarySaleRoyaltyBPS(&_FeralfileExhibitionV3.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FeralfileExhibitionV3.Contract.SupportsInterface(&_FeralfileExhibitionV3.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FeralfileExhibitionV3.Contract.SupportsInterface(&_FeralfileExhibitionV3.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) Symbol() (string, error) {
	return _FeralfileExhibitionV3.Contract.Symbol(&_FeralfileExhibitionV3.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) Symbol() (string, error) {
	return _FeralfileExhibitionV3.Contract.Symbol(&_FeralfileExhibitionV3.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV3.Contract.TokenByIndex(&_FeralfileExhibitionV3.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV3.Contract.TokenByIndex(&_FeralfileExhibitionV3.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV3.Contract.TokenOfOwnerByIndex(&_FeralfileExhibitionV3.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV3.Contract.TokenOfOwnerByIndex(&_FeralfileExhibitionV3.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) TokenURI(tokenId *big.Int) (string, error) {
	return _FeralfileExhibitionV3.Contract.TokenURI(&_FeralfileExhibitionV3.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _FeralfileExhibitionV3.Contract.TokenURI(&_FeralfileExhibitionV3.CallOpts, tokenId)
}

// TotalArtworks is a free data retrieval call binding the contract method 0xe4a233e1.
//
// Solidity: function totalArtworks() view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) TotalArtworks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "totalArtworks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalArtworks is a free data retrieval call binding the contract method 0xe4a233e1.
//
// Solidity: function totalArtworks() view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) TotalArtworks() (*big.Int, error) {
	return _FeralfileExhibitionV3.Contract.TotalArtworks(&_FeralfileExhibitionV3.CallOpts)
}

// TotalArtworks is a free data retrieval call binding the contract method 0xe4a233e1.
//
// Solidity: function totalArtworks() view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) TotalArtworks() (*big.Int, error) {
	return _FeralfileExhibitionV3.Contract.TotalArtworks(&_FeralfileExhibitionV3.CallOpts)
}

// TotalEditionOfArtwork is a free data retrieval call binding the contract method 0xfe2a3bf3.
//
// Solidity: function totalEditionOfArtwork(uint256 artworkID) view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) TotalEditionOfArtwork(opts *bind.CallOpts, artworkID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "totalEditionOfArtwork", artworkID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalEditionOfArtwork is a free data retrieval call binding the contract method 0xfe2a3bf3.
//
// Solidity: function totalEditionOfArtwork(uint256 artworkID) view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) TotalEditionOfArtwork(artworkID *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV3.Contract.TotalEditionOfArtwork(&_FeralfileExhibitionV3.CallOpts, artworkID)
}

// TotalEditionOfArtwork is a free data retrieval call binding the contract method 0xfe2a3bf3.
//
// Solidity: function totalEditionOfArtwork(uint256 artworkID) view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) TotalEditionOfArtwork(artworkID *big.Int) (*big.Int, error) {
	return _FeralfileExhibitionV3.Contract.TotalEditionOfArtwork(&_FeralfileExhibitionV3.CallOpts, artworkID)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) TotalSupply() (*big.Int, error) {
	return _FeralfileExhibitionV3.Contract.TotalSupply(&_FeralfileExhibitionV3.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) TotalSupply() (*big.Int, error) {
	return _FeralfileExhibitionV3.Contract.TotalSupply(&_FeralfileExhibitionV3.CallOpts)
}

// Trustees is a free data retrieval call binding the contract method 0xeee608a4.
//
// Solidity: function trustees(address ) view returns(bool)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Caller) Trustees(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FeralfileExhibitionV3.contract.Call(opts, &out, "trustees", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Trustees is a free data retrieval call binding the contract method 0xeee608a4.
//
// Solidity: function trustees(address ) view returns(bool)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) Trustees(arg0 common.Address) (bool, error) {
	return _FeralfileExhibitionV3.Contract.Trustees(&_FeralfileExhibitionV3.CallOpts, arg0)
}

// Trustees is a free data retrieval call binding the contract method 0xeee608a4.
//
// Solidity: function trustees(address ) view returns(bool)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3CallerSession) Trustees(arg0 common.Address) (bool, error) {
	return _FeralfileExhibitionV3.Contract.Trustees(&_FeralfileExhibitionV3.CallOpts, arg0)
}

// AddTrustee is a paid mutator transaction binding the contract method 0xdc78ac1c.
//
// Solidity: function addTrustee(address _trustee) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Transactor) AddTrustee(opts *bind.TransactOpts, _trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.contract.Transact(opts, "addTrustee", _trustee)
}

// AddTrustee is a paid mutator transaction binding the contract method 0xdc78ac1c.
//
// Solidity: function addTrustee(address _trustee) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) AddTrustee(_trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.AddTrustee(&_FeralfileExhibitionV3.TransactOpts, _trustee)
}

// AddTrustee is a paid mutator transaction binding the contract method 0xdc78ac1c.
//
// Solidity: function addTrustee(address _trustee) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3TransactorSession) AddTrustee(_trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.AddTrustee(&_FeralfileExhibitionV3.TransactOpts, _trustee)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.Approve(&_FeralfileExhibitionV3.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.Approve(&_FeralfileExhibitionV3.TransactOpts, to, tokenId)
}

// AuthorizedTransfer is a paid mutator transaction binding the contract method 0x9fbf39cd.
//
// Solidity: function authorizedTransfer((address,address,uint256,uint256,bytes32,bytes32,uint8)[] transferParams_) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Transactor) AuthorizedTransfer(opts *bind.TransactOpts, transferParams_ []FeralfileExhibitionV3TransferArtworkParam) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.contract.Transact(opts, "authorizedTransfer", transferParams_)
}

// AuthorizedTransfer is a paid mutator transaction binding the contract method 0x9fbf39cd.
//
// Solidity: function authorizedTransfer((address,address,uint256,uint256,bytes32,bytes32,uint8)[] transferParams_) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) AuthorizedTransfer(transferParams_ []FeralfileExhibitionV3TransferArtworkParam) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.AuthorizedTransfer(&_FeralfileExhibitionV3.TransactOpts, transferParams_)
}

// AuthorizedTransfer is a paid mutator transaction binding the contract method 0x9fbf39cd.
//
// Solidity: function authorizedTransfer((address,address,uint256,uint256,bytes32,bytes32,uint8)[] transferParams_) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3TransactorSession) AuthorizedTransfer(transferParams_ []FeralfileExhibitionV3TransferArtworkParam) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.AuthorizedTransfer(&_FeralfileExhibitionV3.TransactOpts, transferParams_)
}

// BatchMint is a paid mutator transaction binding the contract method 0x12d907b9.
//
// Solidity: function batchMint((uint256,uint256,address,address,string)[] mintParams_) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Transactor) BatchMint(opts *bind.TransactOpts, mintParams_ []FeralfileExhibitionV3MintArtworkParam) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.contract.Transact(opts, "batchMint", mintParams_)
}

// BatchMint is a paid mutator transaction binding the contract method 0x12d907b9.
//
// Solidity: function batchMint((uint256,uint256,address,address,string)[] mintParams_) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) BatchMint(mintParams_ []FeralfileExhibitionV3MintArtworkParam) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.BatchMint(&_FeralfileExhibitionV3.TransactOpts, mintParams_)
}

// BatchMint is a paid mutator transaction binding the contract method 0x12d907b9.
//
// Solidity: function batchMint((uint256,uint256,address,address,string)[] mintParams_) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3TransactorSession) BatchMint(mintParams_ []FeralfileExhibitionV3MintArtworkParam) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.BatchMint(&_FeralfileExhibitionV3.TransactOpts, mintParams_)
}

// BurnEditions is a paid mutator transaction binding the contract method 0xfc05ea68.
//
// Solidity: function burnEditions(uint256[] editionIDs_) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Transactor) BurnEditions(opts *bind.TransactOpts, editionIDs_ []*big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.contract.Transact(opts, "burnEditions", editionIDs_)
}

// BurnEditions is a paid mutator transaction binding the contract method 0xfc05ea68.
//
// Solidity: function burnEditions(uint256[] editionIDs_) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) BurnEditions(editionIDs_ []*big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.BurnEditions(&_FeralfileExhibitionV3.TransactOpts, editionIDs_)
}

// BurnEditions is a paid mutator transaction binding the contract method 0xfc05ea68.
//
// Solidity: function burnEditions(uint256[] editionIDs_) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3TransactorSession) BurnEditions(editionIDs_ []*big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.BurnEditions(&_FeralfileExhibitionV3.TransactOpts, editionIDs_)
}

// CreateArtworks is a paid mutator transaction binding the contract method 0x43deaf76.
//
// Solidity: function createArtworks((string,string,string,uint256,uint256,uint256)[] artworks_) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Transactor) CreateArtworks(opts *bind.TransactOpts, artworks_ []FeralfileExhibitionV3Artwork) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.contract.Transact(opts, "createArtworks", artworks_)
}

// CreateArtworks is a paid mutator transaction binding the contract method 0x43deaf76.
//
// Solidity: function createArtworks((string,string,string,uint256,uint256,uint256)[] artworks_) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) CreateArtworks(artworks_ []FeralfileExhibitionV3Artwork) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.CreateArtworks(&_FeralfileExhibitionV3.TransactOpts, artworks_)
}

// CreateArtworks is a paid mutator transaction binding the contract method 0x43deaf76.
//
// Solidity: function createArtworks((string,string,string,uint256,uint256,uint256)[] artworks_) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3TransactorSession) CreateArtworks(artworks_ []FeralfileExhibitionV3Artwork) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.CreateArtworks(&_FeralfileExhibitionV3.TransactOpts, artworks_)
}

// RemoveTrustee is a paid mutator transaction binding the contract method 0x03120506.
//
// Solidity: function removeTrustee(address _trustee) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Transactor) RemoveTrustee(opts *bind.TransactOpts, _trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.contract.Transact(opts, "removeTrustee", _trustee)
}

// RemoveTrustee is a paid mutator transaction binding the contract method 0x03120506.
//
// Solidity: function removeTrustee(address _trustee) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) RemoveTrustee(_trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.RemoveTrustee(&_FeralfileExhibitionV3.TransactOpts, _trustee)
}

// RemoveTrustee is a paid mutator transaction binding the contract method 0x03120506.
//
// Solidity: function removeTrustee(address _trustee) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3TransactorSession) RemoveTrustee(_trustee common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.RemoveTrustee(&_FeralfileExhibitionV3.TransactOpts, _trustee)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) RenounceOwnership() (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.RenounceOwnership(&_FeralfileExhibitionV3.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.RenounceOwnership(&_FeralfileExhibitionV3.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.SafeTransferFrom(&_FeralfileExhibitionV3.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.SafeTransferFrom(&_FeralfileExhibitionV3.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.SafeTransferFrom0(&_FeralfileExhibitionV3.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.SafeTransferFrom0(&_FeralfileExhibitionV3.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.SetApprovalForAll(&_FeralfileExhibitionV3.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.SetApprovalForAll(&_FeralfileExhibitionV3.TransactOpts, operator, approved)
}

// SetRoyaltyPayoutAddress is a paid mutator transaction binding the contract method 0x45aeefde.
//
// Solidity: function setRoyaltyPayoutAddress(address royaltyPayoutAddress_) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Transactor) SetRoyaltyPayoutAddress(opts *bind.TransactOpts, royaltyPayoutAddress_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.contract.Transact(opts, "setRoyaltyPayoutAddress", royaltyPayoutAddress_)
}

// SetRoyaltyPayoutAddress is a paid mutator transaction binding the contract method 0x45aeefde.
//
// Solidity: function setRoyaltyPayoutAddress(address royaltyPayoutAddress_) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) SetRoyaltyPayoutAddress(royaltyPayoutAddress_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.SetRoyaltyPayoutAddress(&_FeralfileExhibitionV3.TransactOpts, royaltyPayoutAddress_)
}

// SetRoyaltyPayoutAddress is a paid mutator transaction binding the contract method 0x45aeefde.
//
// Solidity: function setRoyaltyPayoutAddress(address royaltyPayoutAddress_) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3TransactorSession) SetRoyaltyPayoutAddress(royaltyPayoutAddress_ common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.SetRoyaltyPayoutAddress(&_FeralfileExhibitionV3.TransactOpts, royaltyPayoutAddress_)
}

// SetTokenBaseURI is a paid mutator transaction binding the contract method 0x8ef79e91.
//
// Solidity: function setTokenBaseURI(string baseURI_) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Transactor) SetTokenBaseURI(opts *bind.TransactOpts, baseURI_ string) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.contract.Transact(opts, "setTokenBaseURI", baseURI_)
}

// SetTokenBaseURI is a paid mutator transaction binding the contract method 0x8ef79e91.
//
// Solidity: function setTokenBaseURI(string baseURI_) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) SetTokenBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.SetTokenBaseURI(&_FeralfileExhibitionV3.TransactOpts, baseURI_)
}

// SetTokenBaseURI is a paid mutator transaction binding the contract method 0x8ef79e91.
//
// Solidity: function setTokenBaseURI(string baseURI_) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3TransactorSession) SetTokenBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.SetTokenBaseURI(&_FeralfileExhibitionV3.TransactOpts, baseURI_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.TransferFrom(&_FeralfileExhibitionV3.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.TransferFrom(&_FeralfileExhibitionV3.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.TransferOwnership(&_FeralfileExhibitionV3.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.TransferOwnership(&_FeralfileExhibitionV3.TransactOpts, newOwner)
}

// UpdateArtworkEditionIPFSCid is a paid mutator transaction binding the contract method 0x0cfcb5f1.
//
// Solidity: function updateArtworkEditionIPFSCid(uint256 tokenId, string ipfsCID) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Transactor) UpdateArtworkEditionIPFSCid(opts *bind.TransactOpts, tokenId *big.Int, ipfsCID string) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.contract.Transact(opts, "updateArtworkEditionIPFSCid", tokenId, ipfsCID)
}

// UpdateArtworkEditionIPFSCid is a paid mutator transaction binding the contract method 0x0cfcb5f1.
//
// Solidity: function updateArtworkEditionIPFSCid(uint256 tokenId, string ipfsCID) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Session) UpdateArtworkEditionIPFSCid(tokenId *big.Int, ipfsCID string) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.UpdateArtworkEditionIPFSCid(&_FeralfileExhibitionV3.TransactOpts, tokenId, ipfsCID)
}

// UpdateArtworkEditionIPFSCid is a paid mutator transaction binding the contract method 0x0cfcb5f1.
//
// Solidity: function updateArtworkEditionIPFSCid(uint256 tokenId, string ipfsCID) returns()
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3TransactorSession) UpdateArtworkEditionIPFSCid(tokenId *big.Int, ipfsCID string) (*types.Transaction, error) {
	return _FeralfileExhibitionV3.Contract.UpdateArtworkEditionIPFSCid(&_FeralfileExhibitionV3.TransactOpts, tokenId, ipfsCID)
}

// FeralfileExhibitionV3ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FeralfileExhibitionV3 contract.
type FeralfileExhibitionV3ApprovalIterator struct {
	Event *FeralfileExhibitionV3Approval // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV3ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV3Approval)
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
		it.Event = new(FeralfileExhibitionV3Approval)
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
func (it *FeralfileExhibitionV3ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV3ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV3Approval represents a Approval event raised by the FeralfileExhibitionV3 contract.
type FeralfileExhibitionV3Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*FeralfileExhibitionV3ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV3.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV3ApprovalIterator{contract: _FeralfileExhibitionV3.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV3Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV3.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV3Approval)
				if err := _FeralfileExhibitionV3.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Filterer) ParseApproval(log types.Log) (*FeralfileExhibitionV3Approval, error) {
	event := new(FeralfileExhibitionV3Approval)
	if err := _FeralfileExhibitionV3.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV3ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the FeralfileExhibitionV3 contract.
type FeralfileExhibitionV3ApprovalForAllIterator struct {
	Event *FeralfileExhibitionV3ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV3ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV3ApprovalForAll)
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
		it.Event = new(FeralfileExhibitionV3ApprovalForAll)
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
func (it *FeralfileExhibitionV3ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV3ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV3ApprovalForAll represents a ApprovalForAll event raised by the FeralfileExhibitionV3 contract.
type FeralfileExhibitionV3ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*FeralfileExhibitionV3ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _FeralfileExhibitionV3.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV3ApprovalForAllIterator{contract: _FeralfileExhibitionV3.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV3ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _FeralfileExhibitionV3.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV3ApprovalForAll)
				if err := _FeralfileExhibitionV3.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Filterer) ParseApprovalForAll(log types.Log) (*FeralfileExhibitionV3ApprovalForAll, error) {
	event := new(FeralfileExhibitionV3ApprovalForAll)
	if err := _FeralfileExhibitionV3.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV3BurnArtworkEditionIterator is returned from FilterBurnArtworkEdition and is used to iterate over the raw logs and unpacked data for BurnArtworkEdition events raised by the FeralfileExhibitionV3 contract.
type FeralfileExhibitionV3BurnArtworkEditionIterator struct {
	Event *FeralfileExhibitionV3BurnArtworkEdition // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV3BurnArtworkEditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV3BurnArtworkEdition)
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
		it.Event = new(FeralfileExhibitionV3BurnArtworkEdition)
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
func (it *FeralfileExhibitionV3BurnArtworkEditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV3BurnArtworkEditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV3BurnArtworkEdition represents a BurnArtworkEdition event raised by the FeralfileExhibitionV3 contract.
type FeralfileExhibitionV3BurnArtworkEdition struct {
	EditionID *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBurnArtworkEdition is a free log retrieval operation binding the contract event 0xa5a44c7ed36966786612323ee2cb0cb453d4a9282b90c6befe72cde41d83f488.
//
// Solidity: event BurnArtworkEdition(uint256 indexed editionID)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Filterer) FilterBurnArtworkEdition(opts *bind.FilterOpts, editionID []*big.Int) (*FeralfileExhibitionV3BurnArtworkEditionIterator, error) {

	var editionIDRule []interface{}
	for _, editionIDItem := range editionID {
		editionIDRule = append(editionIDRule, editionIDItem)
	}

	logs, sub, err := _FeralfileExhibitionV3.contract.FilterLogs(opts, "BurnArtworkEdition", editionIDRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV3BurnArtworkEditionIterator{contract: _FeralfileExhibitionV3.contract, event: "BurnArtworkEdition", logs: logs, sub: sub}, nil
}

// WatchBurnArtworkEdition is a free log subscription operation binding the contract event 0xa5a44c7ed36966786612323ee2cb0cb453d4a9282b90c6befe72cde41d83f488.
//
// Solidity: event BurnArtworkEdition(uint256 indexed editionID)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Filterer) WatchBurnArtworkEdition(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV3BurnArtworkEdition, editionID []*big.Int) (event.Subscription, error) {

	var editionIDRule []interface{}
	for _, editionIDItem := range editionID {
		editionIDRule = append(editionIDRule, editionIDItem)
	}

	logs, sub, err := _FeralfileExhibitionV3.contract.WatchLogs(opts, "BurnArtworkEdition", editionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV3BurnArtworkEdition)
				if err := _FeralfileExhibitionV3.contract.UnpackLog(event, "BurnArtworkEdition", log); err != nil {
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

// ParseBurnArtworkEdition is a log parse operation binding the contract event 0xa5a44c7ed36966786612323ee2cb0cb453d4a9282b90c6befe72cde41d83f488.
//
// Solidity: event BurnArtworkEdition(uint256 indexed editionID)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Filterer) ParseBurnArtworkEdition(log types.Log) (*FeralfileExhibitionV3BurnArtworkEdition, error) {
	event := new(FeralfileExhibitionV3BurnArtworkEdition)
	if err := _FeralfileExhibitionV3.contract.UnpackLog(event, "BurnArtworkEdition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV3NewArtworkIterator is returned from FilterNewArtwork and is used to iterate over the raw logs and unpacked data for NewArtwork events raised by the FeralfileExhibitionV3 contract.
type FeralfileExhibitionV3NewArtworkIterator struct {
	Event *FeralfileExhibitionV3NewArtwork // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV3NewArtworkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV3NewArtwork)
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
		it.Event = new(FeralfileExhibitionV3NewArtwork)
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
func (it *FeralfileExhibitionV3NewArtworkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV3NewArtworkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV3NewArtwork represents a NewArtwork event raised by the FeralfileExhibitionV3 contract.
type FeralfileExhibitionV3NewArtwork struct {
	ArtworkID *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewArtwork is a free log retrieval operation binding the contract event 0x22350b25f1b72bb3621199a79abefeb4fcd77bb1e65638cd09350666e4db0891.
//
// Solidity: event NewArtwork(uint256 indexed artworkID)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Filterer) FilterNewArtwork(opts *bind.FilterOpts, artworkID []*big.Int) (*FeralfileExhibitionV3NewArtworkIterator, error) {

	var artworkIDRule []interface{}
	for _, artworkIDItem := range artworkID {
		artworkIDRule = append(artworkIDRule, artworkIDItem)
	}

	logs, sub, err := _FeralfileExhibitionV3.contract.FilterLogs(opts, "NewArtwork", artworkIDRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV3NewArtworkIterator{contract: _FeralfileExhibitionV3.contract, event: "NewArtwork", logs: logs, sub: sub}, nil
}

// WatchNewArtwork is a free log subscription operation binding the contract event 0x22350b25f1b72bb3621199a79abefeb4fcd77bb1e65638cd09350666e4db0891.
//
// Solidity: event NewArtwork(uint256 indexed artworkID)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Filterer) WatchNewArtwork(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV3NewArtwork, artworkID []*big.Int) (event.Subscription, error) {

	var artworkIDRule []interface{}
	for _, artworkIDItem := range artworkID {
		artworkIDRule = append(artworkIDRule, artworkIDItem)
	}

	logs, sub, err := _FeralfileExhibitionV3.contract.WatchLogs(opts, "NewArtwork", artworkIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV3NewArtwork)
				if err := _FeralfileExhibitionV3.contract.UnpackLog(event, "NewArtwork", log); err != nil {
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

// ParseNewArtwork is a log parse operation binding the contract event 0x22350b25f1b72bb3621199a79abefeb4fcd77bb1e65638cd09350666e4db0891.
//
// Solidity: event NewArtwork(uint256 indexed artworkID)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Filterer) ParseNewArtwork(log types.Log) (*FeralfileExhibitionV3NewArtwork, error) {
	event := new(FeralfileExhibitionV3NewArtwork)
	if err := _FeralfileExhibitionV3.contract.UnpackLog(event, "NewArtwork", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV3NewArtworkEditionIterator is returned from FilterNewArtworkEdition and is used to iterate over the raw logs and unpacked data for NewArtworkEdition events raised by the FeralfileExhibitionV3 contract.
type FeralfileExhibitionV3NewArtworkEditionIterator struct {
	Event *FeralfileExhibitionV3NewArtworkEdition // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV3NewArtworkEditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV3NewArtworkEdition)
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
		it.Event = new(FeralfileExhibitionV3NewArtworkEdition)
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
func (it *FeralfileExhibitionV3NewArtworkEditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV3NewArtworkEditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV3NewArtworkEdition represents a NewArtworkEdition event raised by the FeralfileExhibitionV3 contract.
type FeralfileExhibitionV3NewArtworkEdition struct {
	Owner     common.Address
	ArtworkID *big.Int
	EditionID *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewArtworkEdition is a free log retrieval operation binding the contract event 0x4f21e8cd53f1df1da42ec94ba03f881c1185607b26e4dcb81941535157d73dd4.
//
// Solidity: event NewArtworkEdition(address indexed owner, uint256 indexed artworkID, uint256 indexed editionID)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Filterer) FilterNewArtworkEdition(opts *bind.FilterOpts, owner []common.Address, artworkID []*big.Int, editionID []*big.Int) (*FeralfileExhibitionV3NewArtworkEditionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var artworkIDRule []interface{}
	for _, artworkIDItem := range artworkID {
		artworkIDRule = append(artworkIDRule, artworkIDItem)
	}
	var editionIDRule []interface{}
	for _, editionIDItem := range editionID {
		editionIDRule = append(editionIDRule, editionIDItem)
	}

	logs, sub, err := _FeralfileExhibitionV3.contract.FilterLogs(opts, "NewArtworkEdition", ownerRule, artworkIDRule, editionIDRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV3NewArtworkEditionIterator{contract: _FeralfileExhibitionV3.contract, event: "NewArtworkEdition", logs: logs, sub: sub}, nil
}

// WatchNewArtworkEdition is a free log subscription operation binding the contract event 0x4f21e8cd53f1df1da42ec94ba03f881c1185607b26e4dcb81941535157d73dd4.
//
// Solidity: event NewArtworkEdition(address indexed owner, uint256 indexed artworkID, uint256 indexed editionID)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Filterer) WatchNewArtworkEdition(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV3NewArtworkEdition, owner []common.Address, artworkID []*big.Int, editionID []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var artworkIDRule []interface{}
	for _, artworkIDItem := range artworkID {
		artworkIDRule = append(artworkIDRule, artworkIDItem)
	}
	var editionIDRule []interface{}
	for _, editionIDItem := range editionID {
		editionIDRule = append(editionIDRule, editionIDItem)
	}

	logs, sub, err := _FeralfileExhibitionV3.contract.WatchLogs(opts, "NewArtworkEdition", ownerRule, artworkIDRule, editionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV3NewArtworkEdition)
				if err := _FeralfileExhibitionV3.contract.UnpackLog(event, "NewArtworkEdition", log); err != nil {
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

// ParseNewArtworkEdition is a log parse operation binding the contract event 0x4f21e8cd53f1df1da42ec94ba03f881c1185607b26e4dcb81941535157d73dd4.
//
// Solidity: event NewArtworkEdition(address indexed owner, uint256 indexed artworkID, uint256 indexed editionID)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Filterer) ParseNewArtworkEdition(log types.Log) (*FeralfileExhibitionV3NewArtworkEdition, error) {
	event := new(FeralfileExhibitionV3NewArtworkEdition)
	if err := _FeralfileExhibitionV3.contract.UnpackLog(event, "NewArtworkEdition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV3OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FeralfileExhibitionV3 contract.
type FeralfileExhibitionV3OwnershipTransferredIterator struct {
	Event *FeralfileExhibitionV3OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV3OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV3OwnershipTransferred)
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
		it.Event = new(FeralfileExhibitionV3OwnershipTransferred)
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
func (it *FeralfileExhibitionV3OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV3OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV3OwnershipTransferred represents a OwnershipTransferred event raised by the FeralfileExhibitionV3 contract.
type FeralfileExhibitionV3OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FeralfileExhibitionV3OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeralfileExhibitionV3.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV3OwnershipTransferredIterator{contract: _FeralfileExhibitionV3.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV3OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeralfileExhibitionV3.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV3OwnershipTransferred)
				if err := _FeralfileExhibitionV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Filterer) ParseOwnershipTransferred(log types.Log) (*FeralfileExhibitionV3OwnershipTransferred, error) {
	event := new(FeralfileExhibitionV3OwnershipTransferred)
	if err := _FeralfileExhibitionV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeralfileExhibitionV3TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FeralfileExhibitionV3 contract.
type FeralfileExhibitionV3TransferIterator struct {
	Event *FeralfileExhibitionV3Transfer // Event containing the contract specifics and raw log

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
func (it *FeralfileExhibitionV3TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeralfileExhibitionV3Transfer)
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
		it.Event = new(FeralfileExhibitionV3Transfer)
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
func (it *FeralfileExhibitionV3TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeralfileExhibitionV3TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeralfileExhibitionV3Transfer represents a Transfer event raised by the FeralfileExhibitionV3 contract.
type FeralfileExhibitionV3Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*FeralfileExhibitionV3TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV3.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FeralfileExhibitionV3TransferIterator{contract: _FeralfileExhibitionV3.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FeralfileExhibitionV3Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _FeralfileExhibitionV3.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeralfileExhibitionV3Transfer)
				if err := _FeralfileExhibitionV3.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_FeralfileExhibitionV3 *FeralfileExhibitionV3Filterer) ParseTransfer(log types.Log) (*FeralfileExhibitionV3Transfer, error) {
	event := new(FeralfileExhibitionV3Transfer)
	if err := _FeralfileExhibitionV3.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
