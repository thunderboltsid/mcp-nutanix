package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/nutanix-cloud-native/prism-go-client/v3"
)

// Clusters defines the clusters tool interface
func Clusters() mcp.Tool {
	tool := mcp.NewTool("clusters",
		mcp.WithDescription("Perform basic cluster operations"),
		mcp.WithString("operation",
			mcp.Required(),
			mcp.Description("The operation to perform (get, list)"),
			mcp.Enum("get", "list"),
		),
		mcp.WithString("name",
			mcp.Description("Cluster name"),
		),
		mcp.WithString("uuid",
			mcp.Description("Cluster UUID"),
		),
	)

	return tool
}

// ClustersHandler implements the handler for the clusters tool
func ClustersHandler() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		op := request.Params.Arguments["operation"].(string)

		// Get the Prism client
		prismClient := client.GetPrismClient()
		if prismClient == nil {
			return nil, fmt.Errorf("prism client not initialized, please set credentials first")
		}

		switch op {
		case "list":
			// List all clusters
			resp, err := prismClient.V3().ListAllCluster(ctx, "")
			if err != nil {
				return nil, fmt.Errorf("failed to list clusters: %w", err)
			}

			return formatClusterListResponse(resp)

		case "get":
			// Get a specific cluster by name or UUID
			name, hasName := request.Params.Arguments["name"].(string)
			uuid, hasUUID := request.Params.Arguments["uuid"].(string)

			if !hasName && !hasUUID {
				return nil, fmt.Errorf("either name or uuid must be provided to get a cluster")
			}

			// If UUID is provided, get the cluster directly
			if hasUUID && uuid != "" {
				// Return the cluster resource URI
				resourceURI := resources.NutanixURI("cluster", uuid)

				var sb strings.Builder
				sb.WriteString("# Cluster\n\n")
				sb.WriteString(fmt.Sprintf("Accessing cluster with UUID: %s\n\n", uuid))
				sb.WriteString("Detailed cluster information is available at this resource URI:\n\n")
				sb.WriteString(fmt.Sprintf("```\n%s\n```\n", resourceURI))

				return mcp.NewToolResultText(sb.String()), nil
			}

			// If name is provided, find the cluster by name
			if hasName && name != "" {
				return findClusterByName(ctx, prismClient, name)
			}

			return nil, fmt.Errorf("either name or uuid must be provided with valid values")

		default:
			return nil, fmt.Errorf("operation: %s not supported on clusters", op)
		}
	}
}

