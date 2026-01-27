package tools

import (
	"context"
	"reflect"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Subnet defines the Subnet tool
func SubnetList() mcp.Tool {
	return mcp.NewTool("subnet_list",
		mcp.WithDescription("List subnet resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// SubnetListHandler implements the handler for the Subnet list tool
func SubnetListHandler() server.ToolHandlerFunc {
	return CreateListToolHandler(
		resources.ResourceTypeSubnet,
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {
			var allSubnets []interface{}
			page := 0
			limit := 100

			for {
				resp, err := client.V4().SubnetsApiInstance.ListSubnets(&page, &limit, nil, nil, nil, nil)
				if err != nil {
					return nil, err
				}

				if resp.Data != nil {
					if dataValue := resp.Data.GetValue(); dataValue != nil {
						v := reflect.ValueOf(dataValue)
						if v.Kind() == reflect.Slice {
							for i := 0; i < v.Len(); i++ {
								allSubnets = append(allSubnets, v.Index(i).Interface())
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
				"subnets": allSubnets,
				"count":   len(allSubnets),
			}, nil
		},
	)
}

// SubnetCount defines the Subnet count tool
func SubnetCount() mcp.Tool {
	return mcp.NewTool("subnet_count",
		mcp.WithDescription("Count subnet resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// SubnetCountHandler implements the handler for the Subnet count tool
func SubnetCountHandler() server.ToolHandlerFunc {
	return CreateCountToolHandler(
		resources.ResourceTypeSubnet,
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {
			totalCount := 0
			page := 0
			limit := 100

			for {
				resp, err := client.V4().SubnetsApiInstance.ListSubnets(&page, &limit, nil, nil, nil, nil)
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
				"resource_type": "Subnet",
				"count":         totalCount,
			}, nil
		},
	)
}
