package u

import "reflect"

// Type get data type
func Type(data interface{}) reflect.Type {
	return reflect.TypeOf(data)
}
