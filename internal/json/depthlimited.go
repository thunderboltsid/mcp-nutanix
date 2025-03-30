package json

import (
	"encoding/json"
)

type DepthLimitedJSON struct {
	Value     interface{}
	MaxDepth  int
	CurrDepth int
}

func (d DepthLimitedJSON) MarshalJSON() ([]byte, error) {
	if d.CurrDepth >= d.MaxDepth {
		return []byte(`"[depth limit reached]"`), nil
	}

	switch v := d.Value.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})
		for k, val := range v {
			result[k] = DepthLimitedJSON{Value: val, MaxDepth: d.MaxDepth, CurrDepth: d.CurrDepth + 1}
		}
		return json.Marshal(result)
	case []interface{}:
		result := make([]interface{}, len(v))
		for i, val := range v {
			result[i] = DepthLimitedJSON{Value: val, MaxDepth: d.MaxDepth, CurrDepth: d.CurrDepth + 1}
		}
		return json.Marshal(result)
	default:
		return json.Marshal(v)
	}
}
