package omnibus

import (
	"encoding/json"
	"reflect"
	"strconv"
	"time"
)

func DefiniteFloat64(value interface{}) float64 {
	f, _ := ToFloat64(value)
	return f
}

func ToFloat64(value interface{}) (float64, error) {
	v := reflect.ValueOf(value)

	switch value.(type) {
	case nil:
		return 0.0, nil
	case int, int8, int16, int32, int64:
		return float64(v.Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return float64(v.Uint()), nil
	case float32, float64:
		return v.Float(), nil
	case string:
		return strconv.ParseFloat(v.String(), 64)
	case time.Duration:
		return float64(value.(time.Duration)), nil
	case json.Number:
		return value.(json.Number).Float64()
	default:
		return 0.0, ErrConvertType
	}
}
