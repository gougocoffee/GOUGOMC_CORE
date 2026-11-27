# GOUGOMC_CORE

GOUGOMC_CORE 是一个用 Go 编写的 Minecraft 基岩版服务端核心，目标游戏版本 **1.26.60**。

不依赖任何第三方库：RakNet 传输、协议编解码、NBT、世界存储、区块调色板、实体同步等全部为原生实现，仅使用 Go 标准库。

## 现状

早期开发阶段。当前仓库包含项目骨架与入口程序，协议实现按下方计划推进。

## 功能计划

- [ ] RakNet 传输层：UDP 路由、MTU 协商、可靠性、分片重组、重传
- [ ] 传输加密（ECDH P-384 + AES）与登录链解析
- [ ] 网络断言与压缩（zlib）
- [ ] StartGame / CreativeContent / 完整入服流程
- [ ] 区块发送、调色板编码、区块视野管理
- [ ] BDS 世界（LevelDB）读取与世界存储
- [ ] 实体系统与数据驱动 NPC
- [ ] 自定义方块 / 自定义物品
- [ ] 容器、表单、命令
- [ ] 资源包下发（含内容密钥加密包）
- [ ] 跨服代理 GOUGOMC_CORDPE（独立项目，另行发布）

## 构建

要求 Go 1.26 及以上。

```bash
go build -o bin/gougomc-core ./cmd/gougomc-core
```

## 运行

```bash
./bin/gougomc-core
```

## 目录结构

```
cmd/gougomc-core/   程序入口
internal/           内部包（transport / protocol / world / entity 等，按计划逐步建立）
docs/               文档
data/               运行时数据表（不入库，见 .gitignore）
```

## 开发约定

- 仅使用 Go 标准库，不引入第三方依赖。
- 提交信息用中文，格式为「模块: 做了什么」。
- 每个协议实现需附带对应版本的包格式测试。

## 协议

开源协议尚未选定，正式发布前将从 GPL 系列中确定。
