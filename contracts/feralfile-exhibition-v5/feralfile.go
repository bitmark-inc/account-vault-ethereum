package feralfilev5

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"

	ethereum "github.com/bitmark-inc/account-vault-ethereum"
)

const (
	GasLimitPerMint       = 250000
	GasLimitPerBurn       = 50000
	GasLimitApproveForAll = 80000
)

type FeralfileExhibitionV5Contract struct {
	contractAddress string
}

func FeralfileExhibitionV5ContractFactory(contractAddress string) ethereum.Contract {
	return &FeralfileExhibitionV5Contract{
		contractAddress: contractAddress,
	}
}

// Deploy deploys the smart contract to ethereum blockchain
func (c *FeralfileExhibitionV5Contract) Deploy(wallet *ethereum.Wallet, arguments json.RawMessage) (string, string, error) {
	var params struct {
		BaseTokenURI             string         `json:"base_token_uri"`
		ContractURI              string         `json:"contract_uri"`
		Signer                   common.Address `json:"signer"`
		Vault                    common.Address `json:"vault"`
		CostReceiver             common.Address `json:"cost_receiver"`
		IsBurnable               bool           `json:"is_burnable"`
		SeriesIDs                []*big.Int     `json:"series_ids"`
		SeriesMaxSupplies        []*big.Int     `json:"series_max_supplies"`
		SeriesArtworkMaxSupplies []*big.Int     `json:"series_artwork_max_supplies"`
	}

	if err := json.Unmarshal(arguments, &params); err != nil {
		return "", "", err
	}

	t, err := wallet.Transactor()
	if err != nil {
		return "", "", err
	}

	address, tx, _, err := DeployFeralfileExhibitionV5(
		t,
		wallet.RPCClient(),
		params.BaseTokenURI,
		params.ContractURI,
		params.Signer,
		params.Vault,
		params.CostReceiver,
		params.IsBurnable,
		params.SeriesIDs,
		params.SeriesMaxSupplies,
		params.SeriesArtworkMaxSupplies,
	)
	if err != nil {
		return "", "", err
	}
	return address.String(), tx.Hash().String(), nil
}

// Call is the entry function for account vault to interact with a smart contract.
func (c *FeralfileExhibitionV5Contract) Call(wallet *ethereum.Wallet, method, fund string, arguments json.RawMessage, noSend bool, customizeGasPriceInWei *int64, customizedNonce *uint64) (*types.Transaction, error) {
	contract, err := NewFeralfileExhibitionV5(common.HexToAddress(c.contractAddress), wallet.RPCClient())
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
		var params []struct {
			From    string          `json:"from"`
			TokenID ethereum.BigInt `json:"token_id"`
			Amount  ethereum.BigInt `json:"amount"`
		}
		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}
		if len(params) == 0 {
			return nil, errors.New("Invalid token burn parameters")
		}

		t.GasLimit = uint64(GasLimitPerBurn * len(params))

		burnData := make([]FeralfileExhibitionV5BurnData, len(params))
		for i := 0; i < len(params); i++ {
			burnData[i] = FeralfileExhibitionV5BurnData{
				From:    common.HexToAddress(params[i].From),
				TokenId: &params[i].TokenID.Int,
				Amount:  &params[i].Amount.Int,
			}
		}

		return contract.BurnArtworks(t, burnData)
	case "mintArtworks":
		var params []struct {
			SeriesID ethereum.BigInt `json:"series_id"`
			TokenID  ethereum.BigInt `json:"token_id"`
			Owner    common.Address  `json:"owner"`
			Amount   ethereum.BigInt `json:"amount"`
		}
		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}
		if len(params) == 0 {
			return nil, errors.New("Invalid token mint parameters")
		}

		t.GasLimit = uint64(GasLimitPerMint * len(params))

		mintData := make([]FeralfileExhibitionV5MintData, len(params))
		for i := 0; i < len(params); i++ {
			mintData[i] = FeralfileExhibitionV5MintData{
				SeriesId: &params[i].SeriesID.Int,
				TokenId:  &params[i].TokenID.Int,
				Owner:    params[i].Owner,
				Amount:   &params[i].Amount.Int,
			}
		}

		return contract.MintArtworks(t, mintData)
	case "buyArtworks":
		var params struct {
			SaleData struct {
				Price         ethereum.BigInt   `json:"price"`
				Cost          ethereum.BigInt   `json:"cost"`
				ExpiryTime    ethereum.BigInt   `json:"expiryTime"`
				Destination   common.Address    `json:"destination"`
				TokenIds      []ethereum.BigInt `json:"tokenIds"`
				RevenueShares [][]struct {
					Recipient common.Address  `json:"recipient"`
					Bps       ethereum.BigInt `json:"bps"`
				} `json:"revenueShares"`
				PayByVaultContract bool            `json:"payByVaultContract"`
				BiddingUnix        ethereum.BigInt `json:"biddingUnix"`
			} `json:"saleData"`
			R string `json:"r"`
			S string `json:"s"`
			V string `json:"v"`
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

		tokenIDs := make([]*big.Int, 0)
		for _, v := range params.SaleData.TokenIds {
			tokenID := v.Int
			tokenIDs = append(tokenIDs, &tokenID)
		}

		revenueShares := make([][]IFeralfileSaleDataV2RevenueShare, 0)
		for _, v := range params.SaleData.RevenueShares {
			revenueShare := make([]IFeralfileSaleDataV2RevenueShare, 0)
			for _, vv := range v {
				bps := vv.Bps.Int
				revenueShare = append(revenueShare, IFeralfileSaleDataV2RevenueShare{
					Recipient: vv.Recipient,
					Bps:       &bps,
				})
			}
			revenueShares = append(revenueShares, revenueShare)
		}

		saleData := IFeralfileSaleDataV2SaleData{
			Price:              &params.SaleData.Price.Int,
			Cost:               &params.SaleData.Cost.Int,
			ExpiryTime:         &params.SaleData.ExpiryTime.Int,
			Destination:        params.SaleData.Destination,
			TokenIds:           tokenIDs,
			RevenueShares:      revenueShares,
			PayByVaultContract: params.SaleData.PayByVaultContract,
			BiddingUnix:        &params.SaleData.BiddingUnix.Int,
		}

		tx, err := contract.BuyArtworks(t, r32Val, s32Val, uint8(vVal), saleData)
		if err != nil {
			return nil, err
		}
		return tx, nil
	case "approve_for_all":
		var params struct {
			Operator common.Address `json:"operator"`
		}
		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		t.GasLimit = GasLimitApproveForAll

		tx, err := contract.SetApprovalForAll(t, params.Operator, true)
		if err != nil {
			return nil, err
		}
		return tx, nil
	case "startSale":
		return contract.StartSale(t)
	case "setVault":
		var params struct {
			Vault common.Address `json:"vault"`
		}
		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}
		return contract.SetVault(t, params.Vault)
	default:
		return nil, fmt.Errorf("unsupported method")
	}
}

func (c *FeralfileExhibitionV5Contract) ParamEncoder(method string, arguments json.RawMessage) ([]byte, error) {
	return nil, fmt.Errorf("unsupported method")
}

func init() {
	ethereum.RegisterContract("FeralfileExhibitionV5", FeralfileExhibitionV5ContractFactory)
}
