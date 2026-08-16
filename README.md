# multistock

多仓库库存管理核心服务（纯 Go、内存存储），覆盖库存查询、跨仓调拨与库存盘点对账。

## 目录结构
```
cmd/multistock/       程序入口
internal/config/      环境配置
internal/model/       领域模型与状态机
internal/store/       内存库存存储
internal/service/     调拨与盘点业务
internal/worker/      盘点对账 worker
```

## 运行与测试
```bash
go build ./...
go test ./...
```
