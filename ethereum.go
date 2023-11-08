package ethereum

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
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
	chainID := params.SepoliaChainConfig.ChainID
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
	chainID := params.SepoliaChainConfig.ChainID
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
func (w *Wallet) SignABIParameters(ctx context.Context, types []interface{}, arguments ...interface{}) (string, string, string, error) {
	abiTypes, abiTypesArguments, parsedValue, err := parseABIParams(types, arguments)
	if err != nil {
		return "", "", "", err
	}

	args := abi.Arguments{}

	for i, t := range abiTypes {
		var am []abi.ArgumentMarshaling
		if abiTypesArguments[i] != nil {
			am = *abiTypesArguments[i]
		}
		ty, err := abi.NewType(t, "", am)
		if err != nil {
			return "", "", "", err
		}

		args = append(args, abi.Argument{Type: ty})
	}

	argBytes, err := args.Pack(parsedValue...)
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

// SignETHTypedDataV4 sign packed function parameters and returns signature
func (w *Wallet) SignETHTypedDataV4(ctx context.Context, domain map[string]interface{}, types []interface{}, arguments ...interface{}) (string, error) {
	dMessage := []interface{}{
		domain["name"],
		domain["version"],
		domain["chainId"],
		domain["verifyingContract"],
	}

	dArgs := abi.Arguments{}
	dArgBytes, err := dArgs.Pack(dMessage...)
	if err != nil {
		return "", err
	}
	dh := sha3.NewLegacyKeccak256()
	dh.Write(dArgBytes)
	dhash := dh.Sum(nil)

	dh1 := sha3.NewLegacyKeccak256()
	dh1.Write([]byte(domain["contractName"].(string)))
	dhash1 := dh1.Sum(nil)

	dh2 := sha3.NewLegacyKeccak256()
	dh2.Write([]byte("1"))
	dhash2 := dh2.Sum(nil)

	domainMessage := append(dhash, dhash1...)
	domainMessage = append(domainMessage, dhash2...)

	dh3 := sha3.NewLegacyKeccak256()
	dh3.Write(domainMessage)
	dhash3 := dh3.Sum(nil)

	abiTypes, abiTypesArguments, parsedValue, err := parseABIParams(types, arguments)
	if err != nil {
		return "", err
	}

	args := abi.Arguments{}

	for i, t := range abiTypes {
		var am []abi.ArgumentMarshaling
		if abiTypesArguments[i] != nil {
			am = *abiTypesArguments[i]
		}
		ty, err := abi.NewType(t, "", am)
		if err != nil {
			return "", err
		}

		args = append(args, abi.Argument{Type: ty})
	}

	argBytes, err := args.Pack(parsedValue...)
	if err != nil {
		return "", err
	}

	h := sha3.NewLegacyKeccak256()
	h.Write(argBytes)
	hash := h.Sum(nil)

	validationMessage := append([]byte("\x19\x01"), dhash3...)
	validationMessage = append(validationMessage, hash...)

	h2 := sha3.NewLegacyKeccak256()
	h2.Write(validationMessage)
	hash2 := h2.Sum(nil)

	privateKey, err := w.wallet.PrivateKey(w.account)
	if err != nil {
		return "", err
	}
	signature, err := crypto.Sign(hash2, privateKey)
	if err != nil {
		return "", err
	}

	return string(signature), nil
}

// TransferETH performs regular ethereum transferring
func (w *Wallet) TransferETH(ctx context.Context, to string, amount string, customizeGasPriceInWei *int64, customizedNonce *uint64) (string, error) {
	account := w.account

	var nonce uint64
	if customizedNonce == nil {
		n, err := w.rpcClient.PendingNonceAt(ctx, account.Address)
		if err != nil {
			return "", err
		}
		nonce = n
	} else {
		nonce = *customizedNonce
	}

	value, ok := big.NewInt(0).SetString(amount, 0)
	if !ok {
		return "", fmt.Errorf("can not set amount")
	}
	fromAddress := account.Address
	toAddress := common.HexToAddress(to)

	var gasPrice *big.Int
	if customizeGasPriceInWei != nil && *customizeGasPriceInWei != 0 {
		gasPrice = big.NewInt(*customizeGasPriceInWei * params.Wei)
	} else {
		var err error
		gasPrice, err = w.rpcClient.SuggestGasPrice(ctx)
		if err != nil {
			return "", err
		}
	}

	egl, err := w.rpcClient.EstimateGas(ctx, ethereum.CallMsg{
		From:  fromAddress,
		To:    &toAddress,
		Value: value,
	})
	if err != nil {
		return "", err
	}

	tx := types.NewTransaction(nonce, toAddress, value, egl, gasPrice, nil)

	signedTx, err := w.wallet.SignTx(account, tx, nil)
	if err != nil {
		return "", err
	}

	if err := w.rpcClient.SendTransaction(ctx, signedTx); err != nil {
		return "", err
	}

	return signedTx.Hash().String(), nil
}

// GetEthereumBalance return ethereum balance at address
func (w *Wallet) GetEthereumBalance(ctx context.Context) (string, error) {
	balance, err := w.rpcClient.BalanceAt(ctx, w.account.Address, nil)
	if err != nil {
		return "", err
	}
	return balance.String(), nil
}

// GetERC20Balance return ERC20 Token balance at address
func (w *Wallet) GetERC20Balance(ctx context.Context, tokenAddr string) (string, error) {
	if tokenAddr == "" {
		return "", errors.New("invalid address")
	}

	erc20TokenContract, err := NewERC20Token(common.HexToAddress(tokenAddr), w.RPCClient())
	if err != nil {
		return "", err
	}

	balance, err := erc20TokenContract.BalanceOf(&bind.CallOpts{}, w.account.Address)
	if err != nil {
		return "", err
	}
	return balance.String(), nil
}

// GetERC20Allowance return ERC20 token allowance amount at address
func (w *Wallet) GetERC20Allowance(ctx context.Context, tokenAddr, spenderAddr string) (string, error) {
	if tokenAddr == "" {
		return "", errors.New("invalid address")
	}

	erc20TokenContract, err := NewERC20Token(common.HexToAddress(tokenAddr), w.RPCClient())
	if err != nil {
		return "", err
	}

	balance, err := erc20TokenContract.Allowance(&bind.CallOpts{}, w.account.Address, common.HexToAddress(spenderAddr))
	if err != nil {
		return "", err
	}
	return balance.String(), nil
}

// ApproveERC20Token performs approval ERC20 token for spender address
func (w *Wallet) ApproveERC20Token(ctx context.Context, tokenAddr, spender, amount string, customizeGasPriceInWei *int64, customizedNonce *uint64) (*types.Transaction, error) {
	if spender == "" || amount == "" {
		return nil, errors.New("invalid parameter")
	}

	erc20TokenContract, err := NewERC20Token(common.HexToAddress(tokenAddr), w.RPCClient())
	if err != nil {
		return nil, err
	}

	t, err := w.Transactor()
	if err != nil {
		return nil, err
	}

	if customizeGasPriceInWei != nil && *customizeGasPriceInWei != 0 {
		t.GasPrice = big.NewInt(*customizeGasPriceInWei * params.Wei)
	}

	if customizedNonce != nil {
		t.Nonce = big.NewInt(int64(*customizedNonce))
	}

	value, ok := big.NewInt(0).SetString(amount, 0)
	if !ok {
		return nil, fmt.Errorf("can not set amount")
	}

	return erc20TokenContract.Approve(t, common.HexToAddress(spender), value)
}

func (w *Wallet) Transactor() (*bind.TransactOpts, error) {
	k, err := w.wallet.PrivateKey(w.account)
	if err != nil {
		return nil, err
	}

	return bind.NewKeyedTransactorWithChainID(k, w.chainID)
}

// TransferBatchToken performs regular ERC20 Token transferring
func (w *Wallet) TransferBatchToken(ctx context.Context, contractAddress string, arguments json.RawMessage, noSend bool, customizeGasPriceInWei *int64, customizedNonce *uint64) (*types.Transaction, error) {
	transferContract, err := NewTokenBatchTransfer(common.HexToAddress(contractAddress), w.RPCClient())
	if err != nil {
		return nil, err
	}

	t, err := w.Transactor()
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

	var param struct {
		TokenOwner string `json:"token_owner"`
		Recipients []struct {
			Address string `json:"address"`
			Amount  uint64 `json:"amount"`
		} `json:"recipients"`
	}

	if err := json.Unmarshal(arguments, &param); err != nil {
		return nil, err
	}

	emptyOpts := &bind.CallOpts{}

	// Checking allowance
	tokenAddr, err := transferContract.Token(emptyOpts)
	if err != nil {
		return nil, err
	}

	tokenContract, err := NewERC20Token(tokenAddr, w.RPCClient())
	if err != nil {
		return nil, err
	}

	balance, err := tokenContract.BalanceOf(emptyOpts, common.HexToAddress(param.TokenOwner))
	if err != nil {
		return nil, err
	}

	allowance, err := tokenContract.Allowance(emptyOpts, common.HexToAddress(param.TokenOwner), common.HexToAddress(contractAddress))
	if err != nil {
		return nil, err
	}

	amountParams := make([]*big.Int, 0)
	tokenHolderParams := make([]common.Address, 0)
	totalSendingAmount := uint64(0)
	for _, v := range param.Recipients {
		if v.Amount == 0 || v.Address == "" {
			return nil, errors.New("empty amount or empty address parameter")
		}
		tokenHolderParams = append(tokenHolderParams, common.HexToAddress(v.Address))
		amountParams = append(amountParams, big.NewInt(int64(v.Amount)))
		totalSendingAmount += v.Amount
	}

	if balance.Uint64() < totalSendingAmount {
		return nil, errors.New("token balance does not enough for transfer")
	}

	if allowance.Uint64() < totalSendingAmount {
		return nil, errors.New("token allowance does not enough for transfer")
	}

	tx, err := transferContract.BatchTransferFrom(t, common.HexToAddress(param.TokenOwner), tokenHolderParams, amountParams)
	if err != nil {
		return nil, err
	}
	return tx, err
}
