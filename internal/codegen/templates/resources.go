package templates

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

// Resource defines the structure for a Nutanix resource
type Resource struct {
	Name              string
	ResourceType      string
	Description       string
	ClientGetFunc     string
	ClientListFunc    string // Regular List function with DSMetadata parameter
	ClientListAllFunc string // ListAll function with filter string parameter
	HasListFunc       bool   // Whether the service has a ListX function
	HasListAllFunc    bool   // Whether the service has a ListAllX function
}

const resourceTemplate = `package resources

import (
    "context"

    "github.com/thunderboltsid/mcp-nutanix/internal/client"

    "github.com/mark3labs/mcp-go/mcp"
    "github.com/mark3labs/mcp-go/server"
)

// {{.Name}} defines the {{.Name}} resource template
func {{.Name}}() mcp.ResourceTemplate {
    return mcp.NewResourceTemplate(
        string(ResourceURIPrefix(ResourceType{{.Name}})) + "{uuid}",
        string(ResourceType{{.Name}}),
        mcp.WithTemplateDescription("{{.Description}}"),
        mcp.WithTemplateMIMEType("application/json"),
    )
}

// {{.Name}}Handler implements the handler for the {{.Name}} resource
func {{.Name}}Handler() server.ResourceTemplateHandlerFunc {
    return CreateResourceHandler(ResourceType{{.Name}}, func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error) {
        // Get the {{.Name}}
        return client.V3().{{.ClientGetFunc}}(ctx, uuid)
    })
}
`

