# 代码目录

当前仓库是单一 Go CLI，目录按 `cmd`、`internal`、`pkg` 分层，不设置多服务目录。

```text
freshrss-cli/
├── cmd/
│   └── freshrss-cli/
│       └── main.go          # 进程入口
├── internal/
│   ├── cli/
│   │   └── cli.go           # 命令分发和 CLI 输出
│   └── config/
│       └── config.go        # 环境变量合并
├── pkg/
│   ├── fever/
│   │   ├── client.go        # Fever API client
│   │   └── client_test.go   # client 单元测试
│   └── version/
│       ├── version.go       # 构建版本元数据和格式化
│       └── version_test.go  # 版本输出测试
├── docs/                    # 项目文档
├── go.mod
├── README.md
└── ...
```

## 放置规则

- 新增可执行程序放入 `cmd/<name>`。
- 只服务于本项目 CLI 的命令和配置放入 `internal`。
- 已经证明具有跨命令复用价值的 Fever API 或领域能力放入 `pkg`。
- 构建版本元数据放入 `pkg/version`，由构建脚本统一注入和读取。
- 用户可见的使用方式、协议说明和开发约定放入 `docs`。
- 不要为了预留未来功能创建空目录或空抽象。

## 命名规则

- Go package 使用简短、小写、单一职责的名称。
- 命令名称使用 kebab-case，例如 `auth-check`。
- 环境变量使用 `FRESHRSS_` 前缀。
- 对外暴露的类型和函数必须有清晰的 Go doc；内部实现不强制导出。
