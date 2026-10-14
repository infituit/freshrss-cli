# freshrss-cli

FreshRSS Fever API 的命令行工具，用于在终端和脚本中读取、检查和管理 FreshRSS 内容。

## Project Goal

`freshrss-cli` 的目标是提供一个轻量、可脚本化的 Go CLI，围绕 FreshRSS 暴露的 Fever API 完成常见阅读器操作。

首阶段聚焦：

- 连接 FreshRSS Fever API 并验证认证信息。
- 读取 feeds、groups、items、unread/saved item IDs。
- 支持按脚本友好的格式输出数据。
- 提供 mark read/unread 等基础写操作。

非目标：

- 不实现完整 TUI 阅读器。
- 不直接访问 FreshRSS 数据库。
- 不在基础阶段绑定重量级 CLI 框架。

## FreshRSS Fever API

FreshRSS 的 Fever API 通常位于：

```text
https://freshrss.example.net/api/fever.php?api
```

所有请求使用 `POST`，并提交表单字段 `api_key`。`api_key` 是 `username:apiPassword` 的 MD5 值，其中 `apiPassword` 是 FreshRSS API password，不一定等同于网页登录密码。

## Current CLI

当前骨架提供最小可运行命令：

```sh
go run ./cmd/freshrss-cli help
go run ./cmd/freshrss-cli version
go run ./cmd/freshrss-cli auth-check
```

配置可以通过 flag 或环境变量提供：

```sh
FRESHRSS_URL="https://freshrss.example.net" \
FRESHRSS_USER="alice" \
FRESHRSS_API_PASSWORD="secret" \
go run ./cmd/freshrss-cli auth-check
```

也可以直接提供已计算好的 API key：

```sh
FRESHRSS_URL="https://freshrss.example.net" \
FRESHRSS_API_KEY="md5-value" \
go run ./cmd/freshrss-cli auth-check
```

可用 flag：

- `-url`: FreshRSS base URL 或完整 Fever endpoint。
- `-user`: FreshRSS username。
- `-api-password`: FreshRSS API password。
- `-api-key`: 已计算好的 Fever API key。

## Development

```sh
go test ./...
```

使用 `golangci-lint` 执行静态检查：

```sh
make lint
```

lint 配置只启用适用于本项目的规则，避免引入与当前 CLI 无关的检查项。

## Documentation

- [文档导航](docs/README.md)
- [项目目标与边界](docs/concepts/project-goals.md)
- [使用与配置](docs/operation/usage.md)
- [架构设计](docs/developer/architecture.md)
- [代码目录](docs/developer/code-structure.md)
- [开发指南](docs/developer/development.md)
- [Git 协作规范](docs/developer/git-contribution.md)
- [FreshRSS Fever API 参考](docs/reference/fever-api.md)

项目结构：

```text
cmd/freshrss-cli/   CLI 入口
internal/cli/       命令分发
internal/config/    配置读取和校验
pkg/fever/          Fever API client，可复用的 Fever API 能力
```

目录分层约定为：`cmd` 负责启动入口和依赖装配，`internal` 保存项目私有实现，`pkg` 保存可复用的公共能力。当前项目是单一 CLI，因此不会为了形式增加多服务目录。
