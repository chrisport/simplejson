package simplejson

import (
	"encoding/json"
	"log"
)

type JSONObjectS struct {
	innerMap map[string]interface{}
}

type JSONObject interface {
	JSONArray(key string) *JSONArray
	JSONObject(key string) JSONObject
	String(key string) string
	Bool(key string) bool
	Int(key string) int
	Float32(key string) float32
	Float64(key string) float64
	Set(key string, value interface{}) bool
	AsString() (string, error)
	hasError() bool
	failingKey() string
}

// NewJSONObjectFromString returns a new JSONObject, parsed from a string or an error if unsuccessful
func NewJSONObjectFromString(jsonobject string) (JSONObject, error) {
	var resultingMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonobject), &resultingMap)
	if err != nil {
		return nil, err
	}
	return &JSONObjectS{resultingMap}, nil
}

// NewJSONObject creates and returns a new JSONObject
func NewJSONObject() JSONObject {
	return &JSONObjectS{make(map[string]interface{})}
}

func newJSONObjectWithMap(fromMap map[string]interface{}) JSONObject {
	return &JSONObjectS{fromMap}
}

// JSONArray returns JSONArray from specific key
func (j *JSONObjectS) JSONArray(key string) *JSONArray {
	return &JSONArray{interfaceToInterfaceArray(j.innerMap[key])}
}

// JSONObject returns JSONObject from specific key
func (j *JSONObjectS) JSONObject(key string) JSONObject {
	if val, ok := j.innerMap[key]; ok {
		return &JSONObjectS{val.(map[string]interface{})}
	} else {
		return &JSONErrorObject{key}
	}
}

// String returns string from specific key
func (j *JSONObjectS) String(key string) string {
	return j.innerMap[key].(string)
}

// Bool returns bool from specific key
func (j *JSONObjectS) Bool(key string) bool {
	return j.innerMap[key].(bool)
}

// Int returns int from specific key
func (j *JSONObjectS) Int(key string) int {
	return parseInt(j.innerMap[key])
}

// Float32 returns float32 from specific key
func (j *JSONObjectS) Float32(key string) float32 {
	return float32(j.innerMap[key].(float64))
}

// Float64 returns float64 from specific key
func (j *JSONObjectS) Float64(key string) float64 {
	return j.innerMap[key].(float64)
}

// Set sets the value of Key
func (j *JSONObjectS) Set(key string, value interface{}) bool {
	unmarshalled, ok := interfaceToJsonCompatible(value)
	if !ok {
		return ok
	}
	j.innerMap[key] = unmarshalled
	return true
}

// String return json-representation as string
func (j *JSONObjectS) AsString() (string, error) {
	jsonString, err := json.Marshal(j.innerMap)
	return string(jsonString), err
}

func (j *JSONObjectS) failingKey() string {
	return ""
}

func (j *JSONObjectS) hasError() bool {
	return false
}
