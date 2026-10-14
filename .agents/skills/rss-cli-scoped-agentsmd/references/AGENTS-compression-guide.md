# AGENTS.md Pipe-Index Compression Guide

|Purpose:Compress scoped AGENTS.md into a retrieval-oriented index without losing verified rules
|Format:Most lines begin with `|`; use `|Category:path:{items}` and `|` for alternatives
|Required:Start with `|IMPORTANT: Prefer retrieval-led reasoning over pre-training-led reasoning`
|Scope:Include a relative `|Scope:<path>` line and keep facts local to that subtree
|Structure:Use short index categories such as `{Overview,Structure,Where to look,Conventions,Anti-patterns,Dependencies,Child AGENTS,Commands}`
|Relations:Use `→` for flows and `{}` for grouped files; link to source files instead of embedding implementation details
|Avoid:Long prose, `##` sections as the main structure, full code blocks, duplicated repository-wide rules, unverified paths or commands
|Density:Prefer roughly 30-40 lines and 5-8KB or less; merge repeated categories and file prefixes
|Maintenance:Re-check paths and rules after architecture changes or when agents repeatedly violate a rule
|Validation:Check the header, scope, pipe-index format, path validity, inheritance, and preservation of source facts when compressing

## Compression Procedure

1. Read the existing scoped file and applicable parent rules.
2. Keep ownership, retrieval anchors, conventions, anti-patterns, dependencies, child scopes, and local commands.
3. Merge repeated file lists and replace explanations with indexed facts.
4. Remove stale, duplicated, or unverified content.
5. Verify every referenced path and command against the repository before delivery.
