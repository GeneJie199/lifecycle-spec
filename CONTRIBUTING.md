# Contributing to lifecycle-spec

**专业名称：** Contributing Guide  
**通俗解释：** 想改协议或提问题前，先看这里，避免四个工具突然说不同的话。

## 报告 Bug

请开 Issue，并尽量包含：

1. 相关 Schema / 示例文件路径  
2. 期望行为与实际行为  
3. 用于复现的 JSON（**打码所有密钥**）

## 功能或协议提案

1. 先开 Issue 讨论场景与兼容影响  
2. 若涉及 Schema 字段增删改：先提交 `docs/adr/NNNN-*.md`  
3. 同步更新：Schema、`examples/`、`gen/` 类型、`go test ./...`

## 代码贡献流程

1. Fork + 功能分支（一任务一分支）  
2. **禁止**多人同时修改：`schemas/v0.1/defs.json`、`event-envelope.json`、`evidence.json`、`change-event.json`  
3. 用户可见名称使用「专业名 + 通俗解释」  
4. 不要提交真实密钥；示例须通过 `tests/redaction`  

## 内部备忘 vs 对外文档

| 类型 | 路径 | 读者 |
|------|------|------|
| 对外 | `README.md`、`docs/terminology.md` 等 | 访客与集成方 |
| 内部里程碑 | `docs/checkpoint-*.md` | 维护者 |

## 许可

贡献默认按本仓库 Apache-2.0 许可。
