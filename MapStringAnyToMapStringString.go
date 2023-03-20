package maputils

// MapStringAnyToMapStringString converts a map[string]any to map[string]string
func MapStringAnyToMapStringString(data map[string]any) map[string]string {
	result := map[string]string{}
	for k, v := range data {
		result[k] = toString(v)
	}
	return result
}
