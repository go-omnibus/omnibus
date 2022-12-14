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

func DefiniteString(value interface{}) string {
	s, _ := ToString(value)
	return s
}

func ToString(value interface{}) (string, error) {
	if value == nil {
		return "", nil
	}

	v := reflect.ValueOf(value)

	switch value.(type) {
	case float32, float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64), nil
	case int, int8, int16, int32, int64:
		return strconv.FormatInt(v.Int(), 10), nil
	case uint, uint8, uint16, uint32, uint64:
		return strconv.FormatUint(v.Uint(), 10), nil
	case string:
		return v.String(), nil
	case []byte:
		return Bytes2Str(v.Bytes()), nil
	case time.Duration:
		return strconv.FormatUint(uint64(value.(time.Duration).Nanoseconds()), 10), nil
	case json.Number:
		return value.(json.Number).String(), nil
	default:
		return "", ErrConvertType
	}
}

func DefiniteInt64(value interface{}) int64 {
	i, _ := ToInt64(value)
	return i
}

func ToInt64(value interface{}) (int64, error) {
	v := reflect.ValueOf(value)

	switch value.(type) {
	case nil:
		return 0, nil
	case int, int8, int16, int32, int64:
		return v.Int(), nil
	case uint, uint8, uint16, uint32, uint64:
		return int64(v.Uint()), nil
	case float32, float64:
		return int64(v.Float()), nil
	case string:
		return strconv.ParseInt(v.String(), 0, 64)
	case time.Duration:
		return int64(value.(time.Duration)), nil
	case json.Number:
		return value.(json.Number).Int64()
	default:
		return 0, ErrConvertType
	}
}

func DefiniteInt(value interface{}) int {
	i, _ := ToInt(value)
	return i
}

func ToInt(value interface{}) (int, error) {
	v := reflect.ValueOf(value)

	switch value.(type) {
	case nil:
		return 0, nil
	case int, int8, int16, int32, int64:
		return int(v.Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return int(v.Uint()), nil
	case float32, float64:
		return int(v.Float()), nil
	case string:
		return strconv.Atoi(v.String())
	case time.Duration:
		return int(value.(time.Duration)), nil
	case json.Number:
		i, err := value.(json.Number).Int64()
		if err != nil {
			return 0, err
		}
		return int(i), nil
	default:
		return 0, ErrConvertType
	}
}
