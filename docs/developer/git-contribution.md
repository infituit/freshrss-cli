# Git 协作规范

本文档说明从分支、提交到 Pull Request 的协作要求。Issue 模板和 PR 模板是提交格式的直接入口，本文件补充它们背后的判断规则。

## 基本范围

- 一个分支聚焦一个 Issue 或一个行为变化。
- 不要在功能、缺陷或配置改动中混入无关重构、格式化或依赖升级。
- 公开 Issue 不要披露未经确认的安全漏洞细节；敏感问题应使用私下的安全报告渠道。

## 分支

从最新目标分支创建工作分支，并在提交 PR 前确认目标分支没有产生影响本次改动的变化。分支名应能表达 Issue 或改动主题，例如：

```text
feat/items-command
fix/auth-endpoint
docs/contribution-guide
```

## Commit message

提交信息使用简洁、面向动作的英文，并使用 Conventional Commit 风格的类型前缀：

```text
feat: add items command
fix: handle invalid Fever response
docs: document contribution workflow
refactor: simplify config loading
test: cover endpoint normalization
chore: update lint configuration
```

| 类型       | 用途                   |
| ---------- | ---------------------- |
| `feat`     | 新功能或用户可见能力   |
| `fix`      | 缺陷修复               |
| `docs`     | 仅文档变更             |
| `refactor` | 不改变行为的代码重构   |
| `perf`     | 性能改进               |
| `test`     | 新增或修正测试         |
| `chore`    | 构建、工具链或辅助流程 |

需要关联 Issue 时，在提交信息末尾添加：

```text
fix: handle invalid Fever response --issue=#12
```

Commit message 不能替代 PR 正文，也不能据此推断测试、风险或完整影响范围。

## Issue 关联

提交前先搜索现有 Issue，避免重复创建。

- PR 完成 Issue 的全部范围：使用 `closes #N`、`fixes #N` 或 `resolves #N`。
- PR 只完成阶段性范围：使用 `refs #N` 或 `related to #N`，保持 Issue 开放。
- PR 替换旧 PR：保留 `Replaces #N`，并单独判断 Issue 关联关系。
- PR 标题必须关联至少一个真实 Issue；没有合适的 Issue 时，先创建或补充 Issue，再准备 PR。

## PR 标题

PR 标题必须使用英文、面向动作的描述，并关联至少一个 Issue：

```text
<type>: <English subject> --issue=#N
```

例如：

```text
fix: handle invalid Fever response --issue=#12
feat: add items command --issue=#34
fix: fix multiple related bugs --issue=#100, #101, #102
```

不要把 Issue 模板中的 `[Bug]`、`[Feature]` 或 `[Idea]` 前缀带入 PR 标题。
`--issue=#N` 中的 `N` 必须是正整数；多个 Issue 使用逗号分隔。标题描述和 Issue 标记之间必须保留一个空格。

## PR 正文

使用 `.github/PULL_REQUEST_TEMPLATE.md`，并填写以下五部分：

1. 变更摘要：说明改了什么、为什么改，以及评审重点。
2. 测试结果：只填写当前代码实际运行过的命令和结果。
3. 风险与回滚：说明兼容性、配置、认证或行为风险，以及可执行的回滚方式。
4. 关联 Issue：填写正确的关闭或引用关系，并与标题中的 Issue 编号一致。
5. 涉及范围：说明 CLI、Fever API、配置、工具链或文档是否受影响。

历史 PR 的测试结果不能直接当作当前验证结果；替换 PR 时必须重新验证。

## 按改动范围验证

| 改动范围                          | 最低验证要求                                                |
| --------------------------------- | ----------------------------------------------------------- |
| `cmd/**`、`internal/**`、`pkg/**` | `go test ./...`、`go build ./cmd/freshrss-cli`              |
| Fever API 或认证配置              | 上述命令，以及 HTTP 请求和错误场景的相关测试                |
| `.golangci.yml`、`Makefile`       | `go test ./...`、`go build ./cmd/freshrss-cli`、`make lint` |
| `docs/**`                         | 检查链接、命令和配置示例与当前实现一致                      |
| `.github/**`                      | 检查 YAML/Markdown 结构和模板字段是否适用于当前项目         |

跨模块改动至少说明以下链路是否变化：

```text
CLI command -> config -> Fever client -> HTTP API
```

## 提交前清单

- [ ] 分支基于最新目标分支，且没有无关改动。
- [ ] Commit message 使用正确的英文类型前缀。
- [ ] 已复用匹配的 Issue，并选择正确的关联关系。
- [ ] PR 标题使用英文、面向动作的描述，并符合 `<type>: <English subject> --issue=#N`。
- [ ] PR 标题中的 Issue 编号真实存在，且没有 Issue 模板前缀。
- [ ] 使用了当前 `.github/PULL_REQUEST_TEMPLATE.md`。
- [ ] 测试结果只记录当前代码真实执行过的验证。
- [ ] 已说明风险、回滚和影响范围。
- [ ] 已等待自动检查和代码评审完成。
