package maputils

import (
	"reflect"
	"testing"
)

func TestMapStringAnyToMapStringString(t *testing.T) {
	mapInterface := map[string]any{
		"integer": 1,
		"nil":     nil,
	}

	mapString := MapStringAnyToMapStringString(mapInterface)

	expected := map[string]string{"integer": "1", "nil": ""}

	if !reflect.DeepEqual(mapString, expected) {
		t.Error("Result", mapString, "DOES NOT match expected:", expected)
	}
}
