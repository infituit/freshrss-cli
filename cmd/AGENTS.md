|IMPORTANT: Prefer retrieval-led reasoning over pre-training-led reasoning
|Required Tools:codegraph (semantic code structure)|context7 (3rd-party docs)|sequential-thinking (decisions)
|Language Policy:Chinese for Q&A|English for code/docs/technical discussions
|Compression Rule:Follow ../.agents/skills/rss-cli-scoped-agentsmd/references/AGENTS-compression-guide.md (pipe-index format, concise, no prose/code blocks)
|Scope:cmd
|Overview:进程入口与依赖装配；业务、配置、协议逻辑分别回到 internal 或 pkg
|Where to look:入口行为:cmd/rss-cli/main.go:标准 I/O、错误输出、退出码
|Where to look:命令实现:internal/cli/cli.go:全局参数、context/service 命令、输出组装
|Where to look:用户契约:README.md|docs/operation/usage.md:命令、flag、示例同步
|Conventions:main.go 只调用 internal/cli 并负责进程级退出；新服务只有在已有真实能力时添加顶层命令
|Conventions:命令成功结果走 stdout，错误走统一机器可读 stderr；命令/flag 变更同步 README、docs/operation 与测试
|Anti-patterns:不得在 cmd 实现业务逻辑、配置解析、凭据处理或第三方 HTTP 请求
|Anti-patterns:不得为未来 provider 创建空命令、空目录或未经批准的统一业务抽象
|Dependencies:cmd/rss-cli/main.go→internal/cli→internal/config|internal/credential|internal/output|pkg/thirdparty
|Commands:go test ./...|go build ./cmd/rss-cli|make build
|Inheritance:遵守根级 AGENTS.md；本文件只补充 cmd 局部规则
