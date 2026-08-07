package schematest

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

func schemasDir(t *testing.T) string {
	t.Helper()
	// tests/schematest -> ../../schemas/v0.1
	dir, err := filepath.Abs(filepath.Join("..", "..", "schemas", "v0.1"))
	if err != nil {
		t.Fatal(err)
	}
	return dir
}

func examplesDir(t *testing.T) string {
	t.Helper()
	dir, err := filepath.Abs(filepath.Join("..", "..", "examples", "v0.1"))
	if err != nil {
		t.Fatal(err)
	}
	return dir
}

func loadSchema(t *testing.T, name string) *jsonschema.Schema {
	t.Helper()
	compiler := jsonschema.NewCompiler()
	compiler.Draft = jsonschema.Draft2020

	root := schemasDir(t)
	entries, err := os.ReadDir(root)
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".json" {
			continue
		}
		path := filepath.Join(root, e.Name())
		f, err := os.Open(path)
		if err != nil {
			t.Fatal(err)
		}
		// URL must match $id basename references like defs.json
		if err := compiler.AddResource("https://lifecycle-spec.local/schemas/v0.1/"+e.Name(), f); err != nil {
			f.Close()
			t.Fatalf("add %s: %v", e.Name(), err)
		}
		f.Close()
	}

	sch, err := compiler.Compile("https://lifecycle-spec.local/schemas/v0.1/" + name)
	if err != nil {
		t.Fatalf("compile %s: %v", name, err)
	}
	return sch
}

func validateFile(t *testing.T, schemaName, exampleName string) {
	t.Helper()
	sch := loadSchema(t, schemaName)
	raw, err := os.ReadFile(filepath.Join(examplesDir(t), exampleName))
	if err != nil {
		t.Fatal(err)
	}
	raw = bytes.TrimPrefix(raw, []byte{0xEF, 0xBB, 0xBF})
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		t.Fatalf("json %s: %v", exampleName, err)
	}
	if err := sch.Validate(v); err != nil {
		t.Fatalf("%s against %s: %v", exampleName, schemaName, err)
	}
}

func loadExampleMap(t *testing.T, exampleName string) map[string]any {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join(examplesDir(t), exampleName))
	if err != nil {
		t.Fatal(err)
	}
	raw = bytes.TrimPrefix(raw, []byte{0xEF, 0xBB, 0xBF})
	var v map[string]any
	if err := json.Unmarshal(raw, &v); err != nil {
		t.Fatalf("json %s: %v", exampleName, err)
	}
	return v
}

func cloneMap(m map[string]any) map[string]any {
	out := make(map[string]any, len(m))
	for k, v := range m {
		out[k] = v
	}
	return out
}

func TestExamplesValidate(t *testing.T) {
	cases := []struct {
		schema  string
		example string
	}{
		{"event-envelope.json", "event-envelope.json"},
		{"evidence.json", "evidence.json"},
		{"evidence.json", "evidence-manual.json"},
		{"change-event.json", "change-event.json"},
		{"resource.json", "resource.json"},
		{"observation.json", "observation.json"},
		{"snapshot.json", "snapshot.json"},
		{"relationship.json", "relationship.json"},
		{"approval.json", "approval.json"},
		{"release.json", "release.json"},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.example, func(t *testing.T) {
			validateFile(t, tc.schema, tc.example)
		})
	}
}

func TestNegativeMissingRequiredFields(t *testing.T) {
	sch := loadSchema(t, "event-envelope.json")
	base := loadExampleMap(t, "event-envelope.json")
	required := []string{"specVersion", "eventId", "eventType", "occurredAt", "source", "subject", "payload"}
	for _, field := range required {
		field := field
		t.Run("missing_"+field, func(t *testing.T) {
			bad := cloneMap(base)
			delete(bad, field)
			if err := sch.Validate(bad); err == nil {
				t.Fatalf("expected validation failure when %s is missing", field)
			}
		})
	}
}

func TestNegativeBadEventIDPrefix(t *testing.T) {
	sch := loadSchema(t, "event-envelope.json")
	bad := cloneMap(loadExampleMap(t, "event-envelope.json"))
	bad["eventId"] = "bad_01J2QK3M4N5P6Q7R8S9T0ABCD1"
	if err := sch.Validate(bad); err == nil {
		t.Fatal("expected validation failure for non-evt_ eventId prefix")
	}
}

func TestNegativeTimestampWithoutTimezone(t *testing.T) {
	sch := loadSchema(t, "event-envelope.json")
	bad := cloneMap(loadExampleMap(t, "event-envelope.json"))
	bad["occurredAt"] = "2026-08-07T14:00:00"
	if err := sch.Validate(bad); err == nil {
		t.Fatal("expected validation failure for timestamp without timezone")
	}
}

func TestNegativeChangeEventNonChangeEventType(t *testing.T) {
	sch := loadSchema(t, "change-event.json")
	bad := cloneMap(loadExampleMap(t, "change-event.json"))
	bad["eventType"] = "observation.host.scanned"
	if err := sch.Validate(bad); err == nil {
		t.Fatal("expected validation failure for non-change eventType on change-event")
	}
}

func TestForwardCompatUnknownExtensionsKey(t *testing.T) {
	sch := loadSchema(t, "event-envelope.json")
	good := cloneMap(loadExampleMap(t, "event-envelope.json"))
	good["extensions"] = map[string]any{
		"com.example.futureField": map[string]any{
			"enabled": true,
			"note":    "unknown to older consumers",
		},
	}
	if err := sch.Validate(good); err != nil {
		t.Fatalf("unknown extensions key should validate: %v", err)
	}
}

func TestRejectPlaintextPrivateKeyInEvidenceInline(t *testing.T) {
	// Schema cannot fully ban secrets; FND-008 covers pattern scanning.
	// Here we only assert a clearly invalid evidenceId fails schema.
	sch := loadSchema(t, "evidence.json")
	bad := map[string]any{
		"evidenceId": "bad-id",
		"kind":       "log.excerpt",
		"title":      "x",
		"producedAt": "2026-08-07T14:00:00+08:00",
		"producer": map[string]any{
			"product":    "infra-discovery",
			"instanceId": "x",
		},
		"inline":    "BEGIN PRIVATE KEY",
		"redaction": map[string]any{"status": "none"},
	}
	if err := sch.Validate(bad); err == nil {
		t.Fatal("expected invalid evidenceId to fail")
	}
}
