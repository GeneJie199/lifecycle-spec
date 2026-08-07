// Package v0_1 contains hand-synced Go types for lifecycle-spec v0.1.
//
//go:generate echo See scripts/generate-types.md — v0.1 types are checked in and hand-synced from schemas/v0.1
package v0_1

import "encoding/json"

// Shared enums and reusable fragments.

type ProductName string

const (
	ProductLifecycleSpec      ProductName = "lifecycle-spec"
	ProductDevCycle           ProductName = "dev-cycle"
	ProductInfraDiscovery     ProductName = "infra-discovery"
	ProductFleetObservability ProductName = "fleet-observability"
	ProductReleaseValidation  ProductName = "release-validation"
	ProductOther              ProductName = "other"
)

type Environment string

const (
	EnvDev        Environment = "dev"
	EnvTest       Environment = "test"
	EnvStaging    Environment = "staging"
	EnvProduction Environment = "production"
	EnvLocal      Environment = "local"
	EnvOther      Environment = "other"
)

type Classification string

const (
	ClassificationExpected   Classification = "expected"
	ClassificationApproved   Classification = "approved"
	ClassificationTemporary  Classification = "temporary"
	ClassificationUnexpected Classification = "unexpected"
	ClassificationDenied     Classification = "denied"
)

type Severity string

const (
	SeverityInfo     Severity = "info"
	SeverityLow      Severity = "low"
	SeverityMedium   Severity = "medium"
	SeverityHigh     Severity = "high"
	SeverityCritical Severity = "critical"
)

type DigestAlg string

const (
	DigestSHA256 DigestAlg = "sha256"
	DigestSHA512 DigestAlg = "sha512"
	DigestSHA1   DigestAlg = "sha1"
)

type RedactionStatus string

const (
	RedactionNone    RedactionStatus = "none"
	RedactionPartial RedactionStatus = "partial"
	RedactionFull    RedactionStatus = "full"
)

type EvidenceKind string

const (
	EvidenceTestResult         EvidenceKind = "test.result"
	EvidenceScreenshot         EvidenceKind = "screenshot"
	EvidenceLogExcerpt         EvidenceKind = "log.excerpt"
	EvidenceGitCommit          EvidenceKind = "git.commit"
	EvidenceGitDiff            EvidenceKind = "git.diff"
	EvidenceManualConfirmation EvidenceKind = "manual.confirmation"
	EvidenceHTTPResponse       EvidenceKind = "http.response"
	EvidenceArtifactHash       EvidenceKind = "artifact.hash"
	EvidenceCommandOutput      EvidenceKind = "command.output"
	EvidenceOther              EvidenceKind = "other"
)

type ChangeKind string

const (
	ChangeAdded               ChangeKind = "added"
	ChangeRemoved             ChangeKind = "removed"
	ChangeModified            ChangeKind = "modified"
	ChangeMoved               ChangeKind = "moved"
	ChangeExposureChanged     ChangeKind = "exposure.changed"
	ChangeDeployMethodChanged ChangeKind = "deploy.method.changed"
	ChangeRelationshipChanged ChangeKind = "relationship.changed"
	ChangeMetadataChanged     ChangeKind = "metadata.changed"
)

type ObservationStatus string

const (
	ObservationOK       ObservationStatus = "ok"
	ObservationDegraded ObservationStatus = "degraded"
	ObservationError    ObservationStatus = "error"
	ObservationUnknown  ObservationStatus = "unknown"
)

type RelationshipType string

const (
	RelContains  RelationshipType = "contains"
	RelRunsOn    RelationshipType = "runs_on"
	RelListensOn RelationshipType = "listens_on"
	RelDependsOn RelationshipType = "depends_on"
	RelRoutesTo  RelationshipType = "routes_to"
	RelBackedBy  RelationshipType = "backed_by"
	RelSameAs    RelationshipType = "same_as"
	RelOwnedBy   RelationshipType = "owned_by"
	RelOther     RelationshipType = "other"
)

type ApprovalDecision string

const (
	DecisionApproved  ApprovalDecision = "approved"
	DecisionTemporary ApprovalDecision = "temporary"
	DecisionDenied    ApprovalDecision = "denied"
)

type ReleaseStatus string

const (
	ReleaseCandidate  ReleaseStatus = "candidate"
	ReleaseApproved   ReleaseStatus = "approved"
	ReleaseInProgress ReleaseStatus = "in_progress"
	ReleaseCompleted  ReleaseStatus = "completed"
	ReleaseRolledBack ReleaseStatus = "rolled_back"
	ReleaseAborted    ReleaseStatus = "aborted"
)

// Shared objects.

type Source struct {
	Product    ProductName `json:"product"`
	InstanceID string      `json:"instanceId"`
	Version    string      `json:"version,omitempty"`
}

