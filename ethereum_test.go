package ethereum

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

const TestAccountMnemonic = "december neither pumpkin oblige frog cheese worth meadow canyon favorite pair question"

func TestSignABIParameters(t *testing.T) {
	w, err := NewWalletFromMnemonic(TestAccountMnemonic, "testnet", "http://127.0.0.1:7545")
	assert.NoError(t, err)

	bitmarkID, ok := big.NewInt(0).SetString("67ba454e109b5d10b577dc1686338d2ba2759d402814e562920fc6c06df267dc", 16)
	if !ok {
		t.Error("invalid bitmark id")
	}

	r, s, v, err := w.SignABIParameters(context.Background(),
		[]string{"address", "uint256", "uint256"},
		common.HexToAddress("0xa704aa575172f84d917ccaf5ec775da8227ea6c3"),
		bitmarkID,
		big.NewInt(2),
	)
	assert.NoError(t, err)
	assert.Equal(t, "0x9a24d0f023082d20202e790d42c78d1fc7bcf75956722967245be110d3136618", r)
	assert.Equal(t, "0x22bb070a873de3402f4fb498c5f3962a7d9e898ebbed868a2fcce7ee1da81025", s)
	assert.Equal(t, "0x1c", v)
}

func TestSignABIParametersWithInvalidType(t *testing.T) {
	w, err := NewWalletFromMnemonic(TestAccountMnemonic, "testnet", "http://127.0.0.1:7545")
	assert.NoError(t, err)

	_, _, _, err = w.SignABIParameters(context.Background(),
		[]string{"aaddress"},
		common.HexToAddress("0xa704aa575172f84d917ccaf5ec775da8227ea6c3"),
	)

	assert.Error(t, err, "unsupported arg type: aaddress")
}

func TestSignABIParametersWithMismatchedType(t *testing.T) {
	w, err := NewWalletFromMnemonic(TestAccountMnemonic, "testnet", "http://127.0.0.1:7545")
	assert.NoError(t, err)

	_, _, _, err = w.SignABIParameters(context.Background(),
		[]string{"address"},
		big.NewInt(1),
	)

	assert.Error(t, err, "unsupported arg type: aaddress")
}
