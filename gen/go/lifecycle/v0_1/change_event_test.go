package v0_1

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestChangeEventRoundTrip(t *testing.T) {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	example := filepath.Join(filepath.Dir(thisFile), "..", "..", "..", "..", "examples", "v0.1", "change-event.json")
	raw, err := os.ReadFile(example)
	if err != nil {
		t.Fatalf("read example: %v", err)
	}
	raw = bytes.TrimPrefix(raw, []byte{0xEF, 0xBB, 0xBF})

	var ev ChangeEvent
	if err := json.Unmarshal(raw, &ev); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if ev.EventID == "" || ev.EventType == "" || ev.Payload.ChangeKind == "" {
		t.Fatalf("missing required fields after unmarshal: %+v", ev)
	}
	if ev.EventType != "change.resource.modified" {
		t.Fatalf("eventType=%q", ev.EventType)
	}
	if ev.Payload.ChangeKind != ChangeModified {
		t.Fatalf("changeKind=%q", ev.Payload.ChangeKind)
	}

	out, err := json.Marshal(ev)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var again ChangeEvent
	if err := json.Unmarshal(out, &again); err != nil {
		t.Fatalf("re-unmarshal: %v", err)
	}
	if again.EventID != ev.EventID || again.Payload.ChangeKind != ev.Payload.ChangeKind {
		t.Fatalf("round-trip mismatch: got %+v want %+v", again, ev)
	}
}
