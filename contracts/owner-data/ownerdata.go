package owner_data

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	ethereum "github.com/bitmark-inc/account-vault-ethereum"
	ownerdata "github.com/bitmark-inc/feralfile-exhibition-smart-contract/go-binding/owner-data"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

type OwnerDataContract struct {
	contractAddress string
}

func OwnerDataContractFactory(contractAddress string) ethereum.Contract {
	return &OwnerDataContract{
		contractAddress: contractAddress,
	}
}

// Deploy deploys the smart contract to ethereum blockchain
func (c *OwnerDataContract) Deploy(wallet *ethereum.Wallet, arguments json.RawMessage) (string, string, error) {
	return "", "", errors.New("not implemented")
}

// Call is the entry function for account vault to interact with a smart contract.
func (c *OwnerDataContract) Call(wallet *ethereum.Wallet, method, fund string, arguments json.RawMessage, noSend bool, customizeGasPriceInWei *int64, customizedNonce *uint64) (*types.Transaction, error) {
	contract, err := ownerdata.NewOwnerData(common.HexToAddress(c.contractAddress), wallet.RPCClient())
	if err != nil {
		return nil, err
	}

	t, err := wallet.Transactor()
	if err != nil {
		return nil, err
	}

	t.NoSend = noSend
	if customizeGasPriceInWei != nil && *customizeGasPriceInWei != 0 {
		t.GasPrice = big.NewInt(*customizeGasPriceInWei * params.Wei)
	}

	if customizedNonce != nil {
		t.Nonce = big.NewInt(int64(*customizedNonce))
	}

	params, err := c.Parse(wallet, method, arguments)
	if nil != err {
		return nil, err
	}

	switch method {
	case "signedAdd":
		if len(params) != 4 {
			return nil, errors.New("Invalid parameters")
		}

		contractAddress, ok := params[0].(common.Address)
		if !ok {
			return nil, errors.New("Invalid contract address")
		}

		tokenID, ok := params[1].(*big.Int)
		if !ok {
			return nil, errors.New("Invalid token id")
		}

		data, ok := params[2].(ownerdata.OwnerDataData)
		if !ok {
			return nil, errors.New("Invalid data")
		}

		signature, ok := params[3].(ownerdata.OwnerDataSignature)
		if !ok {
			return nil, errors.New("Invalid signature")
		}

		return contract.SignedAdd(t,
			contractAddress,
			tokenID,
			data,
			signature,
		)
	default:
		return nil, fmt.Errorf("unsupported method")
	}
}

func (c *OwnerDataContract) Pack(
	wallet *ethereum.Wallet,
	method string,
	arguments json.RawMessage) ([]byte, error) {
	abi, err := ownerdata.OwnerDataMetaData.GetAbi()
	if nil != err {
		return nil, err
	}

	parsedArgs, err := c.Parse(wallet, method, arguments)
	if nil != err {
		return nil, err
	}

	return abi.Pack(method, parsedArgs...)
}

func (c *OwnerDataContract) Parse(
	wallet *ethereum.Wallet,
	method string,
	arguments json.RawMessage) ([]interface{}, error) {
	switch method {
	case "signedAdd":
		var params struct {
			ContractAddress string          `json:"contract_address"`
			TokenID         ethereum.BigInt `json:"token_id"`
			Data            struct {
				Owner    string `json:"owner"`
				DataHash string `json:"data_hash"`
				Metadata string `json:"metadata"`
			} `json:"data"`
			Signature struct {
				OwnerSignature string          `json:"owner_signature"`
				ExpiryBlock    ethereum.BigInt `json:"expiry_block"`
				R              string          `json:"r"`
				S              string          `json:"s"`
				V              string          `json:"v"`
			} `json:"signature"`
		}

		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		contractAddress := common.HexToAddress(params.ContractAddress)
		rVal, err := hex.DecodeString(strings.Replace(params.Signature.R, "0x", "", -1))
		if err != nil {
			return nil, err
		}
		sVal, err := hex.DecodeString(strings.Replace(params.Signature.S, "0x", "", -1))
		if err != nil {
			return nil, err
		}
		vVal, err := strconv.ParseUint(strings.Replace(params.Signature.V, "0x", "", -1), 16, 64)
		if err != nil {
			return nil, err
		}
		if len(rVal) != 32 || len(sVal) != 32 {
			return nil, errors.New("required signature length is 32")
		}
		var r32Val [32]byte
		var s32Val [32]byte
		copy(r32Val[:], rVal)
		copy(s32Val[:], sVal)

		data := ownerdata.OwnerDataData{
			Owner:    common.HexToAddress(params.Data.Owner),
			DataHash: []byte(params.Data.DataHash),
			Metadata: params.Data.Metadata,
		}
		signature := ownerdata.OwnerDataSignature{
			OwnerSign:   common.Hex2Bytes(strings.Replace(params.Signature.OwnerSignature, "0x", "", -1)),
			ExpiryBlock: &params.Signature.ExpiryBlock.Int,
			R:           r32Val,
			S:           s32Val,
			V:           uint8(vVal),
		}
		return []interface{}{
			contractAddress,
			&params.TokenID.Int,
			data,
			signature}, nil
	default:
		return nil, fmt.Errorf("unsupported method")
	}
}

func (c *OwnerDataContract) EstimateGasLimit(
	wallet *ethereum.Wallet,
	contractAddr common.Address,
	method string,
	arguments json.RawMessage) (uint64, error) {
	return 0, errors.New("not implemented")
}

func init() {
	ethereum.RegisterContract("OwnerData", OwnerDataContractFactory)
}
