package ethereum

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const TestAccountMnemonic = "december neither pumpkin oblige frog cheese worth meadow canyon favorite pair question"

func TestSignABIParameters(t *testing.T) {
	w, err := NewWalletFromMnemonic(TestAccountMnemonic, "testnet", "http://127.0.0.1:7545")
	assert.NoError(t, err)

	r, s, v, err := w.SignABIParameters(context.Background(),
		[]interface{}{"address", "uint256", "uint256"},
		"0xa704aa575172f84d917ccaf5ec775da8227ea6c3",
		"46917335283247895592820561757116052373078415782412711500416590899764819814364",
		"2",
	)
	assert.NoError(t, err)
	assert.Equal(t, "0x9a24d0f023082d20202e790d42c78d1fc7bcf75956722967245be110d3136618", r)
	assert.Equal(t, "0x22bb070a873de3402f4fb498c5f3962a7d9e898ebbed868a2fcce7ee1da81025", s)
	assert.Equal(t, "0x1c", v)
}

func TestSignABIParametersWithComplexStruct(t *testing.T) {
	mnemonic := "easily silver wear length license file final nation people worry replace one"
	w, err := NewWalletFromMnemonic(mnemonic, "testnet", "http://127.0.0.1:7545")
	assert.NoError(t, err)

	r, s, v, err := w.SignABIParameters(context.Background(),
		// structure definition
		[]interface{}{
			"uint256",
			"address",
			map[string]interface{}{
				"tuple": map[string]interface{}{
					"price":       "uint256",
					"cost":        "uint256",
					"expiryTime":  "uint256",
					"destination": "address",
					"tokenIds": map[string]interface{}{
						"slice": "uint256",
					},
					"revenueShares": map[string]interface{}{
						"slice": map[string]interface{}{
							"slice": map[string]interface{}{
								"tuple": map[string]interface{}{
									"recipient": "address",
									"bps":       "uint256",
									"ABI_TUPLE_KEY_ORDER": []interface{}{
										"recipient", "bps",
									},
								},
							},
						},
					},
					"payByVaultContract": "bool",
					"ABI_TUPLE_KEY_ORDER": []interface{}{
						"price", "cost", "expiryTime", "destination", "tokenIds", "revenueShares", "payByVaultContract",
					},
				},
			},
		},
		// value
		"1337",
		"0x1DCe55aab29C5417A15C073057Db30CC287AaF89",
		map[string]interface{}{
			"price":       "250000000000000000",
			"cost":        "20000000000000000",
			"expiryTime":  "1696540204",
			"destination": "0x7b7795353f3d2724CE7EE95085620112f5772038",
			"tokenIds": []interface{}{
				"3000000", "3000001", "4000000", "4000001", "4000002",
			},
			"revenueShares": []interface{}{
				[]interface{}{
					map[string]interface{}{
						"recipient": "0xA77645EC8EA044BB8210c27c7ff31c8464Cf52df",
						"bps":       "8000",
					},
					map[string]interface{}{
						"recipient": "0xc49bdacE62d60e3d6f3f42e9C7a8A5fd72D63906",
						"bps":       "2000",
					},
				},
				[]interface{}{
					map[string]interface{}{
						"recipient": "0xA77645EC8EA044BB8210c27c7ff31c8464Cf52df",
						"bps":       "8000",
					},
					map[string]interface{}{
						"recipient": "0xc49bdacE62d60e3d6f3f42e9C7a8A5fd72D63906",
						"bps":       "2000",
					},
				},
				[]interface{}{
					map[string]interface{}{
						"recipient": "0xA77645EC8EA044BB8210c27c7ff31c8464Cf52df",
						"bps":       "8000",
					},
					map[string]interface{}{
						"recipient": "0xc49bdacE62d60e3d6f3f42e9C7a8A5fd72D63906",
						"bps":       "2000",
					},
				},
				[]interface{}{
					map[string]interface{}{
						"recipient": "0xA77645EC8EA044BB8210c27c7ff31c8464Cf52df",
						"bps":       "8000",
					},
					map[string]interface{}{
						"recipient": "0xc49bdacE62d60e3d6f3f42e9C7a8A5fd72D63906",
						"bps":       "2000",
					},
				},
				[]interface{}{
					map[string]interface{}{
						"recipient": "0xA77645EC8EA044BB8210c27c7ff31c8464Cf52df",
						"bps":       "8000",
					},
					map[string]interface{}{
						"recipient": "0xc49bdacE62d60e3d6f3f42e9C7a8A5fd72D63906",
						"bps":       "2000",
					},
				},
			},
			"payByVaultContract": "false",
		},
	)
	assert.NoError(t, err)
	assert.Equal(t, "0x12dec52a184d8c208dd2643a02f4b1e62ecb69cebf3a8d60baac9d727d962b1a", r)
	assert.Equal(t, "0x15fa50556eac33f472ce1c085dca4171e89baa658e5d23c012cb8d63b747b2b6", s)
	assert.Equal(t, "0x1c", v)
}

