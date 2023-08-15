# nic-feature-discovery

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/Mellanox/nic-feature-discovery)](https://goreportcard.com/report/github.com/Mellanox/nic-feature-discovery)
[![Coverage Status](https://coveralls.io/repos/github/Mellanox/nic-feature-discovery/badge.svg?branch=main)](https://coveralls.io/github/Mellanox/nic-feature-discovery?branch=main)
[![Build, Test, Lint](https://github.com/Mellanox/nic-feature-discovery/actions/workflows/build-test-lint.yaml/badge.svg)](https://github.com/Mellanox/nic-feature-discovery/actions/workflows/build-test-lint.yaml)
[![CodeQL](https://github.com/Mellanox/nic-feature-discovery/actions/workflows/codeql.yaml/badge.svg)](https://github.com/Mellanox/nic-feature-discovery/actions/workflows/codeql.yaml)
[![Image Push](https://github.com/Mellanox/nic-feature-discovery/actions/workflows/image-push-main.yaml/badge.svg)](https://github.com/Mellanox/nic-feature-discovery/actions/workflows/image-push-main.yaml)

- [nic-feature-discovery](#nic-feature-discovery)
  - [Overview](#overview)
  - [Prerequisites](#prerequisites)
  - [Supported Feature Labels](#supported-feature-labels)
  - [Quick Start](#quick-start)
    - [Deploy NFD](#deploy-nfd)
    - [Deploy NVIDIA NIC Feature Discovery](#deploy-nvidia-nic-feature-discovery)
    - [Check That it works](#check-that-it-works)
  - [nic-feature-discovery Command Line Interface](#nic-feature-discovery-command-line-interface)
  - [Build and Run Locally](#build-and-run-locally)
    - [Prerequisites](#prerequisites-1)
    - [Build \& Run](#build--run)
    - [Build \& Run in local K8s cluster](#build--run-in-local-k8s-cluster)
      - [Install Prerequisites](#install-prerequisites)
      - [Create local k8s development environment](#create-local-k8s-development-environment)
      - [Cleanup local k8s development environment](#cleanup-local-k8s-development-environment)

## Overview

NVIDIA NIC Feature Discovery for Kubernetes is a software component that allows
you to automatically generate Node labels for NIC related features available on a K8s Node.
It leverages the [Node Feature Discovery](https://github.com/kubernetes-sigs/node-feature-discovery)
to perform this labeling.

## Prerequisites

- Kubernetes >= `1.24`
- Node Feature Discovery (NFD) >= `0.13.2`
  - Deployed on each node where you want to label with the local source configured
  - To deploy NFD, refer to the project [Official Documentation](https://kubernetes-sigs.github.io/node-feature-discovery/stable/get-started/index.html)

## Supported Feature Labels

| Label Name               | Value Type | Description                                | Example         |
| ------------------------ | ---------- | ------------------------------------------ | --------------- |
| nvidia.com/mofed.version | String     | MOFED driver version if present and loaded | `"23.04-0.5.3"` |

## Quick Start

### Deploy NFD

Refer to [Node Feature Discovery - Quick Start](https://kubernetes-sigs.github.io/node-feature-discovery/v0.13/get-started/quick-start.html#quick-start)

Example deployment using kustomize:

```shell
kubectl apply -k https://github.com/kubernetes-sigs/node-feature-discovery/deployment/overlays/default?ref=v0.13.3
```

### Deploy NVIDIA NIC Feature Discovery

```shell
$ kubectl apply -k https://github.com/Mellanox/nic-feature-discovery/deployment/k8s/overlays/default?ref=main
```

### Check That it works

1. Node Feature Discovery is deployed

```shell
root:~# kubectl get -n node-feature-discovery all
NAME                              READY   STATUS    RESTARTS      AGE
pod/nfd-master-5c56499456-r7g2h   1/1     Running   0             3d23h
pod/nfd-worker-9qbdh              1/1     Running   6 (16h ago)   3d23h
pod/nfd-worker-w6twz              1/1     Running   0             3d23h

NAME                 TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)    AGE
service/nfd-master   ClusterIP   10.96.9.245   <none>        8080/TCP   3d23h

NAME                        DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR   AGE
daemonset.apps/nfd-worker   2         2         2       2            2           <none>          3d23h

NAME                         READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/nfd-master   1/1     1            1           3d23h

NAME                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/nfd-master-5c56499456   1         1         1       3d23h
```

2. NVIDIA NIC Feature discovery is deployed

```shell
root:~# kubectl get -n nic-feature-discovery all
NAME                                 READY   STATUS    RESTARTS   AGE
pod/nic-feature-discovery-ds-ln4r9   1/1     Running   0          3d23h
pod/nic-feature-discovery-ds-rtvjj   1/1     Running   0          3d23h
pod/nic-feature-discovery-ds-tbtr2   1/1     Running   0          3d23h

NAME                                      DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR   AGE
daemonset.apps/nic-feature-discovery-ds   3         3         3       3            3           <none>          3d23h
```

3. Node contain expected labels

```shell
root:~# kubectl describe node my-worker-node
Name:               my-worker-node
Roles:              <none>
Labels:             beta.kubernetes.io/arch=amd64
                    beta.kubernetes.io/os=linux
                    feature.node.kubernetes.io/cpu-cpuid.AESNI=true
                    feature.node.kubernetes.io/cpu-cpuid.AVX=true
                    feature.node.kubernetes.io/cpu-cpuid.CMPXCHG8=true
                    feature.node.kubernetes.io/cpu-cpuid.FLUSH_L1D=true
                    feature.node.kubernetes.io/cpu-cpuid.FXSR=true
                    feature.node.kubernetes.io/cpu-cpuid.FXSROPT=true
                    feature.node.kubernetes.io/cpu-cpuid.IBPB=true
                    ....
                    ...
                    ..
                    .
                    nvidia.com/mofed.version=23.04-0.5.3
```

## nic-feature-discovery Command Line Interface

```text
Usage:
  nic-feature-discovery [flags]

Feature Discovery Daemon flags:

      --features-file-name string                                                                                                                                                                                                              
                features file name (default "nvidia-com-nic-feature-discovery.features")
      --features-scan-interval duration                                                                                                                                                                                                        
                features scan interval (default 1m0s)
      --nfd-features-path string                                                                                                                                                                                                               
                node feature discovery local features path (default "/etc/kubernetes/node-feature-discovery/features.d/")

Logging flags:

      --log-flush-frequency duration                                                                                                                                                                                                           
                Maximum number of seconds between log flushes (default 5s)
      --log-json-info-buffer-size quantity                                                                                                                                                                                                     
                [Alpha] In JSON format with split output streams, the info messages can be buffered for a while to increase performance. The default value of zero bytes disables buffering. The size can be specified as number of bytes (512),
                multiples of 1000 (1K), multiples of 1024 (2Ki), or powers of those (3M, 4G, 5Mi, 6Gi). Enable the LoggingAlphaOptions feature gate to use this.
      --log-json-split-stream                                                                                                                                                                                                                  
                [Alpha] In JSON format, write error messages to stderr and info messages to stdout. The default is to write a single stream to stdout. Enable the LoggingAlphaOptions feature gate to use this.
      --logging-format string                                                                                                                                                                                                                  
                Sets the log format. Permitted formats: "json" (gated by LoggingBetaOptions), "text". (default "text")
  -v, --v Level                                                                                                                                                                                                                                
                number for the log level verbosity
      --vmodule pattern=N,...                                                                                                                                                                                                                  
                comma-separated list of pattern=N settings for file-filtered logging (only works for text log format)

General flags:

  -h, --help                                                                                                                                                                                                                                   
                print help and exit
      --version                                                                                                                                                                                                                                
                print version and exit
```

## Build and Run Locally

### Prerequisites

- The general [prerequitites](#prerequisites)
- golang >= 1.20

### Build & Run

1. install [Task](https://taskfile.dev/installation/)

```shell
go install github.com/go-task/task/v3/cmd/task@latest
```

2. clone project

```shell
git clone https://github.com/Mellanox/nic-feature-discovery.git
```

3. build binary

```shell
cd nic-feature-discovery
task build
```

4. run binary

```shell
./build/nic-feature-discovery
```

> __Note__: To build container image run `task image:build`. to deploy this image in your k8 cluster, you should re-tag and upload to your
> own image registry, then create an overlay for exsiting deployment which overrides the image name with your own image path.


### Build & Run in local K8s cluster

For local development it is possible to use [skaffold](https://skaffold.dev/) to deploy
local changes onto a local k8s cluster.

#### Install Prerequisites

1. [install Docker](https://docs.docker.com/engine/install/)
2. [install minikube](https://minikube.sigs.k8s.io/docs/start/)
3. [install skaffold](https://skaffold.dev/docs/install/#standalone-binary)

#### Create local k8s development environment

```shell
task -o interleaved localdev:start
```

To trigger a rebuild, press any key in the current shell.

#### Cleanup local k8s development environment

```shell
task -o interleaved localdev:clean
```
