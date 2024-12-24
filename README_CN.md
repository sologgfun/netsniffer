# Kyanos 流量抓取工具 - 集群扩展版

此仓库为 **Kyanos** 流量抓取工具的集群扩展版，专为 Kubernetes 集群环境设计。采用 **1 个 Server + n 个 Client** 架构，能够高效地在集群中分布式地抓取流量数据，并进行集中分析。

## 使用说明

### 部署
使用目录下的[yaml文件](https://github.com/sologgfun/netsniffer/tree/main/yaml)直接apply部署，镜像可以使用本地构建的镜像，也可以使用远程仓库的镜像。
```
moppyz/ns-server:v0.1
moppyz/ns-client:v0.1
```

### Client 使用

1. 在 Client Pod 中直接执行 Kyanos 命令行工具，即可开始抓取流量。参数与 Kyanos 保持一致，主要区别在于抓取的数据会被发送至 Server 端的 MySQL 数据库进行存储。
   
   ![Client Example](https://github.com/user-attachments/assets/6c5bd871-a9b3-44e3-9c92-3726599fc090)

### Server 使用

1. 在 Server 端，您可以通过 SQL 查询客户端抓取的流量数据，进行深入分析：
   
   ```bash
   mysql -u root -p
   # 初始密码：rootpwd
   ```

   ![MySQL Example](https://github.com/user-attachments/assets/e78bdbed-3909-4a34-a42d-752ff2ffcf93)

2. 使用 `./ns-ctl` 工具访问 MySQL 数据库，并分析 Kyanos 采集的数据：

   ![ns-ctl Example](https://github.com/user-attachments/assets/490fedc3-abb5-4f62-8d05-06eff65655bf)

## 系统架构概述

本工具采用 **Server-Client 模式** 设计：

- **1 个 Server + n 个 Client**：Server 作为数据汇总和分析的核心，多个 Client 部署在集群各节点上，负责流量的采集。
  
  > 组件架构示意图：
  
  ![Architecture Diagram](https://github.com/user-attachments/assets/9a9b440d-f2d8-4dd7-a8fe-92a10574e222)

  数据通过 Client 采集后，通过部署在 Server 端的 Service（SVC）传输至 Server 进行处理。

- **Kubernetes 集成**：通过 Kubernetes 的 YAML 配置文件，简化集群内应用的部署与管理：
  - **Client 部署**：使用 DaemonSet（DS）在每个节点上部署一个 Client 实例，确保集群内所有节点都能采集到流量数据。
  - **Server 部署**：使用 Deployment 方式部署 Server，支持水平扩展，以应对高负载的需求。
  
- **特权 Pod**：利用 Kubernetes 的特权 Pod，客户端能够抓取集群内所有类型的流量，包括容器流量，确保数据覆盖集群内的每个节点。
- **临时数据库**：Client 采集的流量数据会暂时存储在 Server 的数据库中，便于后续分析。

### 当前版本

**v1.0.0** - 已通过验证并在实际环境中成功部署。

## 开发指南

### 项目结构

1. **server** - `ns-server`
2. **ctl** - `ns-ctl`
3. **client** - 根目录

## 后续开发计划

以下是未来版本的开发计划和功能优化方向：

### 1. **eBPF 功能优化**
   - 当前的 eBPF 功能已基于 Kyanos 库实现，但仍存在性能与灵活性提升空间。我们正在优化 eBPF，以更好地支持大规模集群流量抓取。

### 2. **SQL Web GUI 数据分析界面**
   - 我们计划为 Server 增加一个 SQL Web GUI，用户可以通过 Web 界面查询和分析捕获的流量数据。此功能将提供直观的界面，简化数据分析和操作流程。

### 3. **兼容性提升**
   - 我们将持续增强与不同 Kubernetes 版本和云原生环境的兼容性，确保该工具能够在各类 Kubernetes 集群中无缝运行。

### 4. **性能基准测试**
   - 我们计划进行全面的性能基准测试，评估工具在不同规模集群中的表现，并针对潜在的性能瓶颈进行优化，确保其在生产环境中的高效运行。

## 主要特性

- **可扩展部署架构**：支持在 Kubernetes 集群中轻松部署和扩展，适用于大规模分布式环境，且无侵入性。
- **分布式流量抓取**：Client 利用特权 Pod 捕获每个节点（包括容器）的网络流量，确保所有节点的数据都能被采集。
- **集中式数据收集**：所有流量数据会集中存储在 Server 端的临时数据库中，便于后续的分析和处理。
- **高度可定制与可扩展**：用户可以根据需求定制抓取规则、数据处理方式等，适应不同的业务场景。

## 参考资料

> 本项目基于 Debian 12.0 云主机进行开发，需要具备一定的 eBPF 和 Go 语言基础。相关学习资料如下：

- [Getting Started with eBPF in Go](https://ebpf-go.dev/guides/getting-started/)
- [高级 eBPF 内核特性（FOSDEM, 2021）](https://arthurchiao.art/blog/advanced-bpf-kernel-features-for-container-age-zh/#41-进出宿主机的容器流量host---pod)
- [eBPF 开发实践教程](https://eunomia.dev/zh/tutorials/)