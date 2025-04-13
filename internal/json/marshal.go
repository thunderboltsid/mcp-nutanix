package json

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/itchyny/gojq"
)

// DefaultStripPaths is a list of default paths to strip from the JSON output.
// These paths are used to remove unnecessary or sensitive information from the JSON response.
// guest_customization on VMs is massive and could contain sensitive information.
// spec and status are duplicative info and hence status can be removed.
var DefaultStripPaths = []string{
	"api_version",
	"spec.resources.guest_customization",
	"entities[].spec.resources.guest_customization",
	"entities[].status.resources.guest_customization",
}

func stripProperties(data []byte, paths []string) ([]byte, error) {
	var input interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		return nil, err
	}

	// Start with identity query
	queryStr := "."

	for _, path := range paths {
		if strings.Contains(path, "[]") {
			// Handle array paths (entities[].status)
			parts := strings.Split(path, "[]")
			arrayPath := parts[0]
			fieldPath := parts[1]

			if len(fieldPath) > 0 && fieldPath[0] == '.' {
				fieldPath = fieldPath[1:] // Remove leading dot
			}

			// Correct jq syntax for modifying each element in an array
			queryStr += fmt.Sprintf(" | .%s |= map(del(.%s))", arrayPath, fieldPath)
		} else {
			// Simple path (api_version)
			queryStr += fmt.Sprintf(" | del(.%s)", path)
		}
	}

	// For debugging
	// fmt.Printf("JQ Query: %s\n", queryStr)

	query, err := gojq.Parse(queryStr)
	if err != nil {
		return nil, fmt.Errorf("jq parse error: %v for query: %s", err, queryStr)
	}

	code, err := gojq.Compile(query)
	if err != nil {
		return nil, fmt.Errorf("jq compile error: %v", err)
	}

	iter := code.Run(input)
	result, ok := iter.Next()
	if !ok {
		return nil, fmt.Errorf("jq query returned no results")
	}

	if err, ok := result.(error); ok {
		return nil, fmt.Errorf("jq execution error: %v", err)
	}

	return json.Marshal(result)
}

type CustomJSON struct {
	Value      interface{}
	StripPaths []string
}

type RegularJSON struct {
	Value interface{}
}

func CustomJSONEncoder(value any) *CustomJSON {
	return &CustomJSON{
		Value:      value,
		StripPaths: DefaultStripPaths,
	}
}

func RegularJSONEncoder(value any) *RegularJSON {
	return &RegularJSON{
		Value: value,
	}
}

func (r *RegularJSON) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(r.Value)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *CustomJSON) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(d.Value)
	if err != nil {
		return nil, err
	}

	return stripProperties(data, d.StripPaths)
}
