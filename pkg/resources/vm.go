package resources

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// VM defines the VM resource template
func VM() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		string(ResourceURIPrefix(ResourceTypeVM))+"{uuid}",
		string(ResourceTypeVM),
		mcp.WithTemplateDescription("Virtual Machine resource"),
		mcp.WithTemplateMIMEType("application/json"),
	)
}

// VMHandler implements the handler for the VM resource
func VMHandler() server.ResourceTemplateHandlerFunc {
	return CreateResourceHandler(ResourceTypeVM, func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error) {
		// Use converged API to get VM by ID
		return client.Converged().VMs.Get(ctx, uuid)
	})
}
