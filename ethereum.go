package ethereum

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"golang.org/x/crypto/sha3"
)

const DefaultDerivationPath = "m/44'/60'/0'/0/0"

type Wallet struct {
	chainID   *big.Int
	wallet    *hdwallet.Wallet
	account   accounts.Account
	rpcClient *ethclient.Client
}

// NewWallet creates a wallet from a given seed
func NewWallet(seed []byte, network, rpcURL string) (*Wallet, error) {
	chainID := params.GoerliChainConfig.ChainID
	if network == "livenet" {
		chainID = params.MainnetChainConfig.ChainID
	}

	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		return nil, err
	}

	path := hdwallet.MustParseDerivationPath(DefaultDerivationPath)
	account, err := wallet.Derive(path, true)

	if err != nil {
		return nil, err
	}
	rpcClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	return &Wallet{
		chainID:   chainID,
		wallet:    wallet,
		account:   account,
		rpcClient: rpcClient,
	}, nil
}

// NewWallet creates a wallet from a given mnemonic phrases
func NewWalletFromMnemonic(words, network, rpcURL string) (*Wallet, error) {
	chainID := params.RinkebyChainConfig.ChainID
	if network == "livenet" {
		chainID = params.MainnetChainConfig.ChainID
	}

	wallet, err := hdwallet.NewFromMnemonic(words)
	if err != nil {
		return nil, err
	}

	path := hdwallet.MustParseDerivationPath(DefaultDerivationPath)
	account, err := wallet.Derive(path, true)
	if err != nil {
		return nil, err
	}
	rpcClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	return &Wallet{
		chainID:   chainID,
		wallet:    wallet,
		account:   account,
		rpcClient: rpcClient,
	}, nil

}

// RPCClient returns the RPC client which is bound to the wallet
func (w *Wallet) RPCClient() *ethclient.Client {
	return w.rpcClient
}

// Account returns the ethereum account address
func (w *Wallet) Account() string {
	return w.account.Address.Hex()
}

// SignABIParameters sign packed function parameters and returns (r|v|s) in forms of hex string
func (w *Wallet) SignABIParameters(ctx context.Context, types []string, arguments ...interface{}) (string, string, string, error) {
	args := abi.Arguments{}

	for _, t := range types {
		ty, err := abi.NewType(t, "", nil)
		if err != nil {
			return "", "", "", err
		}

		args = append(args, abi.Argument{Type: ty})
	}

	argBytes, err := args.Pack(arguments...)
	if err != nil {
		return "", "", "", err
	}

	h := sha3.NewLegacyKeccak256()
	h.Write(argBytes)
	hash := h.Sum(nil)

	validationMessage := append([]byte("\x19Ethereum Signed Message:\n32"), hash...)

	h2 := sha3.NewLegacyKeccak256()
	h2.Write(validationMessage)
	hash2 := h2.Sum(nil)

	privateKey, err := w.wallet.PrivateKey(w.account)
	if err != nil {
		return "", "", "", err
	}
	signature, err := crypto.Sign(hash2, privateKey)
	if err != nil {
		return "", "", "", err
	}

	r, s, v :=
		fmt.Sprintf("%#x", signature[0:32]),
		fmt.Sprintf("%#x", signature[32:64]),
		fmt.Sprintf("%#x", signature[64]+27)

	return r, s, v, nil
}

// TransferETH performs regular ethereum transferring
func (w *Wallet) TransferETH(ctx context.Context, to string, amount string, customizeGasPrice int64) (string, error) {
	account := w.account
	nonce, err := w.rpcClient.PendingNonceAt(ctx, account.Address)
	if err != nil {
		return "", err
	}

	value, ok := big.NewInt(0).SetString(amount, 0)
	if !ok {
		return "", fmt.Errorf("can not set amount")
	}
	toAddress := common.HexToAddress(to)

	// fixed gas limit for regular transferring
	gasLimit := uint64(21000)

	var gasPrice *big.Int

	if customizeGasPrice != 0 {
		gasPrice = big.NewInt(customizeGasPrice * params.GWei)
	} else {
		var err error
		gasPrice, err = w.rpcClient.SuggestGasPrice(ctx)
		if err != nil {
			return "", err
		}
	}

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	signedTx, err := w.wallet.SignTx(account, tx, nil)
	if err != nil {
		return "", err
	}

	if err := w.rpcClient.SendTransaction(ctx, signedTx); err != nil {
		return "", err
	}

	return signedTx.Hash().String(), nil
}

func (w *Wallet) Transactor() (*bind.TransactOpts, error) {
	k, err := w.wallet.PrivateKey(w.account)
	if err != nil {
		return nil, err
	}

	return bind.NewKeyedTransactorWithChainID(k, w.chainID)
}
