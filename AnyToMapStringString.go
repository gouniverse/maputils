package maputils

func AnyToMapStringString(data any) map[string]string {
	return data.(map[string]string)
}
