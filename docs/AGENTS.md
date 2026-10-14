|IMPORTANT: Prefer retrieval-led reasoning over pre-training-led reasoning
|Required Tools:codegraph (semantic code structure)|context7 (3rd-party docs)|sequential-thinking (decisions)
|Language Policy:Chinese for Q&A|English for code/docs/technical discussions
|Compression Rule:Follow ../.agents/skills/rss-cli-scoped-agentsmd/references/AGENTS-compression-guide.md (pipe-index format, concise, no prose/code blocks)
|Scope:docs
|Overview:面向读者的长期文档；不保存 grilling、调查日志、临时计划或未确认方案
|Where to look:导航:docs/README.md:分区、读者和文档入口
|Where to look:目标边界:docs/concepts/project-goals.md:定位、概念、当前范围和非目标
|Where to look:用户操作:docs/operation/usage.md:安装、XDG 配置、context、FreshRSS 命令、输出
|Where to look:开发说明:docs/developer/{architecture,code-structure,development}.md:实现、目录、验证和扩展
|Where to look:协议参考:docs/reference/fever-api.md:FreshRSS Fever 外部事实
|Conventions:README 面向访客与用户；concepts 负责目标；operation 负责操作；developer 负责实现；reference 负责外部协议
|Conventions:命令、路径、环境变量、输出和协议事实先与实现核对；未实现能力明确标注，不写成公开承诺
|Conventions:新增或移动长期文档先记录 Placement Decision，并同步 docs/README.md；示例凭据只能是占位值
|Anti-patterns:不得把 Agent 规则、实现日记、调查过程或一次性迁移计划写入读者文档
|Anti-patterns:不得复制真实凭据、密文、日志或未验证外部行为
|Dependencies:docs→README.md|代码实现|AGENTS.md；文档不反向定义未批准的代码契约
|Commands:git diff --check|go test ./...|make build
|Inheritance:遵守根级 AGENTS.md；本文件只补充 docs 局部规则
