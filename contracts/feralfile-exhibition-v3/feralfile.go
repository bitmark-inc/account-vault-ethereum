package feralfilev3

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"unsafe"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	ethereum "github.com/bitmark-inc/account-vault-ethereum"
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
func (c *FeralfileExhibitionV3Contract) Call(wallet *ethereum.Wallet, method, fund string, arguments json.RawMessage, noSend bool) (*types.Transaction, error) {
	contract, err := NewFeralfileExhibitionV3(common.HexToAddress(c.contractAddress), wallet.RPCClient())
	if err != nil {
		return nil, err
	}

	t, err := wallet.Transactor()
	if err != nil {
		return nil, err
	}

	t.NoSend = noSend

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
	case "batch_mint":
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
			value := v
			mintParams = append(mintParams, FeralfileExhibitionV3MintArtworkParam{
				ArtworkID: &value.ArtworkID.Int,
				Edition:   &value.EditionNumber.Int,
				Artist:    value.Artist,
				Owner:     value.Owner,
				IpfsCID:   value.IPFSCID,
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

		t.GasLimit = uint64(400000 * len(params))

		transferParams := make([]FeralfileExhibitionV3TransferArtworkParam, 0)

		for _, v := range params {
			value := v // closure issue in &
			rVal := []byte(value.R)
			sVal := []byte(value.S)
			vVal, _ := strconv.ParseUint(v.V, 10, 8)
			transferParams = append(transferParams, FeralfileExhibitionV3TransferArtworkParam{
				From:      value.From,
				To:        value.To,
				TokenID:   &value.TokenID.Int,
				Timestamp: &value.Timestamp.Int,
				R:         *(*[32]byte)(unsafe.Pointer(&rVal)),
				S:         *(*[32]byte)(unsafe.Pointer(&sVal)),
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
