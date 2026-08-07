# 统一事件信封（Event Envelope）

**任务：** FND-003  
**专业名称：** Event Envelope  
**通俗解释：** 所有“发生了一件事”的通知，都用同一套外壳包装，里面再放具体内容。

Schema：`schemas/v0.1/event-envelope.json`  
示例：`examples/v0.1/event-envelope.json`

## 1. 为何需要信封

四个产品交换的不是数据库行，而是**不可变事件**与**证据引用**。统一外壳保证：

- 能知道谁在何时、对什么资源、说了哪类事；
- 能挂分类、严重级别、发布与证据；
- 不绑定 Kafka 或特定队列——文件 / HTTP / Webhook 均可承载同一 JSON。

## 2. 字段一览

| 字段 | 必需 | 通俗解释 | 说明 |
|------|------|----------|------|
| `specVersion` | 是 | 协议版本 | 语义化版本字符串，如 `0.1.0` |
| `eventId` | 是 | 事件身份证 | `evt_` 前缀，全局唯一；事件不可原地改 |
| `eventType` | 是 | 事件种类 | 点分层，如 `resource.changed` |
| `occurredAt` | 是 | 事情发生时间 | RFC 3339 **带时区** |
| `recordedAt` | 否 | 写入时间 | 默认可等于 occurredAt |
| `source` | 是 | 谁发的 | `product` + `instanceId` |
| `context` | 否 | 业务上下文 | 项目、环境、资源组等 |
| `subject` | 是 | 事件主语资源 | 至少含 `resourceId` |
| `classification` | 否 | 是否预期 | 默认由消费方视为 `unexpected` 若缺失（建议生产者显式填写） |
| `severity` | 否 | 严重程度 | 默认 `info` |
| `releaseId` | 否 | 关联发布 | `rls_` ID |
| `evidence` | 否 | 证据列表 | Evidence 对象或 `evidenceId` 引用 |
| `payload` | 是 | 具体内容 | 对象；种类由 `eventType` 决定 |
| `extensions` | 否 | 扩展袋 | 额外键；消费者可忽略未知键 |

## 3. 不变性

1. 同一 `eventId` 内容不得修改。  
2. 纠错：发新事件，`payload.correctsEventId`（可选约定）指向旧事件。  
3. 删除语义：发 `*.invalidated` 类事件，不物理改历史文件（若存储层支持）。

## 4. eventType 命名

```text
{domain}.{entity}.{verb}
```

首版保留域：

| domain | 用途 |
|--------|------|
| `resource` | 资源发现与属性 |
| `change` | 标准化变化（见 ChangeEvent） |
| `evidence` | 证据产生 |
| `release` | 发布生命周期 |
| `approval` | 批准 |
| `observation` | 观察 |
| `snapshot` | 快照产生 |

示例：`change.resource.modified`、`evidence.created`、`release.candidate.created`。

## 5. 与 ChangeEvent 的关系

`change-event.json` **是**事件信封的一个画像（profile）：

- 必须满足 `event-envelope.json`；
- `eventType` 必须以 `change.` 开头；
- `payload` 必须符合 ChangeEvent payload 定义。

## 6. 安全

信封任意位置（含 `payload`、`extensions`、`evidence.inline`）**禁止**出现明文密码、API Key、私钥、连接串口令。见 `security-and-redaction.md`。

## 7. 验收（FND-003）

- [x] Schema 与文档字段一致  
- [x] 示例通过校验  
- [x] 明确不可变与 eventType 规则  
