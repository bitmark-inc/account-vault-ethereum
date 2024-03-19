package feralfilev2

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	goEthereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"

	ethereum "github.com/bitmark-inc/account-vault-ethereum"
	feralfilev2 "github.com/bitmark-inc/feralfile-exhibition-smart-contract/go-binding/feralfile-exhibition-v2"
)

type FeralfileExhibitionV2Contract struct {
	contractAddress string
}

func FeralfileExhibitionV2ContractFactory(contractAddress string) ethereum.Contract {
	return &FeralfileExhibitionV2Contract{
		contractAddress: contractAddress,
	}
}

// Deploy deploys the smart contract to ethereum blockchain
func (c *FeralfileExhibitionV2Contract) Deploy(
	wallet *ethereum.Wallet,
	arguments json.RawMessage) (string, string, error) {
	var params struct {
		Name                 string          `json:"Name"`
		Symbol               string          `json:"symbol"`
		ExhibitionTitle      string          `json:"exhibition_title"`
		MaxEdition           ethereum.BigInt `json:"max_edition"`
		RoyaltyBPS           ethereum.BigInt `json:"royalty_bps"`
		RoyaltyPayoutAddress common.Address  `json:"royalty_payout_address"`
		ContractURI          string          `json:"contract_uri"`
		TokenBaseURI         string          `json:"token_base_uri"`
	}

	if err := json.Unmarshal(arguments, &params); err != nil {
		return "", "", err
	}

	t, err := wallet.Transactor()
	if err != nil {
		return "", "", err
	}

	address, tx, _, err := feralfilev2.DeployFeralfileExhibitionV2(t, wallet.RPCClient(),
		params.Name, params.Symbol, &params.MaxEdition.Int,
		&params.RoyaltyBPS.Int, params.RoyaltyPayoutAddress,
		params.ContractURI, params.TokenBaseURI)
	if err != nil {
		return "", "", err
	}
	return address.String(), tx.Hash().String(), nil
}

// Call is the entry function for account vault to interact with a smart contract.
func (c *FeralfileExhibitionV2Contract) Call(
	wallet *ethereum.Wallet,
	method,
	fund string,
	arguments json.RawMessage,
	noSend bool,
	customizeGasPriceInWei *int64,
	customizedNonce *uint64) (*types.Transaction, error) {
	contractAddr := common.HexToAddress(c.contractAddress)
	contract, err := feralfilev2.NewFeralfileExhibitionV2(contractAddr, wallet.RPCClient())
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
	case "create_artwork":
		if len(params) != 4 {
			return nil, fmt.Errorf("invalid params")
		}

		fingerprint, ok := params[0].(string)
		if !ok {
			return nil, fmt.Errorf("invalid fingerprint params")
		}

		title, ok := params[1].(string)
		if !ok {
			return nil, fmt.Errorf("invalid title params")
		}

		artistName, ok := params[2].(string)
		if !ok {
			return nil, fmt.Errorf("invalid artist name params")
		}

		editionSize, ok := params[3].(*big.Int)
		if !ok {
			return nil, fmt.Errorf("invalid edition size params")
		}

		return contract.CreateArtwork(
			t,
			fingerprint,
			title,
			artistName,
			editionSize)
	case "swap_artwork_from_bitmark":
		if len(params) != 5 {
			return nil, fmt.Errorf("invalid params")
		}

		artworkID, ok := params[0].(*big.Int)
		if !ok {
			return nil, fmt.Errorf("invalid artwork id params")
		}

		bitmarkID, ok := params[1].(*big.Int)
		if !ok {
			return nil, fmt.Errorf("invalid bitmark id params")
		}

		editionNumber, ok := params[2].(*big.Int)
		if !ok {
			return nil, fmt.Errorf("invalid edition number params")
		}

		to, ok := params[3].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid to params")
		}

		ipfsCID, ok := params[4].(string)
		if !ok {
			return nil, fmt.Errorf("invalid ipfs cid params")
		}

		gasLimit, err := c.EstimateGasLimit(wallet, contractAddr, method, arguments)
		if nil != err {
			return nil, err
		}

		t.GasLimit = gasLimit

		return contract.SwapArtworkFromBitmark(
			t,
			artworkID,
			bitmarkID,
			editionNumber,
			to,
			ipfsCID)
	case "transfer":
		if len(params) != 2 {
			return nil, fmt.Errorf("invalid params")
		}

		to, ok := params[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid to params")
		}

		tokenID, ok := params[1].(*big.Int)
		if !ok {
			return nil, fmt.Errorf("invalid token id params")
		}

		gasLimit, err := c.EstimateGasLimit(wallet, contractAddr, method, arguments)
		if nil != err {
			return nil, err
		}

		t.GasLimit = gasLimit

		return contract.SafeTransferFrom(
			t,
			common.HexToAddress(wallet.Account()),
			to,
			tokenID)
	case "approve_for_all":
		if len(params) != 1 {
			return nil, fmt.Errorf("invalid params")
		}

		operator, ok := params[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid operator params")
		}

		gasLimit, err := c.EstimateGasLimit(wallet, contractAddr, method, arguments)
		if nil != err {
			return nil, err
		}

		t.GasLimit = gasLimit

		return contract.SetApprovalForAll(t, operator, true)
	default:
		return nil, fmt.Errorf("unsupported method")
	}
}

