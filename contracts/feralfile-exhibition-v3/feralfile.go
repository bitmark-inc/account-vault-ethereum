package feralfilev3

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
	feralfilev3 "github.com/bitmark-inc/feralfile-exhibition-smart-contract/go-binding/feralfile-exhibition-v3"
)

type FeralfileExhibitionV3Contract struct {
	contractAddress string
}

func FeralfileExhibitionV3ContractFactory(contractAddress string) ethereum.Contract {
	return &FeralfileExhibitionV3Contract{
		contractAddress: contractAddress,
	}
}

// Deploy deploys the smart contract to ethereum blockchain
func (c *FeralfileExhibitionV3Contract) Deploy(
	wallet *ethereum.Wallet,
	arguments json.RawMessage) (string, string, error) {
	var params struct {
		Name                 string          `json:"name"`
		Symbol               string          `json:"symbol"`
		VersionCode          string          `json:"version_code"`
		RoyaltyBPS           ethereum.BigInt `json:"royalty_bps"`
		RoyaltyPayoutAddress common.Address  `json:"royalty_payout_address"`
		ContractURI          string          `json:"contract_uri"`
		TokenBaseURI         string          `json:"token_base_uri"`
		IsBurnable           bool            `json:"is_burnable"`
		IsBridgeable         bool            `json:"is_bridgeable"`
	}

	if err := json.Unmarshal(arguments, &params); err != nil {
		return "", "", err
	}

	t, err := wallet.Transactor()
	if err != nil {
		return "", "", err
	}

	address, tx, _, err := feralfilev3.DeployFeralfileExhibitionV3(t, wallet.RPCClient(),
		params.Name, params.Symbol, &params.RoyaltyBPS.Int, params.RoyaltyPayoutAddress,
		params.ContractURI, params.TokenBaseURI, params.IsBurnable, params.IsBridgeable)
	if err != nil {
		return "", "", err
	}
	return address.String(), tx.Hash().String(), nil
}

// Call is the entry function for account vault to interact with a smart contract.
func (c *FeralfileExhibitionV3Contract) Call(
	wallet *ethereum.Wallet,
	method,
	fund string,
	arguments json.RawMessage,
	noSend bool,
	gasLimit uint64,
	gasPrice *int64,
	nonce *uint64) (*types.Transaction, error) {
	contractAddr := common.HexToAddress(c.contractAddress)
	contract, err := feralfilev3.NewFeralfileExhibitionV3(contractAddr, wallet.RPCClient())
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
	case "createArtworks":
		if len(params) != 1 {
			return nil, fmt.Errorf("invalid params")
		}

		artworkParams, ok := params[0].([]feralfilev3.FeralfileExhibitionV3Artwork)
		if !ok {
			return nil, fmt.Errorf("invalid artwork params")
		}

		return contract.CreateArtworks(t, artworkParams)
	case "batchMint":
		if len(params) != 1 {
			return nil, fmt.Errorf("invalid params")
		}

		mintParams, ok := params[0].([]feralfilev3.FeralfileExhibitionV3MintArtworkParam)
		if !ok {
			return nil, fmt.Errorf("invalid mint params")
		}

		return contract.BatchMint(t, mintParams)
	case "authorizedTransfer":
		if len(params) != 1 {
			return nil, fmt.Errorf("invalid params")
		}

		transferParams, ok := params[0].([]feralfilev3.FeralfileExhibitionV3TransferArtworkParam)
		if !ok {
			return nil, fmt.Errorf("invalid transfer params")
		}

		return contract.AuthorizedTransfer(t, transferParams)
	case "burnEditions":
		if len(params) != 1 {
			return nil, errors.New("Invalid parameters")
		}

		burnParams, ok := params[0].([]*big.Int)
		if !ok {
			return nil, errors.New("Invalid token burn parameters")
		}

		return contract.BurnEditions(t, burnParams)
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

		tokenID, ok := params[1].(*big.Int)
		if !ok {
			return nil, fmt.Errorf("invalid token id params")
		}

		return contract.SafeTransferFrom(t,
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
	default:
		return nil, fmt.Errorf("unsupported method")
	}
}

