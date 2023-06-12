package ethereum

import (
	"encoding/hex"
	"errors"
	"math/big"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// ABI_TUPLE_KEY_ORDER is a specific key that should be defined in tuple structure definition
const ABI_TUPLE_KEY_ORDER = "ABI_TUPLE_KEY_ORDER"

// parseABIParams parse types to proper []string and []abi.ArgumentMarshaling that could be accepted by abi.NewType(),
// also parse the value into proper golang type that (abi.Arguments).Pack accepted
func parseABIParams(types []interface{}, values []interface{}) ([]string, []*[]abi.ArgumentMarshaling, []interface{}, error) {
	abiTypesResult := make([]string, len(types))
	abiTypesArgumentsResult := make([]*[]abi.ArgumentMarshaling, len(types))
	parseValueResult := make([]interface{}, len(values))
	for i, ty := range types {
		switch typ := ty.(type) {
		case string:
			abiTypesResult[i] = typ
			abiTypesArgumentsResult[i] = nil
			pv, err := parseABIValue(typ, values[i])
			if err != nil {
				return nil, nil, nil, err
			}
			parseValueResult[i] = pv
		case map[string]interface{}:
			for structName, spt := range typ {
				switch structName {
				case "tuple":
					if _, ok := spt.(map[string]interface{}); !ok {
						return nil, nil, nil, errors.New("failed to parse tuple properties")
					}
					structProperties := spt.(map[string]interface{})

					if _, ok := values[i].(map[string]interface{}); !ok {
						return nil, nil, nil, errors.New("failed to parse tuple arguments")
					}
					value := values[i].(map[string]interface{})

					if structProperties[ABI_TUPLE_KEY_ORDER] == nil {
						return nil, nil, nil, errors.New("tuple structure should always define key ABI_TUPLE_KEY_ORDER")
					}

					if _, ok := structProperties[ABI_TUPLE_KEY_ORDER].([]string); !ok {
						return nil, nil, nil, errors.New("failed to parse tuple ABI_TUPLE_KEY_ORDER to []string")
					}
					keyOrder := structProperties[ABI_TUPLE_KEY_ORDER].([]string)

					if len(keyOrder) != len(value) {
						return nil, nil, nil, errors.New("tuple key order param missing")
					}

					tupleArguments := make([]abi.ArgumentMarshaling, len(structProperties)-1)

					// create dynamic struct fields
					sf := make([]reflect.StructField, len(keyOrder))
					// create dynamic struct values
					sv := make([]reflect.Value, len(keyOrder))

					for j, key := range keyOrder {
						structProperty := structProperties[key]
						abiTypes, abiTypesArguments, parsedValue, err := parseABIParams([]interface{}{structProperty}, []interface{}{value[key]})
						if err != nil {
							return nil, nil, nil, err
						}

						abiType := abiTypes[0]
						titleKey := strings.Title(key)

						// assign tuple arguments
						if abiTypesArguments[0] == nil {
							tupleArguments[j] = abi.ArgumentMarshaling{
								Name: titleKey,
								Type: abiType,
							}
						} else {
							tupleArguments[j] = abi.ArgumentMarshaling{
								Name:       titleKey,
								Type:       abiType,
								Components: *abiTypesArguments[0],
							}
						}
						sf[j] = reflect.StructField{
							Name: titleKey,
							Type: reflect.TypeOf(parsedValue[0]),
						}
						sv[j] = reflect.ValueOf(parsedValue[0])
					}

					// create dynamic struct instance
					ds := reflect.StructOf(sf)
					e := reflect.New(ds).Elem()

					// assign dynamic struct value
					for o, f := range sf {
						e.FieldByName(f.Name).Set(sv[o])
					}

					abiTypesResult[i] = structName
					abiTypesArgumentsResult[i] = &tupleArguments
					parseValueResult[i] = e.Interface()
				case "slice":
					sliceProperties := spt
					if _, ok := values[i].([]interface{}); !ok {
						return nil, nil, nil, errors.New("failed to parse slice arguments")
					}
					targetValues := values[i].([]interface{})

					var abiTypes []string
					var abiTypesArguments []*[]abi.ArgumentMarshaling
					var err error

					parsedValues := make([]interface{}, len(targetValues))

					for h, targetValue := range targetValues {
						var parsedValue []interface{}
						abiTypes, abiTypesArguments, parsedValue, err = parseABIParams([]interface{}{sliceProperties}, []interface{}{targetValue})
						if err != nil {
							return nil, nil, nil, err
						}
						parsedValues[h] = parsedValue[0]
					}

					// create dynamic slice
					elemType := reflect.ValueOf(parsedValues[0]).Type()
					elemSlice := reflect.New(reflect.SliceOf(elemType)).Elem()

					for _, dv := range parsedValues {
						elemSlice = reflect.Append(elemSlice, reflect.ValueOf(dv))
					}

					abiTypesResult[i] = abiTypes[0] + "[]"
					abiTypesArgumentsResult[i] = abiTypesArguments[0]
					parseValueResult[i] = elemSlice.Interface()
				default:
					return nil, nil, nil, errors.New("not support definition struct")
				}
			}
		default:
			return nil, nil, nil, errors.New("not supported definition struct")
		}
	}
	return abiTypesResult, abiTypesArgumentsResult, parseValueResult, nil
}

// parseABIValue parse the value into proper golang type that (abi.Arguments).Pack accepted
func parseABIValue(abiType string, value interface{}) (interface{}, error) {
	if _, ok := value.(string); !ok {
		return nil, errors.New("can't parse value not in type string")
	}
	stringValue := value.(string)
	// Todo: check value.(string) ok
	var valuesResult interface{}
	switch abiType {
	case "address":
		valuesResult = common.HexToAddress(stringValue)
	case "uint256", "uint128", "uint64", "uint32", "uint16", "uint8":
		n, ok := big.NewInt(0).SetString(stringValue, 0)
		if !ok {
			return nil, errors.New("fail to parse uint number from hex")
		}
		switch abiType {
		case "uint64":
			valuesResult = n.Uint64()
		case "uint32":
			valuesResult = uint32(n.Uint64())
		case "uint16":
			valuesResult = uint16(n.Uint64())
		case "uint8":
			valuesResult = uint8(n.Uint64())
		default:
			valuesResult = n
		}
	case "bytes32":
		b, err := hex.DecodeString(strings.TrimPrefix(stringValue, "0x"))
		if err != nil {
			return nil, errors.New("fail to decode hex string")
		}

		if len(b) > 32 {
			return nil, errors.New("the length of vaule exceed")
		}

		var valueBytes [32]byte
		copy(valueBytes[:], b[:])

		valuesResult = valueBytes
	case "bytes":
		b, err := hex.DecodeString(strings.TrimPrefix(stringValue, "0x"))
		if err != nil {
			return nil, errors.New("fail to decode hex string")
		}

		valuesResult = b
	case "string":
		valuesResult = value
	case "bool":
		if stringValue == "true" {
			valuesResult = true
		} else if stringValue == "false" {
			valuesResult = false
		} else {
			return nil, errors.New("fail to decode bool")
		}
	default:
		return nil, errors.New("unsupported type")
	}
	return valuesResult, nil
}
