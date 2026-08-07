# 资源 ID 与标识规范

**任务：** FND-002  
**专业名称：** Stable Identifier Conventions  
**通俗解释：** 给每样东西起一个“不会随便变”的身份证号，好让四个工具互相指认。

## 1. 总原则

1. **稳定性优于可读性**：同一逻辑实体在重建进程、重启主机、重新扫描后仍应得到同一 `resourceId`（在约定算法输入不变的前提下）。
2. **禁止用易变字段当 ID**：如 PID、容器短 ID、启动时间、临时端口（除非端口本身是稳定服务契约且写入算法说明）。
3. **所有跨事件引用使用 ID 字符串**，不依赖数组下标。
4. **ID 一经发布用于生产交换，不得原地复用到不同实体**。
5. **大小写敏感**：按 Unicode 码点精确比较。

## 2. 通用 ID 形态

### 2.1 不透明前缀 ID（事件、证据、批准等）

格式：

```text
{prefix}_{ulid}
```

| 前缀 | 对象 | 示例 |
|------|------|------|
| `evt` | 事件 `eventId` | `evt_01J2QK3M4N5P6Q7R8S9T0ABCD` |
| `evd` | 证据 `evidenceId` | `evd_01J2QK3M4N5P6Q7R8S9T0ABCD` |
| `obs` | 观察 `observationId` | `obs_01J2QK3M4N5P6Q7R8S9T0ABCD` |
| `snp` | 快照 `snapshotId` | `snp_01J2QK3M4N5P6Q7R8S9T0ABCD` |
| `apr` | 批准 `approvalId` | `apr_01J2QK3M4N5P6Q7R8S9T0ABCD` |
| `rel` | 关系 `relationshipId` | `rel_01J2QK3M4N5P6Q7R8S9T0ABCD` |
| `rls` | 发布 `releaseId` | `rls_01J2QK3M4N5P6Q7R8S9T0ABCD` |

要求：

- `prefix` 小写 ASCII；
- `ulid` 建议使用 ULID（Crockford Base32），也允许 UUID-v7 去掉连字符；
- 正则：`^[a-z]{2,8}_[0-9A-HJKMNP-TV-Z]{16,32}$`（实现可先按字符串长度与前缀校验）。

> 可读别名（如 `release_2.5.0`）只能放在 `displayName` / `labels`，**不能**替代规范 `releaseId`。若需人类标签，使用 `release.version` 字段。

### 2.2 资源 ID（Resource）

格式：

```text
{resourceType}:{stableLocator}
```

| 部分 | 规则 |
|------|------|
| `resourceType` | 小写，点分层级，如 `host`、`process.service`、`db.postgresql` |
| `:` | 单个冒号分隔 |
| `stableLocator` | 产品定义的稳定定位串，允许 `/`、`.`、`-`；禁止空格与明文密钥 |

示例：

```text
host:shop-prod-api-01
svc.systemd:shop-prod-api-01/nginx.service
db.postgresql:shop-prod/orders
repo.git:github.com/GeneJie199/lifecycle-spec
requirement:proj_shop/req_invite-code
```

正则（宽松）：

```text
^[a-z][a-z0-9.-]*:[A-Za-z0-9._@/-]+$
```

## 3. 稳定 Locator 算法（Linux 主机首版约定）

各产品可扩展，但须文档化。基础设施发现首版建议：

| resourceType | stableLocator 输入 | 说明 |
|--------------|-------------------|------|
| `host` | 主机名（优先）或 machine-id | 主机名冲突时用 `machine-id` |
| `svc.systemd` | `{hostLocator}/{unitName}` | unit 名稳定 |
| `process.bin` | `{hostLocator}/{exePath}` | **不**含 PID；同路径多实例用关系/端口区分 |
| `net.listener` | `{hostLocator}/{proto}/{addr}/{port}` | addr 规范化（去临时接口名噪声另见快照规范） |
| `container.docker` | `{hostLocator}/{composeProject}/{service}` 或镜像+名称策略 | 禁止只用容器短 ID |

进程 PID、启动时间可作为 **属性**，不得进入 ID。

## 4. 跨产品引用

- 开发周期工具导出的 `repo.git:...`、`requirement:...` 可被发布工具引用。
- 发布工具的 `rls_...` 可写入 ChangeEvent 的 `releaseId`。
- 基础设施发现的 `host:...` 可被监控工具作为同一 Resource。

不得要求对方读取本地数据库主键；只通过本规范 ID 交换。

## 5. source.instanceId

标识“哪一个产品实例产生了数据”：

```text
{product}-{stable-instance-key}
```

示例：`infra-discovery-scanner-prod-01`  
不得包含密钥。

## 6. 冲突与迁移

1. 发现两实体撞 ID：视为缺陷；创建新 ID，并用 Relationship `same_as` 或后续迁移事件说明（迁移事件 Schema 另案）。
2. 算法升级导致 ID 变化：必须发 ADR + 主版本或明确 `idMapping` 附件（后续里程碑）。

## 7. 验收（FND-002）

- [x] 文档定义不透明 ID 与资源 ID 两种形态  
- [x] 明确禁止 PID 等噪声进入 ID  
- [x] 给出正则与示例  
- [x] 与事件信封、Evidence、ChangeEvent 字段对齐  
