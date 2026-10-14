# 使用与配置

## 环境要求

- Go 1.25 或兼容的更高版本。
- 可访问 FreshRSS Fever API 的网络环境。
- FreshRSS 用户名和 API password，或已经计算好的 Fever API key。

## 构建与运行

直接运行：

```sh
go run ./cmd/freshrss-cli help
go run ./cmd/freshrss-cli version
```

构建二进制：

```sh
make build VERSION=v0.1.0-alpha.1
./build/v0.1.0-alpha.1/freshrss-cli help
```

`make build` 会把 Git 标记、构建时间和短 commit hash 注入二进制，并将产物写入 `build/<VERSION>/`。也可以覆盖版本号：

```sh
make build VERSION=v0.1.0-alpha.1
./build/v0.1.0-alpha.1/freshrss-cli version
```

直接使用 `go build` 时不会注入发布元数据，`version` 命令会显示 `debug`、`unknown` 等默认值。

跨平台构建：

```sh
make platform-builds VERSION=v0.1.0-alpha.1
make build-matrix VERSION=v0.1.0-alpha.1 GOOS=linux GOARCH=arm64
```

全平台产物使用 `<name>_<GOOS>_<GOARCH>` 命名，Windows 产物增加 `.exe` 后缀，统一输出到 `build/<VERSION>/`。当前矩阵包括 Linux `amd64`、`arm64`、`arm`，以及 Darwin 和 Windows 的 `amd64`、`arm64`。

UPX 压缩默认关闭。启用后，构建完成会自动压缩产物：

```sh
make build VERSION=v0.1.0-alpha.1 UPX_ENABLED=1
make platform-builds VERSION=v0.1.0-alpha.1 UPX_ENABLED=1 UPX_LEVEL=9
```

项目通过 `mise.toml` 固定 UPX 版本，启用压缩前运行 `mise install`。UPX 压缩支持 Linux `amd64`、`arm64`、`arm` 和 Windows `amd64`；Darwin 以及 Windows `arm64` 会跳过压缩。缺少 `upx` 时构建会直接失败。

## 配置方式

### 使用用户名和 API password

```sh
FRESHRSS_URL="https://freshrss.example.net" \
FRESHRSS_USER="alice" \
FRESHRSS_API_PASSWORD="replace-me" \
go run ./cmd/freshrss-cli auth-check
```

### 使用预计算 API key

```sh
FRESHRSS_URL="https://freshrss.example.net" \
FRESHRSS_API_KEY="replace-me" \
go run ./cmd/freshrss-cli auth-check
```

配置优先级为：命令行 flag 高于同名环境变量。

| 配置          | Flag            | 环境变量                | 说明                                    |
| ------------- | --------------- | ----------------------- | --------------------------------------- |
| FreshRSS 地址 | `-url`          | `FRESHRSS_URL`          | 可以是站点根地址或完整 Fever endpoint   |
| 用户名        | `-user`         | `FRESHRSS_USER`         | 与 API password 一起计算 API key        |
| API password  | `-api-password` | `FRESHRSS_API_PASSWORD` | FreshRSS API password                   |
| API key       | `-api-key`      | `FRESHRSS_API_KEY`      | 已计算好的 key，优先于用户名和 password |

## 当前命令

| 命令                       | 状态   | 说明                              |
| -------------------------- | ------ | --------------------------------- |
| `help`                     | 可用   | 显示命令帮助                      |
| `version`                  | 可用   | 显示版本、构建时间和 Git hash     |
| `auth-check`               | 可用   | 请求 Fever API 并检查 `auth` 响应 |
| `feeds`、`groups`、`items` | 规划中 | 尚未实现                          |

## 常见错误

- `FreshRSS URL is required`：设置 `-url` 或 `FRESHRSS_URL`。
- `FreshRSS URL must include scheme and host`：URL 必须包含 `https://` 或 `http://` 以及主机名。
- `Fever API key or username and API password are required`：补充 API key，或同时补充用户名和 API password。
- `authentication failed`：检查 API endpoint、用户名、API password 和 FreshRSS Fever API 配置。

不要把真实 API password 或 API key 写入 shell 历史、脚本仓库或 issue。
