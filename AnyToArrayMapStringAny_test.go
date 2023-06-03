package maputils

import (
	"reflect"
	"testing"
)

func TestAnyToArrayMapAny(t *testing.T) {
	arrayStringAny := []map[string]any{
		{"first_name": "John"},
		{"last_name": "Doe"},
	}

	mapStringAny := AnyToArrayMapStringAny(arrayStringAny)

	expected := []map[string]any{{"first_name": "John"}, {"last_name": "Doe"}}

	if !reflect.DeepEqual(mapStringAny, expected) {
		t.Error("Result", mapStringAny, "DOES NOT match expected:", expected)
	}
}
