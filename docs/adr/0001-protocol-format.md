# ADR-0001：协议主格式选择 JSON Schema

- 状态：Accepted
- 日期：2026-08-07
- 任务：FND-001 前置 / 协议奠基

## 背景

套件需要四个独立产品共享数据模型。总计划同时提到了 JSON Schema 与 Protobuf。需要明确第一版主格式，避免实现分叉。

## 决策

**第一版以版本化 JSON Schema（Draft 2020-12）为唯一规范性源（source of truth）。**

交换载体优先：

1. 本地文件（`.json`）
2. HTTP API / Webhook（JSON body）
3. 可选 YAML 仅作人类可读导出（规范语义以 JSON Schema 为准）

Protobuf / gRPC **不作为** `lifecycle-spec` 第一版规范要求。监控项目内部若使用 gTLS+gRPC，须在边界转换为本协议 JSON 对象，不得要求其他三产品依赖 Protobuf。

## 理由

1. 文件优先、本地优先：JSON 无需代码生成即可人工阅读与 diff。
2. 跨语言：Go / TypeScript / Python 生态均可校验与生成类型。
3. 降低耦合：不强制消息总线或 RPC 运行时。
4. 第一版保持简单：避免双轨维护 JSON + Protobuf。

## 后果

- 正向：示例、文档、CI 校验路径统一。
- 负向：高频二进制遥测仍由各产品内部协议处理，仅在“生命周期事件/证据/快照”边界使用本规范。
- 后续若引入 Protobuf，必须新开 ADR，并定义 JSON ↔ Protobuf 映射与兼容策略。

## 未决（非阻塞）

- 正式对外 `$id` 域名（暂用 `https://lifecycle-spec.local/schemas/v0.1/...` 占位，发布前更换）。
