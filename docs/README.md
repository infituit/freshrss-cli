# 文档导航

本文档面向 `freshrss-cli` 的使用者和贡献者，按目标、运行、开发和协议参考组织。

## 文档分区

### Concepts

- [项目目标与边界](concepts/project-goals.md)：项目解决的问题、首阶段范围和非目标。

### Operation

- [使用与配置](operation/usage.md)：安装、配置、运行当前 CLI 命令和常见错误。

### Developer

- [架构设计](developer/architecture.md)：调用链、模块职责和依赖方向。
- [代码目录](developer/code-structure.md)：`cmd`、`internal`、`pkg` 的目录约定。
- [开发指南](developer/development.md)：本地开发、测试和新增功能的基本流程。
- [Git 协作规范](developer/git-contribution.md)：分支、Commit、Issue、PR 和按范围验证要求。

### Reference

- [FreshRSS Fever API](reference/fever-api.md)：本项目依赖的认证方式、请求约定和响应字段。

## 文档约定

- 文档中的“当前”只描述仓库中已经实现并可验证的行为。
- “规划”表示目标方向，不代表已有 CLI 命令或稳定 API。
- 涉及密码、API key 或 token 的示例使用占位值，不提交真实凭据。
