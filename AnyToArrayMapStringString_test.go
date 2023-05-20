package maputils

import (
	"reflect"
	"testing"
)

func TestAnyToArrayMapString(t *testing.T) {
	arrayStringString := []map[string]string{
		{"first_name": "John"},
		{"last_name": "Doe"},
	}

	mapString := AnyToArrayMapStringString(arrayStringString)

	expected := []map[string]string{{"first_name": "John"}, {"last_name": "Doe"}}

	if !reflect.DeepEqual(mapString, expected) {
		t.Error("Result", mapString, "DOES NOT match expected:", expected)
	}
}
