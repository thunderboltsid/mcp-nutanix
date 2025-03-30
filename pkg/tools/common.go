package tools

import (
	"context"
	"fmt"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/internal/json"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ListResourceFunc defines a function that handles listing a resource type
type ListResourceFunc func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error)

// CreateToolHandler creates a generic tool handler for listing resources
func CreateToolHandler(
	resourceType resources.ResourceType,
	listFunc ListResourceFunc,
) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Get the Prism client
		prismClient := client.GetPrismClient()
		if prismClient == nil {
			return nil, fmt.Errorf("prism client not initialized, please set credentials first")
		}

		// Get filter if provided (for LLM reference only)
		filter, _ := request.Params.Arguments["filter"].(string)

		// List all resources
		resp, err := listFunc(ctx, prismClient, filter)
		if err != nil {
			return nil, fmt.Errorf("failed to list %s: %w", resourceType, err)
		}

		// Convert to JSON
		cjson := json.NewCustomJSONEncoder(resp)
		jsonBytes, err := cjson.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("failed to marshal %s: %w", resourceType, err)
		}

		return mcp.NewToolResultText(string(jsonBytes)), nil
	}
}
