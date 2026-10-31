package lifecyclespec

import "testing"

func TestV2ProjectStatusAndOfflineBoundary(t *testing.T) {
	valid := []byte(`{"spec":"ai-devops.project-status/v2alpha1","projectId":"project","state":"development","links":{"development":"http://127.0.0.1:8766/"}}`)
	if err := ValidateV2("project-status", valid); err != nil {
		t.Fatal(err)
	}
	for _, bad := range []string{`{"spec":"v3"}`, `{"spec":"ai-devops.project-status/v2alpha1","projectId":"project","state":"invented","links":{}}`} {
		if ValidateV2("project-status", []byte(bad)) == nil {
			t.Fatal("malformed status accepted")
		}
	}
	if ValidateV2("../project-status", valid) == nil {
		t.Fatal("path traversal accepted")
	}
	if ValidateV2("https://example.invalid/schema", valid) == nil {
		t.Fatal("remote schema accepted")
	}
	if ValidateV2("ai-team-mission", []byte(`{"state":"completed"}`)) == nil {
		t.Fatal("unsupported mission accepted")
	}
}
