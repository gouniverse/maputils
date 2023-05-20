package maputils

func AnyToArrayMapStringString(data any) []map[string]string {
	return data.([]map[string]string)
}
