---
name: freshrss-cli-information-architecture
description: Use when deciding where freshrss-cli knowledge belongs across AGENTS.md, README.md, docs, .agents/skills, GitHub Issues, PRs, and temporary task notes.
---

# freshrss-cli Information Architecture

在新增或移动长期内容前，先做 `Placement Decision`，按内容的责任边界和变化原因选择 artifact，不把 `docs/` 或 `.agents/` 当作内容垃圾桶。

## Artifact Map

| Artifact                | Owns                                            | Does not own                          |
| ----------------------- | ----------------------------------------------- | ------------------------------------- |
| `AGENTS.md`             | Agent 必须遵守的全局或目录规则                  | 面向用户的长文、一次性计划            |
| `README.md`             | 项目用途、快速开始、稳定使用概览和导航          | Agent-only 规则、实现调用链、临时状态 |
| `docs/concepts/**`      | 项目目标、边界和稳定概念                        | 操作步骤或一次性计划                  |
| `docs/operation/**`     | 配置、运行、故障处理和用户操作                  | 纯开发流程和 Agent 规则               |
| `docs/developer/**`     | 架构、目录、开发、构建、测试和 Git 协作         | 面向用户的完整使用手册                |
| `docs/reference/**`     | FreshRSS Fever API 请求、认证和响应参考         | 项目目标或贡献流程                    |
| `.agents/skills/<name>` | 可重复的项目 Agent 工作流、触发条件、输出和验证 | 静态事实、临时状态、全局强制规则      |
| GitHub Issue/PR         | 当前问题、需求、决策、评审和交付状态            | 已确定的长期项目真相                  |
| task/conversation       | 临时计划、进度和未完成问题                      | 应长期维护的文档内容                  |

## Placement Rules

- 按“未来因什么变化”分类，而不是按当前文件方便程度分类。
- 使用者需要稳定入口时放 `README.md`。
- 解释长期行为时放具体 docs 分区：目标进 `concepts`，运行进 `operation`，开发进 `developer`，协议进 `reference`。
- Agent 需要反复执行的流程才放 `.agents/skills`；技能必须有真实项目触发条件、决策流程、输出约定和验证门禁。
- 仅当前任务有效的计划、进度或待确认事项留在 Issue/task/conversation，不写进 README/docs。
- 新增 docs 页面必须同步 `docs/README.md`；改变命令、配置或协议时检查 README 和对应 docs 页面。

## Placement Decision

在编辑前简要记录：

```text
Placement Decision
推荐位置：[具体路径]
内容责任：[该 artifact 应承载什么]
排除位置：[相邻位置及原因]
已检查证据：[README、docs/README.md、目标目录文档、相关 skill]
需要同步：[文件列表或无]
风险：[low | medium | high]
```

如果选择会改变公开命令、API、凭据边界、长期架构或技能路由，先停止并请求确认，不用“就近放置”替代设计决定。