func (c *FeralfileExhibitionV3Contract) Pack(
	method string,
	arguments json.RawMessage) ([]byte, error) {
	abi, err := feralfilev3.FeralfileExhibitionV3MetaData.GetAbi()
	if nil != err {
		return nil, err
	}

	parsedArgs, err := c.Parse(method, arguments)
	if nil != err {
		return nil, err
	}

	return abi.Pack(method, parsedArgs...)
}

func (c *FeralfileExhibitionV3Contract) Parse(
	method string,
	arguments json.RawMessage) ([]interface{}, error) {
	switch method {
	case "createArtworks":
		var params []struct {
			Fingerprint string `json:"fingerprint"`
			Title       string `json:"title"`
			ArtistName  string `json:"artist_name"`
			EditionSize int64  `json:"edition_size"`
			AEAmount    int64  `json:"ae_amount"`
			PPAmount    int64  `json:"pp_amount"`
		}

		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		artworkParams := make([]feralfilev3.FeralfileExhibitionV3Artwork, 0)
		for _, v := range params {
			artworkParams = append(artworkParams, feralfilev3.FeralfileExhibitionV3Artwork{
				Title:       v.Title,
				ArtistName:  v.ArtistName,
				Fingerprint: v.Fingerprint,
				EditionSize: big.NewInt(v.EditionSize),
				AEAmount:    big.NewInt(v.AEAmount),
				PPAmount:    big.NewInt(v.PPAmount),
			})
		}

		return []interface{}{artworkParams}, nil
	case "batchMint":
		var params []struct {
			ArtworkID     ethereum.BigInt `json:"artwork_id"`
			EditionNumber ethereum.BigInt `json:"edition_number"`
			Artist        common.Address  `json:"artist"`
			Owner         common.Address  `json:"owner"`
			IPFSCID       string          `json:"ipfs_cid"`
		}

		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		mintParams := make([]feralfilev3.FeralfileExhibitionV3MintArtworkParam, 0)

		for _, v := range params {
			artworkID := v.ArtworkID.Int
			editionNumber := v.EditionNumber.Int

			mintParams = append(mintParams, feralfilev3.FeralfileExhibitionV3MintArtworkParam{
				ArtworkID: &artworkID,
				Edition:   &editionNumber,
				Artist:    v.Artist,
				Owner:     v.Owner,
				IpfsCID:   v.IPFSCID,
			})
		}

		return []interface{}{mintParams}, nil
	case "authorizedTransfer":
		var params []struct {
			From      common.Address  `json:"from"`
			To        common.Address  `json:"to"`
			TokenID   ethereum.BigInt `json:"token_id"`
			Timestamp ethereum.BigInt `json:"timestamp"`
			R         string          `json:"r"`
			S         string          `json:"s"`
			V         string          `json:"v"`
		}

		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		transferParams := make([]feralfilev3.FeralfileExhibitionV3TransferArtworkParam, 0)

		for _, v := range params {
			tokenID := v.TokenID.Int
			expireTime := v.Timestamp.Int
			rVal, err := hex.DecodeString(strings.Replace(v.R, "0x", "", -1))
			if err != nil {
				return nil, err
			}
			sVal, err := hex.DecodeString(strings.Replace(v.S, "0x", "", -1))
			if err != nil {
				return nil, err
			}
			vVal, err := strconv.ParseUint(strings.Replace(v.V, "0x", "", -1), 16, 64)
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

			transferParams = append(transferParams, feralfilev3.FeralfileExhibitionV3TransferArtworkParam{
				From:       v.From,
				To:         v.To,
				TokenID:    &tokenID,
				ExpireTime: &expireTime,
				R:          r32Val,
				S:          s32Val,
				V:          uint8(vVal),
			})
		}

		return []interface{}{transferParams}, nil
	case "burnEditions":
		var params []ethereum.BigInt
		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		burnParams := make([]*big.Int, 0)
		for _, v := range params {
			tokenID := v.Int
			burnParams = append(burnParams, &tokenID)
		}

		return []interface{}{burnParams}, nil
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
	default:
		return nil, fmt.Errorf("unsupported method")
	}
}

func init() {
	ethereum.RegisterContract("FeralfileExhibitionV3", FeralfileExhibitionV3ContractFactory)
}
