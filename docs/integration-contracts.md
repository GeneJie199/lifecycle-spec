# 产品集成契约

四个运行产品通过版本化 JSON 或本地 HTTP API 交换数据，不导入彼此的内部代码。

| 生产者 | 消费者 | 契约 | 默认交付方式 |
|---|---|---|---|
| DevCycle | 自动化代理 / ReleaseGuard | `project.json`、`requirement.json`、`acceptance-criterion.json`、`task.json` | 本地 HTTP / JSON |
| DevCycle | ReleaseGuard | `release-candidate.json` | 文件 |
| InfraScout | FleetScope | `fleet-node-report.json` 中的 `inventory` / `drift` | Agent HTTP POST |
| InfraScout | ReleaseGuard | InfraScout drift report | 文件 |
| InfraScout | FleetScope Agent / 运维人员 | `monitoring-plan.json` / `monitoring-recommendation.json` | JSON 文件 |
| FleetScope Agent | FleetScope Center | `telemetry-batch.json` / `event-batch.json` | 带节点凭证的 HTTP POST |
| 主流数据源适配器 | FleetScope Center | 规范化后的 `telemetry-batch.json` / `event-batch.json` | HTTP POST |
| FleetScope | DevCycle / ReleaseGuard | `incident.json`、节点、指标与告警 API | HTTP GET / JSON |
| ReleaseGuard | 审计与 UI | `release-validation-report.json` | 不可变文件 |

## 兼容要求

1. 生产者新增字段前，应放入允许扩展的对象或提升契约版本。
2. 消费者必须拒绝未知的顶层控制字段，证据与指标对象可保留扩展字段。
3. 时间使用带时区的 RFC 3339，摘要使用小写十六进制 SHA-256。
4. 发布候选的 `ready` 不能单独作为依据，消费者应复核验收、证据与任务计数。
5. Fleet 指标允许嵌套探针数组；敏感连接串和环境变量值不得写入报告。
6. FleetScope 原生采集器是默认采集路径。Prometheus、OpenTelemetry 等格式只属于可选兼容输入，不是运行依赖。
7. 发布候选任务的 `dependsOn` 必须引用同一候选中的任务；ReleaseGuard 的观察中间态使用 `decision: HOLD` 和 `observation.status: observing`，恢复完成后写为 `completed`。

## 验证

`go test ./...` 校验本仓库示例。套件根目录的 `scripts/verify-suite.sh` 还会把各产品真实输出串联起来，验证生产者与消费者的运行时兼容性。
