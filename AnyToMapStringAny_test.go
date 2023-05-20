package maputils

import (
	"reflect"
	"testing"
)

func TestAnyToMapStringAny(t *testing.T) {
	mapInterface := map[string]any{
		"integer": 1,
		"nil":     nil,
	}

	mapString := AnyToMapStringAny(mapInterface)

	expected := map[string]any{"integer": 1, "nil": nil}

	if !reflect.DeepEqual(mapString, expected) {
		t.Error("Result", mapString, "DOES NOT match expected:", expected)
	}
}
