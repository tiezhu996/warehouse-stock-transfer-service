package store

import (
	"fmt"
	"sort"
	"sync"

	"multistock/internal/model"
)

type Store struct {
	mu      sync.Mutex
	stocks  map[string]model.Stock
	orders  map[string]*model.TransferOrder
	seq     int
}

func New() *Store {
	return &Store{
		stocks: map[string]model.Stock{},
		orders: map[string]*model.TransferOrder{},
	}
}

func stockKey(warehouseID, sku string) string { return warehouseID + "|" + sku }

func (s *Store) UpsertStock(st model.Stock) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if st.WarehouseID == "" || st.SKU == "" {
		return fmt.Errorf("warehouse and sku required")
	}
	if st.Quantity < 0 {
		return fmt.Errorf("negative quantity")
	}
	s.stocks[stockKey(st.WarehouseID, st.SKU)] = st
	return nil
}

func (s *Store) GetStock(warehouseID, sku string) (model.Stock, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	st, ok := s.stocks[stockKey(warehouseID, sku)]
	return st, ok
}

// Deduct removes quantity from the source warehouse and must fail if insufficient.
func (s *Store) Deduct(warehouseID, sku string, qty int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	st, ok := s.stocks[stockKey(warehouseID, sku)]
	if !ok {
		return fmt.Errorf("stock not found")
	}
	if st.Quantity < qty {
		return fmt.Errorf("insufficient stock")
	}
	st.Quantity -= qty
	s.stocks[stockKey(warehouseID, sku)] = st
	return nil
}

func (s *Store) Add(warehouseID, sku string, qty int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	st, ok := s.stocks[stockKey(warehouseID, sku)]
	if !ok {
		st = model.Stock{WarehouseID: warehouseID, SKU: sku}
	}
	st.Quantity += qty
	s.stocks[stockKey(warehouseID, sku)] = st
	return nil
}

func (s *Store) SaveOrder(o *model.TransferOrder) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if o.ID == "" {
		s.seq++
		o.ID = fmt.Sprintf("TO-%03d", s.seq)
	}
	s.orders[o.ID] = o
}

func (s *Store) GetOrder(id string) (*model.TransferOrder, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	o, ok := s.orders[id]
	return o, ok
}

// AllSKUs returns the union of distinct SKUs sorted ascending.
func (s *Store) AllSKUs() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	set := map[string]struct{}{}
	for _, st := range s.stocks {
		set[st.SKU] = struct{}{}
	}
	out := make([]string, 0, len(set))
	for sku := range set {
		out = append(out, sku)
	}
	sort.Strings(out)
	return out
}

func (s *Store) AllStocks() []model.Stock {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]model.Stock, 0, len(s.stocks))
	for _, st := range s.stocks {
		out = append(out, st)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].WarehouseID != out[j].WarehouseID {
			return out[i].WarehouseID < out[j].WarehouseID
		}
		return out[i].SKU < out[j].SKU
	})
	return out
}
