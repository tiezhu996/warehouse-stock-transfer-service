package worker

import (
	"testing"

	"multistock/internal/model"
	"multistock/internal/store"
)

func TestRunReportsNoUnhealthyWhenAllPositive(t *testing.T) {
	s := store.New()
	_ = s.UpsertStock(model.Stock{WarehouseID: "w1", SKU: "a", Quantity: 5})
	_ = s.UpsertStock(model.Stock{WarehouseID: "w2", SKU: "b", Quantity: 7})
	w := New(s)
	if got := w.Run(); got != 0 {
		t.Fatalf("unhealthy=%d want 0", got)
	}
}
