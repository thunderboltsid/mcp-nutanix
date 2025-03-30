package resources

import (
	"context"
	"fmt"
	"strings"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/internal/json"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ResourceType enum for different resource types
type ResourceType string

const (
	ResourceTypeVM                  ResourceType = "vm"
	ResourceTypeSubnet              ResourceType = "subnet"
	ResourceTypeImage               ResourceType = "image"
	ResourceTypeCluster             ResourceType = "cluster"
	ResourceTypeHost                ResourceType = "host"
	ResourceTypeProject             ResourceType = "project"
	ResourceTypeVolumeGroup         ResourceType = "volumegroup"
	ResourceTypeNetworkSecurityRule ResourceType = "networksecurityrule"
	ResourceTypeServiceGroup        ResourceType = "servicegroup"
	ResourceTypeAddressGroup        ResourceType = "addressgroup"
	ResourceTypeAccessControlPolicy ResourceType = "accesscontrolpolicy"
	ResourceTypeRole                ResourceType = "role"
	ResourceTypeUser                ResourceType = "user"
	ResourceTypeUserGroup           ResourceType = "usergroup"
	ResourceTypePermission          ResourceType = "permission"
	ResourceTypeProtectionRule      ResourceType = "protectionrule"
	ResourceTypeRecoveryPlan        ResourceType = "recoveryplan"
	ResourceTypeRecoveryPlanJob     ResourceType = "recoveryplanjob"
	ResourceTypeCategory            ResourceType = "category"
	ResourceTypeCategoryValue       ResourceType = "categoryvalue"
	ResourceTypeAvailabilityZone    ResourceType = "availabilityzone"
)

// ResourceHandlerFunc defines a function that handles a specific resource get operation
type ResourceHandlerFunc func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error)

// ResourceURIPrefix returns the URI prefix for a resource type
func ResourceURIPrefix(resourceType ResourceType) string {
	return fmt.Sprintf("%s://", resourceType)
}

// NutanixURI returns a URI for a resource type and UUID
func NutanixURI(resourceType ResourceType, uuid string) string {
	return fmt.Sprintf("%s://%s", resourceType, uuid)
}

// ExtractIDFromURI extracts the UUID from a URI
// uri is expected to be in the format of resourceType://uuid
func ExtractIDFromURI(uri string) string {
	parts := strings.Split(uri, "://")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}

// ExtractTypeFromURI extracts the resource type from a URI
// uri is expected to be in the format of resourceType://uuid
func ExtractTypeFromURI(uri string) ResourceType {
	parts := strings.Split(uri, "://")
	if len(parts) != 2 {
		return ""
	}
	return ResourceType(parts[0])
}

// CreateResourceHandler creates a generic resource handler for any Nutanix resource
func CreateResourceHandler(resourceType ResourceType, handlerFunc ResourceHandlerFunc) server.ResourceTemplateHandlerFunc {
	return func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		uuid := ExtractIDFromURI(request.Params.URI)
		if uuid == "" {
			return nil, fmt.Errorf("URI must contain a UUID")
		}

		// Get the Prism client
		prismClient := client.GetPrismClient()
		if prismClient == nil {
			return nil, fmt.Errorf("prism client not initialized, please set credentials first")
		}

		// Call the specific resource handler
		resource, err := handlerFunc(ctx, prismClient, uuid)
		if err != nil {
			return nil, fmt.Errorf("failed to get %s: %w", resourceType, err)
		}

		// Convert to JSON
		cjson := json.NewCustomJSONEncoder(resource)
		jsonBytes, err := cjson.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("failed to marshal %s details: %w", resourceType, err)
		}

		return []mcp.ResourceContents{
			&mcp.TextResourceContents{
				URI:      request.Params.URI,
				MIMEType: "application/json",
				Text:     string(jsonBytes),
			},
		}, nil
	}
}
