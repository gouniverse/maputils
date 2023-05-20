package maputils

func AnyToMapStringAny(data any) map[string]any {
	return data.(map[string]any)
}
