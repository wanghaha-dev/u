package u

import (
	"fmt"
	"strconv"
)

// ToString val to string
func ToString(val interface{}) string {
	switch v := val.(type) {
	case string:
		return v
	case []byte:
		return string(val.([]byte))
	case int:
		return strconv.Itoa(val.(int))
	case int8:
		return strconv.Itoa(int(val.(int8)))
	case int16:
		return strconv.Itoa(int(val.(int16)))
	case int32:
		return strconv.Itoa(int(val.(int32)))
	case int64:
		return strconv.Itoa(int(val.(int64)))
	case uint:
		return strconv.Itoa(int(val.(uint)))
	case uint8:
		return strconv.Itoa(int(val.(uint8)))
	case uint16:
		return strconv.Itoa(int(val.(uint16)))
	case uint32:
		return strconv.Itoa(int(val.(uint32)))
	case uint64:
		return strconv.Itoa(int(val.(uint64)))
	case float32:
		return fmt.Sprintf("%.2f", val.(float32))
	case float64:
		return fmt.Sprintf("%.2f", val.(float64))
	case bool:
		return strconv.FormatBool(val.(bool))
	default:
		return ""
	}
}
