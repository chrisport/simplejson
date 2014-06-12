package simplejson

import "encoding/json"

type JSONArray struct{
	innerArray [] interface{}
}

func NewJSONArrayFromString(jsonarray string) (*JSONArray, error) {
	var resultingArray []interface{}
	err := json.Unmarshal([]byte(jsonarray), &resultingArray)
	if err != nil {
		return nil, err
	}
	return &JSONArray{resultingArray}, nil
}

func NewJSONArray() *JSONArray {
	return &JSONArray{make([]interface{}, 4)}
}

func NewJSONArrayWithArray(array []interface{}) *JSONArray {
	return &JSONArray{array}
}

func (this *JSONArray) GetJSONArray(index int) *JSONArray {
	return &JSONArray{this.innerArray[index].([]interface{})}
}

func (this *JSONArray) GetJSONObject(index int) *JSONObject {
	return &JSONObject{this.innerArray[index].(map[string]interface{})}
}

func (this *JSONArray) GetString(index int) string {
	return this.innerArray[index].(string)
}

func (this *JSONArray) SetString(index int, value string) {
	this.innerArray[index] = value
}

func (this *JSONArray) GetBool(index int) bool {
	return this.innerArray[index].(bool)
}

func (this *JSONArray) GetInt(index int) int {
	return float64ToInt(this.innerArray[index].(float64))
}

func (this *JSONArray) getFloat32(index int) float32 {
	return this.innerArray[index].(float32)
}

func (this *JSONArray) getFloat64(index int) float64 {
	return this.innerArray[index].(float64)
}

func (this *JSONArray) isNil(index int) bool {
	return this.innerArray[index] == nil
}
