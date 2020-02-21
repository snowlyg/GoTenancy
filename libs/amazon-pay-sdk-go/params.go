package amazonpay

// Params API params
type Params map[string]interface{}

// Get get value from params
func (params Params) Get(key string) (interface{}, bool) {
	if params == nil {
		return nil, false
	}

	value, ok := params[key]
	return value, ok
}

// Set set value to params
func (params Params) Set(key string, value interface{}) {
	params[key] = value
}

// Sign sign params
func (params Params) Sign() string {
	return ""
}
