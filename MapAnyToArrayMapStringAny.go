package maputils

func MapAnyToArrayMapStringAny(valueAny any) []map[string]any {
	var valueArrayAny []any

	switch valueType := valueAny.(type) {
	case []map[string]any:
		return valueType
	case []any:
		valueArrayAny = valueType
	}

	out := []map[string]any{}
	for _, v := range valueArrayAny {
		out = append(out, v.(map[string]any))
	}
	return out
}