// GetResourceDefinitions returns all Nutanix resource definitions
func GetResourceDefinitions() []Resource {
	return []Resource{
		{
			Name:              "VM",
			ResourceType:      "vm",
			Description:       "Virtual Machine resource",
			ClientGetFunc:     "GetVM",
			ClientListFunc:    "ListVM",
			ClientListAllFunc: "ListAllVM",
			HasListFunc:       true,
			HasListAllFunc:    true,
		},
		{
			Name:              "Cluster",
			ResourceType:      "cluster",
			Description:       "Cluster resource",
			ClientGetFunc:     "GetCluster",
			ClientListFunc:    "ListCluster",
			ClientListAllFunc: "ListAllCluster",
			HasListFunc:       true,
			HasListAllFunc:    true,
		},
		{
			Name:              "Image",
			ResourceType:      "image",
			Description:       "Image resource",
			ClientGetFunc:     "GetImage",
			ClientListFunc:    "ListImage",
			ClientListAllFunc: "ListAllImage",
			HasListFunc:       true,
			HasListAllFunc:    true,
		},
		{
			Name:              "Subnet",
			ResourceType:      "subnet",
			Description:       "Subnet resource",
			ClientGetFunc:     "GetSubnet",
			ClientListFunc:    "ListSubnet",
			ClientListAllFunc: "ListAllSubnet",
			HasListFunc:       true,
			HasListAllFunc:    true,
		},
		{
			Name:              "Host",
			ResourceType:      "host",
			Description:       "Host resource",
			ClientGetFunc:     "GetHost",
			ClientListFunc:    "ListHost",
			ClientListAllFunc: "ListAllHost",
			HasListFunc:       true,
			HasListAllFunc:    true,
		},
		{
			Name:              "Project",
			ResourceType:      "project",
			Description:       "Project resource",
			ClientGetFunc:     "GetProject",
			ClientListFunc:    "ListProject",
			ClientListAllFunc: "ListAllProject",
			HasListFunc:       true,
			HasListAllFunc:    true,
		},
		{
			Name:              "VolumeGroup",
			ResourceType:      "volumegroup",
			Description:       "Volume Group resource",
			ClientGetFunc:     "GetVolumeGroup",
			ClientListFunc:    "ListVolumeGroup",
			ClientListAllFunc: "",
			HasListFunc:       true,
			HasListAllFunc:    false,
		},
		{
			Name:              "NetworkSecurityRule",
			ResourceType:      "networksecurityrule",
			Description:       "Network Security Rule resource",
			ClientGetFunc:     "GetNetworkSecurityRule",
			ClientListFunc:    "ListNetworkSecurityRule",
			ClientListAllFunc: "ListAllNetworkSecurityRule",
			HasListFunc:       true,
			HasListAllFunc:    true,
		},
		{
			Name:              "Category",
			ResourceType:      "category",
			Description:       "Category resource",
			ClientGetFunc:     "GetCategoryKey",
			ClientListFunc:    "ListCategories",
			ClientListAllFunc: "",
			HasListFunc:       true,
			HasListAllFunc:    false,
		},
		{
			Name:              "AccessControlPolicy",
			ResourceType:      "accesscontrolpolicy",
			Description:       "Access Control Policy resource",
			ClientGetFunc:     "GetAccessControlPolicy",
			ClientListFunc:    "ListAccessControlPolicy",
			ClientListAllFunc: "ListAllAccessControlPolicy",
			HasListFunc:       true,
			HasListAllFunc:    true,
		},
		{
			Name:              "Role",
			ResourceType:      "role",
			Description:       "Role resource",
			ClientGetFunc:     "GetRole",
			ClientListFunc:    "ListRole",
			ClientListAllFunc: "ListAllRole",
			HasListFunc:       true,
			HasListAllFunc:    true,
		},
		{
			Name:              "User",
			ResourceType:      "user",
			Description:       "User resource",
			ClientGetFunc:     "GetUser",
			ClientListFunc:    "ListUser",
			ClientListAllFunc: "ListAllUser",
			HasListFunc:       true,
			HasListAllFunc:    true,
		},
		{
			Name:              "UserGroup",
			ResourceType:      "usergroup",
			Description:       "User Group resource",
			ClientGetFunc:     "GetUserGroup",
			ClientListFunc:    "ListUserGroup",
			ClientListAllFunc: "ListAllUserGroup",
			HasListFunc:       true,
			HasListAllFunc:    true,
		},
		{
			Name:              "Permission",
			ResourceType:      "permission",
			Description:       "Permission resource",
			ClientGetFunc:     "GetPermission",
			ClientListFunc:    "ListPermission",
			ClientListAllFunc: "ListAllPermission",
			HasListFunc:       true,
			HasListAllFunc:    true,
		},
		{
			Name:              "ProtectionRule",
			ResourceType:      "protectionrule",
			Description:       "Protection Rule resource",
			ClientGetFunc:     "GetProtectionRule",
			ClientListFunc:    "ListProtectionRules",
			ClientListAllFunc: "ListAllProtectionRules",
			HasListFunc:       true,
			HasListAllFunc:    true,
		},
		{
			Name:              "RecoveryPlan",
			ResourceType:      "recoveryplan",
			Description:       "Recovery Plan resource",
			ClientGetFunc:     "GetRecoveryPlan",
			ClientListFunc:    "ListRecoveryPlans",
			ClientListAllFunc: "ListAllRecoveryPlans",
			HasListFunc:       true,
			HasListAllFunc:    true,
		},
		{
			Name:              "ServiceGroup",
			ResourceType:      "servicegroup",
			Description:       "Service Group resource",
			ClientGetFunc:     "GetServiceGroup",
			ClientListFunc:    "",
			ClientListAllFunc: "ListAllServiceGroups",
			HasListFunc:       false,
			HasListAllFunc:    true,
		},
		{
			Name:              "AddressGroup",
			ResourceType:      "addressgroup",
			Description:       "Address Group resource",
			ClientGetFunc:     "GetAddressGroup",
			ClientListFunc:    "ListAddressGroups",
			ClientListAllFunc: "ListAllAddressGroups",
			HasListFunc:       true,
			HasListAllFunc:    true,
		},
		{
			Name:              "RecoveryPlanJob",
			ResourceType:      "recoveryplanjob",
			Description:       "Recovery Plan Job resource",
			ClientGetFunc:     "GetRecoveryPlanJob",
			ClientListFunc:    "ListRecoveryPlanJobs",
			ClientListAllFunc: "",
			HasListFunc:       true,
			HasListAllFunc:    false,
		},
		{
			Name:              "AvailabilityZone",
			ResourceType:      "availabilityzone",
			Description:       "Availability Zone resource",
			ClientGetFunc:     "GetAvailabilityZone",
			ClientListFunc:    "",
			ClientListAllFunc: "",
			HasListFunc:       false,
			HasListAllFunc:    false,
		},
	}
}

// GenerateResourceFiles generates resource files for all Nutanix resources
func GenerateResourceFiles(baseDir string) error {
	resources := GetResourceDefinitions()

	// Create the resources directory if it doesn't exist
	resourcesDir := fmt.Sprintf("%s/pkg/resources", baseDir)
	err := os.MkdirAll(resourcesDir, 0755)
	if err != nil {
		return fmt.Errorf("error creating resources directory: %w", err)
	}

	// Parse the resource template
	tmpl, err := template.New("resource").Parse(resourceTemplate)
	if err != nil {
		return fmt.Errorf("error parsing resource template: %w", err)
	}

	// Generate resource files
	for _, res := range resources {
		// Create resource file
		resourceFilePath := fmt.Sprintf("%s/%s.go", resourcesDir, strings.ToLower(res.Name))
		resourceFile, err := os.Create(resourceFilePath)
		if err != nil {
			fmt.Printf("Error creating resource file for %s: %v\n", res.Name, err)
			continue
		}
		defer resourceFile.Close()

		// Execute the template
		err = tmpl.Execute(resourceFile, res)
		if err != nil {
			fmt.Printf("Error executing resource template for %s: %v\n", res.Name, err)
		}
	}

	return nil
}