func (c *FeralfileExhibitionV2Contract) Pack(
	wallet *ethereum.Wallet,
	method string,
	arguments json.RawMessage) ([]byte, error) {
	abi, err := feralfilev2.FeralfileExhibitionV2MetaData.GetAbi()
	if nil != err {
		return nil, err
	}

	parsedArgs, err := c.Parse(wallet, method, arguments)
	if nil != err {
		return nil, err
	}

	switch method {
	case "create_artwork":
		method = "createArtwork"
	case "swap_artwork_from_bitmark":
		method = "swapArtworkFromBitmark"
	case "transfer":
		method = "safeTransferFrom"
	case "approve_for_all":
		method = "setApprovalForAll"
	default:
		return nil, fmt.Errorf("unsupported method")
	}

	return abi.Pack(method, parsedArgs...)
}

func (c *FeralfileExhibitionV2Contract) Parse(
	wallet *ethereum.Wallet,
	method string,
	arguments json.RawMessage) ([]interface{}, error) {
	switch method {
	case "create_artwork":
		var params struct {
			Fingerprint string `json:"fingerprint"`
			Title       string `json:"title"`
			ArtistName  string `json:"artist_name"`
			EditionSize int64  `json:"edition_size"`
		}

		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		return []interface{}{
				params.Fingerprint,
				params.Title,
				params.ArtistName,
				big.NewInt(params.EditionSize)},
			nil
	case "swap_artwork_from_bitmark":
		var params struct {
			ArtworkID      ethereum.BigInt `json:"artwork_id"`
			BitmarkID      ethereum.BigInt `json:"bitmark_id"`
			EditionNumber  ethereum.BigInt `json:"edition_number"`
			To             common.Address  `json:"to"`
			PrevProvenance string          `json:"prev_provenance"`
			IPFSCID        string          `json:"ipfs_cid"`
		}
		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		return []interface{}{
			&params.ArtworkID.Int,
			&params.BitmarkID.Int,
			&params.EditionNumber.Int,
			params.To,
			params.IPFSCID}, nil
	case "transfer":
		var params struct {
			To      common.Address  `json:"to"`
			TokenID ethereum.BigInt `json:"token_id"`
		}
		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		return []interface{}{common.HexToAddress(wallet.Account()), params.To, &params.TokenID.Int}, nil

	case "approve_for_all":
		var params struct {
			Operator common.Address `json:"operator"`
		}
		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		return []interface{}{params.Operator, true}, nil
	default:
		return nil, fmt.Errorf("unsupported method")
	}
}

func (c *FeralfileExhibitionV2Contract) EstimateGasLimit(
	wallet *ethereum.Wallet,
	contractAddr common.Address,
	method string,
	arguments json.RawMessage) (uint64, error) {
	data, err := c.Pack(wallet, method, arguments)
	if nil != err {
		return 0, err
	}

	gas, err := wallet.RPCClient().EstimateGas(context.Background(), goEthereum.CallMsg{
		From: common.HexToAddress(wallet.Account()),
		To:   &contractAddr,
		Data: data,
	})

	if nil != err {
		return 0, err
	}

	return gas * 115 / 100, nil // add 15% buffer
}

func init() {
	ethereum.RegisterContract("FeralfileExhibitionV2", FeralfileExhibitionV2ContractFactory)
}
