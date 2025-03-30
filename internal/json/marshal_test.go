package json

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalJSON(t *testing.T) {
	jsonData := []byte(`
    {
        "api_version": "3.0",
        "entities": [
            {
                "metadata": {"uuid": "b7816400-16c5-47c7-9fcc-474e39594ad5"},
                "status": {"big": "object"},
                "spec": {
                    "name": "vm1",
					"resources": {
                        "guest_customization": {"very": "large", "nested": "object"},
                        "subnets": [{"name": "subnet1"}, {"name": "subnet2"}]
                    }
                }
            },
            {
                "metadata": {"uuid": "b7816400-16c5-47c7-9fcc-474e39594ad6"},
                "status": {"big": "object2"},
                "spec": {
                    "name": "vm2",
					"resources": {
                        "guest_customization": {"very": "large", "nested": "object"},
                        "subnets": [{"name": "subnet1"}, {"name": "subnet2"}]
                    }
                }
            }
        ]
    }`)
	ujson := make(map[string]any)
	err := json.Unmarshal(jsonData, &ujson)
	assert.NoError(t, err)

	cjson := CustomJSON{
		Value: ujson,
		StripPaths: []string{
			"api_version",
			"entities[].status",
			"spec.resources.guest_customization",
			"entities[].spec.resources.guest_customization",
		},
	}

	mdata, err := json.Marshal(cjson)
	assert.NoError(t, err)
	assert.NotNil(t, mdata)
}
