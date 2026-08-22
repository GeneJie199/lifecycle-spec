// Package v0_1 contains checked-in Go wire types for lifecycle-spec v0.1.
//
//go:generate go run ../../../../cmd/lifecycle-gen --check
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

type RelationshipConfirmation string

const (
	RelationshipUnreviewed RelationshipConfirmation = "unreviewed"
	RelationshipConfirmed  RelationshipConfirmation = "confirmed"
	RelationshipRejected   RelationshipConfirmation = "rejected"
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

// Product integration contracts used by the four independently deployable tools.
type FleetNodeReport struct {
	NodeID     string            `json:"node_id"`
	ObservedAt string            `json:"observed_at"`
	Agent      map[string]string `json:"agent,omitempty"`
	Inventory  map[string]any    `json:"inventory,omitempty"`
	Drift      map[string]any    `json:"drift,omitempty"`
	Metrics    map[string]any    `json:"metrics,omitempty"`
	Labels     map[string]string `json:"labels,omitempty"`
}

type TelemetryMetricKind string

const (
	TelemetryGauge   TelemetryMetricKind = "gauge"
	TelemetryCounter TelemetryMetricKind = "counter"
)

type TelemetryPoint struct {
	Metric      string              `json:"metric"`
	Labels      map[string]string   `json:"labels,omitempty"`
	TimestampMS int64               `json:"timestamp_ms"`
	Value       float64             `json:"value"`
	Kind        TelemetryMetricKind `json:"kind,omitempty"`
	Unit        string              `json:"unit,omitempty"`
}

type TelemetryBatch struct {
	Schema   string           `json:"schema"`
	NodeID   string           `json:"node_id"`
	Source   string           `json:"source"`
	Sequence uint64           `json:"sequence"`
	SentAt   string           `json:"sent_at"`
	Points   []TelemetryPoint `json:"points"`
}

type OperationalEvent struct {
	ID          string            `json:"id"`
	TimestampMS int64             `json:"timestamp_ms"`
	Kind        string            `json:"kind"`
	Severity    string            `json:"severity"`
	Service     string            `json:"service,omitempty"`
	Message     string            `json:"message"`
	Attributes  map[string]string `json:"attributes,omitempty"`
}

type EventBatch struct {
	Schema   string             `json:"schema"`
	NodeID   string             `json:"node_id"`
	Source   string             `json:"source"`
	Sequence uint64             `json:"sequence"`
	SentAt   string             `json:"sent_at"`
	Events   []OperationalEvent `json:"events"`
}

type MonitoringRecommendation struct {
	ID         string            `json:"id"`
	TargetID   string            `json:"target_id"`
	Collector  string            `json:"collector"`
	Priority   string            `json:"priority"`
	Reason     string            `json:"reason"`
	Parameters map[string]string `json:"parameters,omitempty"`
}

type MonitoringPlan struct {
	Version         string                     `json:"version"`
	GeneratedAt     string                     `json:"generated_at"`
	Hostname        string                     `json:"hostname"`
	Recommendations []MonitoringRecommendation `json:"recommendations"`
	CoverageGaps    []string                   `json:"coverage_gaps,omitempty"`
}

type ReleaseCheckResult struct {
	Name       string         `json:"name"`
	Type       string         `json:"type"`
	Phase      string         `json:"phase"`
	Status     string         `json:"status"`
	Required   bool           `json:"required"`
	Summary    string         `json:"summary"`
	DurationMS int64          `json:"duration_ms"`
	Evidence   map[string]any `json:"evidence,omitempty"`
}

type DevelopmentRequirement struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type DevelopmentCriterion struct {
	ID            string `json:"id"`
	RequirementID string `json:"requirementId"`
	Description   string `json:"description"`
	Satisfied     bool   `json:"satisfied"`
	CreatedAt     string `json:"createdAt"`
}

type DevelopmentTask struct {
	ID            string   `json:"id"`
	RequirementID string   `json:"requirementId"`
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Status        string   `json:"status"`
	Branch        string   `json:"branch"`
	WorktreePath  string   `json:"worktreePath"`
	DependsOn     []string `json:"dependsOn,omitempty"`
	CreatedAt     string   `json:"createdAt"`
	UpdatedAt     string   `json:"updatedAt"`
}

type DevelopmentEvidence struct {
	ID            string            `json:"id"`
	RequirementID string            `json:"requirementId"`
	CriterionID   string            `json:"criterionId,omitempty"`
	TaskID        string            `json:"taskId,omitempty"`
	Kind          string            `json:"kind"`
	Title         string            `json:"title"`
	Status        string            `json:"status"`
	URI           string            `json:"uri,omitempty"`
	Inline        string            `json:"inline,omitempty"`
	Metadata      map[string]string `json:"metadata,omitempty"`
	CreatedAt     string            `json:"createdAt"`
}

type ReleaseReadiness struct {
	CriteriaTotal        int  `json:"criteriaTotal"`
	CriteriaSatisfied    int  `json:"criteriaSatisfied"`
	TasksTotal           int  `json:"tasksTotal"`
	TasksDone            int  `json:"tasksDone"`
	CriteriaWithEvidence int  `json:"criteriaWithEvidence"`
	SourcesTotal         int  `json:"sourcesTotal"`
	SourcesClean         int  `json:"sourcesClean"`
	Ready                bool `json:"ready"`
}

type DevelopmentDirtyFile struct {
	XY   string `json:"xy"`
	Path string `json:"path"`
}

type DevelopmentSourceSnapshot struct {
	TaskID         string                 `json:"taskId"`
	TaskTitle      string                 `json:"taskTitle"`
	RepositoryPath string                 `json:"repositoryPath"`
	WorktreePath   string                 `json:"worktreePath"`
	Branch         string                 `json:"branch"`
	HeadCommit     string                 `json:"headCommit"`
	Clean          bool                   `json:"clean"`
	DirtyFiles     []DevelopmentDirtyFile `json:"dirtyFiles"`
	CapturedAt     string                 `json:"capturedAt"`
}

type DevelopmentReleaseCandidate struct {
	Spec               string                      `json:"spec"`
	Kind               string                      `json:"kind"`
	GeneratedAt        string                      `json:"generatedAt"`
	Requirement        DevelopmentRequirement      `json:"requirement"`
	AcceptanceCriteria []DevelopmentCriterion      `json:"acceptanceCriteria"`
	Tasks              []DevelopmentTask           `json:"tasks"`
	Evidence           []DevelopmentEvidence       `json:"evidence"`
	Sources            []DevelopmentSourceSnapshot `json:"sources"`
	Readiness          ReleaseReadiness            `json:"readiness"`
}

// Standalone workflow contracts. The Development* types above remain the exact
// release-candidate wire shapes used by DevCycle v1 exports.
type ProjectRepository struct {
	URL           string `json:"url,omitempty"`
	Path          string `json:"path,omitempty"`
	DefaultBranch string `json:"defaultBranch,omitempty"`
}

type Project struct {
	SpecVersion string             `json:"specVersion"`
	ProjectID   string             `json:"projectId"`
	Name        string             `json:"name"`
	Description string             `json:"description,omitempty"`
	Status      string             `json:"status"`
	Repository  *ProjectRepository `json:"repository,omitempty"`
	CreatedAt   string             `json:"createdAt"`
	UpdatedAt   string             `json:"updatedAt"`
	Labels      map[string]string  `json:"labels,omitempty"`
}

type Requirement struct {
	SpecVersion string            `json:"specVersion"`
	ID          string            `json:"id"`
	ProjectID   string            `json:"projectId,omitempty"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      string            `json:"status"`
	Priority    string            `json:"priority,omitempty"`
	Source      *Source           `json:"source,omitempty"`
	CreatedAt   string            `json:"createdAt"`
	UpdatedAt   string            `json:"updatedAt"`
	Labels      map[string]string `json:"labels,omitempty"`
}

type AcceptanceCriterion struct {
	SpecVersion   string   `json:"specVersion"`
	ID            string   `json:"id"`
	RequirementID string   `json:"requirementId"`
	Description   string   `json:"description"`
	Satisfied     bool     `json:"satisfied"`
	EvidenceIDs   []string `json:"evidenceIds,omitempty"`
	CreatedAt     string   `json:"createdAt"`
	UpdatedAt     string   `json:"updatedAt,omitempty"`
}

type Task struct {
	SpecVersion   string   `json:"specVersion"`
	ID            string   `json:"id"`
	RequirementID string   `json:"requirementId"`
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Status        string   `json:"status"`
	Assignee      string   `json:"assignee,omitempty"`
	Branch        string   `json:"branch,omitempty"`
	WorktreePath  string   `json:"worktreePath,omitempty"`
	DependsOn     []string `json:"dependsOn,omitempty"`
	CreatedAt     string   `json:"createdAt"`
	UpdatedAt     string   `json:"updatedAt"`
}

type AgentUsage struct {
	InputTokens   int     `json:"inputTokens,omitempty"`
	OutputTokens  int     `json:"outputTokens,omitempty"`
	EstimatedCost float64 `json:"estimatedCost"`
	Currency      string  `json:"currency,omitempty"`
}

type AgentSession struct {
	SpecVersion  string      `json:"specVersion"`
	ID           string      `json:"id"`
	TaskID       string      `json:"taskId"`
	Provider     string      `json:"provider"`
	Model        string      `json:"model,omitempty"`
	Status       string      `json:"status"`
	WorkingDir   string      `json:"workingDir"`
	PromptDigest *Digest     `json:"promptDigest,omitempty"`
	PID          int         `json:"pid,omitempty"`
	LogURI       string      `json:"logUri,omitempty"`
	StartedAt    string      `json:"startedAt"`
	EndedAt      string      `json:"endedAt,omitempty"`
	Usage        *AgentUsage `json:"usage,omitempty"`
}

type CodeChangeFile struct {
	Path         string `json:"path"`
	PreviousPath string `json:"previousPath,omitempty"`
	Status       string `json:"status"`
	Additions    int    `json:"additions"`
	Deletions    int    `json:"deletions"`
	Binary       bool   `json:"binary"`
}

type CodeChange struct {
	SpecVersion  string           `json:"specVersion"`
	ChangeID     string           `json:"changeId"`
	RepositoryID string           `json:"repositoryId"`
	TaskID       string           `json:"taskId,omitempty"`
	FromRevision string           `json:"fromRevision"`
	ToRevision   string           `json:"toRevision"`
	Summary      string           `json:"summary,omitempty"`
	Risk         Severity         `json:"risk,omitempty"`
	Files        []CodeChangeFile `json:"files"`
	EvidenceIDs  []string         `json:"evidenceIds,omitempty"`
	CreatedAt    string           `json:"createdAt"`
}

type Incident struct {
	SpecVersion string            `json:"specVersion"`
	IncidentID  string            `json:"incidentId"`
	Title       string            `json:"title"`
	Summary     string            `json:"summary,omitempty"`
	Status      string            `json:"status"`
	Severity    Severity          `json:"severity"`
	Owner       string            `json:"owner,omitempty"`
	NodeIDs     []string          `json:"nodeIds,omitempty"`
	ServiceIDs  []string          `json:"serviceIds,omitempty"`
	AlertIDs    []string          `json:"alertIds,omitempty"`
	StartedAt   string            `json:"startedAt"`
	DetectedAt  string            `json:"detectedAt,omitempty"`
	ResolvedAt  string            `json:"resolvedAt,omitempty"`
	UpdatedAt   string            `json:"updatedAt"`
	Labels      map[string]string `json:"labels,omitempty"`
}

// ExpectedChanges declares the reviewed operational changes for one release.
type ExpectedChanges struct {
	Spec        string                      `json:"spec"`
	Kind        string                      `json:"kind"`
	ReleaseID   string                      `json:"release_id"`
	Version     string                      `json:"version"`
	GeneratedAt string                      `json:"generated_at"`
	Changes     []ExpectedChangeDeclaration `json:"changes"`
	Metadata    map[string]string           `json:"metadata,omitempty"`
}

type ExpectedChangeDeclaration struct {
	ID                 string   `json:"id"`
	Source             string   `json:"source"`
	Action             string   `json:"action"`
	ResourceID         string   `json:"resource_id"`
	ResourceType       string   `json:"resource_type,omitempty"`
	NodeID             string   `json:"node_id,omitempty"`
	Fields             []string `json:"fields,omitempty"`
	Fingerprint        string   `json:"fingerprint,omitempty"`
	Summary            string   `json:"summary"`
	EvidenceIDs        []string `json:"evidence_ids,omitempty"`
	VerificationChecks []string `json:"verification_checks,omitempty"`
	MetricPolicies     []string `json:"metric_policies,omitempty"`
	AffectedNodes      []string `json:"affected_nodes,omitempty"`
	Required           *bool    `json:"required,omitempty"`
}

type ReleaseObservedChange struct {
	ID             string   `json:"id"`
	Source         string   `json:"source"`
	Action         string   `json:"action"`
	ResourceID     string   `json:"resource_id"`
	ResourceType   string   `json:"resource_type,omitempty"`
	NodeID         string   `json:"node_id,omitempty"`
	Fields         []string `json:"fields,omitempty"`
	Fingerprint    string   `json:"fingerprint,omitempty"`
	Severity       string   `json:"severity,omitempty"`
	Summary        string   `json:"summary,omitempty"`
	Classification string   `json:"classification,omitempty"`
	ReleaseID      string   `json:"release_id,omitempty"`
}

type ReleaseChangeSourceEvidence struct {
	Source         string `json:"source"`
	ArtifactSHA256 string `json:"artifact_sha256,omitempty"`
	CheckedAt      string `json:"checked_at"`
	Items          int    `json:"items"`
}

type ReleaseChangeCorrelation struct {
	ExpectedID string   `json:"expected_id"`
	Status     string   `json:"status"`
	Required   bool     `json:"required"`
	ObservedID string   `json:"observed_id,omitempty"`
	Reasons    []string `json:"reasons,omitempty"`
}

type ReleaseChangeCoverage struct {
	Spec            string                        `json:"spec"`
	DocumentSHA256  string                        `json:"document_sha256"`
	Declared        []ExpectedChangeDeclaration   `json:"declared"`
	Observed        []ReleaseObservedChange       `json:"observed"`
	Sources         []ReleaseChangeSourceEvidence `json:"sources"`
	Correlations    []ReleaseChangeCorrelation    `json:"correlations"`
	Unexpected      []ReleaseObservedChange       `json:"unexpected"`
	ExpectedTotal   int                           `json:"expected_total"`
	MatchedTotal    int                           `json:"matched_total"`
	MissingRequired int                           `json:"missing_required"`
	MissingOptional int                           `json:"missing_optional"`
	UnexpectedTotal int                           `json:"unexpected_total"`
}

type ReleaseGuidance struct {
	Code       string   `json:"code"`
	Priority   string   `json:"priority"`
	Title      string   `json:"title"`
	Summary    string   `json:"summary"`
	RelatedIDs []string `json:"related_ids,omitempty"`
}

type ReleaseValidationManifest struct {
	ReleaseID   string                  `json:"release_id"`
	Version     string                  `json:"version"`
	CreatedAt   string                  `json:"created_at"`
	Metadata    map[string]any          `json:"metadata,omitempty"`
	Candidate   map[string]any          `json:"candidate,omitempty"`
	Git         map[string]any          `json:"git,omitempty"`
	FleetBefore map[string]any          `json:"fleet_before,omitempty"`
	FleetAfter  map[string]any          `json:"fleet_after,omitempty"`
	Metrics     []ReleaseMetricEvidence `json:"metrics,omitempty"`
	Changes     *ReleaseChangeCoverage  `json:"changes,omitempty"`
}

type ReleaseMetricWindow struct {
	StartMS int64   `json:"start_ms"`
	EndMS   int64   `json:"end_ms"`
	Samples int     `json:"samples"`
	Series  int     `json:"series"`
	Value   float64 `json:"value"`
	Minimum float64 `json:"minimum"`
	Maximum float64 `json:"maximum"`
	Last    float64 `json:"last"`
}

type ReleaseMetricEvidence struct {
	Name              string               `json:"name"`
	Metric            string               `json:"metric"`
	Node              string               `json:"node,omitempty"`
	Aggregate         string               `json:"aggregate"`
	SeriesReduce      string               `json:"series_reduce"`
	Direction         string               `json:"direction"`
	Before            *ReleaseMetricWindow `json:"before,omitempty"`
	After             *ReleaseMetricWindow `json:"after,omitempty"`
	RegressionPercent float64              `json:"regression_percent,omitempty"`
	Pass              bool                 `json:"pass"`
	Summary           string               `json:"summary"`
}

type ReleaseFleetNodeEvidence struct {
	NodeID          string `json:"node_id"`
	Health          string `json:"health"`
	ObservedAt      string `json:"observed_at"`
	ActualVersion   string `json:"actual_version"`
	ExpectedVersion string `json:"expected_version"`
	Match           bool   `json:"match"`
}

type ReleaseFleetEvidence struct {
	CheckedAt      string                     `json:"checked_at"`
	Nodes          []ReleaseFleetNodeEvidence `json:"nodes"`
	CriticalAlerts int                        `json:"critical_alerts"`
}

type ReleaseObservationState struct {
	Status     string                 `json:"status"`
	StartedAt  string                 `json:"started_at"`
	DeadlineAt string                 `json:"deadline_at"`
	Samples    []ReleaseFleetEvidence `json:"samples"`
}

type ReleaseValidationReport struct {
	SchemaVersion  string                    `json:"schema_version"`
	ReleaseID      string                    `json:"release_id"`
	Version        string                    `json:"version"`
	Decision       string                    `json:"decision"`
	DecisionReason string                    `json:"decision_reason"`
	GeneratedAt    string                    `json:"generated_at"`
	PlanSHA256     string                    `json:"plan_sha256"`
	Manifest       ReleaseValidationManifest `json:"manifest"`
	Results        []ReleaseCheckResult      `json:"results"`
	Rollback       []string                  `json:"rollback"`
	Observation    *ReleaseObservationState  `json:"observation,omitempty"`
	Guidance       []ReleaseGuidance         `json:"guidance,omitempty"`
}

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
	SpecVersion      string                   `json:"specVersion,omitempty"`
	RelationshipID   string                   `json:"relationshipId"`
	Type             RelationshipType         `json:"type"`
	From             string                   `json:"from"`
	To               string                   `json:"to"`
	Attributes       map[string]any           `json:"attributes,omitempty"`
	ObservedAt       string                   `json:"observedAt,omitempty"`
	Confidence       float64                  `json:"confidence,omitempty"`
	DiscoverySources []string                 `json:"discoverySources,omitempty"`
	FirstSeenAt      string                   `json:"firstSeenAt,omitempty"`
	LastSeenAt       string                   `json:"lastSeenAt,omitempty"`
	Confirmation     RelationshipConfirmation `json:"confirmation,omitempty"`
	ReviewedBy       string                   `json:"reviewedBy,omitempty"`
	ReviewedAt       string                   `json:"reviewedAt,omitempty"`
	ReviewNote       string                   `json:"reviewNote,omitempty"`
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
