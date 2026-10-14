# 版本与构建规则

本文面向项目维护者和贡献者，说明发布版本、构建版本以及二进制内版本元数据的关系。

## 发布版本

发布版本使用带预发布阶段和序号的版本格式：

```text
v{major}.{minor}.{patch}-{phase}.{serial}[-{mark}]
```

例如：

```text
v0.1.0-alpha.1
v0.1.0-beta.2-test
```

字段含义：

- `major.minor.patch`：基础版本号。
- `phase`：发布阶段，例如 `alpha`、`beta` 或 `release`。
- `serial`：同一基础版本和发布阶段内递增的序号。
- `mark`：可选附加标记，用于测试或其他发布说明。

Git Tag 应直接使用完整发布版本，并与对外发布的版本保持一致。正式版本不使用 `alpha`、`beta` 或 `test` 等预发布标记。

当前项目不会在 Makefile 中自动递增或校验发布版本；创建 Tag 和递增版本序号由发布者负责。

## 默认构建版本

`Makefile` 的 `VERSION` 默认由 Git 描述和构建日期组成：

```make
GITTAG := $(shell git describe --tags --always --dirty 2>/dev/null || printf 'v0.0.0')
VERSION ?= $(GITTAG)-$(shell date +%y.%m.%d)
```

因此，带有发布 Tag 的默认构建版本可能类似于：

```text
v0.1.0-alpha.1-26.10.14
```

其中：

- `GITTAG` 优先取最近的 Git Tag；没有 Tag 时使用 commit 描述；无法读取 Git 信息时使用 `v0.0.0`。
- `-YY.MM.DD` 是构建日期后缀，用于区分当天的构建产物，不代表新的发布版本。
- 工作区存在未提交修改时，`git describe` 可能在版本标识中包含 `-dirty`。
- 通过 `make build VERSION=<version>` 可以显式覆盖默认构建版本。

## 构建元数据

构建时通过 Go `-ldflags -X` 将以下值注入 `pkg/version`：

- `VERSION`：本次构建的版本标识。
- `BUILDTIME`：构建时间，格式为 `YYYY-MM-DDTHH:MM:SS+TZ`。
- `GITHASH`：当前 Git commit 的短 hash。

直接执行 `go build` 不会执行 Makefile 的注入逻辑，`version` 命令会显示 `debug` 和 `unknown` 等开发默认值。

## 构建产物

所有构建产物按版本隔离到：

```text
build/<VERSION>/
```

常用命令：

```sh
make build VERSION=v0.1.0-alpha.1
make platform-builds VERSION=v0.1.0-alpha.1
make build-matrix VERSION=v0.1.0-alpha.1 GOOS=linux GOARCH=amd64
```

构建目录已加入 Git 忽略规则。`make clean` 会删除整个 `build/` 目录。