type Subject struct {
	ResourceID   string `json:"resourceId"`
	ResourceType string `json:"resourceType,omitempty"`
	Path         string `json:"path,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
}

type Context struct {
	ProjectID       string            `json:"projectId,omitempty"`
	Environment     Environment       `json:"environment,omitempty"`
	ResourceGroupID string            `json:"resourceGroupId,omitempty"`
	Labels          map[string]string `json:"labels,omitempty"`
}

type Digest struct {
	Alg   DigestAlg `json:"alg"`
	Value string    `json:"value"`
}

type Redaction struct {
	Status RedactionStatus `json:"status"`
	Notes  string          `json:"notes,omitempty"`
}

// EvidenceRef is an ID-only pointer to Evidence.
type EvidenceRef struct {
	EvidenceID string `json:"evidenceId"`
}

// Evidence is proof material attached to events or other objects.
type Evidence struct {
	SpecVersion string            `json:"specVersion,omitempty"`
	EvidenceID  string            `json:"evidenceId"`
	Kind        EvidenceKind      `json:"kind"`
	Title       string            `json:"title"`
	Summary     string            `json:"summary,omitempty"`
	ProducedAt  string            `json:"producedAt"`
	Producer    Source            `json:"producer"`
	Subject     *Subject          `json:"subject,omitempty"`
	ContentType string            `json:"contentType,omitempty"`
	Inline      any               `json:"inline,omitempty"`
	URI         string            `json:"uri,omitempty"`
	Digest      *Digest           `json:"digest,omitempty"`
	Redaction   Redaction         `json:"redaction"`
	RelatedIDs  []string          `json:"relatedIds,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
}

// EvidenceItem is either a full Evidence object or an ID-only reference.
type EvidenceItem struct {
	Evidence
}

// UnmarshalJSON accepts both EvidenceRef and full Evidence shapes.
func (e *EvidenceItem) UnmarshalJSON(data []byte) error {
	var full Evidence
	if err := json.Unmarshal(data, &full); err != nil {
		return err
	}
	e.Evidence = full
	return nil
}

// MarshalJSON emits a compact ID-only object when only EvidenceID is set.
func (e EvidenceItem) MarshalJSON() ([]byte, error) {
	if e.Kind == "" && e.Title == "" && e.ProducedAt == "" && e.EvidenceID != "" {
		return json.Marshal(EvidenceRef{EvidenceID: e.EvidenceID})
	}
	return json.Marshal(e.Evidence)
}

// EventEnvelope is the common wrapper for every lifecycle event.
type EventEnvelope struct {
	SpecVersion    string         `json:"specVersion"`
	EventID        string         `json:"eventId"`
	EventType      string         `json:"eventType"`
	OccurredAt     string         `json:"occurredAt"`
	RecordedAt     string         `json:"recordedAt,omitempty"`
	Source         Source         `json:"source"`
	Context        *Context       `json:"context,omitempty"`
	Subject        Subject        `json:"subject"`
	Classification Classification `json:"classification,omitempty"`
	Severity       Severity       `json:"severity,omitempty"`
	ReleaseID      string         `json:"releaseId,omitempty"`
	Evidence       []EvidenceItem `json:"evidence,omitempty"`
	Payload        map[string]any `json:"payload"`
	Extensions     map[string]any `json:"extensions,omitempty"`
}

// ChangePayload is the typed body of a ChangeEvent.
type ChangePayload struct {
	ChangeKind          ChangeKind     `json:"changeKind"`
	Summary             string         `json:"summary,omitempty"`
	Paths               []string       `json:"paths,omitempty"`
	Before              map[string]any `json:"before"`
	After               map[string]any `json:"after"`
	NoiseFiltered       *bool          `json:"noiseFiltered,omitempty"`
	BaselineSnapshotID  string         `json:"baselineSnapshotId,omitempty"`
	CandidateSnapshotID string         `json:"candidateSnapshotId,omitempty"`
	CorrectsEventID     string         `json:"correctsEventId,omitempty"`
	ApprovalID          string         `json:"approvalId,omitempty"`
}

// ChangeEvent records what changed versus a prior state.
type ChangeEvent struct {
	SpecVersion    string         `json:"specVersion"`
	EventID        string         `json:"eventId"`
	EventType      string         `json:"eventType"`
	OccurredAt     string         `json:"occurredAt"`
	RecordedAt     string         `json:"recordedAt,omitempty"`
	Source         Source         `json:"source"`
	Context        *Context       `json:"context,omitempty"`
	Subject        Subject        `json:"subject"`
	Classification Classification `json:"classification"`
	Severity       Severity       `json:"severity,omitempty"`
	ReleaseID      string         `json:"releaseId,omitempty"`
	Evidence       []EvidenceItem `json:"evidence,omitempty"`
	Payload        ChangePayload  `json:"payload"`
	Extensions     map[string]any `json:"extensions,omitempty"`
}

