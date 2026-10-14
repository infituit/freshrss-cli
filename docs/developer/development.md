# 开发指南

## 开发前检查

```sh
go version
```

确认未覆盖其他人的工作区变更，并使用当前 Go 版本执行验证。

## 本地验证

格式化：

```sh
gofmt -w ./cmd ./internal ./pkg
```

测试：

```sh
go test ./...
```

构建：

```sh
make build
```

构建产物写入 `build/<VERSION>/`，该目录已加入 Git 忽略规则。`make clean` 会删除整个 `build/`。

跨平台构建：

- `make platform-builds` 构建预设的全部 `GOOS/GOARCH` 组合。
- `make build-matrix GOOS=linux GOARCH=amd64` 构建单个指定组合。
- 交叉构建固定使用 `CGO_ENABLED=0` 和 `-trimpath`，产物名包含目标平台和架构。

版本元数据：

- `pkg/version` 保存 `VERSION`、`BUILDTIME` 和 `GITHASH` 三项构建信息。
- `Makefile` 使用 `-ldflags -X` 注入发布值；默认版本来自 Git tag、日期和当前 commit。
- 不使用 Makefile 时允许直接 `go build`，但得到的是开发默认值。

静态检查：

```sh
make lint
```

`make lint` 需要本地安装 `golangci-lint`。配置文件为仓库根目录的 `.golangci.yml`，使用 v2 格式并显式启用适用于本项目的规则。

## 新增 API 命令

新增 Fever API 命令时遵循以下顺序：

1. 在 `pkg/fever` 增加协议请求、响应类型和单元测试。
2. 在 `internal/cli` 增加命令参数和输出转换。
3. 在 `docs/operation/usage.md` 更新命令状态和示例。
4. 在 `docs/reference/fever-api.md` 补充对应的 Fever API action。
5. 运行 `gofmt`、`go test ./...`、`make build` 和 `make lint`。

不要在 CLI 层重复拼接 Fever API URL 或计算 API key。

## 测试原则

- 不依赖真实 FreshRSS 实例执行单元测试。
- 使用 HTTP transport stub 验证 method、URL、请求体和响应解析。
- 配置错误、非 2xx 响应和非法 JSON 至少各覆盖一个测试场景。
- 涉及写操作的命令必须验证 action 参数，避免默认产生不可逆变更。
