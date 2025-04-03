# dingtalk-mcp
[![smithery badge](https://smithery.ai/badge/@zhaoyunxing92/dingtalk-mcp)](https://smithery.ai/server/@zhaoyunxing92/dingtalk-mcp)
本项目是一个钉钉MCP（Message Connector Protocol）服务，提供了与钉钉企业应用交互的API接口。项目基于Go语言开发，支持员工信息查询和消息发送等功能。

# 安装

### Installing via Smithery

To install DingTalk Message Connector for Claude Desktop automatically via [Smithery](https://smithery.ai/server/@zhaoyunxing92/dingtalk-mcp):

```bash
npx -y @smithery/cli install @zhaoyunxing92/dingtalk-mcp --client claude
```

### Manual Installation
```bash
go install github.com/zhaoyunxing92/dingtalk-mcp@latest
```

## 配置MCP服务

>  [钉钉开放平台](https://open-dev.dingtalk.com) 创建一个应用，并给应用配置权限

```json
{
    "mcpServers": {
       "dingtalk": {
            "command": "dingtalk-mcp", // 如果提示找不到命令，可以将项目编译后的可执行文件放在PATH中
            "args": [],
            "env": {
                "DINGTALK_AGENT_ID": "申请的agentId",
                "DINGTALK_KEY": "应用key",
                "DINGTALK_SECRET": "应用密钥"
            },
            "disabled": false,
            "autoApprove": [
                "get_employees_count",
                "get_simple_employees",
                "recall_corp_conversation",
                "send_corp_conversation",
                "send_markdown_corp_conversation"
            ],
            "timeout": 60
        }
    }
  }
```


## 功能列表

| API名称 | 功能描述 |
|---------|----------|
| get_employees_count | 获取企业员工人数 |
| get_simple_employees | 获取企业的员工基础信息(只获取根部门的人) |
| recall_corp_conversation | 撤回给员工的消息 |
| send_corp_conversation | 企业用户发送文本消息 |
| send_markdown_corp_conversation | 企业用户发送Markdown格式消息 |

