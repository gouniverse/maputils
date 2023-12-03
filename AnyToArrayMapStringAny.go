package maputils

import "reflect"

func AnyToArrayMapStringAny(data any) []map[string]any {
	if data == nil {
		return nil
	}

	t := reflect.TypeOf(data)

	// Is it array or slice?
	kind := t.Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return nil
	}

	// Is the element a map?
	elemKind := t.Elem().Kind()
	if elemKind != reflect.Map {
		return nil
	}

	elemKeyKind := t.Elem().Key().Kind()
	elemKValueKind := t.Elem().Elem().Kind()

	// Is it a match?
	if elemKeyKind == reflect.String && elemKValueKind == reflect.Interface {
		return data.([]map[string]any)
	}

	// Is it map[string]string? Convert to map[string]any
	if elemKeyKind == reflect.String && elemKValueKind == reflect.String {
		arrayStringString := data.([]map[string]string)
		arrayStringAny := []map[string]any{}
		for i := 0; i < len(arrayStringString); i++ {
			mapStringAny := map[string]any{}
			for key, value := range arrayStringString[i] {
				mapStringAny[key] = value
			}
			arrayStringAny = append(arrayStringAny, mapStringAny)
		}

		return arrayStringAny
	}

	// Make last attempt, may panic and need to add another condition
	return data.([]map[string]any)
}
