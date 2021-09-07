package ethereum

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

const DefaultDerivationPath = "m/44'/60'/0'/0/0"

var EthereumRPCURL string

func init() {
	EthereumRPCURL = os.Getenv("ETHEREUM_RPC_URL")
}

type Wallet struct {
	chainID   *big.Int
	wallet    *hdwallet.Wallet
	account   accounts.Account
	rpcClient *ethclient.Client
}

func NewWallet(seed []byte, network string) (*Wallet, error) {
	chainID := params.RinkebyChainConfig.ChainID
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
	rpcClient, err := ethclient.Dial(EthereumRPCURL)
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
