# lifecycle-spec

**专业名称：** Lifecycle Protocol Specification  
**通俗解释：** 四个开源工具用来“说同一种数据语言”的公开协议仓库。

本仓库只包含：

- 版本化 JSON Schema
- 协议文档与 ADR
- 示例数据
- Schema 校验与兼容性测试
- 已提交的 Go / TypeScript 类型（`gen/`）
- 安全脱敏扫描测试

**不包含**任何产品运行时代码。四个产品只能依赖本协议，不能依赖彼此内部实现。

当前协议版本：`0.1.0`（预发布，允许破坏性调整，见 `docs/versioning.md`）。

## 仓库

- GitHub：https://github.com/GeneJie199/lifecycle-spec
- 许可证：Apache-2.0

## 快速浏览

| 文档 | 说明 |
|------|------|
| [docs/terminology.md](docs/terminology.md) | 术语表（FND-001） |
| [docs/id-conventions.md](docs/id-conventions.md) | 稳定 ID 规范（FND-002） |
| [docs/event-envelope.md](docs/event-envelope.md) | 统一事件信封（FND-003） |
| [docs/versioning.md](docs/versioning.md) | 版本与兼容规则 |
| [docs/compatibility.md](docs/compatibility.md) | 0.x 兼容测试约定（FND-006） |
| [docs/security-and-redaction.md](docs/security-and-redaction.md) | 安全与脱敏（FND-008） |
| [docs/architecture.md](docs/architecture.md) | 协议架构 |
| [docs/adr/0001-protocol-format.md](docs/adr/0001-protocol-format.md) | 为何选 JSON Schema |
| [scripts/generate-types.md](scripts/generate-types.md) | 类型维护说明（FND-007） |

## Schema（v0.1）

位于 `schemas/v0.1/`：

- `defs.json` — 公共定义
- `event-envelope.json` — 事件信封
- `resource.json` — 资源
- `evidence.json` — 证据（FND-004）
- `change-event.json` — 变化事件（FND-005）
- `observation.json` — 观察
- `snapshot.json` — 快照
- `relationship.json` — 关系
- `approval.json` — 批准
- `release.json` — 发布

示例位于 `examples/v0.1/`。

## 生成类型（v0.1，已提交）

| 语言 | 路径 |
|------|------|
| Go | `gen/go/lifecycle/v0_1/` |
| TypeScript | `gen/ts/v0.1/index.ts` |

维护方式见 [scripts/generate-types.md](scripts/generate-types.md)。`gen/` 必须入库，勿加入 `.gitignore`。

## 校验示例

```bash
go test ./...
```

## 设计原则（摘要）

1. 独立运行、协议协作  
2. 本地优先、AI 可选  
3. 确定性逻辑优先、证据可追踪  
4. 事件中禁止明文密钥 / 密码 / 私钥  
5. 时间一律 RFC 3339（带时区）  
6. 第一版不引入 Kafka / 消息队列 / 复杂分布式架构  
