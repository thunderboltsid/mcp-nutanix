package resources

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RecoveryPlan defines the RecoveryPlan resource template
func RecoveryPlan() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		string(ResourceURIPrefix(ResourceTypeRecoveryPlan))+"{uuid}",
		string(ResourceTypeRecoveryPlan),
		mcp.WithTemplateDescription("Recovery Plan resource"),
		mcp.WithTemplateMIMEType("application/json"),
	)
}

// RecoveryPlanHandler implements the handler for the RecoveryPlan resource
func RecoveryPlanHandler() server.ResourceTemplateHandlerFunc {
	return CreateResourceHandler(ResourceTypeRecoveryPlan, func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error) {
		// Get the RecoveryPlan
		return client.V3().GetRecoveryPlan(ctx, uuid)
	})
}
