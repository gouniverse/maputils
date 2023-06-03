package maputils

import "reflect"

func AnyToMapStringString(data any) map[string]string {
	if data == nil {
		return nil
	}

	t := reflect.TypeOf(data)

	// Is it a map?
	if t.Kind() != reflect.Map {
		return nil
	}

	elemKeyKind := t.Key().Kind()
	elemKValueKind := t.Elem().Kind()

	// Is it a match?
	if elemKeyKind == reflect.String && elemKValueKind == reflect.String {
		return data.(map[string]string)
	}

	// Is it map[string]any? Convert to map[string]string
	if elemKeyKind == reflect.String && elemKValueKind == reflect.Interface {
		mapStringInterface := data.(map[string]interface{})
		mapStringString := map[string]string{}
		for key, value := range mapStringInterface {
			mapStringString[key] = toString(value)
		}
		return mapStringString
	}

	// Last attempt to convert, may panic and need another condition
	return data.(map[string]string)
}
