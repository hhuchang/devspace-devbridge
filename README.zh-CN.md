# DevBridge

> English: [README.md](README.md)

DevBridge 是华为云的开发隧道服务，通过中继连接 Host（本地服务托管）与 Connect（远端连接），使本地服务无需公网入站端口即可被远程访问。本仓库为 monorepo，各组件使用独立首层目录以避免工具链相互影响。

在线文档：<https://huaweicloud.github.io/devspace-devbridge/>

## Repository structure

```text
.
├── .github/              # 仓库级 CI、发布和同步工作流
├── cli/                  # DevBridge CLI（Go，面向终端用户）
├── go-sdk/               # DevBridge Go SDK（Go，面向三方集成）
├── docs/                 # VitePress 文档站点及其 Node.js 工具链
└── README.md             # 仓库总览
```

各组件使用独立首层目录，构建配置与依赖文件保存在对应目录中，避免工具链相互影响。

## 项目组成

| 组件   | 说明                                                    | 详情                                 |
| ------ | ------------------------------------------------------- | ------------------------------------ |
| CLI    | 华为云 DevStation 远程隧道端口转发命令行工具（Go 1.26） | [cli/README.md](cli/README.md)       |
| Go SDK | 将 CLI 核心能力封装为 Go API，供三方集成（Go 1.22）     | [go-sdk/README.md](go-sdk/README.md) |

## Documentation

文档站点源码位于 `docs/`，由 VitePress 生成并发布到 GitHub Pages。本地预览、新增页面、依赖策略与发布流程见 [docs/README.md](docs/README.md)。
