package u

import (
	"encoding/binary"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// ToFloat32 to float32
func ToFloat32(val interface{}) float32 {
	return float32(ToFloat64(val))
}

// ToFloat64 to float64
func ToFloat64(val interface{}) float64 {
	switch v := val.(type) {
	case string:
		str := val.(string)
		matched, _ := regexp.MatchString(`^[0-9.]+$`, str)
		if !matched {
			return 0
		}

		if strings.Count(val.(string), ".") <= 1 {
			floatVal, err := strconv.ParseFloat(str, 64)
			if err != nil {
				return 0
			}
			return floatVal
		}
		return 0
	case int:
		return float64(val.(int))
	case int8:
		return float64(val.(int8))
	case int16:
		return float64(val.(int16))
	case int32:
		return float64(val.(int32))
	case int64:
		return float64(val.(int64))
	case uint:
		return float64(val.(uint))
	case uint8:
		return float64(val.(uint8))
	case uint16:
		return float64(val.(uint16))
	case uint32:
		return float64(val.(uint32))
	case uint64:
		return float64(val.(int64))
	case float32:
		return float64(val.(float32))
	case float64:
		return v
	case []byte:
		bits := binary.LittleEndian.Uint64(val.([]byte))
		return math.Float64frombits(bits)
	default:
		return 0
	}
}
