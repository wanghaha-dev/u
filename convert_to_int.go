package u

import (
	"bytes"
	"encoding/binary"
	"regexp"
	"strconv"
	"strings"
)

// ToInt to int
func ToInt(val interface{}) int {
	return int(ToInt64(val))
}

// ToInt8 to int8
func ToInt8(val interface{}) int8 {
	return int8(ToInt64(val))
}

// ToInt16 to int16
func ToInt16(val interface{}) int16 {
	return int16(ToInt64(val))
}

// ToInt32 to int32
func ToInt32(val interface{}) int32 {
	return int32(ToInt64(val))
}

// ToInt64 to int64
func ToInt64(val interface{}) int64 {
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
			return int64(floatVal)
		}
		return 0
	case int:
		return int64(val.(int))
	case int8:
		return int64(val.(int8))
	case int16:
		return int64(val.(int16))
	case int32:
		return int64(val.(int32))
	case int64:
		return v
	case uint:
		return int64(val.(uint))
	case uint8:
		return int64(val.(uint8))
	case uint16:
		return int64(val.(uint16))
	case uint32:
		return int64(val.(uint32))
	case uint64:
		return int64(val.(int64))
	case float32:
		return int64(val.(float32))
	case float64:
		return int64(val.(float64))
	case []byte:
		buffer := bytes.NewBuffer(val.([]byte))
		var data int64
		err := binary.Read(buffer, binary.BigEndian, &data)
		if err != nil {
			return 0
		}
		return int64(data)
	default:
		return 0
	}
}

// ToUint to uint
func ToUint(val interface{}) uint {
	return uint(ToUint64(val))
}

// ToUint8 to uint8
func ToUint8(val interface{}) uint8 {
	return uint8(ToUint64(val))
}

// ToUint16 to uint16
func ToUint16(val interface{}) uint16 {
	return uint16(ToUint64(val))
}

// ToUint32 to uint32
func ToUint32(val interface{}) uint32 {
	return uint32(ToUint64(val))
}

// ToUint64 to uint64
func ToUint64(val interface{}) uint64 {
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
			return uint64(floatVal)
		}
		return 0
	case int:
		return uint64(val.(int))
	case int8:
		return uint64(val.(int8))
	case int16:
		return uint64(val.(int16))
	case int32:
		return uint64(val.(int32))
	case int64:
		return uint64(val.(int64))
	case uint:
		return uint64(val.(uint))
	case uint8:
		return uint64(val.(uint8))
	case uint16:
		return uint64(val.(uint16))
	case uint32:
		return uint64(val.(uint32))
	case uint64:
		return v
	case float32:
		return uint64(val.(float32))
	case float64:
		return uint64(val.(float64))
	case []byte:
		buffer := bytes.NewBuffer(val.([]byte))
		var data uint64
		err := binary.Read(buffer, binary.BigEndian, &data)
		if err != nil {
			return 0
		}
		return uint64(data)
	default:
		return 0
	}
}
