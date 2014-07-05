package simplejson

import "encoding/json"

//JSONArray is a struct that represents an Array of JSON-compatible types and provides methods to access them
type JSONArray struct{
	innerArray [] interface{}
}

// NewJSONArrayFromString returns a new JSONArray, parsed from a string or an error if unsuccessful
func NewJSONArrayFromString(jsonarray string) (*JSONArray, error) {
	var resultingArray []interface{}
	err := json.Unmarshal([]byte(jsonarray), &resultingArray)
	if err != nil {
		return nil, err
	}
	return &JSONArray{resultingArray}, nil
}

func NewJSONArrayWithArray(array []interface{}) (*JSONArray, bool) {
	resultingArray, ok := interfaceToJsonCompatible(array)
	if !ok {
		return nil, false
	}
	castedArray := resultingArray.([]interface{})
	return &JSONArray{castedArray}, true
}

// JSONArray returns JSONArray from specific index
func (this *JSONArray) JSONArray(index int) *JSONArray {
	return &JSONArray{this.innerArray[index].([]interface{})}
}

// JSONObject returns JSONArray from specific index
func (this *JSONArray) JSONObject(index int) *JSONObject {
	return &JSONObject{this.innerArray[index].(map[string]interface{})}
}

// String returns String from specific index
func (this *JSONArray) String(index int) string {
	return this.innerArray[index].(string)
}

// Bool returns bool from specific index
func (this *JSONArray) Bool(index int) bool {
	return this.innerArray[index].(bool)
}

// Int returns int from specific index
func (this *JSONArray) Int(index int) int {
	return parseInt(this.innerArray[index])
}

// Float32 returns float32 from specific index
func (this *JSONArray) Float32(index int) float32 {
	float64Representation := float32(this.innerArray[index].(float64))
	return float64Representation
}

// Float64 returns float64 from specific index
func (this *JSONArray) Float64(index int) float64 {
	return this.innerArray[index].(float64)
}

// Length returns the length of this JSONArray
func (this *JSONArray) Length() int {
	theInnerArray := this.innerArray;
	return len(theInnerArray)
}
