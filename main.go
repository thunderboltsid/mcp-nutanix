package main

import (
	"fmt"
	"os"

	"github.com/thunderboltsid/mcp-nutanix/pkg/prompts"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"
	"github.com/thunderboltsid/mcp-nutanix/pkg/tools"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ResourceRegistration represents a resource and tool pairing
type ResourceRegistration struct {
	ToolFunc        func() mcp.Tool
	ToolHandler     server.ToolHandlerFunc
	ResourceFunc    func() mcp.ResourceTemplate
	ResourceHandler server.ResourceTemplateHandlerFunc
}

func main() {
	// Define server hooks for logging and debugging
	hooks := &server.Hooks{}
	hooks.AddOnError(func(id any, method mcp.MCPMethod, message any, err error) {
		fmt.Printf("onError: %s, %v, %v, %v\n", method, id, message, err)
	})

	// Log level based on environment variable
	debugMode := os.Getenv("DEBUG") != ""
	if debugMode {
		hooks.AddBeforeAny(func(id any, method mcp.MCPMethod, message any) {
			fmt.Printf("beforeAny: %s, %v, %v\n", method, id, message)
		})
		hooks.AddOnSuccess(func(id any, method mcp.MCPMethod, message any, result any) {
			fmt.Printf("onSuccess: %s, %v, %v, %v\n", method, id, message, result)
		})
		hooks.AddBeforeInitialize(func(id any, message *mcp.InitializeRequest) {
			fmt.Printf("beforeInitialize: %v, %v\n", id, message)
		})
		hooks.AddAfterInitialize(func(id any, message *mcp.InitializeRequest, result *mcp.InitializeResult) {
			fmt.Printf("afterInitialize: %v, %v, %v\n", id, message, result)
		})
		hooks.AddAfterCallTool(func(id any, message *mcp.CallToolRequest, result *mcp.CallToolResult) {
			fmt.Printf("afterCallTool: %v, %v, %v\n", id, message, result)
		})
		hooks.AddBeforeCallTool(func(id any, message *mcp.CallToolRequest) {
			fmt.Printf("beforeCallTool: %v, %v\n", id, message)
		})
	}

	// Create a new MCP server
	s := server.NewMCPServer(
		"Prism Central",
		"0.0.1",
		server.WithResourceCapabilities(true, true),
		server.WithPromptCapabilities(true),
		server.WithLogging(),
		server.WithHooks(hooks),
	)

	// Add the prompts
	s.AddPrompt(prompts.SetCredentials(), prompts.SetCredentialsResponse())

	// Define all resources and tools
	resourceRegistrations := map[string]ResourceRegistration{
		"vm": {
			ToolFunc:        tools.VM,
			ToolHandler:     tools.VMHandler(),
			ResourceFunc:    resources.VM,
			ResourceHandler: resources.VMHandler(),
		},
		"cluster": {
			ToolFunc:        tools.Cluster,
			ToolHandler:     tools.ClusterHandler(),
			ResourceFunc:    resources.Cluster,
			ResourceHandler: resources.ClusterHandler(),
		},
		"host": {
			ToolFunc:        tools.Host,
			ToolHandler:     tools.HostHandler(),
			ResourceFunc:    resources.Host,
			ResourceHandler: resources.HostHandler(),
		},
		"image": {
			ToolFunc:        tools.Image,
			ToolHandler:     tools.ImageHandler(),
			ResourceFunc:    resources.Image,
			ResourceHandler: resources.ImageHandler(),
		},
		"subnet": {
			ToolFunc:        tools.Subnet,
			ToolHandler:     tools.SubnetHandler(),
			ResourceFunc:    resources.Subnet,
			ResourceHandler: resources.SubnetHandler(),
		},
		"project": {
			ToolFunc:        tools.Project,
			ToolHandler:     tools.ProjectHandler(),
			ResourceFunc:    resources.Project,
			ResourceHandler: resources.ProjectHandler(),
		},
		"volumegroup": {
			ToolFunc:        tools.VolumeGroup,
			ToolHandler:     tools.VolumeGroupHandler(),
			ResourceFunc:    resources.VolumeGroup,
			ResourceHandler: resources.VolumeGroupHandler(),
		},
		"networksecurityrule": {
			ToolFunc:        tools.NetworkSecurityRule,
			ToolHandler:     tools.NetworkSecurityRuleHandler(),
			ResourceFunc:    resources.NetworkSecurityRule,
			ResourceHandler: resources.NetworkSecurityRuleHandler(),
		},
		"category": {
			ToolFunc:        tools.Category,
			ToolHandler:     tools.CategoryHandler(),
			ResourceFunc:    resources.Category,
			ResourceHandler: resources.CategoryHandler(),
		},
		"accesscontrolpolicy": {
			ToolFunc:        tools.AccessControlPolicy,
			ToolHandler:     tools.AccessControlPolicyHandler(),
			ResourceFunc:    resources.AccessControlPolicy,
			ResourceHandler: resources.AccessControlPolicyHandler(),
		},
		"role": {
			ToolFunc:        tools.Role,
			ToolHandler:     tools.RoleHandler(),
			ResourceFunc:    resources.Role,
			ResourceHandler: resources.RoleHandler(),
		},
		"user": {
			ToolFunc:        tools.User,
			ToolHandler:     tools.UserHandler(),
			ResourceFunc:    resources.User,
			ResourceHandler: resources.UserHandler(),
		},
		"usergroup": {
			ToolFunc:        tools.UserGroup,
			ToolHandler:     tools.UserGroupHandler(),
			ResourceFunc:    resources.UserGroup,
			ResourceHandler: resources.UserGroupHandler(),
		},
		"permission": {
			ToolFunc:        tools.Permission,
			ToolHandler:     tools.PermissionHandler(),
			ResourceFunc:    resources.Permission,
			ResourceHandler: resources.PermissionHandler(),
		},
		"protectionrule": {
			ToolFunc:        tools.ProtectionRule,
			ToolHandler:     tools.ProtectionRuleHandler(),
			ResourceFunc:    resources.ProtectionRule,
			ResourceHandler: resources.ProtectionRuleHandler(),
		},
		"recoveryplan": {
			ToolFunc:        tools.RecoveryPlan,
			ToolHandler:     tools.RecoveryPlanHandler(),
			ResourceFunc:    resources.RecoveryPlan,
			ResourceHandler: resources.RecoveryPlanHandler(),
		},
		"servicegroup": {
			ToolFunc:        tools.ServiceGroup,
			ToolHandler:     tools.ServiceGroupHandler(),
			ResourceFunc:    resources.ServiceGroup,
			ResourceHandler: resources.ServiceGroupHandler(),
		},
		"addressgroup": {
			ToolFunc:        tools.AddressGroup,
			ToolHandler:     tools.AddressGroupHandler(),
			ResourceFunc:    resources.AddressGroup,
			ResourceHandler: resources.AddressGroupHandler(),
		},
		"recoveryplanjob": {
			ToolFunc:        tools.RecoveryPlanJob,
			ToolHandler:     tools.RecoveryPlanJobHandler(),
			ResourceFunc:    resources.RecoveryPlanJob,
			ResourceHandler: resources.RecoveryPlanJobHandler(),
		},
	}

	// Register all tools and resources
	for name, registration := range resourceRegistrations {
		// Add the tool
		s.AddTool(registration.ToolFunc(), registration.ToolHandler)

		// Add the resource
		s.AddResourceTemplate(registration.ResourceFunc(), registration.ResourceHandler)

		if debugMode {
			fmt.Printf("Registered %s resource and tool\n", name)
		}
	}

	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
