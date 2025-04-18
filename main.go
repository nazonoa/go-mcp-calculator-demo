package main

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/server"
	"go-mcp-calculator-demo/config"
	"go-mcp-calculator-demo/tool"
	"log"
	"net/http"
	"os/signal"
	"syscall"
)

func main() {

	fmt.Println("mcp服务启动:" + config.SSE_ADDRESS)

	// 新建mcp server
	s := server.NewMCPServer(
		"go-mcp-calculator-demo",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
		server.WithRecovery(),
	)

	//注册工具
	tool.RegisterAllTools(s)

	//SSE发送
	if err := serveSSE(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}

	/*
		//Stdio发送
		if err := server.ServeStdio(s); err != nil {
			fmt.Printf("Server error: %v\n", err)
		}
	*/

}

func serveSSE(s *server.MCPServer) error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := server.NewSSEServer(s)

	mux := http.NewServeMux()

	mux.Handle("/sse", authMiddleware(srv))

	mux.Handle("/message", srv)

	httpServer := &http.Server{
		Addr:    config.SSE_ADDRESS,
		Handler: mux,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	return httpServer.Shutdown(context.Background())
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		//鉴权
		if key != config.API_KEY {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
