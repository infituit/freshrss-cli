# 架构设计

## 总体调用链

```text
命令行参数
    |
    v
cmd/freshrss-cli/main.go
    |
    v
internal/cli.Run
    |--------------------|
    v                    v
internal/config      pkg/fever.Client
                         |
                         v
                    FreshRSS Fever API
```

## 模块职责

### `cmd/freshrss-cli`

进程入口只负责获取命令行参数、创建标准输出/错误输出，并调用 `internal/cli.Run`。这里不放业务逻辑。

### `internal/cli`

负责命令分发、flag 定义和 CLI 输出。它可以依赖 `internal/config` 与 `pkg/fever`，但 API 协议细节应留在 `pkg/fever`。

### `internal/config`

负责把环境变量合并到 CLI 配置。配置结构是当前应用的内部实现，不作为公共库 API 暴露。

### `pkg/fever`

负责 Fever API endpoint 处理、API key 生成、HTTP POST、响应 JSON 解码和协议错误。当前它接收 `internal/config.Config`，但不依赖 `internal/cli`；如果未来需要将 Fever client 提供给其他模块，再评估是否把配置类型移到更低层。

## 依赖方向

```text
cmd -> internal/cli -> internal/config
                   \-> pkg/fever -> internal/config.Config
```

禁止 `pkg` 反向依赖 `cmd` 或 `internal/cli`。如果某个类型确实需要跨多个应用复用，优先放在 `pkg`，否则保持在 `internal`。当前 `pkg/fever` 对 `internal/config` 的依赖是已知的基础框架简化，后续扩展前应避免继续扩大该耦合。

## 错误处理

- 配置错误在创建 Fever client 前返回。
- HTTP 非 2xx 响应转换为包含状态码的错误。
- JSON 解码失败保留 `decode Fever API response` 上下文。
- CLI 入口把错误写入 stderr，并以非零状态退出。
