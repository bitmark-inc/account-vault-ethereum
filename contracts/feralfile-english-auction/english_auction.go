package english_auction

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
	english_auction "github.com/bitmark-inc/feralfile-exhibition-smart-contract/go-binding/feralfile-english-auction"
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
func (c *FeralfileEnglishAuctionContract) Deploy(
	wallet *ethereum.Wallet,
	arguments json.RawMessage) (string, string, error) {
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

	address, tx, _, err := english_auction.DeployFeralfileEnglishAuction(t, wallet.RPCClient(), params.Signer)
	if err != nil {
		return "", "", err
	}
	return address.String(), tx.Hash().String(), nil
}

// Call is the entry function for account vault to interact with a smart contract.
func (c *FeralfileEnglishAuctionContract) Call(
	wallet *ethereum.Wallet,
	method,
	fund string,
	arguments json.RawMessage,
	noSend bool,
	customizeGasPriceInWei *int64,
	customizedNonce *uint64) (*types.Transaction, error) {
	contract, err := english_auction.NewFeralfileEnglishAuction(common.HexToAddress(c.contractAddress), wallet.RPCClient())
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
	case "registerAuctions":
		if len(params) != 1 {
			return nil, fmt.Errorf("invalid params")
		}

		auctions, ok := params[0].([]english_auction.FeralfileEnglishAuctionAuction)
		if !ok {
			return nil, fmt.Errorf("invalid auctions params")
		}

		return contract.RegisterAuctions(t, auctions)
	case "placeSignedBid":
		if len(params) != 7 {
			return nil, fmt.Errorf("invalid params")
		}

		auctionID, ok := params[0].(big.Int)
		if !ok {
			return nil, fmt.Errorf("invalid auction id")
		}

		bidder, ok := params[1].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid bidder")
		}

		amount, ok := params[2].(big.Int)
		if !ok {
			return nil, fmt.Errorf("invalid amount")
		}

		expiryTime, ok := params[3].(big.Int)
		if !ok {
			return nil, fmt.Errorf("invalid expiry time")
		}

		r32Val, ok := params[4].([32]byte)
		if !ok {
			return nil, fmt.Errorf("invalid r")
		}

		s32Val, ok := params[5].([32]byte)
		if !ok {
			return nil, fmt.Errorf("invalid s")
		}

		vVal, ok := params[6].(uint8)
		if !ok {
			return nil, fmt.Errorf("invalid v")
		}

		return contract.PlaceSignedBid(
			t,
			&auctionID,
			bidder,
			&amount,
			&expiryTime,
			r32Val,
			s32Val,
			vVal)
	case "settleAuction":
		if len(params) != 7 {
			return nil, fmt.Errorf("invalid params")
		}

		auctionID, ok := params[0].(big.Int)
		if !ok {
			return nil, fmt.Errorf("invalid auction id")
		}

		contractAddress, ok := params[1].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid contract address")
		}

		vaultAddress, ok := params[2].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid vault address")
		}

		saleData, ok := params[3].(english_auction.IFeralfileSaleDataSaleData)
		if !ok {
			return nil, fmt.Errorf("invalid sale data")
		}

		r32Val, ok := params[4].([32]byte)
		if !ok {
			return nil, fmt.Errorf("invalid r")
		}

		s32Val, ok := params[5].([32]byte)
		if !ok {
			return nil, fmt.Errorf("invalid s")
		}

		vVal, ok := params[6].(uint8)
		if !ok {
			return nil, fmt.Errorf("invalid v")
		}

		tx, err := contract.SettleAuction(
			t,
			&auctionID,
			contractAddress,
			vaultAddress,
			saleData,
			r32Val,
			s32Val,
			vVal)
		if err != nil {
			return nil, err
		}
		return tx, nil
	default:
		return nil, fmt.Errorf("unsupported method")
	}
}

func (c *FeralfileEnglishAuctionContract) Pack(
	method string,
	arguments json.RawMessage) ([]byte, error) {
	parsedABI, err := abi.JSON(strings.NewReader(english_auction.FeralfileEnglishAuctionABI))
	if nil != err {
		return nil, err
	}

	parsedArgs, err := c.Parse(method, arguments)
	if nil != err {
		return nil, err
	}

	return parsedABI.Pack(method, parsedArgs...)
}

