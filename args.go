package i18n

func Args(args ...interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	argCount := len(args)
	for i := 0; i < argCount; i += 2 {
		if keyVal, ok := args[i].(string); ok {
			var val interface{}
			if i+2 > argCount {
				val = nil
			} else {
				val = args[i+1]
			}
			res[keyVal] = val
		}
	}
	return res
}
