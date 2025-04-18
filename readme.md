# go mcp-server demo

### 直接启动即可，默认使用sse模式

1. cline 配置：

```json
{
  "mcpServers": {
    "calculator": {
      "autoApprove": [
        "calculate"
      ],
      "timeout": 60,
      "url": "http://localhost:8087/sse?key=1234",
      "transportType": "sse"
    }
  }
}
```
2. cursor 配置:
```json
{
  "mcpServers": {
    "calculator": {
      "url": "http://localhost:8087/sse?key=1234"
    }
  }
}
```