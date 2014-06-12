package simplejson

import (
	"testing"
	"fmt"
	"strconv"
)

var jsonString = `{
	"keyString":"stringValue",
	"keyInt": 123,
	"keyFloat64": 1.23,
	"keyFloat32": 1.23,
	"keyBool": true,
	"keyJSONObject": {
						"keyString":"stringValue",
						"keyInt": 123
					},
	"keyArray":[
		"string1",
		"string2"
	]
}`

var jsonArrayAsRoot = `[
	{
		"elemKeyString": "stringValue",
		"elemKeyInt": 0
	},
	123,
	"stringElement"
]
`

func TestJSONObjectWithoutArrays(t *testing.T) {
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

func TestParsingStringToJSONArray(t *testing.T) {
	jsonArray, err := NewJSONArrayFromString(jsonArrayAsRoot);
	fmt.Printf("%d",len(jsonArray))
	if err != nil {
		t.Error("Parsing failed with error: " + err.Error())
		t.FailNow()
	}

	object0 := jsonArray.GetJSONObject(0)
	if stringValue := object0.GetString("elemKeyString"); stringValue != "stringValue" {
		t.Error("object[0]::elemKeyString was " + stringValue + " instead of \"stringValue\" ")
	}
	if intValue := object0.GetInt("elemKeyInt"); intValue != 0 {
		t.Error("object[0]::elemKeyInt was " + strconv.Itoa(intValue) + " instead of 0 ")
	}

	object1 := jsonArray.GetInt(1)
	if object1 != 123 {
		t.Error("object1 was %d instead of 123 ",object1)
	}

	object2 := jsonArray.GetString(2)
	if object2 != "stringElement" {
		t.Error("object1 was " + object2 + " instead of \"stringElement\"")
	}
}
