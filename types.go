package ethereum

import (
	"fmt"
	"math/big"
	"strconv"
)

type BigInt struct {
	big.Int
}

// UnmarshalJSON parses numbers and string into BigInt
func (i *BigInt) UnmarshalJSON(input []byte) error {
	inputString := string(input)

	// if
	if s, err := strconv.Unquote(inputString); err == nil {
		if _, ok := i.SetString(s, 0); !ok {
			return fmt.Errorf("not a valid big integer: %s", input)
		}
		return nil
	} else {
		if _, ok := i.SetString(inputString, 0); !ok {
			return fmt.Errorf("not a valid big integer: %s", input)
		}
	}
	return nil
}

// MarshalText converts BigInt back to string value
func (i BigInt) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}
