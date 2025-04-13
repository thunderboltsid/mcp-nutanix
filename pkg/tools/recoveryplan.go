package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RecoveryPlan defines the RecoveryPlan tool
func RecoveryPlanList() mcp.Tool {
	return mcp.NewTool("recoveryplan_list",
		mcp.WithDescription("List recoveryplan resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// RecoveryPlanListHandler implements the handler for the RecoveryPlan list tool
func RecoveryPlanListHandler() server.ToolHandlerFunc {
	return CreateListToolHandler(
		resources.ResourceTypeRecoveryPlan,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllRecoveryPlans(ctx, "")

		},
	)
}

// RecoveryPlanCount defines the RecoveryPlan count tool
func RecoveryPlanCount() mcp.Tool {
	return mcp.NewTool("recoveryplan_count",
		mcp.WithDescription("Count recoveryplan resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// RecoveryPlanCountHandler implements the handler for the RecoveryPlan count tool
func RecoveryPlanCountHandler() server.ToolHandlerFunc {
	return CreateCountToolHandler(
		resources.ResourceTypeRecoveryPlan,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			resp, err := client.V3().ListAllRecoveryPlans(ctx, "")

			if err != nil {
				return nil, err
			}

			res := map[string]interface{}{
				"resource_type": "RecoveryPlan",
				"count":         len(resp.Entities),
				"metadata":      resp.Metadata,
			}

			return res, nil
		},
	)
}
