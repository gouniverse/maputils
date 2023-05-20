package maputils

func AnyToArrayMapStringString(data any) []map[string]string {
	if data == nil {
		return []map[string]string{}
	}

	arrayMapStringAny := AnyToArrayMapStringAny(data)

	arrayMapStringString := []map[string]string{}

	for _, v := range arrayMapStringAny {
		mapStringString := MapStringAnyToMapStringString(v)
		arrayMapStringString = append(arrayMapStringString, mapStringString)
	}

	return arrayMapStringString
}
