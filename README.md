# Kyanos 流量抓取工具 - 集群扩展版

本仓库是 **Kyanos** 流量抓取工具的集群扩展版本，专为 Kubernetes 集群环境设计。该工具采用 **1 个 Server + n 个 Client** 的架构，能够高效地在集群中分布式地抓取流量数据并进行集中分析。

## 系统架构概述

本工具的架构设计为 **Server-Client 模式**：

- **1 个 Server + n 个 Client**：Server 作为数据汇总和分析的核心，多个 Client 部署在集群的各个节点上，负责流量的采集。
> 组件参考图
<img width="538" alt="image" src="https://github.com/user-attachments/assets/60d1b56a-8cd4-40dc-9433-1671b308f8e5">

> server数据库参考图
<img width="674" alt="image" src="https://github.com/user-attachments/assets/c5a8acda-299c-4629-8fc5-5edb7d8ce74d">

通过client采集数据通过server部署的svc发给server端处理

- **Kubernetes 集成**：通过 Kubernetes 的 YAML 配置文件进行自动化部署，简化集群内应用的部署和管理。
  - **Client 部署**：使用 DaemonSet（DS）在集群的每个节点上部署 Client，确保每个节点上都有一个 Client 实例。
  - **Server 部署**：使用 Deployment 部署 Server，支持水平扩展以应对高负载需求。
- **特权 Pod**：利用 Kubernetes 特权 Pod 的特性，客户端能够抓取包括容器在内的各类流量，确保集群内所有节点的流量都能被采集。
- **临时数据库**：Client 抓取的流量数据会暂时存储到 Server 的数据库中，进行后续分析。

### 当前版本

**v1.0.0** - 已验证可行性，并在实际环境中部署成功。

## 后续开发计划

以下是未来版本的开发计划和功能优化方向：

### 1. **eBPF 功能优化**
   - 当前 eBPF 功能基于 Kyanos 库实现，但在性能和灵活性上还有提升空间。我们正在探索如何优化 eBPF 以更好地支持大规模集群的流量抓取。

### 2. **SQL Web GUI 数据分析界面**
   - 后续将为 Server 增加一个 SQL Web GUI，允许用户通过 Web 界面查询和分析捕获的流量数据。此功能将使用户可以通过直观的界面进行实时数据分析，简化操作和使用。

### 3. **兼容性提升**
   - 我们将持续增强与不同 Kubernetes 版本以及云原生环境的兼容性，以确保工具能够在各种 Kubernetes 集群中无缝运行。

### 4. **Benchmark 性能分析**
   - 我们计划对工具的性能进行全面的基准测试，评估其在不同规模集群中的表现，并针对性能瓶颈进行优化，确保在生产环境下的高效运行。

## 主要特性

- **可扩展的部署架构**：可以在 Kubernetes 集群中轻松部署和扩展，支持大规模分布式部署，无侵入性。
- **分布式流量抓取**：客户端利用特权 Pod 捕获每个节点（包括容器）的网络流量，确保集群内所有节点的数据都能被捕获。
- **集中式数据收集**：所有抓取的流量数据都集中存储在 Server 的临时数据库中，便于后续的数据分析。
- **可定制和可扩展**：支持用户根据需求自定义抓取规则、数据处理方式等，适应不同的业务场景。

## 开发指南

1. server和client分别见server文件和根目录
2. 个人使用云主机基于debian12.0远程开发，进行开发需要一定ebpf及golang的基础知识，可参考下方文档

## 参考资料
- [Getting Started with eBPF in Go](https://ebpf-go.dev/guides/getting-started/)
- [[译] 为容器时代设计的高级 eBPF 内核特性（FOSDEM, 2021）](https://arthurchiao.art/blog/advanced-bpf-kernel-features-for-container-age-zh/#41-进出宿主机的容器流量host---pod)
- [eBPF 开发实践教程](https://eunomia.dev/zh/tutorials/)
