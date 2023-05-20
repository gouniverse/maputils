package maputils

import (
	"reflect"
	"testing"
)

func TestMapAnyToArrayMapStringAny(t *testing.T) {
	mapInterface := []map[string]any{
		{
			"integer": 1,
			"nil":     nil,
		},
	}

	mapString := MapAnyToArrayMapStringAny(mapInterface)

	expected := []map[string]any{{"integer": 1, "nil": nil}}

	if !reflect.DeepEqual(mapString, expected) {
		t.Error("Result", mapString, "DOES NOT match expected:", expected)
	}
}
