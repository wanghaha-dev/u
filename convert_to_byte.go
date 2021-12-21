package u

import (
	"bytes"
	"encoding/binary"
	"math"
)

// ToBytes to []byte
func ToBytes(val interface{}) []byte {
	switch v := val.(type) {
	case []byte:
		return v
	case string:
		return []byte(val.(string))
	case int:
		return int64ToBytes(int64(val.(int)))
	case int8:
		return int64ToBytes(int64(val.(int8)))
	case int16:
		return int64ToBytes(int64(val.(int16)))
	case int32:
		return int64ToBytes(int64(val.(int32)))
	case int64:
		return int64ToBytes(val.(int64))
	case uint:
		return uint64ToBytes(uint64(val.(uint)))
	case uint8:
		return uint64ToBytes(uint64(val.(uint8)))
	case uint16:
		return uint64ToBytes(uint64(val.(uint16)))
	case uint32:
		return uint64ToBytes(uint64(val.(uint32)))
	case uint64:
		return uint64ToBytes(val.(uint64))
	case float32:
		return float32ToBytes(val.(float32))
	case float64:
		return float64ToBytes(val.(float64))
	default:
		return []byte{}
	}
}

// int64ToBytes int64 to []byte
func int64ToBytes(val int64) []byte {
	buffer := bytes.NewBuffer(nil)
	err := binary.Write(buffer, binary.BigEndian, val)
	if err != nil {
		return []byte{}
	}
	return buffer.Bytes()
}

// uint64ToBytes uint64 to []byte
func uint64ToBytes(val uint64) []byte {
	buffer := bytes.NewBuffer(nil)
	err := binary.Write(buffer, binary.BigEndian, val)
	if err != nil {
		return []byte{}
	}
	return buffer.Bytes()
}

// float32ToBytes float32 to []byte
func float32ToBytes(val float32) []byte {
	bits := math.Float32bits(val)
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, bits)
	return b
}

// float64ToBytes float64 to []byte
func float64ToBytes(val float64) []byte {
	bits := math.Float64bits(val)
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, bits)
	return b
}
