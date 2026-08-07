# 协议架构说明

**专业名称：** Protocol Architecture  
**通俗解释：** 这套“数据语言”怎么支撑四个独立工具协作，同时谁也能单独干活。

## 1. 边界

```text
┌─────────────────────────────────────────────┐
│                 lifecycle-spec              │
│  docs + JSON Schema + examples + tests      │
└─────────────────────────────────────────────┘
          ▲           ▲           ▲
          │文件/HTTP  │           │
   ┌──────┴──┐  ┌─────┴────┐  ┌───┴──────────┐
   │dev-cycle│  │infra-disc│  │fleet / release│
   └─────────┘  └──────────┘  └──────────────┘
```

规则：

1. 产品之间**禁止** import 对方内部包。  
2. 只通过本协议对象交换。  
3. 无消息队列要求；可用目录投递、CLI 输出、Webhook。

## 2. 核心对象关系

```text
Release ──expects──► ChangeEvent
ChangeEvent ──subject──► Resource
ChangeEvent ──evidence──► Evidence
Observation ──about──► Resource
Snapshot ──contains──► Resource 状态集合
Relationship ──links──► Resource × Resource
Approval ──covers──► ChangeEvent 或例外窗口
```

## 3. 本地优先交换路径（第一版）

| 路径 | 用途 |
|------|------|
| `inventory.json` / `snapshot.json` | 发现工具输出 |
| `events/*.json` | 变化与发布事件 |
| `evidence/*` | 证据材料与清单 |
| HTTP GET/POST JSON | 可选本地 API |

## 4. AI 可选

协议字段可被 AI 用于解释，但：

- AI 输出不得标为确定性事实，除非绑定 Evidence；  
- 未配置 AI 时，Diff / 扫描 / Git 等确定性功能不受影响。

## 5. 非目标

- 不规定统一数据库  
- 不规定 Kafka / NATS / Pulsar  
- 不规定单一 monorepo 运行时  
- 不在协议内实现 RBAC 全集（Approval 仅表达结果）

## 6. 与总计划书关系

总体产品愿景见上级目录 `ai-devops-open-source-suite-plan-v0.1.md`。  
本仓库是其 **M0 协议地基** 的规范性实现起点。
