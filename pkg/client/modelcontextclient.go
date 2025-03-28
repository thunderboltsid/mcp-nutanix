package client

import (
	"errors"
	"sync"

	"github.com/nutanix-cloud-native/prism-go-client/environment/providers/mcp"
)

// mcpModelContextClient is an example implementation of ModelContextClient.
// It stores key–value pairs in an in-memory map.
type mcpModelContextClient struct {
	mu   sync.RWMutex
	data map[string]string
}

var PrismClientProvider = &mcpModelContextClient{
	data: make(map[string]string),
}

var _ mcp.ModelContextClient = &mcpModelContextClient{}

// NewMCPModelContextClient creates a new instance with the provided initial data.
func NewMCPModelContextClient(initialData map[string]string) mcp.ModelContextClient {
	if initialData == nil {
		initialData = make(map[string]string)
	}

	return &mcpModelContextClient{
		data: initialData,
	}
}

// GetValue retrieves the value for a given key from the model context.
// It returns an error if the key is not found.
func (c *mcpModelContextClient) GetValue(key string) (string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if val, exists := c.data[key]; exists {
		return val, nil
	}
	return "", errors.New("model context key not found")
}

// Optionally, you can add a method to update values in the model context.
func (c *mcpModelContextClient) UpdateValue(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}
