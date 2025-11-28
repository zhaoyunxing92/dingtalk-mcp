package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/mark3labs/mcp-go/server"

	"github.com/zhaoyunxing92/dingtalk-mcp/internal/service"
	"github.com/zhaoyunxing92/dingtalk-mcp/pkg/dingtalk"
)

// Parse flags
var (
	// 定义指针变量
	idFlag     = flag.Int("id", -1, "钉钉小程序agentId")
	keyFlag    = flag.String("key", "", "钉钉小程序key")
	secretFlag = flag.String("secret", "", "钉钉小程序secret")
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered in main", err)
		}
	}()
	flag.Parse()

	var (
		id     = *idFlag
		key    = *keyFlag
		secret = *secretFlag
	)

	if id == -1 {
		if envAgentId := os.Getenv("DINGTALK_AGENT_ID"); envAgentId != "" {
			if eId, err := strconv.Atoi(envAgentId); err == nil {
				id = eId
			} else {
				fmt.Println("agentId must be set")
				os.Exit(1)
			}
		}
	}

	if key == "" {
		key = os.Getenv("DINGTALK_KEY")
	}
	if secret == "" {
		secret = os.Getenv("DINGTALK_SECRET")
	}
	// 在解析参数后解引用指针
	client := dingtalk.NewClient(id, key, secret)

	// Create MCP server
	svc := server.NewMCPServer(
		"dingtalk",
		"1.0.0",
		server.WithLogging(),
		server.WithResourceCapabilities(true, true),
	)
	service.NewEmployeeService(client).AddTools(svc)
	service.NewMessageService(client).AddTools(svc)

	// sse model
	//sse := server.NewSSEServer(svc)
	//if err := sse.Start(":8100"); err != nil {
	//	fmt.Println("SSE Server error", err)
	//}

	// stdio model
	if err := server.ServeStdio(svc); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