// Resource is anything managed with a stable ID.
type Resource struct {
	SpecVersion      string            `json:"specVersion,omitempty"`
	ResourceID       string            `json:"resourceId"`
	ResourceType     string            `json:"resourceType"`
	DisplayName      string            `json:"displayName"`
	Attributes       map[string]any    `json:"attributes,omitempty"`
	Labels           map[string]string `json:"labels,omitempty"`
	FirstSeenAt      string            `json:"firstSeenAt,omitempty"`
	LastSeenAt       string            `json:"lastSeenAt,omitempty"`
	ParentResourceID string            `json:"parentResourceId,omitempty"`
}

// Observation is one look at one or more resources.
type Observation struct {
	SpecVersion   string            `json:"specVersion,omitempty"`
	ObservationID string            `json:"observationId"`
	ObservedAt    string            `json:"observedAt"`
	Source        Source            `json:"source"`
	Subjects      []Subject         `json:"subjects"`
	Status        ObservationStatus `json:"status"`
	Summary       string            `json:"summary,omitempty"`
	Metrics       map[string]any    `json:"metrics,omitempty"`
	Evidence      []EvidenceItem    `json:"evidence,omitempty"`
	Raw           map[string]any    `json:"raw,omitempty"`
}

// NoisePolicy describes fields excluded from snapshot diffs.
type NoisePolicy struct {
	FilteredFields []string `json:"filteredFields,omitempty"`
	Notes          string   `json:"notes,omitempty"`
}

// Snapshot is a normalized photo of state used for diffs.
type Snapshot struct {
	SpecVersion   string         `json:"specVersion,omitempty"`
	SnapshotID    string         `json:"snapshotId"`
	CapturedAt    string         `json:"capturedAt"`
	Source        Source         `json:"source"`
	Context       *Context       `json:"context,omitempty"`
	NoisePolicy   *NoisePolicy   `json:"noisePolicy,omitempty"`
	Resources     []Resource     `json:"resources"`
	Relationships []Relationship `json:"relationships,omitempty"`
	Digest        *Digest        `json:"digest,omitempty"`
}

// Relationship describes how two resources connect.
type Relationship struct {
	SpecVersion    string           `json:"specVersion,omitempty"`
	RelationshipID string           `json:"relationshipId"`
	Type           RelationshipType `json:"type"`
	From           string           `json:"from"`
	To             string           `json:"to"`
	Attributes     map[string]any   `json:"attributes,omitempty"`
	ObservedAt     string           `json:"observedAt,omitempty"`
}

// Actor is a human decision-maker on an Approval.
type Actor struct {
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

// ApprovalScope bounds what an approval covers.
type ApprovalScope struct {
	EventIDs    []string `json:"eventIds,omitempty"`
	ResourceIDs []string `json:"resourceIds,omitempty"`
	ValidFrom   string   `json:"validFrom,omitempty"`
	ValidUntil  string   `json:"validUntil,omitempty"`
}

// Approval records a human decision about a change or exception.
type Approval struct {
	SpecVersion string           `json:"specVersion,omitempty"`
	ApprovalID  string           `json:"approvalId"`
	Decision    ApprovalDecision `json:"decision"`
	DecidedAt   string           `json:"decidedAt"`
	Actor       Actor            `json:"actor"`
	Scope       *ApprovalScope   `json:"scope,omitempty"`
	Reason      string           `json:"reason,omitempty"`
	Evidence    []EvidenceItem   `json:"evidence,omitempty"`
}

// ExpectedChange declares an anticipated change for release matching.
type ExpectedChange struct {
	ChangeKind string `json:"changeKind"`
	ResourceID string `json:"resourceId"`
	Path       string `json:"path,omitempty"`
	Summary    string `json:"summary,omitempty"`
}

// Release is a planned or executed go-live package.
type Release struct {
	SpecVersion     string           `json:"specVersion,omitempty"`
	ReleaseID       string           `json:"releaseId"`
	Version         string           `json:"version"`
	Status          ReleaseStatus    `json:"status"`
	CreatedAt       string           `json:"createdAt"`
	StartedAt       string           `json:"startedAt,omitempty"`
	CompletedAt     string           `json:"completedAt,omitempty"`
	Source          Source           `json:"source"`
	Context         *Context         `json:"context,omitempty"`
	Subject         *Subject         `json:"subject,omitempty"`
	ExpectedChanges []ExpectedChange `json:"expectedChanges,omitempty"`
	Evidence        []EvidenceItem   `json:"evidence,omitempty"`
	Notes           string           `json:"notes,omitempty"`
}
