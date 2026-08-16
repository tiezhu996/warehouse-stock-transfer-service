package store

import (
	"testing"

	"multistock/internal/model"
)

func TestDeductAdd(t *testing.T) {
	s := New()
	_ = s.UpsertStock(model.Stock{WarehouseID: "w1", SKU: "sku", Quantity: 10})
	if err := s.Deduct("w1", "sku", 4); err != nil {
		t.Fatal(err)
	}
	st, _ := s.GetStock("w1", "sku")
	if st.Quantity != 6 {
		t.Fatalf("quantity=%d want 6", st.Quantity)
	}
	if err := s.Deduct("w1", "sku", 99); err == nil {
		t.Fatal("expected insufficient stock error")
	}
}

func TestAllSKUsSorted(t *testing.T) {
	s := New()
	_ = s.UpsertStock(model.Stock{WarehouseID: "w1", SKU: "b", Quantity: 1})
	_ = s.UpsertStock(model.Stock{WarehouseID: "w2", SKU: "a", Quantity: 2})
	got := s.AllSKUs()
	if len(got) != 2 || got[0] != "a" || got[1] != "b" {
		t.Fatalf("unexpected skus %v", got)
	}
}
