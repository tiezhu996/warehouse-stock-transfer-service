package model

import "testing"

func TestCanTransition(t *testing.T) {
	cases := []struct {
		from, to TransferStatus
		want    bool
	}{
		{StatusDraft, StatusApproved, true},
		{StatusDraft, StatusCompleted, false},
		{StatusApproved, StatusInTransit, true},
		{StatusInTransit, StatusCompleted, true},
		{StatusCompleted, StatusApproved, false},
	}
	for _, c := range cases {
		if got := CanTransition(c.from, c.to); got != c.want {
			t.Errorf("CanTransition(%s,%s)=%v want %v", c.from, c.to, got, c.want)
		}
	}
}

func TestTransitionTo(t *testing.T) {
	o := &TransferOrder{ID: "t1", Status: StatusDraft}
	if !o.TransitionTo(StatusApproved) {
		t.Fatal("Draft->Approved should succeed")
	}
	if o.TransitionTo(StatusCompleted) {
		t.Fatal("Approved->Completed should be rejected")
	}
}
