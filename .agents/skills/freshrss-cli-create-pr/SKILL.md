---
name: freshrss-cli-create-pr
description: Use when preparing, creating, updating, replacing, or splitting a GitHub Pull Request for freshrss-cli, including CLI commands, Fever API, credentials, build tooling, tests, and documentation changes.
---

# Create a freshrss-cli Pull Request

This is the project delivery workflow for `freshrss-cli`. It binds PR scope, validation, Issue relations, and remote-write safety to this repository's CLI, Fever API, configuration, tooling, and documentation contracts.

## Operating Modes

| User wording                             | Mode          | Default remote behavior                                 |
| ---------------------------------------- | ------------- | ------------------------------------------------------- |
| “准备 PR”“整理 PR”“draft”                | `preflight`   | local inspection and draft only                         |
| “创建/提交 PR”                           | `create`      | draft first; confirm each remote write separately       |
| “更新 PR”“修标题/描述”                   | `update`      | read existing PR; confirm the requested edit            |
| “替代旧 PR”                              | `replacement` | preserve old facts and revalidate current HEAD          |
| independent behaviors or staged delivery | `split`       | keep PRs reviewable and use non-closing Issue relations |

Ambiguous wording defaults to `preflight`. Reviewing a draft never authorizes push, Issue creation, PR creation, labels, reviewers, or milestones.

## 1. Determine Scope

Read before drafting:

- `docs/developer/git-contribution.md`
- `.github/PULL_REQUEST_TEMPLATE.md`
- `.github/ISSUE_TEMPLATE/bug_report.yml`
- `.github/ISSUE_TEMPLATE/feature_request.yml`
- `.github/ISSUE_TEMPLATE/discussion.yml`
- `.github/ISSUE_TEMPLATE/config.yml`
- `.github/workflows/pr-title-lint.yml`
- `.golangci.yml` and `Makefile` when tooling is touched

Record branch, target base, HEAD, commits, changed paths, Issue/PR references, and unrelated dirty-worktree changes. Do not include unrelated changes in PR scope.

Classify affected areas using the project template: CLI command or output; FreshRSS Fever API contract; authentication or configuration; build, lint, or test tooling; documentation only.

For cross-layer changes, trace:

```text
CLI command -> config -> Fever client -> HTTP API
```

## 2. Issue Relation

Search for an existing Issue before proposing a duplicate. Every PR title must include at least one real Issue using `--issue=#N`; do not invent Issue numbers. Use `closes`, `fixes`, or `resolves` only when the PR completes the Issue; use `refs` or `related to` for staged work; preserve `Replaces` for a replacement PR. Never invent reproduction, credentials, impact, acceptance, or environment facts.

## 3. Verification Matrix

| Changed paths                     | Required validation                                       |
| --------------------------------- | --------------------------------------------------------- |
| `cmd/**`, `internal/**`, `pkg/**` | `go test ./...`; `go build ./cmd/freshrss-cli`            |
| Fever API or auth/config          | above plus relevant HTTP/error tests                      |
| `.golangci.yml`, `Makefile`       | above plus `make lint`                                    |
| `docs/**`                         | links, commands, and config examples match implementation |
| `.github/**`                      | YAML/Markdown structure and project fields                |

Only check commands that ran on the current HEAD. If `golangci-lint` is unavailable, record that exact limitation instead of claiming `make lint` passed.

## 4. Draft The PR

Start from `.github/PULL_REQUEST_TEMPLATE.md` and fill its five sections: summary, current test results, risks and rollback, related Issues, and affected areas. Include Fever compatibility, credential, output, build, or documentation risks when applicable.

Use an English action-oriented title in `<type>: <subject> --issue=#N` form, such as `feat: add items command --issue=#12` or `fix: handle invalid Fever response --issue=#34`. For multiple Issues, use `--issue=#12, #34`. Validate the title against `.github/workflows/pr-title-lint.yml` before drafting. Do not copy Issue form prefixes into the PR title.

## 5. Remote Writes

Authorize separately: branch push; Issue create/edit; PR create/edit; labels/reviewers/assignees/milestones. Use explicit repository, base, head, and body-file arguments with `gh`. Read resources back after writes; after an ambiguous failure, query state before retrying.

For `preflight`, finish with `No GitHub writes performed` and list pending actions. Delivery reports must include mode, scope/base/head, title, Issue relation, affected areas, actual checks, unverified checks, risks/rollback, and remote actions.
