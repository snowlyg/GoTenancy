package amazonpay

import (
	"encoding/json"
	"fmt"
)

func updateParams(params *Params, namespace string, value interface{}) error {
	resultMap := map[string]interface{}{}

	result, err := json.Marshal(value)
	if err == nil {
		err = json.Unmarshal(result, &resultMap)
	}

	if err != nil {
		return err
	}

	var addToParam func(key string, value interface{})

	addToParam = func(key string, value interface{}) {
		switch v := value.(type) {
		case []interface{}:
			for i, vv := range v {
				addToParam(fmt.Sprintf("%v[%v]", key, i), vv)
			}
		case map[string]interface{}:
			for k, vv := range v {
				addToParam(fmt.Sprintf("%v.%v", key, k), vv)
			}
		default:
			params.Set(key, fmt.Sprint(v))
		}
	}

	for key, value := range resultMap {
		if namespace != "" {
			addToParam(fmt.Sprintf("%v.%v", namespace, key), value)
		} else {
			addToParam(key, value)
		}
	}

	return nil
}
