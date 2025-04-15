package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/internal/json"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ApiNamespacesList defines the API namespaces list tool
func ApiNamespacesList() mcp.Tool {
	return mcp.NewTool("api_namespaces_list",
		mcp.WithDescription("List available API namespaces and their routes in Prism Central"),
	)
}

// ApiNamespacesListHandler implements the handler for the API namespaces list tool
func ApiNamespacesListHandler() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Get the Prism client
		prismClient := client.GetPrismClient()

		// Call the Actuator API to get version routes
		response, err := prismClient.V4().ActuatorApiInstance.GetVersionRoutes(ctx)
		if err != nil {
			return nil, err
		}

		// Convert to JSON using regular JSON encoder
		cjson := json.RegularJSONEncoder(response)
		jsonBytes, err := cjson.MarshalJSON()
		if err != nil {
			return nil, err
		}

		return mcp.NewToolResultText(string(jsonBytes)), nil
	}
}
