# lifecycle-spec

**一句话定位：** 让开发、发现、监控、发布四类工具说同一种数据语言的开放协议。

**专业名称：** Lifecycle Protocol Specification  
**通俗解释：** 这里没有会「跑起来」的业务程序，只有大家约定好的数据格式、例子和校验——这样工具之间不用绑死对方的代码。

[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)

套件总览：[project-docs](https://github.com/GeneJie199/project-docs) · 当前协议版本：**0.1.0**（预发布，可能有破坏性调整）

---

## 为什么需要协议仓

四个产品必须能**单独安装**。若互相 import 内部包，会变成单体。  
用版本化 JSON Schema 交换「资源、证据、快照、变化」等对象，才能既独立又协作。

---

## 核心概念（专业名 + 通俗解释）

| 专业名称 | 通俗解释 |
|----------|----------|
| Resource | 被管理的对象：主机、服务、端口、仓库等 |
| Snapshot | 某一刻的标准化状态「照片」，用来对比 |
| ChangeEvent | 一次可追踪的变化记录（相对之前状态） |
| Evidence | 证明某事发生或通过的材料（测试、日志摘录、人工确认等） |
| Event Envelope | 所有事件共用的外包装（谁、何时、关于什么） |
| Infrastructure Drift | 基础设施漂移：实际状态偏离基线或预期 |

完整词典见 [docs/terminology.md](docs/terminology.md)。

---

## 快速开始

```bash
git clone https://github.com/GeneJie199/lifecycle-spec.git
cd lifecycle-spec
go test ./...
```

- Schema：`schemas/v0.1/`  
- 示例：`examples/v0.1/`  
- Go 类型：`gen/go/lifecycle/v0_1/`  
- TypeScript 类型：`gen/ts/v0.1/`  

---

## 设计原则

1. 本地文件 / HTTP JSON 即可交换，不强制消息队列  
2. 时间：RFC 3339（带时区）  
3. 稳定 ID；禁止用 PID 等易变值当身份证  
4. 事件中禁止明文密码、Token、私钥  
5. AI 可选；协议本身不绑定模型厂商  

---

## 文档索引

| 文档 | 内容 |
|------|------|
| [terminology.md](docs/terminology.md) | 术语表 |
| [id-conventions.md](docs/id-conventions.md) | ID 规范 |
| [event-envelope.md](docs/event-envelope.md) | 事件信封 |
| [versioning.md](docs/versioning.md) | 版本与兼容 |
| [compatibility.md](docs/compatibility.md) | 兼容测试约定 |
| [security-and-redaction.md](docs/security-and-redaction.md) | 安全与脱敏 |
| [architecture.md](docs/architecture.md) | 协议架构 |
| [adr/0001-protocol-format.md](docs/adr/0001-protocol-format.md) | 为何选 JSON Schema |
| [scripts/generate-types.md](scripts/generate-types.md) | 类型维护说明 |

> `docs/checkpoint-*.md` 为里程碑内部备忘，不是面向访客的产品说明。

---

## 贡献

见 [CONTRIBUTING.md](CONTRIBUTING.md)。修改核心 Schema 前请先补 ADR。安全问题见 [SECURITY.md](SECURITY.md)。

## 许可

Apache-2.0 — [LICENSE](LICENSE)
