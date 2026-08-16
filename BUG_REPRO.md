# BUG_REPRO

## Bug 是什么
配置读取忽略环境变量返回空；状态迁移表错乱；审批跳过状态校验且扣库存改为加库存；盘点直接返回全部记录数。

## 如何触发
`go test ./...`

## 错误信息
- config.TestLoadDefaults 失败（AppName 为空）。
- model.TestCanTransition / TestTransitionTo 失败（迁移表错乱）。
- service.TestTransferLifecycle 失败（库存反向）。
- store.TestDeductAdd 失败（扣减变增加）。
- worker.TestRunReportsNoUnhealthyWhenAllPositive 失败（全部计异常）。
