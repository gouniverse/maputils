package maputils

import (
	"reflect"
	"testing"
)

func TestAnyToMapStringString(t *testing.T) {
	mapInterface := map[string]string{
		"first_name": "John",
		"last_name":  "Doe",
	}

	mapString := AnyToMapStringString(mapInterface)

	expected := map[string]string{"first_name": "John", "last_name": "Doe"}

	if !reflect.DeepEqual(mapString, expected) {
		t.Error("Result", mapString, "DOES NOT match expected:", expected)
	}
}

func TestAnyMapStringInterfaceToMapStringString(t *testing.T) {
	mapInterface := map[string]any{
		"first_name": "John",
		"last_name":  "Doe",
	}

	mapString := AnyToMapStringString(mapInterface)

	expected := map[string]string{"first_name": "John", "last_name": "Doe"}

	if !reflect.DeepEqual(mapString, expected) {
		t.Error("Result", mapString, "DOES NOT match expected:", expected)
	}
}
