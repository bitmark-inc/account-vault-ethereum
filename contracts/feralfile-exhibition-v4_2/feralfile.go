package feralfilev4_2

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
	feralfilev4 "github.com/bitmark-inc/feralfile-exhibition-smart-contract/go-binding/feralfile-exhibition-v4_2"
)

type FeralfileExhibitionV4_2Contract struct {
	contractAddress string
}

func FeralfileExhibitionV4_2ContractFactory(contractAddress string) ethereum.Contract {
	return &FeralfileExhibitionV4_2Contract{
		contractAddress: contractAddress,
	}
}

// Deploy deploys the smart contract to ethereum blockchain
func (c *FeralfileExhibitionV4_2Contract) Deploy(
	wallet *ethereum.Wallet,
	arguments json.RawMessage) (string, string, error) {
	var params struct {
		Name               string         `json:"name"`
		Symbol             string         `json:"symbol"`
		Signer             common.Address `json:"signer"`
		Vault              common.Address `json:"vault"`
		CostReceiver       common.Address `json:"cost_receiver"`
		ContractURI        string         `json:"contract_uri"`
		IsBurnable         bool           `json:"is_burnable"`
		IsBridgeable       bool           `json:"is_bridgeable"`
		SeriesIDs          []*big.Int     `json:"series_ids"`
		SeriesMaxSupplies  []*big.Int     `json:"series_max_supplies"`
		SeriesNextTokenIds []*big.Int     `json:"series_next_token_ids"`
	}

	if err := json.Unmarshal(arguments, &params); err != nil {
		return "", "", err
	}

	t, err := wallet.Transactor()
	if err != nil {
		return "", "", err
	}

	address, tx, _, err := feralfilev4.DeployFeralfileExhibitionV42(
		t,
		wallet.RPCClient(),
		params.Name,
		params.Symbol,
		params.IsBurnable,
		params.IsBridgeable,
		params.Signer,
		params.Vault,
		params.CostReceiver,
		params.ContractURI,
		params.SeriesIDs,
		params.SeriesMaxSupplies,
		params.SeriesNextTokenIds,
	)
	if err != nil {
		return "", "", err
	}
	return address.String(), tx.Hash().String(), nil
}

