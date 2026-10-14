# FreshRSS Fever API

本项目基于 FreshRSS 提供的 Fever API。具体部署可能修改 API 路径，运行前应以实例配置为准。

## Endpoint

默认 endpoint：

```text
https://freshrss.example.net/api/fever.php?api
```

程序接受 FreshRSS 根地址，例如 `https://freshrss.example.net`，并补全为 `/api/fever.php`。如果传入的路径已经以 `/api/fever.php` 结尾，则保留该路径。

## Authentication

每个请求都使用 `POST`，并在 `application/x-www-form-urlencoded` body 中提交：

```text
api_key=md5(username:apiPassword)
```

示例：

```sh
printf '%s' 'alice:replace-me' | md5sum
```

实际使用时不要把真实 password 或 API key 提交到公共日志。

## Authentication response

基础认证响应为 JSON：

```json
{
  "api_version": 3,
  "auth": 1,
  "last_refreshed_on_time": "1520013061"
}
```

字段说明：

- `api_version`：Fever API 版本。
- `auth`：`1` 表示认证成功，`0` 表示认证失败。
- `last_refreshed_on_time`：最近一次刷新时间，通常是 Unix timestamp 字符串。

## 其他 Fever actions

Fever API 常见 action 包括：

- `feeds`
- `groups`
- `items`
- `unread_item_ids`
- `saved_item_ids`
- `mark=item&as=read&id=<id>`
- `mark=item&as=unread&id=<id>`

这些 action 属于 Fever 协议能力；当前 CLI 暴露的是 FreshRSS `auth check`，其他资源命令尚未实现。

## 参考资料

- [FreshRSS Fever API documentation](https://github.com/FreshRSS/FreshRSS/blob/edge/docs/en/developers/06_Fever_API.md)
- [Fever API archive reference](https://web.archive.org/web/20230616124016/https://feedafever.com/api)
