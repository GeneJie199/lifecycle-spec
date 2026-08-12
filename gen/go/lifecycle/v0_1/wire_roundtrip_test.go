package v0_1

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	lifecyclespec "github.com/GeneJie199/lifecycle-spec"
)

func TestEveryPublicWireTypeRoundTripsItsExample(t *testing.T) {
	cases := map[string]func() any{
		"acceptance-criterion":      func() any { return new(AcceptanceCriterion) },
		"agent-session":             func() any { return new(AgentSession) },
		"approval":                  func() any { return new(Approval) },
		"change-event":              func() any { return new(ChangeEvent) },
		"code-change":               func() any { return new(CodeChange) },
		"event-batch":               func() any { return new(EventBatch) },
		"event-envelope":            func() any { return new(EventEnvelope) },
		"evidence":                  func() any { return new(Evidence) },
		"fleet-node-report":         func() any { return new(FleetNodeReport) },
		"incident":                  func() any { return new(Incident) },
		"monitoring-plan":           func() any { return new(MonitoringPlan) },
		"monitoring-recommendation": func() any { return new(MonitoringRecommendation) },
		"observation":               func() any { return new(Observation) },
		"project":                   func() any { return new(Project) },
		"relationship":              func() any { return new(Relationship) },
		"release-candidate":         func() any { return new(DevelopmentReleaseCandidate) },
		"release-validation-report": func() any { return new(ReleaseValidationReport) },
		"release":                   func() any { return new(Release) },
		"requirement":               func() any { return new(Requirement) },
		"resource":                  func() any { return new(Resource) },
		"snapshot":                  func() any { return new(Snapshot) },
		"task":                      func() any { return new(Task) },
		"telemetry-batch":           func() any { return new(TelemetryBatch) },
	}
	_, source, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("locate test source")
	}
	root := filepath.Join(filepath.Dir(source), "..", "..", "..", "..")
	for name, makeValue := range cases {
		name, makeValue := name, makeValue
		t.Run(name, func(t *testing.T) {
			raw, err := os.ReadFile(filepath.Join(root, "examples", "v0.1", name+".json"))
			if err != nil {
				t.Fatal(err)
			}
			raw = bytes.TrimPrefix(raw, []byte{0xEF, 0xBB, 0xBF})
			value := makeValue()
			if err := json.Unmarshal(raw, value); err != nil {
				t.Fatalf("unmarshal into generated type: %v", err)
			}
			out, err := json.Marshal(value)
			if err != nil {
				t.Fatalf("marshal generated type: %v", err)
			}
			if err := lifecyclespec.Validate(name, out); err != nil {
				t.Fatalf("generated type output fails schema: %v", err)
			}
			var before, after any
			if err := json.Unmarshal(raw, &before); err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal(out, &after); err != nil {
				t.Fatal(err)
			}
			assertNoMeaningfulDataLoss(t, "$", before, after)
		})
	}
}

func assertNoMeaningfulDataLoss(t *testing.T, path string, before, after any) {
	t.Helper()
	switch before := before.(type) {
	case map[string]any:
		afterMap, ok := after.(map[string]any)
		if !ok {
			t.Fatalf("%s changed object shape", path)
		}
		for key, value := range before {
			got, exists := afterMap[key]
			if !exists {
				if !isEmptyJSON(value) {
					t.Fatalf("%s.%s was lost by generated type", path, key)
				}
				continue
			}
			assertNoMeaningfulDataLoss(t, path+"."+key, value, got)
		}
	case []any:
		afterSlice, ok := after.([]any)
		if !ok || len(before) != len(afterSlice) {
			t.Fatalf("%s changed array length", path)
		}
		for index := range before {
			assertNoMeaningfulDataLoss(t, path, before[index], afterSlice[index])
		}
	default:
		if !reflect.DeepEqual(before, after) {
			t.Fatalf("%s changed from %#v to %#v", path, before, after)
		}
	}
}

func isEmptyJSON(value any) bool {
	switch value := value.(type) {
	case string:
		return value == ""
	case []any:
		return len(value) == 0
	case map[string]any:
		return len(value) == 0
	case nil:
		return true
	default:
		return false
	}
}