func TestSignETHTypedDataV4(t *testing.T) {
	// NOTE: the signature validation is from https://weijiekoh.github.io/eip712-signing-demo/index.html
	mnemonic := "exotic syrup achieve seven dial idle isolate vintage very harbor adult oxygen"
	w, err := NewWalletFromMnemonic(mnemonic, "testnet", "http://127.0.0.1:7545")
	assert.NoError(t, err)

	js := json.RawMessage(`{
		"domain": {
		  "name": "Seaport",
			"version": "1.5",
			"chainId": 5,
			"verifyingContract": "0x00000000000000adc04c56bf30ac9d3c0aaf14dc"
		},
		"message": {
			"offerer": "0x00021a5d2a0ef7c84959991ff79c357cb737d1a7",
			"zone": "0x0000000000000000000000000000000000000000",
			"offer": [
				{
					"itemType": 2,
					"token": "0xb9a0ab98e8457ded503d6d3ecdbf2bcc06e6ab3f",
					"identifierOrCriteria": "300221895159563365540897365039676187549271169",
					"startAmount": "1",
					"endAmount": "1"
				}
			],
			"consideration": [
				{
					"itemType": 0,
					"token": "0x0000000000000000000000000000000000000000",
					"identifierOrCriteria": "0",
					"startAmount": "1170000000000000000",
					"endAmount": "1170000000000000000",
					"recipient": "0x00021a5d2a0ef7c84959991ff79c357cb737d1a7"
				},
				{
					"itemType": 0,
					"token": "0x0000000000000000000000000000000000000000",
					"identifierOrCriteria": "0",
					"startAmount": "30000000000000000",
					"endAmount": "30000000000000000",
					"recipient": "0x0000a26b00c1f0df003000390027140000faa719"
				}
			],
			"orderType": 0,
			"startTime": 1698737840,
			"endTime": 1714289900,
			"zoneHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
			"salt": "0x1d4da48b0000000000000000000000009e1da4edfe34e4d35451a438021bdac2",
			"conduitKey": "0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000",
			"counter": "0"
		},
		"primaryType": "OrderComponents",
		"types": {
		   "EIP712Domain": [
			  {
				 "name":"name",
				 "type":"string"
			  },
			  {
				 "name":"version",
				 "type":"string"
			  },
			  {
				 "name":"chainId",
				 "type":"uint256"
			  },
			  {
				 "name":"verifyingContract",
				 "type":"address"
			  }
		   ],
		   "OrderComponents": [
					{
						"name": "offerer",
						"type": "address"
					},
					{
						"name": "zone",
						"type": "address"
					},
					{
						"name": "offer",
						"type": "OfferItem[]"
					},
					{
						"name": "consideration",
						"type": "ConsiderationItem[]"
					},
					{
						"name": "orderType",
						"type": "uint8"
					},
					{
						"name": "startTime",
						"type": "uint256"
					},
					{
						"name": "endTime",
						"type": "uint256"
					},
					{
						"name": "zoneHash",
						"type": "bytes32"
					},
					{
						"name": "salt",
						"type": "uint256"
					},
					{
						"name": "conduitKey",
						"type": "bytes32"
					},
					{
						"name": "counter",
						"type": "uint256"
					}
				],
				"OfferItem": [
					{
						"name": "itemType",
						"type": "uint8"
					},
					{
						"name": "token",
						"type": "address"
					},
					{
						"name": "identifierOrCriteria",
						"type": "uint256"
					},
					{
						"name": "startAmount",
						"type": "uint256"
					},
					{
						"name": "endAmount",
						"type": "uint256"
					}
				],
				"ConsiderationItem": [
					{
						"name": "itemType",
						"type": "uint8"
					},
					{
						"name": "token",
						"type": "address"
					},
					{
						"name": "identifierOrCriteria",
						"type": "uint256"
					},
					{
						"name": "startAmount",
						"type": "uint256"
					},
					{
						"name": "endAmount",
						"type": "uint256"
					},
					{
						"name": "recipient",
						"type": "address"
					}
				]
		}
	 }`)

	signature, err := w.SignETHTypedDataV4(context.Background(), js)
	assert.NoError(t, err, "SignETHTypedDataV4 error")
	assert.Equal(t, "0xe72b657a96fcee4dbd7d5d36c935ddf1ae51af6b7547d6ca3a271d7fda9da0477bba4d555c07ee24b7c94ec85cb37862e840923fb46021b2572d9f72a9995f3301", signature)
}

func TestSendTransaction(t *testing.T) {
	// NOTE: the signature validation is from https://weijiekoh.github.io/eip712-signing-demo/index.html
	mnemonic := "exotic syrup achieve seven dial idle isolate vintage very harbor adult oxygen"
	w, err := NewWalletFromMnemonic(mnemonic, "testnet", "http://127.0.0.1:7545")
	assert.NoError(t, err)

	txID, err := w.TransferTransaction(
		context.Background(),
		"0xC766Ff6854008270D768195Cb93b7D7452B01f63",
		"0xa22cb4650000000000000000000000001e0049783f008a0085193e00003d00cd54003c710000000000000000000000000000000000000000000000000000000000000001",
		nil,
		nil,
	)
	assert.NoError(t, err, "Transfer Transaction error")
	assert.Equal(t, "", txID)
}

func TestSignABIParametersWithInvalidType(t *testing.T) {
	w, err := NewWalletFromMnemonic(TestAccountMnemonic, "testnet", "http://127.0.0.1:7545")
	assert.NoError(t, err)

	_, _, _, err = w.SignABIParameters(context.Background(),
		[]interface{}{"aaddress"},
		"0xa704aa575172f84d917ccaf5ec775da8227ea6c3",
	)
	assert.Equal(t, errors.New("unsupported type"), err)
}

func TestSignABIParametersWithMismatchedType(t *testing.T) {
	w, err := NewWalletFromMnemonic(TestAccountMnemonic, "testnet", "http://127.0.0.1:7545")
	assert.NoError(t, err)

	_, _, _, err = w.SignABIParameters(context.Background(),
		[]interface{}{"address"},
		1,
	)
	assert.Equal(t, errors.New("can't parse value not in type string"), err)
}
