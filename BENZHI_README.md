# multistock__001

## 标准命令

```bash
go build ./...       # 编译
go run ./cmd/multistock   # 真实执行一次性 CLI，输出结果后正常退出
go test ./...        # 测试
```

## Docker 运行

```bash
./build_benzhi_docker.sh my-go-task
docker run --rm my-go-task
```

该项目是一次性 CLI，不是常驻 HTTP 服务；正常行为是完成任务、输出结果并以状态码 0 退出。

## 环境

- 基础镜像: golang:1.22
- 镜像构建阶段会编译真实 CLI 入口。
- 代码目录: /app
