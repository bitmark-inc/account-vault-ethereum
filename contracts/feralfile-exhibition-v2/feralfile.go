package feralfilev2

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"

	ethereum "github.com/bitmark-inc/account-vault-ethereum"
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
func (c *FeralfileExhibitionV2Contract) Deploy(wallet *ethereum.Wallet, arguments json.RawMessage) (string, string, error) {
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

	address, tx, _, err := DeployFeralfileExhibitionV2(t, wallet.RPCClient(),
		params.Name, params.Symbol, &params.MaxEdition.Int,
		&params.RoyaltyBPS.Int, params.RoyaltyPayoutAddress,
		params.ContractURI, params.TokenBaseURI)
	if err != nil {
		return "", "", err
	}
	return address.String(), tx.Hash().String(), nil
}

// Call is the entry function for account vault to interact with a smart contract.
func (c *FeralfileExhibitionV2Contract) Call(wallet *ethereum.Wallet, method, fund string, arguments json.RawMessage, noSend bool, customizeGasPriceInWei *int64) (*types.Transaction, error) {
	contract, err := NewFeralfileExhibitionV2(common.HexToAddress(c.contractAddress), wallet.RPCClient())
	if err != nil {
		return nil, err
	}

	t, err := wallet.Transactor()
	if err != nil {
		return nil, err
	}

	t.NoSend = noSend
	if customizeGasPriceInWei != nil {
		t.GasPrice = big.NewInt(*customizeGasPriceInWei * params.Wei)
	}

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

		tx, err := contract.CreateArtwork(t, params.Fingerprint,
			params.Title, params.ArtistName,
			big.NewInt(params.EditionSize))
		if err != nil {
			return nil, err
		}
		return tx, err
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

		t.GasLimit = 400000

		tx, err := contract.SwapArtworkFromBitmark(t,
			&params.ArtworkID.Int, &params.BitmarkID.Int, &params.EditionNumber.Int,
			params.To, params.IPFSCID)
		if err != nil {
			return nil, err
		}
		return tx, nil
	case "transfer":
		var params struct {
			To      common.Address  `json:"to"`
			TokenID ethereum.BigInt `json:"token_id"`
		}
		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		t.GasLimit = 120000

		tx, err := contract.SafeTransferFrom(t,
			common.HexToAddress(wallet.Account()), params.To, &params.TokenID.Int)
		if err != nil {
			return nil, err
		}
		return tx, nil
	default:
		return nil, fmt.Errorf("unsupported method")
	}
}

func init() {
	ethereum.RegisterContract("FeralfileExhibitionV2", FeralfileExhibitionV2ContractFactory)
}
