package simplejson

import (
	"math"
	"fmt"
	"encoding/json"
)

// tries to parse an int from an interface and returns the resulting int
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

// casts interface to a json-compatible format and returns the resulting value as interface
func interfaceToJsonCompatible(aInterface interface{}) (interface{}, bool) {
	//TODO research for better solution (e.g. direct encoding for every possible type as Marshal is doing
	marshalled, err := json.Marshal(aInterface)
	if err != nil {
		return nil, false
	}
	var unmarshalled interface{}
	err = json.Unmarshal(marshalled, &unmarshalled)
	if (err != nil) {
		return nil, false
	}
	return unmarshalled, true
}

// parses an interface to an interface array
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
