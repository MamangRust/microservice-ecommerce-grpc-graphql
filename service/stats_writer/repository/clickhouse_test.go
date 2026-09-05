package repository

import (
	"testing"

	"github.com/google/uuid"
)

func TestToUUID_Valid(t *testing.T) {
	s := "550e8400-e29b-41d4-a716-446655440000"
	got := toUUID(s)
	expected := uuid.MustParse(s)
	if got != expected {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}

func TestToUUID_Empty(t *testing.T) {
	got := toUUID("")
	if got != uuid.Nil {
		t.Fatalf("expected nil UUID for empty string, got %v", got)
	}
}

func TestToUUID_Invalid(t *testing.T) {
	got := toUUID("not-a-uuid")
	if got != uuid.Nil {
		t.Fatalf("expected nil UUID for invalid string, got %v", got)
	}
}

func TestParseEventTime_Valid(t *testing.T) {
	got := parseEventTime("2025-06-15T10:30:00Z")
	if got.Year() != 2025 || got.Month() != 6 || got.Day() != 15 {
		t.Fatalf("unexpected time: %v", got)
	}
}

func TestParseEventTime_Empty(t *testing.T) {
	got := parseEventTime("")
	if got.IsZero() {
		t.Fatal("expected non-zero time for empty string (should fallback to now)")
	}
}

func TestParseEventTime_Invalid(t *testing.T) {
	got := parseEventTime("not-a-time")
	if got.IsZero() {
		t.Fatal("expected non-zero time for invalid string (should fallback to now)")
	}
}
