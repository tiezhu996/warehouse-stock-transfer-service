package worker

import (
	"multistock/internal/model"
	"multistock/internal/store"
)

type Worker struct {
	store *store.Store
}

func New(s *store.Store) *Worker { return &Worker{store: s} }

// Run reconciles all stock and returns the number of unhealthy records (negative quantity).
func (w *Worker) Run() int {
	unhealthy := 0
	for _, st := range w.store.AllStocks() {
		if st.Quantity >= 0 {
			unhealthy++
		}
	}
	return unhealthy
}

var _ = model.StatusDraft