// formatClusterListResponse formats a ClusterListIntentResponse into a user-friendly output
func formatClusterListResponse(resp *v3.ClusterListIntentResponse) (*mcp.CallToolResult, error) {
	// Build a markdown table for better readability
	var sb strings.Builder
	sb.WriteString("# Clusters\n\n")
	sb.WriteString("| Name | Service IP | Nodes | State | Version | RF | Services |\n")
	sb.WriteString("|------|------------|-------|-------|---------|----|---------|\n")

	for _, entity := range resp.Entities {
		// Default values
		name := entity.Spec.Name
		state := entity.Status.State
		serviceIP := ""
		numNodes := 0
		version := ""
		rf := 0
		serviceStr := ""

		// Extract network information
		if entity.Status.Resources != nil && entity.Status.Resources.Network != nil {
			serviceIP = entity.Status.Resources.Network.ExternalIP
		}

		// Extract node count
		if entity.Status.Resources != nil && entity.Status.Resources.Nodes != nil {
			numNodes = len(entity.Status.Resources.Nodes.HypervisorServerList)
		}

		// Extract version and services
		if entity.Status.Resources != nil && entity.Status.Resources.Config != nil {
			if entity.Status.Resources.Config.Build != nil {
				version = *entity.Status.Resources.Config.Build.Version
			}

			// Get redundancy factor
			rf = int(entity.Status.Resources.Config.RedundancyFactor)

			// Get first service
			if entity.Status.Resources.Config.ServiceList != nil && len(entity.Status.Resources.Config.ServiceList) > 0 {
				serviceStr = *entity.Status.Resources.Config.ServiceList[0]
				if len(entity.Status.Resources.Config.ServiceList) > 1 {
					serviceStr += ", ..."
				}
			}
		}

		sb.WriteString(fmt.Sprintf("| %s | %s | %d | %s | %s | %d | %s |\n",
			name, serviceIP, numNodes, state, version, rf, serviceStr))
	}

	sb.WriteString("\n\nTotal clusters: " + fmt.Sprintf("%d", len(resp.Entities)))
	sb.WriteString("\n\nDetails:\n\n")

	for _, entity := range resp.Entities {
		sb.WriteString(fmt.Sprintf("### %s\n", entity.Spec.Name))
		sb.WriteString(fmt.Sprintf("- **UUID**: %s\n", entity.Metadata.UUID))

		if entity.Status.Resources != nil && entity.Status.Resources.Network != nil {
			sb.WriteString(fmt.Sprintf("- **Service IP**: %s\n", entity.Status.Resources.Network.ExternalIP))
			if entity.Status.Resources.Network.ExternalDataServicesIP != "" {
				sb.WriteString(fmt.Sprintf("- **Data Services IP**: %s\n", entity.Status.Resources.Network.ExternalDataServicesIP))
			}
		}

		if entity.Status.Resources != nil && entity.Status.Resources.Nodes != nil {
			sb.WriteString(fmt.Sprintf("- **Nodes**: %d\n", len(entity.Status.Resources.Nodes.HypervisorServerList)))
		}

		sb.WriteString(fmt.Sprintf("- **State**: %s\n", entity.Status.State))

		if entity.Status.Resources != nil && entity.Status.Resources.Config != nil {
			sb.WriteString(fmt.Sprintf("- **RF**: %d\n", entity.Status.Resources.Config.RedundancyFactor))

			if entity.Status.Resources.Config.ServiceList != nil && len(entity.Status.Resources.Config.ServiceList) > 0 {
				sb.WriteString("- **Services**: ")
				for i, service := range entity.Status.Resources.Config.ServiceList {
					if i > 0 {
						sb.WriteString(", ")
					}
					sb.WriteString(*service)
				}
				sb.WriteString("\n")
			}
		}

		sb.WriteString("\n")
	}

	sb.WriteString("\nTo get details for a specific cluster, use one of the following commands:\n")
	sb.WriteString("```\nclusters get name=<cluster_name>\n```\n")
	sb.WriteString("\nor\n")
	sb.WriteString("```\nclusters get uuid=<cluster_uuid>\n```\n")
	sb.WriteString("\nOr access the resource directly:\n")
	sb.WriteString("```\ncluster://<cluster_uuid>\n```\n")

	return mcp.NewToolResultText(sb.String()), nil
}

// findClusterByName finds a cluster by name by filtering the list client-side
func findClusterByName(ctx context.Context, prismClient *client.NutanixClient, name string) (*mcp.CallToolResult, error) {
	// List all clusters and filter by name client-side
	resp, err := prismClient.V3().ListAllCluster(ctx, "")
	if err != nil {
		return nil, fmt.Errorf("failed to list clusters: %w", err)
	}

	// Find the cluster by name
	var clusterUUID string
	var clusterName string
	var matchingCluster *v3.ClusterIntentResponse

	for _, entity := range resp.Entities {
		if entity.Spec.Name == name {
			clusterUUID = *entity.Metadata.UUID
			clusterName = entity.Spec.Name
			matchingCluster = entity
			break
		}
	}

	if clusterUUID == "" {
		return nil, fmt.Errorf("no cluster found with name: %s", name)
	}

	// Return the cluster resource URI and basic info
	resourceURI := resources.NutanixURI("cluster", clusterUUID)

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("# Cluster: %s\n\n", clusterName))
	sb.WriteString(fmt.Sprintf("Found cluster with UUID: %s\n\n", clusterUUID))

	// Add some basic information from the cluster
	if matchingCluster.Status.Resources != nil {
		if matchingCluster.Status.Resources.Network != nil {
			sb.WriteString(fmt.Sprintf("**Service IP**: %s\n", matchingCluster.Status.Resources.Network.ExternalIP))
		}

		if matchingCluster.Status.Resources.Config != nil && matchingCluster.Status.Resources.Config.Build != nil {
			sb.WriteString(fmt.Sprintf("**Version**: %s\n", matchingCluster.Status.Resources.Config.Build.Version))
		}
	}

	sb.WriteString(fmt.Sprintf("**State**: %s\n\n", matchingCluster.Status.State))
	sb.WriteString("Detailed cluster information is available at this resource URI:\n\n")
	sb.WriteString(fmt.Sprintf("```\n%s\n```\n", resourceURI))

	return mcp.NewToolResultText(sb.String()), nil
}
