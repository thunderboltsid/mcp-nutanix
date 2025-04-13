package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// AccessControlPolicy defines the AccessControlPolicy tool
func AccessControlPolicyList() mcp.Tool {
	return mcp.NewTool("accesscontrolpolicy_list",
		mcp.WithDescription("List accesscontrolpolicy resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// AccessControlPolicyListHandler implements the handler for the AccessControlPolicy list tool
func AccessControlPolicyListHandler() server.ToolHandlerFunc {
	return CreateListToolHandler(
		resources.ResourceTypeAccessControlPolicy,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllAccessControlPolicy(ctx, "")

		},
	)
}

// AccessControlPolicyCount defines the AccessControlPolicy count tool
func AccessControlPolicyCount() mcp.Tool {
	return mcp.NewTool("accesscontrolpolicy_count",
		mcp.WithDescription("Count accesscontrolpolicy resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// AccessControlPolicyCountHandler implements the handler for the AccessControlPolicy count tool
func AccessControlPolicyCountHandler() server.ToolHandlerFunc {
	return CreateCountToolHandler(
		resources.ResourceTypeAccessControlPolicy,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			resp, err := client.V3().ListAllAccessControlPolicy(ctx, "")

			if err != nil {
				return nil, err
			}

			res := map[string]interface{}{
				"resource_type": "AccessControlPolicy",
				"count":         len(resp.Entities),
				"metadata":      resp.Metadata,
			}

			return res, nil
		},
	)
}
