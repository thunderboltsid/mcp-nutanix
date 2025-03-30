package resources

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Subnet defines the Subnet resource template
func Subnet() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		string(ResourceURIPrefix(ResourceTypeSubnet))+"{uuid}",
		string(ResourceTypeSubnet),
		mcp.WithTemplateDescription("Subnet resource"),
		mcp.WithTemplateMIMEType("application/json"),
	)
}

// SubnetHandler implements the handler for the Subnet resource
func SubnetHandler() server.ResourceTemplateHandlerFunc {
	return CreateResourceHandler(ResourceTypeSubnet, func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error) {
		// Get the Subnet
		return client.V3().GetSubnet(ctx, uuid)
	})
}