func (c *FeralfileEnglishAuctionContract) Parse(
	method string,
	arguments json.RawMessage) ([]interface{}, error) {
	switch method {
	case "registerAuctions":
		var params []struct {
			ID                ethereum.BigInt `json:"id"`
			StartAt           ethereum.BigInt `json:"startAt"`
			EndAt             ethereum.BigInt `json:"endAt"`
			ExtendDuration    ethereum.BigInt `json:"extendDuration"`
			ExtendThreshold   ethereum.BigInt `json:"extendThreshold"`
			MinIncreaseFactor ethereum.BigInt `json:"minIncreaseFactor"`
			MinIncreaseAmount ethereum.BigInt `json:"minIncreaseAmount"`
			MinPrice          ethereum.BigInt `json:"minPrice"`
			IsSettled         bool            `json:"isSettled"`
		}
		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}
		if len(params) == 0 {
			return nil, errors.New("Invalid params")
		}

		auctions := make([]english_auction.FeralfileEnglishAuctionAuction, len(params))
		for i, v := range params {
			k := v
			auctions[i] = english_auction.FeralfileEnglishAuctionAuction{
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

		return []interface{}{auctions}, nil
	case "placeSignedBid":
		var params struct {
			AuctionID  ethereum.BigInt `json:"auctionID"`
			Bidder     common.Address  `json:"bidder"`
			Amount     ethereum.BigInt `json:"amount"`
			ExpiryTime ethereum.BigInt `json:"expiryTime"`
			R          string          `json:"r"`
			S          string          `json:"s"`
			V          string          `json:"v"`
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

		return []interface{}{
				params.AuctionID.Int,
				params.Bidder,
				params.Amount.Int,
				params.ExpiryTime.Int,
				r32Val,
				s32Val,
				uint8(vVal)},
			nil
	case "settleAuction":
		var params struct {
			AuctionID       ethereum.BigInt `json:"auctionID"`
			ContractAddress common.Address  `json:"contractAddress"`
			VaultAddress    common.Address  `json:"vaultAddress"`
			SaleData        struct {
				Price         ethereum.BigInt   `json:"price"`
				Cost          ethereum.BigInt   `json:"cost"`
				ExpiryTime    ethereum.BigInt   `json:"expiryTime"`
				Destination   common.Address    `json:"destination"`
				TokenIds      []ethereum.BigInt `json:"tokenIds"`
				RevenueShares [][]struct {
					Recipient common.Address  `json:"recipient"`
					Bps       ethereum.BigInt `json:"bps"`
				} `json:"revenueShares"`
				PayByVaultContract bool `json:"payByVaultContract"`
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

		revenueShares := make([][]english_auction.IFeralfileSaleDataRevenueShare, 0)
		for _, v := range params.SaleData.RevenueShares {
			revenueShare := make([]english_auction.IFeralfileSaleDataRevenueShare, 0)
			for _, vv := range v {
				bps := vv.Bps.Int
				revenueShare = append(revenueShare, english_auction.IFeralfileSaleDataRevenueShare{
					Recipient: vv.Recipient,
					Bps:       &bps,
				})
			}
			revenueShares = append(revenueShares, revenueShare)
		}

		saleData := english_auction.IFeralfileSaleDataSaleData{
			Price:              &params.SaleData.Price.Int,
			Cost:               &params.SaleData.Cost.Int,
			ExpiryTime:         &params.SaleData.ExpiryTime.Int,
			Destination:        params.SaleData.Destination,
			TokenIds:           tokenIDs,
			RevenueShares:      revenueShares,
			PayByVaultContract: params.SaleData.PayByVaultContract,
		}

		return []interface{}{
				params.AuctionID.Int,
				params.ContractAddress,
				params.VaultAddress,
				saleData,
				r32Val,
				s32Val,
				uint8(vVal)},
			nil
	default:
		return nil, fmt.Errorf("unsupported method")
	}
}

func init() {
	ethereum.RegisterContract("FeralfileEnglishAuction", FeralfileEnglishAuctionFactory)
}
