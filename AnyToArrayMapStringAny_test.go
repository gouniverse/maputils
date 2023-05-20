package maputils

import (
	"reflect"
	"testing"
)

func TestAnyToArrayMapAny(t *testing.T) {
	arrayStringString := []map[string]any{
		{"first_name": "John"},
		{"last_name": "Doe"},
	}

	mapString := AnyToArrayMapStringAny(arrayStringString)

	expected := []map[string]any{{"first_name": "John"}, {"last_name": "Doe"}}

	if !reflect.DeepEqual(mapString, expected) {
		t.Error("Result", mapString, "DOES NOT match expected:", expected)
	}
}
