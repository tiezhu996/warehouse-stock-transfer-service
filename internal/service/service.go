package service

import (
	"fmt"

	"multistock/internal/model"
	"multistock/internal/store"
)

type Service struct {
	store *store.Store
}

func New(s *store.Store) *Service { return &Service{store: s} }

func (s *Service) CreateTransfer(src, dst, sku string, qty int) (*model.TransferOrder, error) {
	if src == "" || dst == "" || sku == "" {
		return nil, fmt.Errorf("source, dest and sku required")
	}
	if src == dst {
		return nil, fmt.Errorf("source and dest warehouses must differ")
	}
	if qty <= 0 {
		return nil, fmt.Errorf("quantity must be positive")
	}
	if _, ok := s.store.GetStock(src, sku); !ok {
		return nil, fmt.Errorf("source stock not found")
	}
	o := &model.TransferOrder{
		SourceWarehouseID: src,
		DestWarehouseID:   dst,
		SKU:               sku,
		Quantity:          qty,
		Status:            model.StatusDraft,
	}
	s.store.SaveOrder(o)
	return o, nil
}

func (s *Service) Approve(id string) error {
	o, ok := s.store.GetOrder(id)
	if !ok {
		return fmt.Errorf("order not found")
	}
	if !o.TransitionTo(model.StatusApproved) {
		return fmt.Errorf("invalid transition to Approved")
	}
	if err := s.store.Deduct(o.SourceWarehouseID, o.SKU, o.Quantity); err != nil {
		return err
	}
	return nil
}

func (s *Service) Complete(id string) error {
	o, ok := s.store.GetOrder(id)
	if !ok {
		return fmt.Errorf("order not found")
	}
	if !o.TransitionTo(model.StatusCompleted) {
		return fmt.Errorf("invalid transition to Completed")
	}
	if err := s.store.Add(o.DestWarehouseID, o.SKU, o.Quantity); err != nil {
		return err
	}
	return nil
}

func (s *Service) StockOf(warehouseID, sku string) int {
	st, ok := s.store.GetStock(warehouseID, sku)
	if !ok {
		return 0
	}
	return st.Quantity
}
