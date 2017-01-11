package simplejson

import "errors"

const (
	failingKeyPanicPrefix = "Error while unmarshalling JSON. Failing key: "
)

type JSONErrorObject struct {
	firstFailingKey string
}

func (j *JSONErrorObject) JSONArray(key string) *JSONArray {
	return nil
}

func (j *JSONErrorObject) JSONObject(key string) JSONObject {
	return &JSONErrorObject{j.firstFailingKey}
}

func (j *JSONErrorObject) String(key string) string {
	panic(failingKeyPanicPrefix + j.firstFailingKey)
}

func (j *JSONErrorObject) Bool(key string) bool {
	panic(failingKeyPanicPrefix + j.firstFailingKey)
}

func (j *JSONErrorObject) Int(key string) int {
	panic(failingKeyPanicPrefix + j.firstFailingKey)
}

func (j *JSONErrorObject) Float32(key string) float32 {
	panic(failingKeyPanicPrefix + j.firstFailingKey)
}

func (j *JSONErrorObject) Float64(key string) float64 {
	panic(failingKeyPanicPrefix + j.firstFailingKey)
}

func (j *JSONErrorObject) Set(key string, value interface{}) bool {
	panic(failingKeyPanicPrefix + j.firstFailingKey)
}

func (j *JSONErrorObject) AsString() (string, error) {
	return "", errors.New("Object is JSONErrorObject")
}

func (j *JSONErrorObject) failingKey() string {
	return j.firstFailingKey
}

func (j *JSONErrorObject) hasError() bool {
	return true
}