// Call is the entry function for account vault to interact with a smart contract.
func (c *FeralfileExhibitionV4_2Contract) Call(
	wallet *ethereum.Wallet,
	method,
	fund string,
	arguments json.RawMessage,
	noSend bool,
	gasLimit uint64,
	gasPrice *int64,
	nonce *uint64) (*types.Transaction, error) {
	contractAddr := common.HexToAddress(c.contractAddress)
	contract, err := feralfilev4.NewFeralfileExhibitionV42(contractAddr, wallet.RPCClient())
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
	case "burnArtworks":
		if len(params) != 1 {
			return nil, errors.New("Invalid parameters")
		}

		tokenIDs, ok := params[0].([]*big.Int)
		if !ok {
			return nil, errors.New("Invalid token burn parameters")
		}

		return contract.BurnArtworks(t, tokenIDs)
	case "mintArtworks":
		if len(params) != 1 {
			return nil, errors.New("Invalid parameters")
		}

		mintData, ok := params[0].([]feralfilev4.FeralfileExhibitionV4MintData)
		if !ok {
			return nil, errors.New("Invalid token mint parameters")
		}

		return contract.MintArtworks(t, mintData)
	case "setTokenBaseURI":
		if len(params) != 1 {
			return nil, errors.New("Invalid parameters")
		}

		baseURI, ok := params[0].(string)
		if !ok {
			return nil, errors.New("Invalid token base URI")
		}

		return contract.SetTokenBaseURI(t, baseURI)
	case "startSale":
		return contract.StartSale(t)
	case "transferOwnership":
		if len(params) != 1 {
			return nil, errors.New("Invalid parameters")
		}

		newOwner, ok := params[0].(common.Address)
		if !ok {
			return nil, errors.New("Invalid new owner")
		}

		return contract.TransferOwnership(t, newOwner)
	case "buyArtworks":
		if len(params) != 4 {
			return nil, errors.New("Invalid parameters")
		}

		r, ok := params[0].([32]byte)
		if !ok {
			return nil, fmt.Errorf("invalid r")
		}

		s, ok := params[1].([32]byte)
		if !ok {
			return nil, fmt.Errorf("invalid s")
		}

		v, ok := params[2].(uint8)
		if !ok {
			return nil, fmt.Errorf("invalid v")
		}

		saleData, ok := params[3].(feralfilev4.IFeralfileSaleDataSaleData)
		if !ok {
			return nil, fmt.Errorf("invalid sale data")
		}

		return contract.BuyArtworks(t, r, s, v, saleData)
	case "buyBulkArtworks":
		if len(params) != 4 {
			return nil, errors.New("Invalid parameters")
		}

		r, ok := params[0].([32]byte)
		if !ok {
			return nil, fmt.Errorf("invalid r")
		}

		s, ok := params[1].([32]byte)
		if !ok {
			return nil, fmt.Errorf("invalid s")
		}

		v, ok := params[2].(uint8)
		if !ok {
			return nil, fmt.Errorf("invalid v")
		}

		saleData, ok := params[3].(feralfilev4.IFeralfileSaleDataV2SaleDataV2)
		if !ok {
			return nil, fmt.Errorf("invalid sale data")
		}

		return contract.BuyBulkArtworks(t, r, s, v, saleData)
	case "setApprovalForAll":
		if len(params) != 2 {
			return nil, errors.New("Invalid parameters")
		}

		operator, ok := params[0].(common.Address)
		if !ok {
			return nil, errors.New("Invalid operator")
		}

		approved, ok := params[1].(bool)
		if !ok {
			return nil, errors.New("Invalid approved")
		}

		t.GasLimit = gasLimit

		return contract.SetApprovalForAll(t, operator, approved)
	case "setAdvanceSetting":
		if len(params) != 2 {
			return nil, errors.New("Invalid parameters")
		}

		return contract.SetAdvanceSetting(t, params[0].([]common.Address), params[1].([]*big.Int))
	default:
		return nil, fmt.Errorf("unsupported method")
	}
}

func (c *FeralfileExhibitionV4_2Contract) Pack(
	method string,
	arguments json.RawMessage) ([]byte, error) {
	abi, err := feralfilev4.FeralfileExhibitionV42MetaData.GetAbi()
	if nil != err {
		return nil, err
	}

	parsedArgs, err := c.Parse(method, arguments)
	if nil != err {
		return nil, err
	}

	return abi.Pack(method, parsedArgs...)
}

