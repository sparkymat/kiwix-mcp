package handler

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

var ErrRequestFailed = errors.New("request failed")

type Config interface {
	BaseURL() string
}

func Fetch(cfg Config) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name, err := request.RequireString("name")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		namePath := strings.ReplaceAll(name, " ", "_")

		url := cfg.BaseURL() + namePath

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			return nil, ErrRequestFailed
		}

		defer resp.Body.Close()

		respBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return mcp.NewToolResultText(string(respBytes)), nil
	}
}
