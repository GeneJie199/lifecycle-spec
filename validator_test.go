package lifecyclespec

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestGeneratedSchemaManifestIsCurrent(t *testing.T) {
	command := exec.Command("go", "run", "./cmd/lifecycle-gen", "--check")
	if output, err := command.CombinedOutput(); err != nil {
		t.Fatalf("generated contract manifest is stale: %v\n%s", err, output)
	}
	data, err := os.ReadFile(filepath.Join("gen", "schema-manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	var manifest struct {
		Schemas map[string]json.RawMessage `json:"schemas"`
	}
	if err := json.Unmarshal(data, &manifest); err != nil {
		t.Fatal(err)
	}
	names, err := SchemaNames()
	if err != nil {
		t.Fatal(err)
	}
	if len(manifest.Schemas) != len(names) {
		t.Fatalf("manifest covers %d schemas, want %d", len(manifest.Schemas), len(names))
	}
}

func TestBundledValidator(t *testing.T) {
	b, err := os.ReadFile(filepath.Join("examples", "v0.1", "release-candidate.json"))
	if err != nil {
		t.Fatal(err)
	}
	if err = Validate("release-candidate.json", b); err != nil {
		t.Fatal(err)
	}
	if err = Validate("release-candidate", b); err != nil {
		t.Fatal(err)
	}
	if err = Validate("missing.json", b); err == nil {
		t.Fatal("missing schema should fail")
	}
}

func TestSchemaNames(t *testing.T) {
	names, err := SchemaNames()
	if err != nil {
		t.Fatal(err)
	}
	if len(names) < 10 || names[0] == "defs.json" {
		t.Fatalf("schema names = %v", names)
	}
}
