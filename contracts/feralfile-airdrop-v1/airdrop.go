package airdropv1

import (
	"encoding/json"
	"fmt"
	"math/big"

	ethereum "github.com/bitmark-inc/account-vault-ethereum"
	airdropv1 "github.com/bitmark-inc/feralfile-exhibition-smart-contract/go-binding/feralfile-airdrop-v1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

type FeralFileAirdropV1Contract struct {
	contractAddress string
}

func FeralFileAirdropV1ContractFactory(contractAddress string) ethereum.Contract {
	return &FeralFileAirdropV1Contract{
		contractAddress: contractAddress,
	}
}

func (c *FeralFileAirdropV1Contract) Deploy(
	wallet *ethereum.Wallet,
	arguments json.RawMessage) (string, string, error) {
	var params struct {
		TokenType   uint8  `json:"token_type"`
		TokenURI    string `json:"token_uri"`
		ContractURI string `json:"contract_uri"`
		Burnable    bool   `json:"burnable"`
		Bridgeable  bool   `json:"bridgeable"`
	}

	if err := json.Unmarshal(arguments, &params); err != nil {
		return "", "", err
	}

	t, err := wallet.Transactor()
	if err != nil {
		return "", "", err
	}

	address, tx, _, err := airdropv1.DeployFeralFileAirdropV1(
		t,
		wallet.RPCClient(),
		params.TokenType,
		params.TokenURI,
		params.ContractURI,
		params.Burnable,
		params.Bridgeable,
	)
	if err != nil {
		return "", "", err
	}

	return address.Hex(), tx.Hash().Hex(), nil
}

func (c *FeralFileAirdropV1Contract) Call(
	wallet *ethereum.Wallet,
	method,
	fund string,
	arguments json.RawMessage,
	noSend bool,
	customizeGasPriceInWei *int64,
	customizedNonce *uint64) (*types.Transaction, error) {
	contract, err := airdropv1.NewFeralFileAirdropV1(
		common.HexToAddress(c.contractAddress),
		wallet.RPCClient())
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

	params, err := c.Parse(method, arguments)
	if nil != err {
		return nil, err
	}

	switch method {
	case "mint":
		if len(params) != 2 {
			return nil, fmt.Errorf("invalid parameters")
		}

		tokenID, ok := params[0].(*big.Int)
		if !ok {
			return nil, fmt.Errorf("invalid token id")
		}

		amount, ok := params[1].(*big.Int)
		if !ok {
			return nil, fmt.Errorf("invalid amount")
		}

		t.GasLimit = 100000
		return contract.Mint(t, tokenID, amount)
	case "airdrop":
		if len(params) != 2 {
			return nil, fmt.Errorf("invalid parameters")
		}

		tokenID, ok := params[0].(*big.Int)
		if !ok {
			return nil, fmt.Errorf("invalid token id")
		}

		to, ok := params[1].([]common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid to address")
		}

		const gasLimitPerItem = 150000
		t.GasLimit = uint64(gasLimitPerItem * len(to))

		return contract.Airdrop(t, tokenID, to)
	}

	return nil, fmt.Errorf("unsupported method")
}

func (c *FeralFileAirdropV1Contract) Pack(
	method string,
	arguments json.RawMessage) ([]byte, error) {
	abi, err := airdropv1.FeralFileAirdropV1MetaData.GetAbi()
	if nil != err {
		return nil, err
	}

	parsedArgs, err := c.Parse(method, arguments)
	if nil != err {
		return nil, err
	}

	return abi.Pack(method, parsedArgs...)
}

func (c *FeralFileAirdropV1Contract) Parse(
	method string,
	arguments json.RawMessage) ([]interface{}, error) {
	switch method {
	case "mint":
		var params struct {
			TokenID *ethereum.BigInt `json:"token_id"`
			Amount  *ethereum.BigInt `json:"amount"`
		}

		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		return []interface{}{&params.TokenID.Int, &params.Amount.Int}, nil
	case "airdrop":
		var params struct {
			TokenID *ethereum.BigInt `json:"token_id"`
			To      []common.Address `json:"to"`
		}

		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}
		if len(params.To) == 0 {
			return nil, fmt.Errorf("invalid token airdrop parameters (to)")
		}

		return []interface{}{&params.TokenID.Int, params.To}, nil
	default:
		return nil, fmt.Errorf("unsupported method")
	}
}

func init() {
	ethereum.RegisterContract("FeralFileAirdropV1", FeralFileAirdropV1ContractFactory)
}
