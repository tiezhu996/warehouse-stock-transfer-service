package model

type Stock struct {
	WarehouseID string
	SKU         string
	Quantity    int
}

type TransferStatus string

const (
	StatusDraft     TransferStatus = "Draft"
	StatusApproved  TransferStatus = "Approved"
	StatusInTransit TransferStatus = "InTransit"
	StatusCompleted TransferStatus = "Completed"
	StatusCancelled TransferStatus = "Cancelled"
)

type TransferOrder struct {
	ID                string
	SourceWarehouseID string
	DestWarehouseID   string
	SKU               string
	Quantity          int
	Status            TransferStatus
}

var validTransitions = map[TransferStatus][]TransferStatus{
	StatusDraft:     {StatusApproved, StatusCancelled},
	StatusApproved:  {StatusInTransit, StatusCancelled},
	StatusInTransit: {StatusCompleted},
}

func CanTransition(from, to TransferStatus) bool {
	for _, next := range validTransitions[from] {
		if next == to {
			return true
		}
	}
	return false
}

func (o *TransferOrder) TransitionTo(to TransferStatus) bool {
	if !CanTransition(o.Status, to) {
		return false
	}
	o.Status = to
	return true
}
