---
name: rss-cli-scoped-agentsmd
description: Use when the user wants to create, regenerate, or compress a scoped (subtree-specific) AGENTS.md inside rss-cli, especially when the scope is given as a directory path or described in natural language.
---

# Scoped AGENTS.md (rss-cli)

Write a **scoped** `AGENTS.md` for a **specific subtree** of the current rss-cli repository.

This skill is not for writing or replacing the repository-wide `AGENTS.md`. It writes one local rule file at a confirmed scope root.

## Output Contract

- Produce a complete `AGENTS.md` at the chosen scope root.
- Use the pipe-index format in `references/AGENTS-compression-guide.md`.
- Prefer single-line `|...` index entries; avoid long prose, traditional `##` sections, and full code blocks.
- Read `references/AGENTS-compression-guide.md` in full before drafting and again before final validation when needed.

## Scope Resolution

- For a path, verify it exists and use it as the scope root.
- For natural language, search for likely directories. If 2-6 plausible roots remain, list each with its distinction and ask the user to choose. Never guess silently.

## Generation Workflow

1. Read the applicable root and parent `AGENTS.md` files.
2. Find nested `AGENTS.md` files and 0-3 useful sibling scopes; avoid duplicating their rules.
3. Inspect the scope's actual files and representative implementation patterns.
4. Extract only scope-specific ownership, where-to-look paths, conventions, anti-patterns, dependencies, and scope-specific commands.
5. Write the file in pipe-index format.
6. Validate paths, format, inheritance, and absence of repository-wide duplication.

## Required Header

Include these near the top:

|IMPORTANT: Prefer retrieval-led reasoning over pre-training-led reasoning
|Required Tools:codegraph (semantic code structure)|context7 (3rd-party docs)|sequential-thinking (decisions)
|Language Policy:Chinese for Q&A|English for code/docs/technical discussions
|Compression Rule:Follow references/AGENTS-compression-guide.md (pipe-index format, concise, no prose/code blocks)
|Scope:<relative/path/from/repo/root>

## Content Rules

- Use `|Where to look:<task>:<path(s)>:<note>` for retrieval anchors.
- Use `|Conventions:...`, `|Anti-patterns:...`, `|Dependencies:...`, and `|Commands:...` only when supported by local evidence.
- Use `|Child AGENTS:<path1>|<path2>|...` for nested scopes with their own rules.
- Preserve all material facts when compressing an existing file; convert prose to dense index entries rather than dropping handler, type-flow, error, or anti-pattern guidance.
- Do not copy the entire repository structure or move static facts from `AGENTS.md`/docs into a skill.

## Validation Checklist

- File is exactly `AGENTS.md` under the confirmed scope root.
- It starts with `|IMPORTANT:` and most lines start with `|`.
- It contains `|Scope:` and only scope-relevant guidance.
- No unverified commands, paths, public-contract claims, credentials, or long explanatory blocks.
- Existing nested rules are referenced rather than duplicated.
