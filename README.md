# DevBridge

> 中文文档：[README.zh-CN.md](README.zh-CN.md)

DevBridge is a Huawei Cloud development tunneling service. It relays between a Host (local
service hosting) and a Connect (remote connection) so that local services can be reached
remotely without opening any public inbound port. This repository is a monorepo; each
component lives in its own top-level directory to avoid toolchain conflicts.

Online documentation: <https://huaweicloud.github.io/devspace-devbridge/>

## Repository structure

```text
.
├── .github/              # repo-level CI, release and sync workflows
├── cli/                  # DevBridge CLI (Go, for end users)
├── go-sdk/               # DevBridge Go SDK (Go, for third-party integration)
├── docs/                 # VitePress docs site and its Node.js toolchain
└── README.md             # repository overview
```

Each component uses its own top-level directory, keeping build configs and dependency files
within that directory to avoid toolchain conflicts.

## Components

| Component | Description                                                           | Details                              |
| --------- | --------------------------------------------------------------------- | ------------------------------------ |
| CLI       | Huawei Cloud DevStation remote tunnel port-forwarding CLI (Go 1.26)   | [cli/README.md](cli/README.md)       |
| Go SDK    | Wraps the CLI core capabilities as a Go API for integration (Go 1.22) | [go-sdk/README.md](go-sdk/README.md) |

## Documentation

The docs site source lives in `docs/`, built with VitePress and published to GitHub Pages.
For local preview, adding pages, dependency policy and publishing, see
[docs/README.md](docs/README.md).
