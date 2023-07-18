# Account Vault Ethereum

The goal of this package is to enable bitmark account vault supporting Ethereum.

This package includes:

- An ethereum wallet
- A contract interface which defines how account vault interact with smart contracts
- A singleton for contract registration

## Wrap a contract

With the contract ABI file, we can generate a contract module through `abigen`. By creating a wrapper structure which implements the `Call` and `Deploy` function, a client can interact with all the different contract in the same interfaces.

The `Call` function determines which method could be exposed and how a client is going to invoke it. Whereas the `Deploy` function is obviously for deploying a new contract.

Here is a simple example of wrapping a contract:

```go
type TestContract struct {
	contractAddress string
}

// Deploy deploys the smart contract to ethereum blockchain
func (c *TestContract) Deploy(wallet *ethereum.Wallet, arguments json.RawMessage) (string, string, error) {
	t, err := wallet.Transactor()
	if err != nil {
		return "", "", err
	}

	address, tx, _, err := DeployTestContract(t, wallet.RPCClient())
	if err != nil {
		return "", "", err
	}
	return address.String(), tx.Hash().String(), nil
}

// Call is the entry function for account vault to interact with a smart contract.
func (c *TestContract) Call(wallet *ethereum.Wallet, method, fund string, arguments json.RawMessage) (string, error) {
	contract, err := NewTestContract(common.HexToAddress(c.contractAddress), wallet.RPCClient())
	if err != nil {
		return "", err
	}

	t, err := wallet.Transactor()
	if err != nil {
		return "", err
	}

	switch method {
	case "set_tester":
		var params struct {
		    Tester common.Address `json:"tester"`
		}
		if err := json.Unmarshal(arguments, &params); err != nil {
			return "", err
		}

		tx, err := contract.CreateArtwork(t, params.Tester)
		if err != nil {
			return "", err
		}
		return tx.Hash().String(), err
	default:
		return "", fmt.Errorf("unsupported method")
	}
}
```

## Register contracts

To simply the contract importing, the package creates a sigleton of registered contracts. Each wrapped contract needs to create a factory function which takes an address as parameter and returns the instance of it. For example,

```go
func TestContractFactory(contractAddress string) ethereum.Contract {
	return &TestContract{
		contractAddress: contractAddress,
	}
}
```

For each contract modules, it registers itself using the `RegisterContract` call. For example,

```go
func init() {
	ethereum.RegisterContract("TestContract", TestContractFactory)
}
```

A client can easily query contracts by using `GetContract` function.

## Build abi.go

### Required package

- `truffle`
- `jq`
- `abigen`
- `npm`
- `make`

### Scripts

```shell
make build
```
