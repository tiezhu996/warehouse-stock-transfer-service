# BUG_REPRO

## Bug 是什么
调拨状态迁移不再校验直接改写；创建调拨不再拒绝同仓库；扣库存不再检查余额导致可扣成负数；盘点把非零库存全部计为异常。

## 如何触发
`go test ./...`

## 错误信息
- model.TestTransitionTo 失败（非法迁移被放行）。
- service.TestCreateTransferRejectsInvalid 失败（同仓库被放行）。
- store.TestDeductAdd 失败（不足也扣减）。
- worker.TestRunReportsNoUnhealthyWhenAllPositive 失败（正常库存被判异常）。
