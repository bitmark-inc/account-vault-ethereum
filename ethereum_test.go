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
	mnemonic := "exotic syrup achieve seven dial idle isolate vintage very harbor adult oxygen"
	w, err := NewWalletFromMnemonic(mnemonic, "livenet", "http://127.0.0.1:7545")
	assert.NoError(t, err)
	js := json.RawMessage(`{
		"domain":{
		   "chainId":0,
		   "name":"My amazing dApp",
		   "verifyingContract":"0x1C56346CD2A2Bf3202F771f50d3D14a367B48070",
		   "version":"2",
		   "salt":"0xf2d857f4a3edcb9b78b4d503bfe733db1e3f6cdc2b7971ee739626c97e86a558"
		},
		"message":{
		   "amount":100,
		   "bidder":{
			  "userId":323,
			  "wallet":"0x3333333333333333333333333333333333333333"
		   }
		},
		"primaryType":"Bid",
		"types":{
		   "EIP712Domain":[
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
			  },
			  {
				 "name":"salt",
				 "type":"bytes32"
			  }
		   ],
		   "Bid":[
			  {
				 "name":"amount",
				 "type":"uint256"
			  },
			  {
				 "name":"bidder",
				 "type":"Identity"
			  }
		   ],
		   "Identity":[
			  {
				 "name":"userId",
				 "type":"uint256"
			  },
			  {
				 "name":"wallet",
				 "type":"address"
			  }
		   ]
		}
	 }`)

	r, s, v, err := w.SignETHTypedDataV4(context.Background(), js)
	assert.NoError(t, err, "SignETHTypedDataV4 error")
	assert.Equal(t, "0x602f1f0627198c2071c4a7f2dca96b047923d42517d1d26b20a71589970cdc81", r)
	assert.Equal(t, "0x57363d053782350fea0d7d2196720333b48997e21c113ef76f25aa8ba1c65605", s)
	assert.Equal(t, "0x1c", v)
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
