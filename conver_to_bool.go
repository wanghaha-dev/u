package u

import (
	"strings"
)

// ToBool to bool
func ToBool(val interface{}) bool {
	switch v := val.(type) {
	case []byte:
		if len(val.([]byte)) > 0 {
			return true
		}
		return false
	case string:
		if strings.ToLower(val.(string)) == "true" {
			return true
		} else if strings.ToLower(val.(string)) == "false" {
			return false
		} else if len(val.(string)) > 0 {
			return true
		}
		return false
	case bool:
		return v
	case int:
		if val.(int) == 0 {return false}
		return true
	case int8:
		if val.(int8) == 0 {return false}
		return true
	case int16:
		if val.(int16) == 0 {return false}
		return true
	case int32:
		if val.(int32) == 0 {return false}
		return true
	case int64:
		if val.(int64) == 0 {return false}
		return true
	case uint:
		if val.(uint) == 0 {return false}
		return true
	case uint8:
		if val.(uint8) == 0 {return false}
		return true
	case uint16:
		if val.(uint16) == 0 {return false}
		return true
	case uint32:
		if val.(uint32) == 0 {return false}
		return true
	case uint64:
		if val.(uint64) == 0 {return false}
		return true
	case float32:
		if val.(float32) == 0 {return false}
		return true
	case float64:
		if val.(float64) == 0 {return false}
		return true
	default:
		return false
	}
}
