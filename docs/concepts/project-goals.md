# 项目目标与边界

## 项目定位

`rss-cli` 是一个面向个人重度用户和 AI agent 的 RSS 服务命令行工具。它提供统一的命令行入口、context 管理、结构化输出和服务能力发现，同时保留不同 RSS 应用的实际差异。

FreshRSS 是当前首个服务命令。后续服务接入时，每种服务通过自己的顶层子命令提供能力，不要求所有服务使用相同的业务资源模型。

## 核心概念

- **context**：一套工作环境，保存多个服务配置和 CLI 偏好。
- **service**：context 内的一种 RSS 服务类型，也是 CLI 的顶层子命令。
- **provider**：service 对应的服务适配实现。
- **credentials**：某个 context 内某个 service 的唯一认证信息。

一个 context 可以只配置部分服务；同一 context 内同一 service 最多一份认证。需要使用同一服务的不同环境时，应创建不同 context。

## 当前范围

- `context` 的创建、列出、切换、查看和删除。
- FreshRSS 服务配置和加密认证凭据管理。
- FreshRSS Fever API 认证检查。
- JSON 默认输出、table 输出和结构化 capabilities。
- XDG 配置目录和 `RSS_CLI_CONFIG_DIR` 覆盖。

## 非目标

- 不直接访问 RSS 服务数据库。
- 不把不同服务强制映射为一个完整的通用业务模型。
- 不在尚无复用证据时提前创建跨服务领域抽象。
- 不实现完整 TUI 阅读器。
- 不兼容旧的 `freshrss-cli` 命令、配置或环境变量。
