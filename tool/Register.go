package tool

import "github.com/mark3labs/mcp-go/server"

// 工具注册
func RegisterAllTools(s *server.MCPServer) {
	s.AddTool(calculatorTool(), handleCalculator)
}
