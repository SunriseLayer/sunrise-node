# Sunrise Node

[![Go Reference](https://pkg.go.dev/badge/github.com/sunriselayer/sunrise-da.svg)](https://pkg.go.dev/github.com/sunriselayer/sunrise-da)
[![GitHub release (latest by date including pre-releases)](https://img.shields.io/github/v/release/celestiaorg/celestia-node)](https://github.com/sunriselayer/sunrise-da/releases/latest)
[![Go CI](https://github.com/sunriselayer/sunrise-da/actions/workflows/go-ci.yml/badge.svg)](https://github.com/sunriselayer/sunrise-da/actions/workflows/go-ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/sunriselayer/sunrise-da)](https://goreportcard.com/report/github.com/sunriselayer/sunrise-da)
[![codecov](https://codecov.io/gh/celestiaorg/celestia-node/branch/main/graph/badge.svg?token=CWGA4RLDS9)](https://codecov.io/gh/celestiaorg/celestia-node)

Golang implementation of Sunrise's data availability node types (`light` | `full` | `bridge`).

The sunrise-node types described above comprise the sunrise data availability (DA) network.

The DA network wraps the sunrise-core consensus network by listening for blocks from the consensus network and making them digestible for data availability sampling (DAS).

Continue reading [here](https://blog.celestia.org/celestia-mvp-release-data-availability-sampling-light-clients) if you want to learn more about DAS and how it enables secure and scalable access to Sunrise chain data.

## Table of Contents

- [Sunrise Node](#sunrise-node)
  - [Table of Contents](#table-of-contents)
  - [Minimum requirements](#minimum-requirements)
  - [System Requirements](#system-requirements)
  - [Installation](#installation)
  - [API docs](#api-docs)
  - [Node types](#node-types)
  - [Run a node](#run-a-node)
  - [Environment variables](#environment-variables)
  - [Package-specific documentation](#package-specific-documentation)
  - [Code of Conduct](#code-of-conduct)

## Minimum requirements

| Requirement | Notes          |
| ----------- |----------------|
| Go version  | 1.22 or higher |

## System Requirements

See the official docs page for system requirements per node type:

- [Bridge](https://docs.celestia.org/nodes/bridge-node#hardware-requirements)
- [Light](https://docs.celestia.org/nodes/light-node#hardware-requirements)
- [Full](https://docs.celestia.org/nodes/full-storage-node#hardware-requirements)

## Installation

```sh
git clone https://github.com/sunriselayer/sunrise-da.git
cd sunrise-node
make build
sudo make install
```

For more information on setting up a node and the hardware requirements needed, go visit our docs at <https://docs.celestia.org>.

## API docs

The sunrise-node public API is documented [here](https://node-rpc-docs.celestia.org/).

## Node types

- **Bridge** nodes - relay blocks from the sunrise consensus network to the sunrise data availability (DA) network
- **Full** nodes - fully reconstruct and store blocks by sampling the DA network for shares
- **Light** nodes - verify the availability of block data by sampling the DA network for shares

More information can be found [here](https://github.com/sunriselayer/sunrise-da/blob/main/docs/adr/adr-003-march2022-testnet.md#legend).

## Run a node

`<node_type>` can be: `bridge`, `full` or `light`.

```sh
sunrise <node_type> init
```

```sh
sunrise <node_type> start
```

## Network

- **Mainnet** - the main network of Sunrise. Use tokens with real value
- **Testnet** - This network is used to test operations on the mainnet. Normally, the same environment as the mainnet is provided, but tokens have no value
- **Private** - a network for testing new features. It may contain critical bugs. Do NOT normally use

### Run a bridge node

Bridge nodes connect the data availability layer and the consensus layer.

```sh
sunrise bridge init --core.ip <URI> --p2p.network <NETWORK>
```

The `--core.ip` gRPC port defaults to 9090. Normally, set up a RPC node running sunrise-app.

```sh
sunrise bridge start --core.ip <URI> --p2p.network <NETWORK>
```

### Run a full node

Full storage nodes are Sunrise nodes that store all the data. Full storage nodes send block shares, headers, and fraud proofs to light nodes.

```sh
sunrise full init --p2p.network <NETWORK>
sunrise full start --core.ip <URI> --p2p.network <NETWORK>
```

Start a full node using RPC. The bridge and full nodes refer to [bootstrap.go](./nodebuilder/p2p/bootstrap.go).

### Run a light node

Light nodes ensure data availability. This is the most common way to interact with Sunrise networks. It does NOT require large storage or high-speed connections.

```sh
sunrise light init --p2p.network <NAME>
sunrise light start --core.ip <URI> --p2p.network <NAME>
```

## Environment variables

| Variable                | Explanation                         | Default value | Required |
| ----------------------- | ----------------------------------- | ------------- | -------- |
| `CELESTIA_BOOTSTRAPPER` | Start the node in bootstrapper mode | `false`       | Optional |

## Package-specific documentation

- [Header](./header/doc.go)
- [Share](./share/doc.go)
- [DAS](./das/doc.go)

## Code of Conduct

See our Code of Conduct [here](https://docs.celestia.org/community/coc).
