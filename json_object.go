package simplejson

import "encoding/json"

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

func newJSONObjectWithMap(fromMap map[string] interface{}) *JSONObject {
	return &JSONObject{fromMap}
}

// GetJSONArray returns JSONArray from specific key
func (this *JSONObject) GetJSONArray(key string) *JSONArray {
	return &JSONArray{interfaceToInterfaceArray(this.innerMap[key])}
}

// GetJSONObject returns JSONObject from specific key
func (this *JSONObject) GetJSONObject(key string) *JSONObject {
	return &JSONObject{this.innerMap[key].(map[string]interface{})}
}

// GetString returns string from specific key
func (this *JSONObject) GetString(key string) string {
	return this.innerMap[key].(string)
}

// GetBool returns bool from specific key
func (this *JSONObject) GetBool(key string) bool {
	return this.innerMap[key].(bool)
}

// GetInt returns int from specific key
func (this *JSONObject) GetInt(key string) int {
	return parseInt(this.innerMap[key])
}

// GetFloat32 returns float32 from specific key
func (this *JSONObject) GetFloat32(key string) float32 {
	return float32(this.innerMap[key].(float64))
}

// GetFloat64 returns float64 from specific key
func (this *JSONObject) GetFloat64(key string) float64 {
	return this.innerMap[key].(float64)
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
func (this *JSONObject) String() (string, error) {
	jsonString, err := json.Marshal(this.innerMap)
	return string(jsonString), err
}
