package v0_1

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestIntegrationTypeJSONFieldNames(t *testing.T) {
	candidate := DevelopmentReleaseCandidate{Spec: "lifecycle-spec/release-candidate/v1", Kind: "release-candidate", GeneratedAt: "2026-08-12T01:00:00Z", Requirement: DevelopmentRequirement{ID: "req-1", Title: "Ship", Description: "", CreatedAt: "2026-08-12T00:00:00Z", UpdatedAt: "2026-08-12T01:00:00Z"}, AcceptanceCriteria: []DevelopmentCriterion{}, Tasks: []DevelopmentTask{}, Evidence: []DevelopmentEvidence{}, Readiness: ReleaseReadiness{CriteriaTotal: 1, CriteriaSatisfied: 1, CriteriaWithEvidence: 1, Ready: true}}
	b, err := json.Marshal(candidate)
	if err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{`"generatedAt"`, `"acceptanceCriteria"`, `"criteriaWithEvidence"`} {
		if !strings.Contains(string(b), key) {
			t.Fatalf("missing %s in %s", key, b)
		}
	}
	report := ReleaseValidationReport{SchemaVersion: "releaseguard.report/v1", ReleaseID: "rel-1", Version: "1", Decision: "GO", GeneratedAt: "2026-08-12T01:00:00Z", PlanSHA256: strings.Repeat("a", 64), Manifest: ReleaseValidationManifest{ReleaseID: "rel-1", Version: "1", CreatedAt: "2026-08-12T01:00:00Z"}, Results: []ReleaseCheckResult{}, Rollback: []string{"restore"}}
	b, err = json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{`"schema_version"`, `"release_id"`, `"plan_sha256"`} {
		if !strings.Contains(string(b), key) {
			t.Fatalf("missing %s in %s", key, b)
		}
	}
}
