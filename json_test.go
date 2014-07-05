package simplejson

import (
	"testing"
	"strconv"
	"fmt"
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

var jsonArrayAllTypes = `{
	"keyString":["stringValue"],
	"keyInt": [123],
	"keyFloat64": [1.23],
	"keyFloat32": [1.23],
	"keyBool": [true],
	"keyJSONObject": [{
						"keyString":"stringValue",
						"keyInt": 123
					}],
	"keyArray":[[
		"string1",
		"string2"
	]]
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
var jsonArrayAsRoot2 = `[
	0,
	1,
	2
]
`

func TestJsonArrayWithInts(t *testing.T) {
	jsonArray, err := NewJSONArrayFromString(jsonArrayAsRoot2);
	if err != nil {
		t.Error("Parsing failed with error: " + err.Error())
		t.FailNow()
	}
	if jsonArray.GetInt(0) != 0 || jsonArray.GetInt(1) != 1 || jsonArray.GetInt(2) != 2 {
		t.Error("Parsing failed with error: " + err.Error())
		t.FailNow()
	}
}

func TestParseInt(t *testing.T) {
	var someFloat64 float64
	var someInt int
	var someInterface interface{}
	someInt = 123
	someFloat64 = 123
	someInterface = 123
	if parseInt(someInt) != 123 || parseInt(someFloat64) != 123 || parseInt(someInterface) != 123 {
		t.Error("parseInt could not parse interface")
		t.FailNow()
	}
}

func TestJSONObjectAllTypesExceptArray(t *testing.T) {
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

	float64Value := jsonObject.GetFloat64("keyFloat64")
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

func TestJSONArrayAllTypes(t *testing.T) {
	jsonObject, err := NewJSONObjectFromString(jsonArrayAllTypes);
	if err != nil {
		t.Error("Parsing failed with error: " + err.Error())
		t.FailNow()
	}

	stringValue := jsonObject.GetJSONArray("keyString").GetString(0)
	if (stringValue != "stringValue") {
		t.Error("keyString was " + stringValue + " instead of \"stringValue\" ")
	}

	intValue := jsonObject.GetJSONArray("keyInt").GetInt(0)
	if (intValue != 123) {
		t.Errorf("keyInt was %d instead of 123 ", intValue)
	}

	float32Value := jsonObject.GetJSONArray("keyFloat32").GetFloat32(0)
	if (float32Value != 1.23) {
		t.Errorf("keyFloat32 was %d instead of 1.23 ", float32Value)
	}

	float64Value := jsonObject.GetJSONArray("keyFloat64").GetFloat64(0)
	if (float64Value != 1.23) {
		t.Errorf("keyFloat64 was %d instead of 1.23 ", float64Value)
	}

	boolValue := jsonObject.GetJSONArray("keyBool").GetBool(0)
	if (boolValue != true) {
		t.Errorf("keyBool was %b instead of true ", boolValue)
	}

	jsonValue := jsonObject.GetJSONArray("keyJSONObject").GetJSONObject(0)
	if (jsonValue.GetString("keyString") != "stringValue") {
		t.Error("keyJSONObject didn't include the string strin")
	}
}

func TestParsingStringToJSONArray(t *testing.T) {
	jsonArray, err := NewJSONArrayFromString(jsonArrayAsRoot);
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
		t.Error("object1 was %d instead of 123 ", object1)
	}

	object2 := jsonArray.GetString(2)
	if object2 != "stringElement" {
		t.Error("object1 was " + object2 + " instead of \"stringElement\"")
	}
}


func TestSet(t *testing.T) {
	jsonobject := NewJSONObject();
	jsonobject.Set("object1", "stringValue")
	if value := jsonobject.GetString("object1"); value != "stringValue" {
		t.Error("object1 was " + value + " instead of \"stringValue\"")
	}

	object2 := make([]float32, 5, 5)
	object2[3] = 19.88
	jsonobject.Set("object2", object2)
	fmt.Println(jsonobject.String())
	valueArray := jsonobject.GetJSONArray("object2")
	if current := valueArray.GetInt(0); current != 0 {
		t.Error("object2[0] was %d instead of 0", current)
	}else if current := valueArray.GetInt(1); current != 0 {
		t.Error("object2[1] was %d instead of 0", current)
	}else if current := valueArray.GetInt(2); current != 0 {
		t.Error("object2[2] was %d instead of 0", current)
	}else if current := valueArray.GetFloat32(3); current != 19.88 {
		t.Error("object2[3] was %d instead of 1988", current)
	}else if current := valueArray.GetInt(4); current != 0 {
		t.Error("object2[4] was %d instead of 0", current)
	}
}
