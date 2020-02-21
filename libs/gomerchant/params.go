package gomerchant

type Params map[string]interface{}

func (params Params) Get(key string) (interface{}, bool) {
	if params == nil {
		return nil, false
	}

	value, ok := params[key]
	return value, ok
}

func (params Params) Set(key string, value interface{}) {
	params[key] = value
}

func (params Params) IgnoreBlankFields() Params {
	var result = Params{}
	for key, value := range params {
		switch value.(type) {
		case string:
			if value != "" {
				result[key] = value
			}
		default:
			result[key] = value
		}
	}

	return result
}
