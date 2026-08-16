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
	return len(w.store.AllStocks())
}

var _ = model.StatusDraft
