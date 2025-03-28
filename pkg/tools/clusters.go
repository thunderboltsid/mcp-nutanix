package tools

import (
	"context"
	"fmt"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func Clusters() mcp.Tool {
	tool := mcp.NewTool("clusters",
		mcp.WithDescription("Perform basic cluster operations"),
		mcp.WithString("operation",
			mcp.Required(),
			mcp.Description("The operation to perform (get, list)"),
			mcp.Enum("get", "list"),
		),
		mcp.WithString("name",
			mcp.Description("Cluster name"),
		),
	)

	return tool
}

func ClustersHandler() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		op := request.Params.Arguments["operation"].(string)
		name := request.Params.Arguments["name"].(string)

		_, err := client.GetPrismClient().V3().ListAllCluster(ctx, "")
		if err != nil {
			return nil, err
		}

		var result string
		switch op {
		case "list":
		case "get":
			if name == "" {
				return nil, fmt.Errorf("no name provided to get the cluster")
			}
		default:
			return nil, fmt.Errorf("operation: %s not supported on clusters", op)
		}

		return mcp.NewToolResultText(result), nil
	}
}
