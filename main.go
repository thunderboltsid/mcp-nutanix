package main

import (
	"fmt"

	"github.com/thunderboltsid/mcp-nutanix/pkg/prompts"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"
	"github.com/thunderboltsid/mcp-nutanix/pkg/tools"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	hooks := &server.Hooks{}
	hooks.AddBeforeAny(func(id any, method mcp.MCPMethod, message any) {
		fmt.Printf("beforeAny: %s, %v, %v\n", method, id, message)
	})
	hooks.AddOnSuccess(func(id any, method mcp.MCPMethod, message any, result any) {
		fmt.Printf("onSuccess: %s, %v, %v, %v\n", method, id, message, result)
	})
	hooks.AddOnError(func(id any, method mcp.MCPMethod, message any, err error) {
		fmt.Printf("onError: %s, %v, %v, %v\n", method, id, message, err)
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

	// Add the handlers
	s.AddTool(tools.Clusters(), tools.ClustersHandler())

	// Add the resources
	s.AddResourceTemplate(resources.Cluster(), resources.ClusterHandler())

	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
