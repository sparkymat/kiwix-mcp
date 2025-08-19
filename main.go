package main

import (
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/sparkymat/kiwix-mcp/internal/config"
	"github.com/sparkymat/kiwix-mcp/internal/handler"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	s := server.NewMCPServer(
		"kiwix-mcp",
		"1.0.0",
		server.WithToolCapabilities(false),
	)

	// Add tool
	tool := mcp.NewTool("fetch_article",
		mcp.WithDescription("Fetch article on requested topic"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name of the topic to fetch"),
		),
	)

	// Add tool handler
	s.AddTool(tool, handler.Fetch(cfg))

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
