---
name: freshrss-cli-audience-boundary
description: Use when writing or reviewing freshrss-cli README files, usage guides, developer docs, Fever API references, Issues, Pull Requests, or other reader-facing artifacts.
---

# freshrss-cli Audience Boundary

## Reader And Artifact

先确定最终读者、读者要完成的动作和目标文件：

| Artifact            | Reader                   | Required decision or action      |
| ------------------- | ------------------------ | -------------------------------- |
| `README.md`         | CLI 使用者和项目访客     | 判断用途并完成最短启动路径       |
| `docs/operation/**` | CLI 使用者               | 配置 FreshRSS 并运行命令         |
| `docs/developer/**` | 贡献者                   | 修改、构建、测试和提交代码       |
| `docs/reference/**` | API 集成或 client 开发者 | 理解 Fever 请求、认证和响应约定  |
| `docs/concepts/**`  | 项目维护者和贡献者       | 理解目标、边界和长期方向         |
| GitHub Issue        | 维护者和贡献者           | 判断问题、需求或讨论的范围       |
| GitHub PR           | Reviewer 和维护者        | 判断行为、风险、验证和是否可合并 |
| `.agents/skills/**` | Agent                    | 执行可复用的项目工作流           |

## Content Boundary

- README 只放稳定入口、用途、快速开始和导航，不放 Agent 指令、临时状态或实现日记。
- `docs/**` 放长期有效的目标、操作、开发和 Fever API 说明，不放临时计划。
- `.agents/skills/**` 只放可重复的 Agent 工作流，不把静态协议事实或一次性笔记伪装成 skill。
- Issue/PR 写读者需要的事实、影响、风险、验证和行动，不写作者调查过程。
- 未确认的行为、版本、兼容性或安全结论不能写成公开承诺。
- 密码、API key 和 token 只能使用占位值。

## Delivery Check

交付前按目标读者通读文档，确认命令、路径、配置名和 Fever API 事实与当前实现一致；实现、命令或配置变更应同步检查 README 和对应 docs 分区。
