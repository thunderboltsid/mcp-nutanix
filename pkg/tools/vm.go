package tools

import (
	"context"
	"reflect"

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
			var allVMs []interface{}
			page := 0
			limit := 100

			for {
				resp, err := client.V4().VmApiInstance.ListVms(&page, &limit, nil, nil, nil)
				if err != nil {
					return nil, err
				}

				if resp.Data != nil {
					if dataValue := resp.Data.GetValue(); dataValue != nil {
						v := reflect.ValueOf(dataValue)
						if v.Kind() == reflect.Slice {
							for i := 0; i < v.Len(); i++ {
								allVMs = append(allVMs, v.Index(i).Interface())
							}
						}
					}
				}

				if resp.Metadata != nil && resp.Metadata.Links != nil {
					hasNext := false
					for _, link := range resp.Metadata.Links {
						if link.Rel != nil && *link.Rel == "next" {
							hasNext = true
							break
						}
					}
					if !hasNext {
						break
					}
					page++
				} else {
					break
				}
			}

			return map[string]interface{}{
				"vms":   allVMs,
				"count": len(allVMs),
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
			totalCount := 0
			page := 0
			limit := 100

			for {
				resp, err := client.V4().VmApiInstance.ListVms(&page, &limit, nil, nil, nil)
				if err != nil {
					return nil, err
				}

				if resp.Data != nil {
					if dataValue := resp.Data.GetValue(); dataValue != nil {
						v := reflect.ValueOf(dataValue)
						if v.Kind() == reflect.Slice {
							totalCount += v.Len()
						}
					}
				}

				if resp.Metadata != nil && resp.Metadata.Links != nil {
					hasNext := false
					for _, link := range resp.Metadata.Links {
						if link.Rel != nil && *link.Rel == "next" {
							hasNext = true
							break
						}
					}
					if !hasNext {
						break
					}
					page++
				} else {
					break
				}
			}

			return map[string]interface{}{
				"resource_type": "VM",
				"count":         totalCount,
			}, nil
		},
	)
}
