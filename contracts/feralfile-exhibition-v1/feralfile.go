package feralfile

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"

	ethereum "github.com/bitmark-inc/account-vault-ethereum"
)

type FeralfileExhibitionV1Contract struct {
	contractAddress string
}

func FeralfileExhibitionV1ContractFactory(contractAddress string) ethereum.Contract {
	return &FeralfileExhibitionV1Contract{
		contractAddress: contractAddress,
	}
}

// Deploy deploys the smart contract to ethereum blockchain
func (c *FeralfileExhibitionV1Contract) Deploy(wallet *ethereum.Wallet, arguments json.RawMessage) (string, string, error) {
	var params struct {
		Title        string          `json:"title"`
		Symbol       string          `json:"symbol"`
		Curator      common.Address  `json:"curator"`
		MaxEdition   ethereum.BigInt `json:"max_edition"`
		BasePrice    ethereum.BigInt `json:"base_price"`
		RoyaltyBPS   ethereum.BigInt `json:"royalty_bps"`
		ContractURI  string          `json:"contract_uri"`
		TokenBaseURI string          `json:"token_base_uri"`
	}

	if err := json.Unmarshal(arguments, &params); err != nil {
		return "", "", err
	}

	t, err := wallet.Transactor()
	if err != nil {
		return "", "", err
	}

	address, tx, _, err := DeployFeralfileExhibition(t, wallet.RPCClient(),
		params.Title, params.Symbol, params.Curator,
		&params.MaxEdition.Int, &params.BasePrice.Int, &params.RoyaltyBPS.Int,
		params.ContractURI, params.TokenBaseURI)
	if err != nil {
		return "", "", err
	}
	return address.String(), tx.Hash().String(), nil
}

// Call is the entry function for account vault to interact with a smart contract.
func (c *FeralfileExhibitionV1Contract) Call(wallet *ethereum.Wallet, method, fund string, arguments json.RawMessage, noSend bool, customizeGasPriceInWei *int64) (*types.Transaction, error) {
	contract, err := NewFeralfileExhibition(common.HexToAddress(c.contractAddress), wallet.RPCClient())
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
			Description string `json:"description"`
			Medium      string `json:"medium"`
			EditionSize int64  `json:"edition_size"`
		}

		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		tx, err := contract.CreateArtwork(t, params.Fingerprint,
			params.Title, params.Description, t.From,
			params.Medium,
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

		tx, err := contract.SwapArtworkFromBitmark(t,
			&params.ArtworkID.Int, &params.BitmarkID.Int, &params.EditionNumber.Int,
			params.To, params.PrevProvenance, params.IPFSCID)
		if err != nil {
			return nil, err
		}
		return tx, nil
	default:
		return nil, fmt.Errorf("unsupported method")
	}
}

func init() {
	ethereum.RegisterContract("FeralfileExhibitionV1", FeralfileExhibitionV1ContractFactory)
}
