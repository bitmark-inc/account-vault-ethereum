package feralfiletoken

import (
	"encoding/json"
	"fmt"
	"math/big"

	ethereum "github.com/bitmark-inc/account-vault-ethereum"
	feralfiletoken "github.com/bitmark-inc/feralfile-exhibition-smart-contract/go-binding/feralfile-token"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

type FeralFileTokenContract struct {
	contractAddress string
}

func FeralFileTokenContractFactory(contractAddress string) ethereum.Contract {
	return &FeralFileTokenContract{
		contractAddress: contractAddress,
	}
}

func (c *FeralFileTokenContract) Deploy(
	wallet *ethereum.Wallet,
	arguments json.RawMessage) (string, string, error) {
	var params struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	}

	if err := json.Unmarshal(arguments, &params); err != nil {
		return "", "", err
	}

	t, err := wallet.Transactor()
	if err != nil {
		return "", "", err
	}

	address, tx, _, err := feralfiletoken.DeployFeralfileToken(t, wallet.RPCClient(), params.Name, params.Symbol)
	if err != nil {
		return "", "", err
	}
	return address.String(), tx.Hash().String(), nil
}

func (c *FeralFileTokenContract) Call(
	wallet *ethereum.Wallet,
	method,
	fund string,
	arguments json.RawMessage,
	noSend bool,
	gasLimit uint64,
	gasPrice *int64,
	nonce *uint64) (*types.Transaction, error) {
	contractAddr := common.HexToAddress(c.contractAddress)
	contract, err := feralfiletoken.NewFeralfileToken(
		contractAddr,
		wallet.RPCClient())
	if err != nil {
		return nil, err
	}

	t, err := wallet.Transactor()
	if err != nil {
		return nil, err
	}

	t.NoSend = noSend
	t.GasLimit = gasLimit
	if gasPrice != nil && *gasPrice != 0 {
		t.GasPrice = big.NewInt(*gasPrice * params.Wei)
	}

	if nonce != nil {
		t.Nonce = big.NewInt(int64(*nonce))
	}

	params, err := c.Parse(method, arguments)
	if nil != err {
		return nil, err
	}

	switch method {
	case "mint":
		if len(params) != 2 {
			return nil, fmt.Errorf("invalid parameters")
		}

		owner, ok := params[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid owner")
		}

		amount, ok := params[1].(*big.Int)
		if !ok {
			return nil, fmt.Errorf("invalid amount")
		}

		return contract.Mint(t, owner, amount)
	case "batchMint":
		if len(params) != 2 {
			return nil, fmt.Errorf("invalid parameters")
		}

		owners, ok := params[0].([]common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid owners")
		}

		amounts, ok := params[1].([]*big.Int)
		if !ok {
			return nil, fmt.Errorf("invalid amounts")
		}

		return contract.BatchMint(t, owners, amounts)
	}

	return nil, fmt.Errorf("unsupported method")
}

func (c *FeralFileTokenContract) Pack(
	method string,
	arguments json.RawMessage) ([]byte, error) {
	abi, err := feralfiletoken.FeralfileTokenMetaData.GetAbi()
	if nil != err {
		return nil, err
	}

	parsedArgs, err := c.Parse(method, arguments)
	if nil != err {
		return nil, err
	}

	return abi.Pack(method, parsedArgs...)
}

func (c *FeralFileTokenContract) Parse(
	method string,
	arguments json.RawMessage) ([]interface{}, error) {
	switch method {
	case "mint":
		var params struct {
			Owner  *common.Address  `json:"owner"`
			Amount *ethereum.BigInt `json:"amount"`
		}

		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		return []interface{}{&params.Owner, &params.Amount.Int}, nil
	case "batchMint":
		var params struct {
			Owners  []common.Address   `json:"owners"`
			Amounts []*ethereum.BigInt `json:"amounts"`
		}

		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		amounts := make([]*big.Int, 0)
		for _, a := range params.Amounts {
			v := a.Int
			amounts = append(amounts, &v)
		}

		return []interface{}{params.Owners, amounts}, nil
	default:
		return nil, fmt.Errorf("unsupported method")
	}
}

func init() {
	ethereum.RegisterContract("FeralfileToken", FeralFileTokenContractFactory)
}
