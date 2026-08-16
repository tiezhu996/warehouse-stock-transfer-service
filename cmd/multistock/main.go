package main

import (
	"fmt"

	"multistock/internal/config"
	"multistock/internal/model"
	"multistock/internal/service"
	"multistock/internal/store"
	"multistock/internal/worker"
)

func main() {
	cfg := config.Load()
	st := store.New()
	svc := service.New(st)
	w := worker.New(st)

	if err := st.UpsertStock(model.Stock{WarehouseID: "wh-1", SKU: "sku-a", Quantity: 100}); err != nil {
		panic(err)
	}
	if _, err := svc.CreateTransfer("wh-1", "wh-2", "sku-a", 30); err != nil {
		panic(err)
	}
	fmt.Printf("%s worker=%d\n", cfg.AppName, w.Run())
}
