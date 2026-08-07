# 安全与脱敏规范

**任务关联：** FND-008  
**专业名称：** Security & Redaction  
**通俗解释：** 协议数据可以给人看、给 AI 看，但绝不能把密码原样写进去。

## 1. 硬性禁止

以下内容**不得**出现在任何协议对象的字段值中（含 `payload`、`extensions`、Evidence `inline`、日志摘录）：

1. 密码、口令、通行短语  
2. API Key、Token、Session Cookie 值  
3. TLS/SSH 私钥、证书私钥材料  
4. 数据库连接串中的口令部分  
5. 云厂商永久访问密钥 Secret  

允许：

- 密钥**名称**/环境变量**名**（不含值）  
- 指纹、哈希、后四位掩码  
- `redacted://` 占位符

## 2. 推荐脱敏替换

| 原值类型 | 替换 |
|----------|------|
| 通用秘密 | `***REDACTED***` |
| 连接串 | `postgres://user:***REDACTED***@host:5432/db` |
| Token | 保留前缀类型 + `***`，如 `gho_***REDACTED***` |
| 私钥块 | `***REDACTED PRIVATE KEY***` |

## 3. Evidence 规则

1. 默认最小化：能给摘要就不要给全文。  
2. `redaction.status`：`none` | `partial` | `full`。  
3. 截图须避免含秘密；含秘密则 `full` 且不附原始 URI，或 URI 指向已处理副本。  
4. 发送给 AI 前必须可预览，并默认启用脱敏（产品侧责任，本协议要求字段可表达脱敏状态）。

## 4. 事件与快照

- 环境变量：默认只采集名称列表。  
- 进程命令行：若匹配秘密模式，须脱敏后再入 Snapshot。  
- Diff 结果同样禁止回填明文秘密。

## 5. 传输与存储（产品责任摘要）

协议不规定存储引擎，但要求：

- 本地优先；  
- 远程传输使用 TLS；  
- 监控 Agent 与中心 mTLS（产品文档）；  
- 本仓库测试样例不得含真实秘密。

## 6. 负面验收与自动化扫描

应拒绝或告警的样例：

- `password=...`、`BEGIN PRIVATE KEY` / `BEGIN RSA PRIVATE KEY`、`aws_secret`、GitHub token（`ghp_` / `gho_` + 长秘密）等。

自动化：`tests/redaction`（包名 `redaction`）扫描 `examples/` 与 `schemas/` 下全部 JSON，断言不得命中上述模式；并含正向用例，确认构造的秘密字符串会被扫描器命中。运行：`go test ./tests/redaction/...`。
