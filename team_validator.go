package lifecyclespec

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

//go:embed schemas/v2/*.json
var teamSchemas embed.FS

// ValidateV2 validates an additive development contract without changing v0.1 lookup.
// All references are embedded; the compiler cannot fetch remote schemas.
func ValidateV2(name string, document []byte) error {
	if name == "" || filepath.Base(name) != name || strings.ContainsAny(name, "/\\") {
		return fmt.Errorf("invalid v2 schema name")
	}
	if !strings.HasSuffix(name, ".json") {
		name += ".json"
	}
	compiler := jsonschema.NewCompiler()
	compiler.Draft = jsonschema.Draft2020
	compiler.LoadURL = func(string) (io.ReadCloser, error) { return nil, fmt.Errorf("remote schema loading disabled") }
	entries, err := fs.ReadDir(teamSchemas, "schemas/v2")
	if err != nil {
		return err
	}
	const base = "https://lifecycle-spec.local/schemas/v2/"
	for _, entry := range entries {
		raw, err := teamSchemas.ReadFile("schemas/v2/" + entry.Name())
		if err != nil {
			return err
		}
		if err = compiler.AddResource(base+entry.Name(), bytes.NewReader(raw)); err != nil {
			return err
		}
	}
	schema, err := compiler.Compile(base + name)
	if err != nil {
		return err
	}
	var value any
	if err = json.Unmarshal(bytes.TrimPrefix(document, []byte{0xef, 0xbb, 0xbf}), &value); err != nil {
		return err
	}
	return schema.Validate(value)
}
