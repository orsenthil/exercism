package flatten

func Flatten(nested interface{}) []interface{} {
	var result []interface{}

	result = make([]interface{}, 0)

	switch nested.(type) {

	case []interface{}:
		for _, v := range nested.([]interface{}) {
			result = append(result, Flatten(v)...)
		}
	case interface{}:
		if nested != nil {
			result = append(result, nested)
		}
	case nil:
		// do nothing
	}

	return result
}
