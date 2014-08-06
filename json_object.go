package simplejson

import (
	"encoding/json"
	"strings"
	"fmt"
	"strconv"
)

type JSONObject struct{
	innerMap map[string] interface{}
}

// NewJSONObjectFromString returns a new JSONObject, parsed from a string or an error if unsuccessful
func NewJSONObjectFromString(jsonobject string) (*JSONObject, error) {
	var resultingMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonobject), &resultingMap)
	if err != nil {
		return nil, err
	}
	return &JSONObject{resultingMap}, nil
}

// NewJSONObject creates and returns a new JSONObject
func NewJSONObject() *JSONObject {
	return &JSONObject{make(map[string] interface{})}
}

func NewJSONObjectWithMap(fromMap map[string] interface{}) *JSONObject {
	return &JSONObject{fromMap}
}

// JSONArray returns JSONArray from specific key
func (this *JSONObject) JSONArray(key string) *JSONArray {
	return &JSONArray{interfaceToInterfaceArray(this.resolveSegmentation(key))}
}

// JSONObject returns JSONObject from specific key
func (this *JSONObject) JSONObject(key string) *JSONObject {
	return &JSONObject{this.resolveSegmentation(key).(map[string]interface{})}
}

// String returns string from specific key
func (this *JSONObject) String(key string) string {
	return this.resolveSegmentation(key).(string)
}

// Bool returns bool from specific key
func (this *JSONObject) Bool(key string) bool {
	return this.resolveSegmentation(key).(bool)
}

// Int returns int from specific key
func (this *JSONObject) Int(key string) int {
	return parseInt(this.resolveSegmentation(key))
}

// Float32 returns float32 from specific key
func (this *JSONObject) Float32(key string) float32 {
	return float32(this.resolveSegmentation(key).(float64))
}

// Float64 returns float64 from specific key
func (this *JSONObject) Float64(key string) float64 {
	return this.resolveSegmentation(key).(float64)
}

// Set sets the value of Key
func (this *JSONObject) Set(key string, value interface{}) bool {
	unmarshalled, ok := interfaceToJsonCompatible(value)
	if !ok {
		return ok
	}
	this.innerMap[key] = unmarshalled
	return true
}

// String return json-representation as string
func (this *JSONObject) AsString() (string, error) {
	jsonString, err := json.Marshal(this.innerMap)
	return string(jsonString), err
}

func (this *JSONObject) GetMap() (map[string]interface{}) {
	return this.innerMap
}

func (this *JSONObject) resolveSegmentation(key string) (value interface{}) {
	keys := strings.Split(key, "::")
	nrOfKeys := len(keys)
	value = this.innerMap[keys[0]]
	keys = keys[1:nrOfKeys]

	for _, currentKey := range keys {

		switch v := value.(type){
		default:
			fmt.Printf("unexpected type %T", v)
		case []interface {}:
			index, err := strconv.Atoi(currentKey)
			if err != nil {
				panic(err)
			}
			value = v[index]
		case map[string] interface {}:
			value = v[currentKey]
		}
	}
	return value
}
