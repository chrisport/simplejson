package simplejson

import "encoding/json"

//JSONArrayS is a struct that represents an Array of JSON-compatible types and provides methods to access them
type JSONArrayS struct {
	innerArray []interface{}
}

// NewJSONArrayFromString returns a new JSONArray, parsed from a string or an error if unsuccessful
func NewJSONArrayFromString(jsonarray string) (JSONArray, error) {
	var resultingArray []interface{}
	err := json.Unmarshal([]byte(jsonarray), &resultingArray)
	if err != nil {
		return nil, err
	}
	return &JSONArrayS{resultingArray}, nil
}

func NewJSONArrayWithArray(array []interface{}) (JSONArray, bool) {
	resultingArray, ok := interfaceToJsonCompatible(array)
	if !ok {
		return nil, false
	}
	castedArray := resultingArray.([]interface{})
	return &JSONArrayS{castedArray}, true
}

// JSONArray returns JSONArray from specific index
func (j *JSONArrayS) JSONArray(index int) JSONArray {
	return &JSONArrayS{j.innerArray[index].([]interface{})}
}

// JSONObject returns JSONArray from specific index
func (j *JSONArrayS) JSONObject(index int) JSONObject {
	return &JSONObjectS{j.innerArray[index].(map[string]interface{})}
}

// String returns String from specific index
func (j *JSONArrayS) String(index int) string {
	return j.innerArray[index].(string)
}

// Bool returns bool from specific index
func (j *JSONArrayS) Bool(index int) bool {
	return j.innerArray[index].(bool)
}

// Int returns int from specific index
func (j *JSONArrayS) Int(index int) int {
	return parseInt(j.innerArray[index])
}

// Float32 returns float32 from specific index
func (j *JSONArrayS) Float32(index int) float32 {
	float64Representation := float32(j.innerArray[index].(float64))
	return float64Representation
}

// Float64 returns float64 from specific index
func (j *JSONArrayS) Float64(index int) float64 {
	return j.innerArray[index].(float64)
}

// Length returns the length of j JSONArray
func (j *JSONArrayS) Length() int {
	theInnerArray := j.innerArray
	return len(theInnerArray)
}

func (j *JSONArrayS) failingKey() string {
	return ""
}

func (j *JSONArrayS) hasError() bool {
	return false
}
