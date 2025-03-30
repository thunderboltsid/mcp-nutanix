package resources

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RecoveryPlanJob defines the RecoveryPlanJob resource template
func RecoveryPlanJob() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		string(ResourceURIPrefix(ResourceTypeRecoveryPlanJob))+"{uuid}",
		string(ResourceTypeRecoveryPlanJob),
		mcp.WithTemplateDescription("Recovery Plan Job resource"),
		mcp.WithTemplateMIMEType("application/json"),
	)
}

// RecoveryPlanJobHandler implements the handler for the RecoveryPlanJob resource
func RecoveryPlanJobHandler() server.ResourceTemplateHandlerFunc {
	return CreateResourceHandler(ResourceTypeRecoveryPlanJob, func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error) {
		// Get the RecoveryPlanJob
		return client.V3().GetRecoveryPlanJob(ctx, uuid)
	})
}
