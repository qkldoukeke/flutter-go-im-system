# flutter-go-im-system

即时通讯系统（Flutter + Go + PostgreSQL + Redis）项目样板。

## 目录结构

```
backend/           # Go 后端按微服务拆分
client/flutter_im/ # Flutter 客户端
db/                 # 数据库结构和 key 说明
docs/               # 设计文档
```

## 快速启动
1. 启动 PostgreSQL/Redis
2. backend 各服务 go run/main.go
3. Flutter 运行 client/flutter_im

详细见各服务 README。