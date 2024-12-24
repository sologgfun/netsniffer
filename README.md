# Kyanos Traffic Capture Tool - Cluster Edition

This repository contains the **Cluster Edition** of the **Kyanos** traffic capture tool, specifically designed for Kubernetes cluster environments. It follows a **1 Server + n Clients** architecture, enabling efficient distributed traffic data capture across the cluster with centralized analysis.

## Usage Instructions

### Deployment

You can directly apply and deploy using the [yaml files](https://github.com/sologgfun/netsniffer/tree/main/yaml) in the directory. The image can either be a locally built image or one from a remote repository.
```
moppyz/ns-server:v0.1
moppyz/ns-client:v0.1
```

### Client Usage

1. In the Client Pod, simply run the Kyanos command-line tool to start capturing traffic. The parameters are the same as the original Kyanos tool, with the main difference being that the captured data is sent to the MySQL database on the server for storage.

   ![Client Example](https://github.com/user-attachments/assets/6c5bd871-a9b3-44e3-9c92-3726599fc090)

### Server Usage

1. On the Server side, you can query the traffic data captured by the Client using SQL statements for analysis:

   ```bash
   mysql -u root -p
   # Initial password: rootpwd
   ```

   ![MySQL Example](https://github.com/user-attachments/assets/e78bdbed-3909-4a34-a42d-752ff2ffcf93)

2. Use the `./ns-ctl` tool to access the MySQL database and analyze the data collected by Kyanos:

   ![ns-ctl Example](https://github.com/user-attachments/assets/490fedc3-abb5-4f62-8d05-06eff65655bf)

## System Architecture Overview

The tool is designed using a **Server-Client** architecture:

- **1 Server + n Clients**: The Server acts as the central hub for data aggregation and analysis, while multiple Clients are deployed across the cluster nodes to capture traffic.

  > Component architecture diagram:

  ![Architecture Diagram](https://github.com/user-attachments/assets/9a9b440d-f2d8-4dd7-a8fe-92a10574e222)

  The captured data is sent from the Clients to the Server via the service (SVC) deployed on the Server.

- **Kubernetes Integration**: Simplified deployment and management through Kubernetes YAML configuration files:
  - **Client Deployment**: A DaemonSet (DS) is used to deploy a Client on each node, ensuring traffic capture on every node in the cluster.
  - **Server Deployment**: The Server is deployed using a Deployment, supporting horizontal scaling to handle high traffic loads.

- **Privileged Pods**: The Clients leverage Kubernetes privileged Pods to capture traffic from all types of traffic, including container traffic, ensuring data collection across all nodes in the cluster.
- **Temporary Database**: The traffic data captured by the Clients is temporarily stored in the Server's database for subsequent analysis.

### Current Version

**v1.0.0** - Successfully validated and deployed in a real environment.

## Development Guide

### Project Structure

1. **server** - `ns-server`
2. **ctl** - `ns-ctl`
3. **client** - Root directory

## Future Development Plans

The following features and improvements are planned for future versions:

### 1. **eBPF Functionality Optimization**
   - The current eBPF functionality is based on the Kyanos library but still has room for performance and flexibility improvements. We are working on optimizing eBPF to better support traffic capture in large-scale clusters.

### 2. **SQL Web GUI for Data Analysis**
   - A SQL Web GUI will be added to the Server in future versions, allowing users to query and analyze captured traffic data through a web interface. This feature will provide an intuitive, user-friendly experience for real-time data analysis.

### 3. **Compatibility Enhancements**
   - We will continue to improve compatibility with different Kubernetes versions and cloud-native environments, ensuring the tool runs seamlessly across various Kubernetes clusters.

### 4. **Benchmark Performance Analysis**
   - We plan to conduct comprehensive performance benchmarking, evaluating the tool's performance across different cluster sizes, and optimize for potential bottlenecks to ensure efficient production use.

## Key Features

- **Scalable Deployment Architecture**: Easily deployable and scalable in Kubernetes clusters, supporting large-scale distributed deployments with zero intrusion.
- **Distributed Traffic Capture**: Clients use privileged Pods to capture network traffic from each node (including containers), ensuring comprehensive data collection across the entire cluster.
- **Centralized Data Collection**: All captured traffic data is stored in a temporary database on the Server for easy access and analysis.
- **Customizable and Extensible**: The tool allows users to define custom traffic capture rules and data processing methods to fit specific use cases.

## References

> This project was developed on a cloud-hosted Debian 12.0 instance and requires a basic understanding of eBPF and Go programming. The following resources may be helpful:

- [Getting Started with eBPF in Go](https://ebpf-go.dev/guides/getting-started/)
- [Advanced eBPF Kernel Features for Container Age (FOSDEM, 2021)](https://arthurchiao.art/blog/advanced-bpf-kernel-features-for-container-age-zh/#41-进出宿主机的容器流量host---pod)
- [eBPF Development Tutorial](https://eunomia.dev/zh/tutorials/)