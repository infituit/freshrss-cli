# 项目目标与边界

## 项目定位

`freshrss-cli` 是使用 Go 编写的 FreshRSS Fever API 命令行工具。它的主要使用场景是：

- 在终端检查 FreshRSS Fever API 是否可用。
- 通过脚本读取订阅、分组和文章数据。
- 用统一的命令行接口完成常见阅读操作。

项目优先考虑可脚本化、依赖少和行为可预测，不以完整 TUI 阅读器为目标。

## 首阶段目标

当前已经具备：

- Go CLI 入口和基础命令分发。
- `help`、`version`、`auth-check` 命令。
- FreshRSS URL、用户名、API password 和预计算 API key 配置。
- Fever API 认证请求和响应解析。

下一阶段计划：

- 读取 feeds、groups、items。
- 读取 unread/saved item IDs。
- 支持标记文章为 read/unread。
- 增加适合脚本处理的 JSON 输出选项。

## 非目标

- 不直接访问 FreshRSS 数据库。
- 不复制 FreshRSS Web UI 的全部功能。
- 不在没有实际复用需求时扩展 `pkg` 公共 API。
- 不为了目录形式引入多服务部署结构。

## 设计原则

1. **标准库优先**：基础 CLI 和 HTTP 能力优先使用 Go 标准库。
2. **协议边界清晰**：Fever API 的 HTTP 细节集中在 `pkg/fever`。
3. **配置显式**：敏感配置从 flag 或环境变量进入，不写入仓库文件。
4. **增量扩展**：先完成可验证的单个命令，再增加资源和写操作。
