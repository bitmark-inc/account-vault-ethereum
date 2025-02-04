package feralfilev2

import (
	"encoding/json"
	"fmt"
	"math/big"

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
	gasLimit uint64,
	gasPrice *int64,
	nonce *uint64) (*types.Transaction, error) {
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
	case "createArtwork":
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
	case "swapArtworkFromBitmark":
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

		return contract.SwapArtworkFromBitmark(
			t,
			artworkID,
			bitmarkID,
			editionNumber,
			to,
			ipfsCID)
	case "safeTransferFrom":
		if len(params) != 3 {
			return nil, fmt.Errorf("invalid params")
		}

		from, ok := params[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid from params")
		}

		to, ok := params[1].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid to params")
		}

		tokenID, ok := params[2].(*big.Int)
		if !ok {
			return nil, fmt.Errorf("invalid token id params")
		}

		return contract.SafeTransferFrom(
			t,
			from,
			to,
			tokenID)
	case "setApprovalForAll":
		if len(params) != 2 {
			return nil, fmt.Errorf("invalid params")
		}

		operator, ok := params[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid operator params")
		}

		approved, ok := params[1].(bool)
		if !ok {
			return nil, fmt.Errorf("invalid approved params")
		}

		return contract.SetApprovalForAll(t, operator, approved)
	case "updateArtworkEditionIPFSCid":
		if len(params) != 2 {
			return nil, fmt.Errorf("invalid params")
		}

		tokenID, ok := params[0].(*big.Int)
		if !ok {
			return nil, fmt.Errorf("invalid token id params")
		}

		ipfsCID, ok := params[1].(string)
		if !ok {
			return nil, fmt.Errorf("invalid ipfs cid params")
		}

		return contract.UpdateArtworkEditionIPFSCid(t, tokenID, ipfsCID)
	default:
		return nil, fmt.Errorf("unsupported method")
	}
}

func (c *FeralfileExhibitionV2Contract) Pack(
	method string,
	arguments json.RawMessage) ([]byte, error) {
	abi, err := feralfilev2.FeralfileExhibitionV2MetaData.GetAbi()
	if nil != err {
		return nil, err
	}

	parsedArgs, err := c.Parse(method, arguments)
	if nil != err {
		return nil, err
	}

	return abi.Pack(method, parsedArgs...)
}

func (c *FeralfileExhibitionV2Contract) Parse(
	method string,
	arguments json.RawMessage) ([]interface{}, error) {
	switch method {
	case "createArtwork":
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
	case "swapArtworkFromBitmark":
		var params struct {
			ArtworkID     ethereum.BigInt `json:"artwork_id"`
			BitmarkID     ethereum.BigInt `json:"bitmark_id"`
			EditionNumber ethereum.BigInt `json:"edition_number"`
			To            common.Address  `json:"to"`
			IPFSCID       string          `json:"ipfs_cid"`
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
	case "safeTransferFrom":
		var params struct {
			From    common.Address  `json:"from"`
			To      common.Address  `json:"to"`
			TokenID ethereum.BigInt `json:"token_id"`
		}
		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		return []interface{}{params.From, params.To, &params.TokenID.Int}, nil

	case "setApprovalForAll":
		var params struct {
			Operator common.Address `json:"operator"`
			Approved bool           `json:"approved"`
		}
		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		return []interface{}{params.Operator, params.Approved}, nil
	case "updateArtworkEditionIPFSCid":
		var params struct {
			TokenID string `json:"token_id"`
			IPFSCID string `json:"ipfs_cid"`
		}
		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		tokenIDInt, ok := new(big.Int).SetString(params.TokenID, 10)
		if !ok {
			return nil, fmt.Errorf("invalid token id")
		}
		return []interface{}{tokenIDInt, params.IPFSCID}, nil
	default:
		return nil, fmt.Errorf("unsupported method")
	}
}

func init() {
	ethereum.RegisterContract("FeralfileExhibitionV2", FeralfileExhibitionV2ContractFactory)
}
