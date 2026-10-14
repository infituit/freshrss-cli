# 架构设计

## 运行链路

```text
cmd/rss-cli/main.go
        |
        v
internal/cli
   |          |
   v          v
internal/  internal/
config     credential/output
   |
   v
pkg/thirdparty/freshrss
        |
        v
FreshRSS Fever API
```

## 模块职责

### `cmd/rss-cli`

只负责进程入口、标准 I/O 和退出码，不放业务逻辑、配置解析或 HTTP 请求。

### `internal/cli`

负责全局参数、context 命令、服务命令、帮助和 capabilities。它把服务配置和凭据组装成 provider client，并将结果交给统一输出层。

### `internal/config`

负责 XDG 配置目录解析、context/service YAML 文件、当前 context 和严格字段校验。context 管理只理解环境和服务目录，不理解 provider 专属认证字段。

### `internal/credential`

负责 `credentials.enc` 的应用层加密、解密和文件权限。凭据不进入普通 YAML、日志或输出。

### `internal/output`

负责 JSON 默认输出、table 输出和 stderr 错误 envelope。

### `pkg/thirdparty/freshrss`

负责 FreshRSS Fever endpoint、API key 生成、POST 表单、响应 JSON 和协议错误。该包不依赖 CLI、context 或输出实现。

## 依赖方向

```text
cmd -> internal/cli
internal/cli -> internal/config
internal/cli -> internal/credential
internal/cli -> internal/output
internal/cli -> pkg/thirdparty/freshrss
```

`pkg/thirdparty` 不得反向依赖 `cmd` 或 `internal/cli`。不同服务的协议差异保留在各自 provider 包中；只有出现真实复用证据时才新增跨服务领域抽象。

## 错误处理

- 业务命令返回 Go error，由入口转换为机器可读 JSON 并写入 stderr。
- 成功结果只写入 stdout。
- provider 负责保留 HTTP 状态和协议解析上下文。
- 凭据解密失败不得回显密文或认证内容。