func (c *FeralfileExhibitionV4_2Contract) Parse(
	method string,
	arguments json.RawMessage) ([]interface{}, error) {
	switch method {
	case "startSale":
		return nil, nil
	case "burnArtworks":
		var params []ethereum.BigInt
		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}
		if len(params) == 0 {
			return nil, errors.New("Invalid token burn parameters")
		}

		tokenIDs := make([]*big.Int, len(params))
		for i, v := range params {
			tokenID := v.Int
			tokenIDs[i] = &tokenID
		}

		return []interface{}{tokenIDs}, nil
	case "mintArtworks":
		var params []struct {
			SeriesID ethereum.BigInt `json:"series_id"`
			TokenID  ethereum.BigInt `json:"token_id"`
			Owner    common.Address  `json:"owner"`
		}
		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}
		if len(params) == 0 {
			return nil, errors.New("Invalid token mint parameters")
		}

		mintData := make([]feralfilev4.FeralfileExhibitionV4MintData, len(params))
		for i := 0; i < len(params); i++ {
			mintData[i] = feralfilev4.FeralfileExhibitionV4MintData{
				SeriesId: &params[i].SeriesID.Int,
				TokenId:  &params[i].TokenID.Int,
				Owner:    params[i].Owner,
			}
		}

		return []interface{}{mintData}, nil
	case "setTokenBaseURI":
		baseURI := strings.Trim(string(arguments), "\"")
		if baseURI == "" {
			return nil, errors.New("Invalid token base URI")
		}
		return []interface{}{baseURI}, nil
	case "transferOwnership":
		var newOwner common.Address
		if err := json.Unmarshal(arguments, &newOwner); err != nil {
			return nil, err
		}
		return []interface{}{newOwner}, nil
	case "setApprovalForAll":
		var params struct {
			Operator common.Address `json:"operator"`
			Approved bool           `json:"approved"`
		}
		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		return []interface{}{params.Operator, params.Approved}, nil
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

		revenueShares := make([][]feralfilev4.IFeralfileSaleDataRevenueShare, 0)
		for _, v := range params.SaleData.RevenueShares {
			revenueShare := make([]feralfilev4.IFeralfileSaleDataRevenueShare, 0)
			for _, vv := range v {
				bps := vv.Bps.Int
				revenueShare = append(revenueShare, feralfilev4.IFeralfileSaleDataRevenueShare{
					Recipient: vv.Recipient,
					Bps:       &bps,
				})
			}
			revenueShares = append(revenueShares, revenueShare)
		}

		saleData := feralfilev4.IFeralfileSaleDataSaleData{
			Price:              &params.SaleData.Price.Int,
			Cost:               &params.SaleData.Cost.Int,
			ExpiryTime:         &params.SaleData.ExpiryTime.Int,
			Destination:        params.SaleData.Destination,
			TokenIds:           tokenIDs,
			RevenueShares:      revenueShares,
			PayByVaultContract: params.SaleData.PayByVaultContract,
		}

		return []interface{}{r32Val, s32Val, uint8(vVal), saleData}, nil
	case "buyBulkArtworks":
		var params struct {
			SaleData struct {
				Price         ethereum.BigInt `json:"price"`
				Cost          ethereum.BigInt `json:"cost"`
				ExpiryTime    ethereum.BigInt `json:"expiryTime"`
				Destination   common.Address  `json:"destination"`
				Nonce         ethereum.BigInt `json:"nonce"`
				SeriesID      ethereum.BigInt `json:"seriesId"`
				Quantity      uint16          `json:"quantity"`
				RevenueShares []struct {
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

		revenueShares := make([]feralfilev4.IFeralfileSaleDataRevenueShare, 0)
		for _, v := range params.SaleData.RevenueShares {
			bps := v.Bps.Int
			revenueShares = append(revenueShares, feralfilev4.IFeralfileSaleDataRevenueShare{
				Recipient: v.Recipient,
				Bps:       &bps,
			})
		}

		saleData := feralfilev4.IFeralfileSaleDataV2SaleDataV2{
			Price:              &params.SaleData.Price.Int,
			Cost:               &params.SaleData.Cost.Int,
			ExpiryTime:         &params.SaleData.ExpiryTime.Int,
			Destination:        params.SaleData.Destination,
			Nonce:              &params.SaleData.Nonce.Int,
			SeriesID:           &params.SaleData.SeriesID.Int,
			Quantity:           params.SaleData.Quantity,
			RevenueShares:      revenueShares,
			PayByVaultContract: params.SaleData.PayByVaultContract,
		}

		return []interface{}{r32Val, s32Val, uint8(vVal), saleData}, nil
	case "setAdvanceSetting":
		var params struct {
			AdvanceAddresses []common.Address  `json:"advance_addresses"`
			AdvanceAmounts   []ethereum.BigInt `json:"advance_amounts"`
		}
		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		advanceAmounts := make([]*big.Int, len(params.AdvanceAmounts))
		for i, v := range params.AdvanceAmounts {
			amount := v.Int // closure issue
			advanceAmounts[i] = &amount
		}

		return []interface{}{params.AdvanceAddresses, advanceAmounts}, nil
	default:
		return nil, fmt.Errorf("unsupported method")
	}
}

func init() {
	ethereum.RegisterContract("FeralfileExhibitionV4_2", FeralfileExhibitionV4_2ContractFactory)
}
