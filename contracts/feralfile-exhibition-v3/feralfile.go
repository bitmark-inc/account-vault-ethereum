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
)

const (
	GasLimitForAuthTransfer = 130000
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
func (c *FeralfileExhibitionV3Contract) Deploy(wallet *ethereum.Wallet, arguments json.RawMessage) (string, string, error) {
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

	address, tx, _, err := DeployFeralfileExhibitionV3(t, wallet.RPCClient(),
		params.Name, params.Symbol, &params.MaxEdition.Int,
		&params.RoyaltyBPS.Int, params.RoyaltyPayoutAddress,
		params.ContractURI, params.TokenBaseURI)
	if err != nil {
		return "", "", err
	}
	return address.String(), tx.Hash().String(), nil
}

// Call is the entry function for account vault to interact with a smart contract.
func (c *FeralfileExhibitionV3Contract) Call(wallet *ethereum.Wallet, method, fund string, arguments json.RawMessage, noSend bool, customizeGasPriceInWei *int64, customizedNonce *uint64) (*types.Transaction, error) {
	contract, err := NewFeralfileExhibitionV3(common.HexToAddress(c.contractAddress), wallet.RPCClient())
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
	case "register_artworks":
		var params []struct {
			Fingerprint string `json:"fingerprint"`
			Title       string `json:"title"`
			ArtistName  string `json:"artist_name"`
			EditionSize int64  `json:"edition_size"`
		}

		if err := json.Unmarshal(arguments, &params); err != nil {
			return nil, err
		}

		artworkParams := make([]FeralfileExhibitionV3Artwork, 0)

		for _, v := range params {
			artworkParams = append(artworkParams, FeralfileExhibitionV3Artwork{
				Title:       v.Title,
				ArtistName:  v.ArtistName,
				Fingerprint: v.Fingerprint,
				EditionSize: big.NewInt(v.EditionSize),
			})
		}

		tx, err := contract.CreateArtworks(t, artworkParams)
		if err != nil {
			return nil, err
		}
		return tx, err
	case "mint_editions":
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

		t.GasLimit = uint64(400000 * len(params))

		mintParams := make([]FeralfileExhibitionV3MintArtworkParam, 0)

		for _, v := range params {
			artworkID := v.ArtworkID.Int
			editionNumber := v.EditionNumber.Int

			mintParams = append(mintParams, FeralfileExhibitionV3MintArtworkParam{
				ArtworkID: &artworkID,
				Edition:   &editionNumber,
				Artist:    v.Artist,
				Owner:     v.Owner,
				IpfsCID:   v.IPFSCID,
			})
		}

		tx, err := contract.BatchMint(t, mintParams)
		if err != nil {
			return nil, err
		}
		return tx, nil
	case "authorized_transfer":
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

		t.GasLimit = uint64(GasLimitForAuthTransfer * len(params))

		transferParams := make([]FeralfileExhibitionV3TransferArtworkParam, 0)

		for _, v := range params {
			tokenID := v.TokenID.Int
			timestamp := v.Timestamp.Int
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

			transferParams = append(transferParams, FeralfileExhibitionV3TransferArtworkParam{
				From:      v.From,
				To:        v.To,
				TokenID:   &tokenID,
				Timestamp: &timestamp,
				R:         r32Val,
				S:         s32Val,
				V:         uint8(vVal),
			})
		}

		tx, err := contract.AuthorizedTransfer(t, transferParams)
		if err != nil {
			return nil, err
		}
		return tx, nil
	default:
		return nil, fmt.Errorf("unsupported method")
	}
}

func init() {
	ethereum.RegisterContract("FeralfileExhibitionV3", FeralfileExhibitionV3ContractFactory)
}
