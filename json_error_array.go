package simplejson

type JSONArray interface {
	JSONArray(index int) JSONArray
	JSONObject(index int) JSONObject
	String(index int) string
	Bool(index int) bool
	Int(index int) int
	Float32(index int) float32
	Float64(index int) float64
	Length() int
	hasError() bool
	failingKey() string
}

type JSONErrorArray struct {
	firstFailingKey string
}

// JSONArray returns JSONArray from specific index
func (j *JSONErrorArray) JSONArray(index int) JSONArray {
	return &JSONErrorArray{j.firstFailingKey}
}

// JSONObject returns JSONArray from specific index
func (j *JSONErrorArray) JSONObject(index int) JSONObject {
	return &JSONErrorObject{j.firstFailingKey}
}

// String returns String from specific index
func (j *JSONErrorArray) String(index int) string {
	panic(failingKeyPanicPrefix + j.firstFailingKey)
}

// Bool returns bool from specific index
func (j *JSONErrorArray) Bool(index int) bool {
	panic(failingKeyPanicPrefix + j.firstFailingKey)
}

// Int returns int from specific index
func (j *JSONErrorArray) Int(index int) int {
	panic(failingKeyPanicPrefix + j.firstFailingKey)
}

// Float32 returns float32 from specific index
func (j *JSONErrorArray) Float32(index int) float32 {
	panic(failingKeyPanicPrefix + j.firstFailingKey)
}

// Float64 returns float64 from specific index
func (j *JSONErrorArray) Float64(index int) float64 {
	panic(failingKeyPanicPrefix + j.firstFailingKey)
}

// Length returns the length of j JSONArray
func (j *JSONErrorArray) Length() int {
	panic(failingKeyPanicPrefix + j.firstFailingKey)
}

func (j *JSONErrorArray) failingKey() string {
	return j.firstFailingKey
}

func (j *JSONErrorArray) hasError() bool {
	return true
}
