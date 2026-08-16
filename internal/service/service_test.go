package service

import (
	"testing"

	"multistock/internal/model"
	"multistock/internal/store"
)

func TestTransferLifecycle(t *testing.T) {
	st := store.New()
	_ = st.UpsertStock(model.Stock{WarehouseID: "wh-src", SKU: "sku", Quantity: 50})
	_ = st.UpsertStock(model.Stock{WarehouseID: "wh-dst", SKU: "sku", Quantity: 0})
	s := New(st)

	o, err := s.CreateTransfer("wh-src", "wh-dst", "sku", 20)
	if err != nil {
		t.Fatal(err)
	}
	if err := s.Approve(o.ID); err != nil {
		t.Fatal(err)
	}
	if got := s.StockOf("wh-src", "sku"); got != 30 {
		t.Fatalf("source stock=%d want 30", got)
	}
	// Draft->Completed must be rejected; go through InTransit.
	if err := s.Complete(o.ID); err == nil {
		t.Fatal("expected invalid transition Approved->Completed")
	}
	ap, _ := st.GetOrder(o.ID)
	_ = ap.TransitionTo(model.StatusInTransit)
	if err := s.Complete(o.ID); err != nil {
		t.Fatal(err)
	}
	if got := s.StockOf("wh-dst", "sku"); got != 20 {
		t.Fatalf("dest stock=%d want 20", got)
	}
}

func TestCreateTransferRejectsInvalid(t *testing.T) {
	st := store.New()
	_ = st.UpsertStock(model.Stock{WarehouseID: "a", SKU: "sku", Quantity: 10})
	s := New(st)
	if _, err := s.CreateTransfer("a", "a", "sku", 5); err == nil {
		t.Fatal("expected same-warehouse error")
	}
	if _, err := s.CreateTransfer("a", "b", "sku", 0); err == nil {
		t.Fatal("expected non-positive quantity error")
	}
	if _, err := s.CreateTransfer("missing", "b", "sku", 5); err == nil {
		t.Fatal("expected missing stock error")
	}
}
