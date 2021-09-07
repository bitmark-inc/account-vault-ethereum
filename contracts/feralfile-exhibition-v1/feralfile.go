package feralfile

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

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
		&params.MaxEdition.Int, &params.BasePrice.Int, params.TokenBaseURI)
	if err != nil {
		return "", "", err
	}
	return address.String(), tx.Hash().String(), nil
}

// Call is the entry function for account vault to interact with a smart contract.
func (c *FeralfileExhibitionV1Contract) Call(wallet *ethereum.Wallet, method, fund string, arguments json.RawMessage) (string, error) {
	contract, err := NewFeralfileExhibition(common.HexToAddress(c.contractAddress), wallet.RPCClient())
	if err != nil {
		return "", err
	}

	t, err := wallet.Transactor()
	if err != nil {
		return "", err
	}

	switch method {
	case "create_artwork":
		var params struct {
			Fingerprint  ethereum.BigInt `json:"fingerprint"`
			Title        string          `json:"title"`
			Medium       string          `json:"medium"`
			EditionSize  int64           `json:"edition_size"`
			InitialPrice int64           `json:"initial_price"`
		}

		if err := json.Unmarshal(arguments, &params); err != nil {
			return "", err
		}

		fingerprintBytes := params.Fingerprint.Bytes()
		if len(fingerprintBytes) > 32 {
			return "", fmt.Errorf("the fingerprint can not be more than 32 bytes")
		}

		var fingerprintArray [32]byte
		if n := copy(fingerprintArray[:], fingerprintBytes); n != len(fingerprintBytes) {
			return "", fmt.Errorf("incorrect size of fingerprint copied")
		}

		tx, err := contract.CreateArtwork(t, fingerprintArray,
			params.Title, t.From,
			params.Medium, "",
			big.NewInt(params.EditionSize), big.NewInt(params.InitialPrice))
		if err != nil {
			return "", err
		}
		return tx.Hash().String(), err
	case "mint_artwork":
		var params struct {
			ArtworkID ethereum.BigInt `json:"artwork_id"`
			IPFSCID   []string        `json:"ipfs_cids"`
		}
		if err := json.Unmarshal(arguments, &params); err != nil {
			return "", err
		}

		tx, err := contract.MintArtwork(t, &params.ArtworkID.Int, params.IPFSCID)
		if err != nil {
			return "", err
		}
		return tx.Hash().String(), nil
	case "swap_artwork_from_bitmark":
		var params struct {
			ArtworkID     ethereum.BigInt `json:"artwork_id"`
			BitmarkID     ethereum.BigInt `json:"bitmark_id"`
			EditionNumber ethereum.BigInt `json:"edition_number"`
			To            common.Address  `json:"to"`
			IPFSCID       string          `json:"ipfs_cid"`
		}
		if err := json.Unmarshal(arguments, &params); err != nil {
			return "", err
		}

		tx, err := contract.SwapArtworkFromBitmarks(t,
			&params.ArtworkID.Int, &params.BitmarkID.Int, &params.EditionNumber.Int,
			params.To, params.IPFSCID)
		if err != nil {
			return "", err
		}
		return tx.Hash().String(), nil
	default:
		return "", fmt.Errorf("unsupported method")
	}
}

func init() {
	ethereum.RegisterContract("FeralfileExhibitionV1", FeralfileExhibitionV1ContractFactory)
}
