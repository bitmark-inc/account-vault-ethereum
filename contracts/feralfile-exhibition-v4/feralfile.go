package feralfilev4

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"

	ethereum "github.com/bitmark-inc/account-vault-ethereum"
)

const (
	GasLimitPerMint         = 450000
	GasLimitPerAuthTransfer = 150000
)

type FeralfileExhibitionV4Contract struct {
	contractAddress string
}

func FeralfileExhibitionV4ContractFactory(contractAddress string) ethereum.Contract {
	return &FeralfileExhibitionV4Contract{
		contractAddress: contractAddress,
	}
}

// Deploy deploys the smart contract to ethereum blockchain
func (c *FeralfileExhibitionV4Contract) Deploy(wallet *ethereum.Wallet, arguments json.RawMessage) (string, string, error) {
	var params struct {
		Name              string         `json:"name"`
		Symbol            string         `json:"symbol"`
		Signer            common.Address `json:"signer"`
		Vault             common.Address `json:"vault"`
		CostReceiver      common.Address `json:"cost_receiver"`
		ContractURI       string         `json:"contract_uri"`
		TokenBaseURI      string         `json:"token_base_uri"`
		IsBurnable        bool           `json:"is_burnable"`
		IsBridgeable      bool           `json:"is_bridgeable"`
		SeriesIDs         []*big.Int     `json:"seriesIDs"`
		SeriesMaxSupplies []*big.Int     `json:"series_max_supplies"`
	}

	if err := json.Unmarshal(arguments, &params); err != nil {
		return "", "", err
	}

	t, err := wallet.Transactor()
	if err != nil {
		return "", "", err
	}

	address, tx, _, err := DeployFeralfileExhibitionV4(t, wallet.RPCClient(),
		params.Name, params.Symbol, params.Signer, params.IsBurnable,
		params.IsBridgeable, params.Vault, params.CostReceiver, params.TokenBaseURI,
		params.ContractURI, params.SeriesIDs, params.SeriesMaxSupplies)
	if err != nil {
		return "", "", err
	}
	return address.String(), tx.Hash().String(), nil
}

// Call is the entry function for account vault to interact with a smart contract.
func (c *FeralfileExhibitionV4Contract) Call(wallet *ethereum.Wallet, method, fund string, arguments json.RawMessage, noSend bool, customizeGasPriceInWei *int64, customizedNonce *uint64) (*types.Transaction, error) {
	contract, err := NewFeralfileExhibitionV4(common.HexToAddress(c.contractAddress), wallet.RPCClient())
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

	switch method {
	case "burnArtworks":
		var params []ethereum.BigInt

		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		burnParams := make([]*big.Int, 0)
		for _, v := range params {
			tokenID := v.Int
			burnParams = append(burnParams, &tokenID)
		}

		tx, err := contract.BurnArtworks(t, burnParams)
		if err != nil {
			return nil, err
		}
		return tx, nil
	default:
		return nil, fmt.Errorf("unsupported method")
	}
}

func (c *FeralfileExhibitionV4Contract) ParamEncoder(method string, arguments json.RawMessage) ([]byte, error) {
	parsed, err := abi.JSON(strings.NewReader(FeralfileExhibitionV4ABI))
	if err != nil {
		return nil, err
	}

	switch method {
	case "buyArtworks":
		var params struct {
			SaleData IFeralfileSaleDataSaleData
			R        string
			S        string
			V        string
		}

		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		rVal, err := hex.DecodeString(strings.Replace(params.R, "0x", "", -1))
		if err != nil {
			return nil, err
		}
		sVal, err := hex.DecodeString(strings.Replace(params.S, "0x", "", -1))
		if err != nil {
			return nil, err
		}
		vVal, err := strconv.ParseUint(strings.Replace(params.V, "0x", "", -1), 16, 64)
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

		input, err := parsed.Pack(method, r32Val, s32Val, uint8(vVal), params.SaleData)
		if err != nil {
			return nil, err
		}
		return input, nil
	default:
		return nil, fmt.Errorf("unsupported method")
	}
}

func init() {
	ethereum.RegisterContract("FeralfileExhibitionV4", FeralfileExhibitionV4ContractFactory)
}
