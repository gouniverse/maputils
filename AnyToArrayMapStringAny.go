package maputils

func AnyToArrayMapStringAny(data any) []map[string]any {
	if data == nil {
		return []map[string]any{}
	}

	return data.([]map[string]any)
}
