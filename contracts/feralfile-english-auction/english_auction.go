package english_auction

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

type FeralfileEnglishAuctionContract struct {
	contractAddress string
}

func FeralfileEnglishAuctionFactory(contractAddress string) ethereum.Contract {
	return &FeralfileEnglishAuctionContract{
		contractAddress: contractAddress,
	}
}

// Deploy deploys the smart contract to ethereum blockchain
func (c *FeralfileEnglishAuctionContract) Deploy(wallet *ethereum.Wallet, arguments json.RawMessage) (string, string, error) {
	var params struct {
		Signer common.Address `json:"signer"`
	}

	if err := json.Unmarshal(arguments, &params); err != nil {
		return "", "", err
	}

	t, err := wallet.Transactor()
	if err != nil {
		return "", "", err
	}

	address, tx, _, err := DeployFeralfileEnglishAuction(t, wallet.RPCClient(), params.Signer)
	if err != nil {
		return "", "", err
	}
	return address.String(), tx.Hash().String(), nil
}

// Call is the entry function for account vault to interact with a smart contract.
func (c *FeralfileEnglishAuctionContract) Call(wallet *ethereum.Wallet, method, fund string, arguments json.RawMessage, noSend bool, customizeGasPriceInWei *int64, customizedNonce *uint64) (*types.Transaction, error) {
	contract, err := NewFeralfileEnglishAuction(common.HexToAddress(c.contractAddress), wallet.RPCClient())
	if err != nil {
		return nil, err
	}
	fmt.Println(c.contractAddress)
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
	case "registerAuctions":
		var params []struct {
			ID                ethereum.BigInt `json:"id"`
			StartAt           ethereum.BigInt `json:"start_at"`
			EndAt             ethereum.BigInt `json:"end_at"`
			ExtendDuration    ethereum.BigInt `json:"extend_duration"`
			ExtendThreshold   ethereum.BigInt `json:"extend_threshold"`
			MinIncreaseFactor ethereum.BigInt `json:"min_increase_factor"`
			MinIncreaseAmount ethereum.BigInt `json:"min_increase_amount"`
			MinPrice          ethereum.BigInt `json:"min_price"`
			IsSettled         bool            `json:"is_settled"`
		}
		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}
		if len(params) == 0 {
			return nil, errors.New("Invalid params")
		}

		fmt.Println(params)

		auctions := make([]FeralfileEnglishAuctionAuction, len(params))
		for i, v := range params {
			k := v
			auctions[i] = FeralfileEnglishAuctionAuction{
				Id:                &k.ID.Int,
				StartAt:           &k.StartAt.Int,
				EndAt:             &k.EndAt.Int,
				ExtendDuration:    &k.ExtendDuration.Int,
				ExtendThreshold:   &k.ExtendThreshold.Int,
				MinIncreaseFactor: &k.MinIncreaseFactor.Int,
				MinIncreaseAmount: &k.MinIncreaseAmount.Int,
				MinPrice:          &k.MinPrice.Int,
				IsSettled:         k.IsSettled,
			}
		}

		fmt.Println(auctions)
		tx, err := contract.RegisterAuctions(t, auctions)
		if err != nil {
			return nil, err
		}
		return tx, nil
	case "placeSignedBid":
		var params struct {
			AuctionID ethereum.BigInt `json:"auction_id"`
			Bidder    common.Address  `json:"bidder"`
			Amount    ethereum.BigInt `json:"amount"`
			Timestamp ethereum.BigInt `json:"timestamp"`
			R         string          `json:"r"`
			S         string          `json:"s"`
			V         string          `json:"v"`
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

		return contract.PlaceSignedBid(t, &params.AuctionID.Int, params.Bidder, &params.Amount.Int, &params.Timestamp.Int, r32Val, s32Val, uint8(vVal))
	case "settleAuction":
		var params struct {
			AuctionID       ethereum.BigInt `json:"auction_id"`
			ContractAddress common.Address  `json:"contract_address"`
			VaultAddress    common.Address  `json:"vault_address"`
			SaleData        struct {
				Price         ethereum.BigInt   `json:"price"`
				Cost          ethereum.BigInt   `json:"cost"`
				ExpiryTime    ethereum.BigInt   `json:"expiry_time"`
				Destination   common.Address    `json:"destination"`
				TokenIds      []ethereum.BigInt `json:"token_ids"`
				RevenueShares [][]struct {
					Recipient common.Address  `json:"recipient"`
					Bps       ethereum.BigInt `json:"bps"`
				} `json:"revenue_shares"`
				PayByVaultContract bool
			} `json:"sale_data"`
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

		revenueShares := make([][]IFeralfileSaleDataRevenueShare, 0)
		for _, v := range params.SaleData.RevenueShares {
			revenueShare := make([]IFeralfileSaleDataRevenueShare, 0)
			for _, vv := range v {
				bps := vv.Bps.Int
				revenueShare = append(revenueShare, IFeralfileSaleDataRevenueShare{
					Recipient: vv.Recipient,
					Bps:       &bps,
				})
			}
			revenueShares = append(revenueShares, revenueShare)
		}

		saleData := IFeralfileSaleDataSaleData{
			Price:              &params.SaleData.Price.Int,
			Cost:               &params.SaleData.Cost.Int,
			ExpiryTime:         &params.SaleData.ExpiryTime.Int,
			Destination:        params.SaleData.Destination,
			TokenIds:           tokenIDs,
			RevenueShares:      revenueShares,
			PayByVaultContract: params.SaleData.PayByVaultContract,
		}

		tx, err := contract.SettleAuction(t, &params.AuctionID.Int, params.ContractAddress, params.VaultAddress, saleData, r32Val, s32Val, uint8(vVal))
		if err != nil {
			return nil, err
		}
		return tx, nil
	default:
		return nil, fmt.Errorf("unsupported method")
	}
}

func (c *FeralfileEnglishAuctionContract) ParamEncoder(method string, arguments json.RawMessage) ([]byte, error) {
	return nil, fmt.Errorf("unsupported method")
}

func init() {
	ethereum.RegisterContract("FeralfileEnglishAuction", FeralfileEnglishAuctionFactory)
}
