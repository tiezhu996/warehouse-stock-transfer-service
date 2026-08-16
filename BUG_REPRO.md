# BUG_REPRO

## Bug 是什么
跨仓调拨状态机允许 Draft 直接跳到 Completed；创建调拨时把同仓库校验反向、数量只拦负数；审批/完成时源仓库与目标仓库用反；库存盘点把非负库存全部统计为异常。

## 如何触发
`go test ./...`

## 错误信息
- model.TestCanTransition 失败（Draft->Completed 被判为合法）。
- service.TestTransferLifecycle / TestCreateTransferRejectsInvalid 失败（同仓库、零数量被放行，库存反向）。
- store.TestDeductAdd 失败（充足库存被判为不足）。
- worker.TestRunReportsNoUnhealthyWhenAllPositive 失败（正常库存被判异常）。
