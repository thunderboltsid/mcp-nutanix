package resources

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// VolumeGroup defines the VolumeGroup resource template
func VolumeGroup() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		string(ResourceURIPrefix(ResourceTypeVolumeGroup))+"{uuid}",
		string(ResourceTypeVolumeGroup),
		mcp.WithTemplateDescription("Volume Group resource"),
		mcp.WithTemplateMIMEType("application/json"),
	)
}

// VolumeGroupHandler implements the handler for the VolumeGroup resource
func VolumeGroupHandler() server.ResourceTemplateHandlerFunc {
	return CreateResourceHandler(ResourceTypeVolumeGroup, func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error) {
		// Get the VolumeGroup
		return client.V3().GetVolumeGroup(ctx, uuid)
	})
}
