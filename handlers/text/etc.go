package text

func normalizeInterface(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = normalizeInterface(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = normalizeInterface(v)
		}
	}
	return i
}
