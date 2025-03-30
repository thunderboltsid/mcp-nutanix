package resources

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// AvailabilityZone defines the AvailabilityZone resource template
func AvailabilityZone() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		string(ResourceURIPrefix(ResourceTypeAvailabilityZone))+"{uuid}",
		string(ResourceTypeAvailabilityZone),
		mcp.WithTemplateDescription("Availability Zone resource"),
		mcp.WithTemplateMIMEType("application/json"),
	)
}

// AvailabilityZoneHandler implements the handler for the AvailabilityZone resource
func AvailabilityZoneHandler() server.ResourceTemplateHandlerFunc {
	return CreateResourceHandler(ResourceTypeAvailabilityZone, func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error) {
		// Get the AvailabilityZone
		return client.V3().GetAvailabilityZone(ctx, uuid)
	})
}
