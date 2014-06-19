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

// GetJSONArray returns JSONArray from specific index
func (this *JSONArray) GetJSONArray(index int) *JSONArray {
	return &JSONArray{this.innerArray[index].([]interface{})}
}

// GetJSONObject returns JSONArray from specific index
func (this *JSONArray) GetJSONObject(index int) *JSONObject {
	return &JSONObject{this.innerArray[index].(map[string]interface{})}
}

// GetString returns String from specific index
func (this *JSONArray) GetString(index int) string {
	return this.innerArray[index].(string)
}

// GetBool returns bool from specific index
func (this *JSONArray) GetBool(index int) bool {
	return this.innerArray[index].(bool)
}

// GetInt returns int from specific index
func (this *JSONArray) GetInt(index int) int {
	return parseInt(this.innerArray[index])
}

// GetFloat32 returns float32 from specific index
func (this *JSONArray) GetFloat32(index int) float32 {
	float64Representation := float32(this.innerArray[index].(float64))
	return float64Representation
}

// GetFloat64 returns float64 from specific index
func (this *JSONArray) GetFloat64(index int) float64 {
	return this.innerArray[index].(float64)
}

// Length returns the length of this JSONArray
func (this *JSONArray) Length() int {
	theInnerArray := this.innerArray;
	return len(theInnerArray)
}
