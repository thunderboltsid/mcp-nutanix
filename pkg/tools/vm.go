package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// VM defines the VM tool
func VMList() mcp.Tool {
	return mcp.NewTool("vm_list",
		mcp.WithDescription("List vm resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// VMListHandler implements the handler for the VM list tool
func VMListHandler() server.ToolHandlerFunc {
	return CreateListToolHandler(
		resources.ResourceTypeVM,
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {
			vms, err := client.Converged().VMs.List(ctx)
			if err != nil {
				return nil, err
			}

			return map[string]interface{}{
				"vms":   vms,
				"count": len(vms),
			}, nil
		},
	)
}

// VMCount defines the VM count tool
func VMCount() mcp.Tool {
	return mcp.NewTool("vm_count",
		mcp.WithDescription("Count vm resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// VMCountHandler implements the handler for the VM count tool
func VMCountHandler() server.ToolHandlerFunc {
	return CreateCountToolHandler(
		resources.ResourceTypeVM,
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {
			vms, err := client.Converged().VMs.List(ctx)
			if err != nil {
				return nil, err
			}

			return map[string]interface{}{
				"resource_type": "VM",
				"count":         len(vms),
			}, nil
		},
	)
}
