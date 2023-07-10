package ethereum

import (
	"context"
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
