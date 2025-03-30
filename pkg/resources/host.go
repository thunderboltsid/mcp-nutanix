package resources

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Host defines the Host resource template
func Host() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		string(ResourceURIPrefix(ResourceTypeHost))+"{uuid}",
		string(ResourceTypeHost),
		mcp.WithTemplateDescription("Host resource"),
		mcp.WithTemplateMIMEType("application/json"),
	)
}

// HostHandler implements the handler for the Host resource
func HostHandler() server.ResourceTemplateHandlerFunc {
	return CreateResourceHandler(ResourceTypeHost, func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error) {
		// Get the Host
		return client.V3().GetHost(ctx, uuid)
	})
}
