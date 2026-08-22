package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
)

type schemaContract struct {
	GoType     string `json:"go_type"`
	TypeScript string `json:"typescript_type"`
	SHA256     string `json:"sha256"`
}

type manifest struct {
	Version   int                       `json:"version"`
	Schemas   map[string]schemaContract `json:"schemas"`
	Artifacts map[string]string         `json:"artifacts"`
}

var contractTypes = map[string][2]string{
	"acceptance-criterion.json":      {"AcceptanceCriterion", "AcceptanceCriterion"},
	"agent-session.json":             {"AgentSession", "AgentSession"},
	"approval.json":                  {"Approval", "Approval"},
	"change-event.json":              {"ChangeEvent", "ChangeEvent"},
	"code-change.json":               {"CodeChange", "CodeChange"},
	"event-batch.json":               {"EventBatch", "EventBatch"},
	"event-envelope.json":            {"EventEnvelope", "EventEnvelope"},
	"evidence.json":                  {"Evidence", "Evidence"},
	"expected-changes.json":          {"ExpectedChanges", "ExpectedChanges"},
	"fleet-node-report.json":         {"FleetNodeReport", "FleetNodeReport"},
	"incident.json":                  {"Incident", "Incident"},
	"monitoring-plan.json":           {"MonitoringPlan", "MonitoringPlan"},
	"monitoring-recommendation.json": {"MonitoringRecommendation", "MonitoringRecommendation"},
	"observation.json":               {"Observation", "Observation"},
	"project.json":                   {"Project", "Project"},
	"relationship.json":              {"Relationship", "Relationship"},
	"release-candidate.json":         {"DevelopmentReleaseCandidate", "DevelopmentReleaseCandidate"},
	"release-validation-report.json": {"ReleaseValidationReport", "ReleaseValidationReport"},
	"release.json":                   {"Release", "Release"},
	"requirement.json":               {"Requirement", "Requirement"},
	"resource.json":                  {"Resource", "Resource"},
	"snapshot.json":                  {"Snapshot", "Snapshot"},
	"task.json":                      {"Task", "Task"},
	"telemetry-batch.json":           {"TelemetryBatch", "TelemetryBatch"},
}

func main() {
	write := flag.Bool("write", false, "write gen/schema-manifest.json")
	check := flag.Bool("check", false, "verify the checked-in manifest")
	flag.Parse()
	if !*write && !*check {
		*check = true
	}
	root, err := repositoryRoot()
	if err != nil {
		fatal(err)
	}
	want, err := buildManifest(root)
	if err != nil {
		fatal(err)
	}
	manifestPath := filepath.Join(root, "gen", "schema-manifest.json")
	if *write {
		data, err := json.MarshalIndent(want, "", "  ")
		if err != nil {
			fatal(err)
		}
		if err := os.WriteFile(manifestPath, append(data, '\n'), 0o644); err != nil {
			fatal(err)
		}
	}
	if *check {
		data, err := os.ReadFile(manifestPath)
		if err != nil {
			fatal(fmt.Errorf("read manifest: %w; run go run ./cmd/lifecycle-gen --write", err))
		}
		var got manifest
		if err := json.Unmarshal(data, &got); err != nil {
			fatal(fmt.Errorf("decode manifest: %w", err))
		}
		if !reflect.DeepEqual(got, want) {
			fatal(fmt.Errorf("schema or generated types changed; update both wire types, run tests, then run go run ./cmd/lifecycle-gen --write"))
		}
	}
}

func repositoryRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		data, readErr := os.ReadFile(filepath.Join(dir, "go.mod"))
		if readErr == nil && strings.Contains(string(data), "module github.com/GeneJie199/lifecycle-spec") {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("lifecycle-spec repository root not found")
		}
		dir = parent
	}
}

func buildManifest(root string) (manifest, error) {
	goPath := filepath.Join(root, "gen", "go", "lifecycle", "v0_1", "types.go")
	tsPath := filepath.Join(root, "gen", "ts", "v0.1", "index.ts")
	goData, err := os.ReadFile(goPath)
	if err != nil {
		return manifest{}, err
	}
	tsData, err := os.ReadFile(tsPath)
	if err != nil {
		return manifest{}, err
	}
	entries, err := os.ReadDir(filepath.Join(root, "schemas", "v0.1"))
	if err != nil {
		return manifest{}, err
	}
	result := manifest{Version: 1, Schemas: map[string]schemaContract{}, Artifacts: map[string]string{
		"gen/go/lifecycle/v0_1/types.go": digest(goData),
		"gen/ts/v0.1/index.ts":           digest(tsData),
	}}
	seen := make([]string, 0, len(entries))
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || filepath.Ext(name) != ".json" || name == "defs.json" {
			continue
		}
		mapping, ok := contractTypes[name]
		if !ok {
			return manifest{}, fmt.Errorf("public schema %s has no generated type mapping", name)
		}
		if !strings.Contains(string(goData), "type "+mapping[0]+" ") {
			return manifest{}, fmt.Errorf("go type %s for %s is missing", mapping[0], name)
		}
		if !strings.Contains(string(tsData), "interface "+mapping[1]+" ") && !strings.Contains(string(tsData), "type "+mapping[1]+" ") {
			return manifest{}, fmt.Errorf("typescript type %s for %s is missing", mapping[1], name)
		}
		data, err := os.ReadFile(filepath.Join(root, "schemas", "v0.1", name))
		if err != nil {
			return manifest{}, err
		}
		result.Schemas[name] = schemaContract{GoType: mapping[0], TypeScript: mapping[1], SHA256: digest(data)}
		seen = append(seen, name)
	}
	sort.Strings(seen)
	if len(seen) != len(contractTypes) {
		return manifest{}, fmt.Errorf("type map covers %d schemas but repository contains %d", len(contractTypes), len(seen))
	}
	return result, nil
}

func digest(data []byte) string {
	data = bytes.ReplaceAll(data, []byte("\r\n"), []byte("\n"))
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "lifecycle-gen:", err)
	os.Exit(1)
}
