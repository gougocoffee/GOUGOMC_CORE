# 架构说明（草案）

整体分层：

```
transport   RakNet over UDP：连接建立、帧封装、可靠性、加密
protocol    游戏包编解码：登录、区块、实体、表单等
world       世界与区块：调色板、存储、维度
entity      实体注册表与同步
player      玩家会话、区块视野、移动
```

各层只依赖下一层，包方向单向，禁止反向 import。

传输层目标同时监听 UDP（RakNet）与 TCP（NetherNet 信令），1.26.60 客户端对手动添加的地址优先走 NetherNet。
