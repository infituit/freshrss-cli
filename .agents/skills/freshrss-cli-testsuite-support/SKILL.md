---
name: freshrss-cli-testsuite-support
description: Use when adding or changing real FreshRSS integration tests that use testsuite/support or the project Docker Compose fixture.
---

# freshrss-cli testsuite support

## Overview

Use this skill for integration tests that exercise the real FreshRSS Fever API. The project-owned fixture is intentionally narrow: it starts one disposable FreshRSS instance through `testsuite/docker-compose.yml`, creates the test user with Fever enabled, waits for readiness, and removes the Compose project and volumes after the test.

This skill is specific to `freshrss-cli`; do not copy MongoDB, Redis, or unrelated service helpers from other projects.

## When to Use

Use this skill when a test:

- imports `testsuite/support`;
- needs a real FreshRSS HTTP server or Fever endpoint;
- changes `testsuite/docker-compose.yml`;
- changes the `integration-test` Make target;
- changes FreshRSS integration credentials, port selection, readiness, or cleanup.

Do not use it for:

- unit tests in `pkg/fever` using an HTTP transport stub;
- CLI tests that do not need a real server;
- documentation-only changes unrelated to integration-test execution.

## Source Anchors

Read these before changing the integration path:

- `testsuite/docker-compose.yml`: FreshRSS image, first-run install, Fever enablement, port and healthcheck.
- `testsuite/support/freshrss.go`: fixture lifecycle, temporary Compose environment, project isolation and cleanup.
- `testsuite/fever_integration_test.go`: real client usage and current integration assertion.
- `pkg/fever/client.go`: endpoint construction, API-key derivation and HTTP contract.
- `docs/developer/development.md`: unit-test boundary and integration command.
- `Makefile`: `integration-test` and formatting targets.

## Core Rules

1. Add `//go:build integration` to every test file that starts or imports the real FreshRSS fixture.
2. Keep `go test ./...` free of Docker startup and external-service access.
3. Start FreshRSS through `support.RequireFreshRSS(t)`; do not call `docker`, create HTTP servers, or create credentials directly in package tests.
4. Use the returned `FreshRSS.BaseURL`, `Username`, and `APIPassword`; never hardcode a host port or derive the API key in the test.
5. Keep the Compose project name unique per test and remove containers, volumes, and orphans with `t.Cleanup`.
6. Keep generated credential files in a temporary directory with restrictive permissions. Never commit local credential files or put real credentials in Compose YAML, docs, logs, or test failure output.
7. Pin the FreshRSS image unless a deliberate version change is required and documented.
8. Extend the fixture only when a real integration test needs the capability; do not add speculative services or abstractions.

## Configuration

The fixture uses these process-level settings:

| Setting              | Default | Purpose                                                       |
| -------------------- | ------- | ------------------------------------------------------------- |
| `FRESHRSS_TEST_PORT` | `18080` | Localhost port published by the disposable FreshRSS container |

The fixture writes the remaining Compose variables to a temporary file. The test user is `freshrss`; its API password is fixture-owned and must not be reused outside the test.

If parallel tests need different FreshRSS instances, assign distinct `FRESHRSS_TEST_PORT` values. Do not make the shared default port dynamic without also preserving a deterministic base URL for the FreshRSS installation.

## Test Commands

Default tests:

```bash
go test ./...
```

Integration tests:

```bash
make integration-test
```

Compile the integration path without starting the service:

```bash
go test -tags=integration ./testsuite/... -run '^$' -count=1
```

When Docker is available, run the real test with verbose output:

```bash
go test -tags=integration ./testsuite/... -run '^TestFreshRSSFeverAuthCheck$' -count=1 -v
```

After changing Go files, run `gofmt` through `make fmt` or format the affected `testsuite` paths directly. Run `git diff --check` before delivery.

## Common Mistakes

| Mistake                                                    | Correction                                                                                 |
| ---------------------------------------------------------- | ------------------------------------------------------------------------------------------ |
| Importing `testsuite/support` from a default test          | Add the `integration` build tag.                                                           |
| Starting Docker from `init`, package setup, or a unit test | Start only through `RequireFreshRSS`; the default suite must remain external-service-free. |
| Reusing a fixed Compose project name                       | Use the generated per-test project name so stale containers cannot collide.                |
| Leaving volumes after a test                               | Keep `down --volumes --remove-orphans` in cleanup, including failure paths.                |
| Reimplementing Fever API-key hashing in the test           | Pass credentials to `fever.NewClient`; production client code owns the protocol contract.  |
| Logging generated passwords or the Compose env file        | Return concise errors without secrets and keep the env file under `t.TempDir()`.           |
| Adding a second generic Docker helper                      | Extend this FreshRSS-specific helper only for a demonstrated test need.                    |

## Lightweight Eval Scenarios

1. `我要给 pkg/fever 增加一个真实的 auth-check integration test，应该如何启动和清理 FreshRSS？`
2. `普通 go test ./... 会不会启动 Docker？`
3. `FRESHRSS_TEST_PORT 冲突时应该在哪里配置？`

A good answer loads this skill, uses `support.RequireFreshRSS(t)`, preserves the `integration` tag, keeps credentials temporary, explains cleanup, and gives both default and integration verification commands.
