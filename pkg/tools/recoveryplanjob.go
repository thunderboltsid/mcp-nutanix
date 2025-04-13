package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	v3 "github.com/nutanix-cloud-native/prism-go-client/v3"
)

// RecoveryPlanJob defines the RecoveryPlanJob tool
func RecoveryPlanJobList() mcp.Tool {
	return mcp.NewTool("recoveryplanjob_list",
		mcp.WithDescription("List recoveryplanjob resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// RecoveryPlanJobListHandler implements the handler for the RecoveryPlanJob list tool
func RecoveryPlanJobListHandler() server.ToolHandlerFunc {
	return CreateListToolHandler(
		resources.ResourceTypeRecoveryPlanJob,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Create DSMetadata without filter
			metadata := &v3.DSMetadata{}

			return client.V3().ListRecoveryPlanJobs(ctx, metadata)

		},
	)
}

// RecoveryPlanJobCount defines the RecoveryPlanJob count tool
func RecoveryPlanJobCount() mcp.Tool {
	return mcp.NewTool("recoveryplanjob_count",
		mcp.WithDescription("Count recoveryplanjob resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// RecoveryPlanJobCountHandler implements the handler for the RecoveryPlanJob count tool
func RecoveryPlanJobCountHandler() server.ToolHandlerFunc {
	return CreateCountToolHandler(
		resources.ResourceTypeRecoveryPlanJob,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Create DSMetadata without filter
			metadata := &v3.DSMetadata{}

			resp, err := client.V3().ListRecoveryPlanJobs(ctx, metadata)

			if err != nil {
				return nil, err
			}

			res := map[string]interface{}{
				"resource_type": "RecoveryPlanJob",
				"count":         len(resp.Entities),
				"metadata":      resp.Metadata,
			}

			return res, nil
		},
	)
}
