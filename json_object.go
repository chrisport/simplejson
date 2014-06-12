package simplejson

import "encoding/json"

type JSONObject struct{
	innerMap map[string] interface{}
}

func NewJSONObjectFromString(jsonobject string) (*JSONObject, error) {
	var resultingMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonobject), &resultingMap)
	if err != nil {
		return nil, err
	}
	return &JSONObject{resultingMap}, nil
}

func NewJSONObject() *JSONObject {
	return &JSONObject{make(map[string] interface{})}
}

func NewJSONObjectWithMap(fromMap map[string] interface{}) *JSONObject {
	return &JSONObject{fromMap}
}

func (this *JSONObject) GetJSONArray(key string) *JSONArray {
	return &JSONArray{interfaceToInterfaceArray(this.innerMap[key])}
}

func (this *JSONObject) GetJSONObject(key string) *JSONObject {
	return &JSONObject{this.innerMap[key].(map[string]interface{})}
}

func (this *JSONObject) GetString(key string) string {
	return this.innerMap[key].(string)
}

func (this *JSONObject) GetBool(key string) bool {
	return this.innerMap[key].(bool)
}

func (this *JSONObject) GetInt(key string) int {
	return parseInt(this.innerMap[key])
}

func (this *JSONObject) GetFloat32(key string) float32 {
	return float32(this.innerMap[key].(float64))
}

func (this *JSONObject) GetFloat64(key string) float64 {
	return this.innerMap[key].(float64)
}

func (this *JSONObject) Set(key string, value interface{}) {
	this.innerMap[key] = value
}

func (this *JSONObject) String() (string, error) {
	jsonString, err := json.Marshal(this.innerMap)
	return string(jsonString), err
}
