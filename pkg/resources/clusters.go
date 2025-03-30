package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Cluster defines the cluster resource template
func Cluster() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		"cluster://{uuid}",
		"cluster",
		mcp.WithTemplateDescription("Prism Element (AOS) Cluster resource"),
		mcp.WithTemplateMIMEType("application/json"),
	)
}

// ClusterHandler implements the handler for the cluster resource
func ClusterHandler() server.ResourceTemplateHandlerFunc {
	return func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		clusterUUID := extractIDFromURI(request.Params.URI)
		if clusterUUID == "" {
			return nil, fmt.Errorf("URI must contain a UUID")
		}

		// Get the Prism client
		prismClient := client.GetPrismClient()
		if prismClient == nil {
			return nil, fmt.Errorf("prism client not initialized, please set credentials first")
		}

		// Get the cluster
		resp, err := prismClient.V3().GetCluster(ctx, clusterUUID)
		if err != nil {
			return nil, fmt.Errorf("failed to get cluster: %w", err)
		}

		// Convert the ClusterIntentResponse to JSON
		jsonBytes, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("failed to marshal cluster response: %w", err)
		}

		// Create a more human-readable markdown version
		var markdownBuilder strings.Builder
		markdownBuilder.WriteString(fmt.Sprintf("# Cluster: %s\n\n", resp.Spec.Name))
		markdownBuilder.WriteString(fmt.Sprintf("**UUID**: %s\n\n", resp.Metadata.UUID))

		markdownBuilder.WriteString("## Network Information\n\n")
		if resp.Status.Resources != nil && resp.Status.Resources.Network != nil {
			network := resp.Status.Resources.Network
			markdownBuilder.WriteString(fmt.Sprintf("- **External IP**: %s\n", network.ExternalIP))
			markdownBuilder.WriteString(fmt.Sprintf("- **External Data Services IP**: %s\n", network.ExternalDataServicesIP))
			markdownBuilder.WriteString(fmt.Sprintf("- **FQDN**: %s\n", network.FullyQualifiedDomainName))

			if len(network.NameServerIPList) > 0 {
				markdownBuilder.WriteString("\n**Name Servers**:\n")
				for _, server := range network.NameServerIPList {
					markdownBuilder.WriteString(fmt.Sprintf("- %s\n", server))
				}
			}

			if len(network.NtpServerIPList) > 0 {
				markdownBuilder.WriteString("\n**NTP Servers**:\n")
				for _, server := range network.NtpServerIPList {
					markdownBuilder.WriteString(fmt.Sprintf("- %s\n", server))
				}
			}
		}

		markdownBuilder.WriteString("\n## Cluster Information\n\n")
		markdownBuilder.WriteString(fmt.Sprintf("- **State**: %s\n", resp.Status.State))

		if resp.Status.Resources != nil {
			if resp.Status.Resources.Nodes != nil {
				markdownBuilder.WriteString(fmt.Sprintf("- **Number of Nodes**: %d\n",
					len(resp.Status.Resources.Nodes.HypervisorServerList)))
			}

			if resp.Status.Resources.Config != nil {
				config := resp.Status.Resources.Config
				markdownBuilder.WriteString(fmt.Sprintf("- **Redundancy Factor**: %d\n", config.RedundancyFactor))
				markdownBuilder.WriteString(fmt.Sprintf("- **Operation Mode**: %s\n", config.OperationMode))
				markdownBuilder.WriteString(fmt.Sprintf("- **Timezone**: %s\n", config.Timezone))

				markdownBuilder.WriteString("\n## Services and Features\n\n")
				if len(config.ServiceList) > 0 {
					markdownBuilder.WriteString("**Services**:\n")
					for _, service := range config.ServiceList {
						markdownBuilder.WriteString(fmt.Sprintf("- %s\n", service))
					}
				}

				if len(config.EnabledFeatureList) > 0 {
					markdownBuilder.WriteString("\n**Enabled Features**:\n")
					for _, feature := range config.EnabledFeatureList {
						markdownBuilder.WriteString(fmt.Sprintf("- %s\n", feature))
					}
				}

				markdownBuilder.WriteString("\n## Software Information\n\n")
				if config.Build != nil {
					markdownBuilder.WriteString(fmt.Sprintf("- **Version**: %s\n", config.Build.Version))
					markdownBuilder.WriteString(fmt.Sprintf("- **Build Type**: %s\n", config.Build.BuildType))
				}
			}
		}

		// Create and return both JSON and markdown versions
		return []mcp.ResourceContents{
			&mcp.TextResourceContents{
				URI:      request.Params.URI,
				MIMEType: "application/json",
				Text:     string(jsonBytes),
			},
			&mcp.TextResourceContents{
				URI:      request.Params.URI + "/markdown",
				MIMEType: "text/markdown",
				Text:     markdownBuilder.String(),
			},
		}, nil
	}
}
