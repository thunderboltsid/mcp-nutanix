package resources

import (
	"context"
	"fmt"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"

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
// Note: V4 API requires both cluster ID and host ID to get a host.
// Since the resource URI only provides the host UUID, we cannot directly call GetHostById.
// This handler returns an error advising users to use the host_list tool instead.
func HostHandler() server.ResourceTemplateHandlerFunc {
	return CreateResourceHandler(ResourceTypeHost, func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error) {
		// V4 API GetHostById requires both clusterExtId and extId
		// We don't have cluster ID from the resource URI
		return nil, fmt.Errorf("host resource access by ID is not supported in v4 API (requires cluster ID). Please use the host_list tool to find host information")
	})
}
