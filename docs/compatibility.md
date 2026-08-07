# 兼容性说明（0.x）

**任务关联：** FND-006  
**专业名称：** Compatibility  
**通俗解释：** 0.x 还在定型，怎么测、怎么扩，才不会 silently 把四个产品弄挂。

完整规则见 [versioning.md](versioning.md)。本文只摘要测试侧约定。

## 0.x 规则摘要

- `0.x` 为预发布：允许破坏性变更，但须 bump 版本、更新 Schema/示例/测试，并写清变更。
- 产品集成应锁定次版本（如 `0.1`），并在 CI 校验 `specVersion`。
- 目标（1.0+）向后兼容：可新增可选字段与枚举值；不可删字段、改语义、或把可选改成必需。
- `extensions` 中的未知键：默认须被忽略且**不得**导致 Schema 校验失败（严格模式由产品可选开启）。

## 本仓库测试覆盖

`tests/schematest/`：

1. **正例**：`examples/v0.1/*` 对对应 Schema 全部通过。
2. **负例**：缺必需字段、错误 `eventId` 前缀、无时区时间戳、`change-event` 使用非 `change.*` 的 `eventType`。
3. **正向兼容**：事件带未知 `extensions` 键仍通过校验。

Schema 是规范性源；类型与文档须与之同步（见 `scripts/generate-types.md`）。
