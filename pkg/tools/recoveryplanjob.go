package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	v3 "github.com/nutanix-cloud-native/prism-go-client/v3"
)

// RecoveryPlanJob defines the RecoveryPlanJob tool
func RecoveryPlanJob() mcp.Tool {
	return mcp.NewTool("recoveryplanjobs",
		mcp.WithDescription("List recoveryplanjob resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// RecoveryPlanJobHandler implements the handler for the RecoveryPlanJob tool
func RecoveryPlanJobHandler() server.ToolHandlerFunc {
	return CreateToolHandler(
		resources.ResourceTypeRecoveryPlanJob,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Create DSMetadata without filter
			var length int64 = 100
			metadata := &v3.DSMetadata{
				Length: &length,
			}

			return client.V3().ListRecoveryPlanJobs(ctx, metadata)

		},
	)
}
