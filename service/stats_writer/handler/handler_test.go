package handler

import (
	"encoding/json"
	"testing"
)

func TestTryUnwrap_ValidEnvelope(t *testing.T) {
	env := statEnvelope{
		EventID: "evt-123",
		Payload: json.RawMessage(`{"order_id":1}`),
	}
	raw, _ := json.Marshal(env)

	h := NewStatsHandler(nil, nil)
	got := h.tryUnwrap(raw)
	if got == nil {
		t.Fatal("expected non-nil envelope")
	}
	if got.EventID != "evt-123" {
		t.Fatalf("expected EventID=evt-123, got %q", got.EventID)
	}
}

func TestTryUnwrap_EmptyEventID(t *testing.T) {
	env := statEnvelope{
		EventID: "",
		Payload: json.RawMessage(`{"order_id":1}`),
	}
	raw, _ := json.Marshal(env)

	h := NewStatsHandler(nil, nil)
	got := h.tryUnwrap(raw)
	if got != nil {
		t.Fatalf("expected nil for empty eventID, got %+v", got)
	}
}

func TestTryUnwrap_InvalidJSON(t *testing.T) {
	h := NewStatsHandler(nil, nil)
	got := h.tryUnwrap([]byte("not-json"))
	if got != nil {
		t.Fatalf("expected nil for invalid JSON, got %+v", got)
	}
}

func TestTryUnwrap_RawPayloadNoEnvelope(t *testing.T) {
	h := NewStatsHandler(nil, nil)
	got := h.tryUnwrap([]byte(`{"order_id":1}`))
	if got != nil {
		t.Fatalf("expected nil for raw payload without event_id, got %+v", got)
	}
}

func TestIsDuplicate_FirstCall(t *testing.T) {
	h := NewStatsHandler(nil, nil)
	if h.isDuplicate("evt-001") {
		t.Fatal("first call should not be duplicate")
	}
}

func TestIsDuplicate_SecondCall(t *testing.T) {
	h := NewStatsHandler(nil, nil)
	h.isDuplicate("evt-002")
	if !h.isDuplicate("evt-002") {
		t.Fatal("second call should be duplicate")
	}
}

func TestIsDuplicate_DifferentIDs(t *testing.T) {
	h := NewStatsHandler(nil, nil)
	h.isDuplicate("evt-a")
	if h.isDuplicate("evt-b") {
		t.Fatal("different ID should not be duplicate")
	}
}

func TestIsDuplicate_EmptyID(t *testing.T) {
	h := NewStatsHandler(nil, nil)
	if h.isDuplicate("") {
		t.Fatal("empty eventID should never be duplicate")
	}
	if h.isDuplicate("") {
		t.Fatal("empty eventID should never be duplicate on second call either")
	}
}

func TestStatsTopics(t *testing.T) {
	topics := StatsTopics()
	if len(topics) != 3 {
		t.Fatalf("expected 3 topics, got %d", len(topics))
	}
	expected := map[string]bool{
		"stats.ecommerce.order.event":         false,
		"stats.ecommerce.order_item.event":    false,
		"stats.ecommerce.transaction.event":   false,
	}
	for _, tp := range topics {
		if _, ok := expected[tp]; !ok {
			t.Fatalf("unexpected topic: %q", tp)
		}
		expected[tp] = true
	}
	for tp, found := range expected {
		if !found {
			t.Fatalf("missing topic: %q", tp)
		}
	}
}
