package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RecoveryPlan defines the RecoveryPlan tool
func RecoveryPlan() mcp.Tool {
	return mcp.NewTool("recoveryplans",
		mcp.WithDescription("List recoveryplan resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// RecoveryPlanHandler implements the handler for the RecoveryPlan tool
func RecoveryPlanHandler() server.ToolHandlerFunc {
	return CreateToolHandler(
		resources.ResourceTypeRecoveryPlan,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllRecoveryPlans(ctx, "")

		},
	)
}
