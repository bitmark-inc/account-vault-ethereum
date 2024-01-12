package airdropv1

import (
	"encoding/json"
	"fmt"

	ethereum "github.com/bitmark-inc/account-vault-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type FeralFileAirdropV1Contract struct {
	contractAddress string
}

func FeralFileAirdropV1ContractFactory(contractAddress string) ethereum.Contract {
	return &FeralFileAirdropV1Contract{
		contractAddress: contractAddress,
	}
}

func (c *FeralFileAirdropV1Contract) Deploy(wallet *ethereum.Wallet, arguments json.RawMessage) (string, string, error) {
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

	address, tx, _, err := DeployFeralFileAirdropV1(
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

func (c *FeralFileAirdropV1Contract) Call(wallet *ethereum.Wallet, method, fund string, arguments json.RawMessage, noSend bool, customizeGasPriceInWei *int64, customizedNonce *uint64) (*types.Transaction, error) {
	contract, err := NewFeralFileAirdropV1(common.HexToAddress(c.contractAddress), wallet.RPCClient())
	if err != nil {
		return nil, err
	}

	switch method {
	case "mint":
		var params struct {
			TokenID *ethereum.BigInt `json:"token_id"`
			Amount  *ethereum.BigInt `json:"amount"`
		}

		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		t, err := wallet.Transactor()
		if err != nil {
			return nil, err
		}

		t.GasLimit = 100000
		return contract.Mint(t, &params.TokenID.Int, &params.Amount.Int)
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

		t, err := wallet.Transactor()
		if err != nil {
			return nil, err
		}

		const gasLimitPerItem = 150000
		t.GasLimit = uint64(gasLimitPerItem * len(params.To))

		return contract.Airdrop(t, &params.TokenID.Int, params.To)
	}

	return nil, fmt.Errorf("unsupported method")
}

func (c *FeralFileAirdropV1Contract) ParamEncoder(method string, arguments json.RawMessage) ([]byte, error) {
	return nil, fmt.Errorf("unsupported method")
}

func init() {
	ethereum.RegisterContract("FeralFileAirdropV1", FeralFileAirdropV1ContractFactory)
}
