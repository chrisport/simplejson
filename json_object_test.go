package simplejson

import (
	"testing"
)

func TestParsingStringToJSONObject(t *testing.T) {
	jsonString :=
	`{
		"keyString":"stringValue",
		"keyInt": 123,
		"keyFloat64": 1.23,
		"keyFloat32": 1.23,
		"keyBool": true,
		"keyJSONObject": {
							"keyString":"stringValue",
							"keyInt": 123
						}
	}`
	jsonObject, err := NewJSONObjectFromString(jsonString);
	if err != nil {
		t.Error("Parsing failed with error: " + err.Error())
		t.FailNow()
	}

	stringValue := jsonObject.GetString("keyString")
	if (stringValue != "stringValue") {
		t.Error("keyString was " + stringValue + " instead of \"stringValue\" ")
	}

	intValue := jsonObject.GetInt("keyInt")
	if (intValue != 123) {
		t.Errorf("keyInt was %d instead of 123 ", intValue)
	}

	float32Value := jsonObject.GetFloat32("keyFloat32")
	if (float32Value != 1.23) {
		t.Errorf("keyFloat32 was %d instead of 1.23 ", float32Value)
	}

	float64Value := jsonObject.GetFloat32("keyFloat64")
	if (float64Value != 1.23) {
		t.Errorf("keyFloat64 was %d instead of 1.23 ", float64Value)
	}

	boolValue := jsonObject.GetBool("keyBool")
	if (boolValue != true) {
		t.Errorf("keyBool was %b instead of true ", boolValue)
	}

	jsonValue := jsonObject.GetJSONObject("keyJSONObject")
	if (jsonValue.GetString("keyString") != "stringValue") {
		t.Error("keyJSONObject didn't include the string strin")
	}
}

