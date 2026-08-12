package lifecyclespec

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

//go:embed schemas/v0.1/*.json
var schemaFiles embed.FS

const schemaBaseURL = "https://lifecycle-spec.local/schemas/v0.1/"

// Validate checks a JSON document against one bundled v0.1 schema.
func Validate(schemaName string, document []byte) error {
	schemaName = strings.TrimSpace(schemaName)
	if schemaName == "" || filepath.Base(schemaName) != schemaName {
		return fmt.Errorf("invalid schema name %q", schemaName)
	}
	if !strings.HasSuffix(schemaName, ".json") {
		schemaName += ".json"
	}
	compiler := jsonschema.NewCompiler()
	compiler.Draft = jsonschema.Draft2020
	entries, err := fs.ReadDir(schemaFiles, "schemas/v0.1")
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		b, err := schemaFiles.ReadFile("schemas/v0.1/" + entry.Name())
		if err != nil {
			return err
		}
		if err = compiler.AddResource(schemaBaseURL+entry.Name(), bytes.NewReader(b)); err != nil {
			return err
		}
	}
	schema, err := compiler.Compile(schemaBaseURL + schemaName)
	if err != nil {
		return fmt.Errorf("compile schema: %w", err)
	}
	var value any
	if err = json.Unmarshal(bytes.TrimPrefix(document, []byte{0xEF, 0xBB, 0xBF}), &value); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}
	if err = schema.Validate(value); err != nil {
		return fmt.Errorf("document does not match %s: %w", schemaName, err)
	}
	return nil
}

// SchemaNames returns the bundled public v0.1 schema file names.
func SchemaNames() ([]string, error) {
	entries, err := fs.ReadDir(schemaFiles, "schemas/v0.1")
	if err != nil {
		return nil, err
	}
	names := []string{}
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".json" && entry.Name() != "defs.json" {
			names = append(names, entry.Name())
		}
	}
	sort.Strings(names)
	return names, nil
}
