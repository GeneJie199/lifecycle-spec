# 术语表（Terminology）

**任务：** FND-001  
**专业名称：** Shared Terminology  
**通俗解释：** 四个工具共用的“词典”，避免同名异义。

本表是规范性文档。Schema 字段名与 `eventType` 取值不得与本表冲突；新增术语须同步更新本文件。

## 1. 总则

| 术语（中/英） | 通俗解释 | 规范性含义 |
|---------------|----------|------------|
| 协议套件 / Suite | 一整套可分可合的工具 | 四个独立产品 + `lifecycle-spec` |
| 产品 / Product | 某一个可安装工具 | 如开发周期、基础设施发现、监控、发布验收 |
| 协议 / Spec | 大家约定的数据格式 | 本仓库中的 Schema 与文档 |
| 实例 / Instance | 某次安装或某台扫描器 | `source.instanceId` 标识运行中的一个产品副本 |

## 2. 核心对象

| 对象 | 通俗解释 | 定义 |
|------|----------|------|
| **Resource**（资源） | “系统里被管理的一样东西” | 具有稳定 `resourceId` 的实体：主机、进程、服务、容器、数据库、仓库、需求等 |
| **Relationship**（关系） | “谁依赖谁 / 谁包含谁” | 两个 Resource 之间的有向或无向关联，含关系类型 |
| **Evidence**（证据） | “证明某事发生过的材料” | 可核验的附属材料：测试结果、截图、日志摘录、提交哈希、人工确认等 |
| **Observation**（观察） | “某时刻看了一眼的结果” | 针对一个或多个 Resource 的一次观测记录，可引用 Evidence |
| **Snapshot**（快照） | “某一刻的标准照片” | 已去除噪声字段的规范化状态集合，用于对比 |
| **ChangeEvent**（变化事件） | “和上次比，哪里变了” | 描述资源状态差异的不可变事件，必须可挂到信封上 |
| **Approval**（批准） | “人工说这次变化可以” | 对变化或例外的授权记录（含临时窗口） |
| **Release**（发布） | “一次打算上线的版本动作” | 发布候选或已执行发布的声明，可关联预期变化 |

## 3. 开发周期相关（跨产品引用）

| 对象 | 通俗解释 | 定义 |
|------|----------|------|
| Project | 一个代码仓库或产品线 | 逻辑项目边界，常对应 Git 仓库 |
| Requirement | 要做什么 | 用户可理解的需求描述 |
| AcceptanceCriterion | 怎样算做完 | 可验证的验收条件，应绑定 Evidence |
| Task | 拆出来的开发任务 | 通常对应 branch / worktree |
| AgentSession | AI 干活的一次会话 | 外部 Agent 的一次执行记录 |
| CodeChange | 代码改了什么 | 基于 Git Diff / commit 的变更摘要对象 |

> 上述开发对象的完整 Schema 可在后续里程碑补充；第一版跨产品协作以 Resource / Evidence / ChangeEvent / Release 为主。

## 4. 事件与分类

| 术语 | 通俗解释 | 定义 |
|------|----------|------|
| Event Envelope（事件信封） | 所有事件外面那一层统一包装 | 见 `event-envelope.md` |
| classification | 这次变化算不算“预料之中” | `expected` / `approved` / `temporary` / `unexpected` / `denied` |
| severity | 严重程度 | `info` / `low` / `medium` / `high` / `critical` |
| occurredAt | 事情发生的时间 | RFC 3339，**必须带时区** |
| recordedAt | 系统记下来的时间 | RFC 3339，可与 occurredAt 不同 |

### classification 取值

| 值 | 通俗解释 |
|----|----------|
| `expected` | 发布声明里就写了会变 |
| `approved` | 人确认可以接受 |
| `temporary` | 只在时间窗内允许 |
| `unexpected` | 找不到来源的未知变化 |
| `denied` | 违反禁止策略 |

## 5. 禁止混淆的概念

| 不要把… | 当成… | 原因 |
|---------|--------|------|
| Observation | Snapshot | 观察可以是局部；快照是标准化全集/子集并用于 Diff |
| Evidence | ChangeEvent | 证据是材料；变化事件是结论性差异记录 |
| Approval | Release | 批准可针对单次变化；发布是版本级动作 |
| resourceId | 临时 PID | PID 不稳定，不得作为稳定 ID |

## 6. 产品内部代号（非正式对外名）

计划书中的 DevCycle / InfraScout / FleetScope / ReleaseGuard 仅为内部代号。公开命名前不得写入规范性 `product` 枚举的唯一合法值。第一版 `source.product` 使用稳定短名：

| product 值 | 指代 |
|------------|------|
| `lifecycle-spec` | 本协议（测试/示例） |
| `dev-cycle` | AI 开发周期管理工具 |
| `infra-discovery` | 基础设施发现与漂移工具 |
| `fleet-observability` | 多节点监控与拓扑工具 |
| `release-validation` | 发布检查与验收工具 |
| `other` | 第三方或未列出实现 |

## 7. 修订规则

1. 新增术语：补本表 + 如有 Schema 则同步。  
2. 重命名字段：视为破坏性变更，走 `versioning.md`。  
3. 仅补充通俗解释：非破坏性。  
