package simplejson

import (
	"math"
	"fmt"
)

func parseInt(rawValue interface{}) int {
	switch t := rawValue.(type) {
	default:
		panic(fmt.Sprintf("interface{} could not be converted to int. Value was: %d", rawValue))
	case float64:
		if intValue := math.Trunc(t); intValue != rawValue {
			panic(fmt.Sprintf("float64 could not be converted to int. Value was: %d", rawValue))
		}else {
			return int(intValue)
		}
	case int:
		return t

	}
}

func interfaceToInterfaceArray(aInterface interface{}) (interfaceSlice []interface{}) {
	// these are the default array types, there should never be a different type stored in an JSONArray/JSONObject
	// the method Set(Key, Value) has to ensure that the types are converted to one of those.
	switch dataSlice := aInterface.(type) {
	default:
		panic(fmt.Sprintf("interface{} could not be converted to []interface{}. Value was: %t", aInterface))
	case []bool:
		interfaceSlice = make([]interface{}, len(dataSlice))
	for i, d := range dataSlice {
		interfaceSlice[i] = d
	}
	case []string:
		interfaceSlice = make([]interface{}, len(dataSlice))
	for i, d := range dataSlice {
		interfaceSlice[i] = d
	}
	case []float64:
		interfaceSlice = make([]interface{}, len(dataSlice))
	for i, d := range dataSlice {
		interfaceSlice[i] = d
	}
	case []interface{}:
		interfaceSlice = make([]interface{}, len(dataSlice))
	for i, d := range dataSlice {
		interfaceSlice[i] = d
	}
	}
	return
}
