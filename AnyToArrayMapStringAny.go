package maputils

func AnyToArrayMapStringAny(data any) []map[string]any {
	return data.([]map[string]any)
}
