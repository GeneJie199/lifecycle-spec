/**
 * Hand-synced TypeScript types for lifecycle-spec v0.1.
 * Source of truth: schemas/v0.1/*.json
 * See scripts/generate-types.md
 */

export type ProductName =
  | "lifecycle-spec"
  | "dev-cycle"
  | "infra-discovery"
  | "fleet-observability"
  | "release-validation"
  | "other";

export type Environment =
  | "dev"
  | "test"
  | "staging"
  | "production"
  | "local"
  | "other";

export type Classification =
  | "expected"
  | "approved"
  | "temporary"
  | "unexpected"
  | "denied";

export type Severity = "info" | "low" | "medium" | "high" | "critical";

export type DigestAlg = "sha256" | "sha512" | "sha1";

export type RedactionStatus = "none" | "partial" | "full";

export type EvidenceKind =
  | "test.result"
  | "screenshot"
  | "log.excerpt"
  | "git.commit"
  | "git.diff"
  | "manual.confirmation"
  | "http.response"
  | "artifact.hash"
  | "command.output"
  | "other";

export type ChangeKind =
  | "added"
  | "removed"
  | "modified"
  | "moved"
  | "exposure.changed"
  | "deploy.method.changed"
  | "relationship.changed"
  | "metadata.changed";

export type ObservationStatus = "ok" | "degraded" | "error" | "unknown";

export type RelationshipType =
  | "contains"
  | "runs_on"
  | "listens_on"
  | "depends_on"
  | "routes_to"
  | "backed_by"
  | "same_as"
  | "owned_by"
  | "other";

export type ApprovalDecision = "approved" | "temporary" | "denied";

export type ReleaseStatus =
  | "candidate"
  | "approved"
  | "in_progress"
  | "completed"
  | "rolled_back"
  | "aborted";

export interface FleetNodeReport {
  node_id: string;
  observed_at: string;
  agent?: Record<string, string>;
  inventory?: Record<string, unknown>;
  drift?: Record<string, unknown>;
  metrics?: Record<string, unknown>;
  labels?: Record<string, string>;
}

export type TelemetryMetricKind = "gauge" | "counter";

export interface TelemetryPoint {
  metric: string;
  labels?: Record<string, string>;
  timestamp_ms: number;
  value: number;
  kind?: TelemetryMetricKind;
  unit?: string;
}

export interface TelemetryBatch {
  schema: "telemetry.batch/v1";
  node_id: string;
  source: string;
  sequence: number;
  sent_at: string;
  points: TelemetryPoint[];
}

export interface OperationalEvent {
  id: string;
  timestamp_ms: number;
  kind: string;
  severity: "debug" | "info" | "warning" | "error" | "critical";
  service?: string;
  message: string;
  attributes?: Record<string, string>;
}

export interface EventBatch {
  schema: "event.batch/v1";
  node_id: string;
  source: string;
  sequence: number;
  sent_at: string;
  events: OperationalEvent[];
}

export interface MonitoringRecommendation {
  id: string;
  target_id: string;
  collector: string;
  priority: "required" | "recommended" | "optional";
  reason: string;
  parameters?: Record<string, string>;
}

export interface MonitoringPlan {
  version: "infrascout.monitoring/v1";
  generated_at: string;
  hostname: string;
  recommendations: MonitoringRecommendation[];
  coverage_gaps?: string[];
}

export interface DevelopmentRequirement {
  id: string;
  title: string;
  description: string;
  createdAt: string;
  updatedAt: string;
}

export interface DevelopmentCriterion {
  id: string;
  requirementId: string;
  description: string;
  satisfied: boolean;
  createdAt: string;
}

export interface DevelopmentTask {
  id: string;
  requirementId: string;
  title: string;
  description: string;
  status: "todo" | "in_progress" | "done";
  branch: string;
  worktreePath: string;
  dependsOn?: string[];
  createdAt: string;
  updatedAt: string;
}

export interface DevelopmentEvidence {
  id: string;
  requirementId: string;
  criterionId?: string;
  taskId?: string;
  kind: string;
  title: string;
  status: "passed" | "failed" | "informational";
  uri?: string;
  inline?: string;
  metadata?: Record<string, string>;
  createdAt: string;
}

export interface ReleaseReadiness {
  criteriaTotal: number;
  criteriaSatisfied: number;
  tasksTotal: number;
  tasksDone: number;
  criteriaWithEvidence: number;
  sourcesTotal: number;
  sourcesClean: number;
  ready: boolean;
}

export interface DevelopmentSourceSnapshot {
  taskId: string;
  taskTitle: string;
  repositoryPath: string;
  worktreePath: string;
  branch: string;
  headCommit: string;
  clean: boolean;
  dirtyFiles: Array<{ xy: string; path: string }>;
  capturedAt: string;
}

export interface DevelopmentReleaseCandidate {
  spec: "lifecycle-spec/release-candidate/v1";
  kind: "release-candidate";
  generatedAt: string;
  requirement: DevelopmentRequirement;
  acceptanceCriteria: DevelopmentCriterion[];
  tasks: DevelopmentTask[];
  evidence: DevelopmentEvidence[];
  sources: DevelopmentSourceSnapshot[];
  readiness: ReleaseReadiness;
}

export interface ProjectRepository {
  url?: string;
  path?: string;
  defaultBranch?: string;
}

export interface Project {
  specVersion: string;
  projectId: string;
  name: string;
  description?: string;
  status: "active" | "paused" | "archived";
  repository?: ProjectRepository;
  createdAt: string;
  updatedAt: string;
  labels?: Record<string, string>;
}

export interface Requirement {
  specVersion: string;
  id: string;
  projectId?: string;
  title: string;
  description: string;
  status: "draft" | "ready" | "in_progress" | "accepted" | "rejected";
  priority?: "low" | "medium" | "high" | "critical";
  source?: Source;
  createdAt: string;
  updatedAt: string;
  labels?: Record<string, string>;
}

export interface AcceptanceCriterion {
  specVersion: string;
  id: string;
  requirementId: string;
  description: string;
  satisfied: boolean;
  evidenceIds?: string[];
  createdAt: string;
  updatedAt?: string;
}

export interface Task {
  specVersion: string;
  id: string;
  requirementId: string;
  title: string;
  description: string;
  status: "todo" | "in_progress" | "blocked" | "done" | "canceled";
  assignee?: string;
  branch?: string;
  worktreePath?: string;
  dependsOn?: string[];
  createdAt: string;
  updatedAt: string;
}

export interface AgentUsage {
  inputTokens?: number;
  outputTokens?: number;
  estimatedCost?: number;
  currency?: string;
}

export interface AgentSession {
  specVersion: string;
  id: string;
  taskId: string;
  provider: string;
  model?: string;
  status: "queued" | "running" | "completed" | "failed" | "stopped" | "interrupted";
  workingDir: string;
  promptDigest?: Digest;
  pid?: number;
  logUri?: string;
  startedAt: string;
  endedAt?: string;
  usage?: AgentUsage;
}

export interface CodeChangeFile {
  path: string;
  previousPath?: string;
  status: "added" | "modified" | "deleted" | "renamed" | "copied" | "unmerged";
  additions: number;
  deletions: number;
  binary?: boolean;
}

export interface CodeChange {
  specVersion: string;
  changeId: string;
  repositoryId: string;
  taskId?: string;
  fromRevision: string;
  toRevision: string;
  summary?: string;
  risk?: Severity;
  files: CodeChangeFile[];
  evidenceIds?: string[];
  createdAt: string;
}

export interface Incident {
  specVersion: string;
  incidentId: string;
  title: string;
  summary?: string;
  status: "investigating" | "identified" | "monitoring" | "resolved" | "closed";
  severity: Severity;
  owner?: string;
  nodeIds?: string[];
  serviceIds?: string[];
  alertIds?: string[];
  startedAt: string;
  detectedAt?: string;
  resolvedAt?: string;
  updatedAt: string;
  labels?: Record<string, string>;
}

export interface ReleaseCheckResult {
  name: string;
  type: string;
  phase: "plan" | "recovery" | "delivery" | "verification" | "infrastructure" | "observation" | "internal";
  status: "pass" | "fail";
  required: boolean;
  summary: string;
  duration_ms: number;
  evidence?: Record<string, unknown>;
}

export interface ReleaseMetricWindow {
  start_ms: number;
  end_ms: number;
  samples: number;
  series: number;
  value: number;
  minimum: number;
  maximum: number;
  last: number;
}

export interface ReleaseMetricEvidence {
  name: string;
  metric: string;
  node?: string;
  aggregate: string;
  series_reduce: "avg" | "min" | "max" | "last" | "sum";
  direction: string;
  before?: ReleaseMetricWindow;
  after?: ReleaseMetricWindow;
  regression_percent?: number;
  pass: boolean;
  summary: string;
}

export interface ReleaseFleetNodeEvidence {
  node_id: string;
  health: string;
  observed_at: string;
  actual_version: string;
  expected_version: string;
  match: boolean;
}

export interface ReleaseFleetEvidence {
  checked_at: string;
  nodes: ReleaseFleetNodeEvidence[];
  critical_alerts: number;
}

export interface ReleaseObservationState {
  status: "observing" | "completed";
  started_at: string;
  deadline_at: string;
  samples: ReleaseFleetEvidence[];
}

export interface ReleaseValidationReport {
  schema_version: "releaseguard.report/v1";
  release_id: string;
  version: string;
  decision: "GO" | "HOLD" | "NO-GO";
  decision_reason: string;
  generated_at: string;
  plan_sha256: string;
  manifest: Record<string, unknown> & { release_id: string; version: string; created_at: string };
  results: ReleaseCheckResult[];
  rollback: string[];
  observation?: ReleaseObservationState;
}

export interface Source {
  product: ProductName;
  instanceId: string;
  version?: string;
}

export interface Subject {
  resourceId: string;
  resourceType?: string;
  path?: string;
  displayName?: string;
}

export interface Context {
  projectId?: string;
  environment?: Environment;
  resourceGroupId?: string;
  labels?: Record<string, string>;
}

export interface Digest {
  alg: DigestAlg;
  value: string;
}

export interface Redaction {
  status: RedactionStatus;
  notes?: string;
}

export interface EvidenceRef {
  evidenceId: string;
}

export interface Evidence {
  specVersion?: string;
  evidenceId: string;
  kind: EvidenceKind;
  title: string;
  summary?: string;
  producedAt: string;
  producer: Source;
  subject?: Subject;
  contentType?: string;
  inline?: string | Record<string, unknown> | unknown[];
  uri?: string;
  digest?: Digest;
  redaction: Redaction;
  relatedIds?: string[];
  labels?: Record<string, string>;
}

export type EvidenceItem = Evidence | EvidenceRef;

export interface EventEnvelope {
  specVersion: string;
  eventId: string;
  eventType: string;
  occurredAt: string;
  recordedAt?: string;
  source: Source;
  context?: Context;
  subject: Subject;
  classification?: Classification;
  severity?: Severity;
  releaseId?: string;
  evidence?: EvidenceItem[];
  payload: Record<string, unknown>;
  extensions?: Record<string, unknown>;
}

export interface ChangePayload {
  changeKind: ChangeKind;
  summary?: string;
  paths?: string[];
  before?: Record<string, unknown> | null;
  after?: Record<string, unknown> | null;
  noiseFiltered?: boolean;
  baselineSnapshotId?: string;
  candidateSnapshotId?: string;
  correctsEventId?: string;
  approvalId?: string;
}

export interface ChangeEvent {
  specVersion: string;
  eventId: string;
  eventType: string;
  occurredAt: string;
  recordedAt?: string;
  source: Source;
  context?: Context;
  subject: Subject;
  classification: Classification;
  severity?: Severity;
  releaseId?: string;
  evidence?: EvidenceItem[];
  payload: ChangePayload;
  extensions?: Record<string, unknown>;
}

export interface Resource {
  specVersion?: string;
  resourceId: string;
  resourceType: string;
  displayName: string;
  attributes?: Record<string, unknown>;
  labels?: Record<string, string>;
  firstSeenAt?: string;
  lastSeenAt?: string;
  parentResourceId?: string;
}

export interface Observation {
  specVersion?: string;
  observationId: string;
  observedAt: string;
  source: Source;
  subjects: Subject[];
  status: ObservationStatus;
  summary?: string;
  metrics?: Record<string, number | string | boolean>;
  evidence?: EvidenceItem[];
  raw?: Record<string, unknown>;
}

export interface NoisePolicy {
  filteredFields?: string[];
  notes?: string;
}

export interface Snapshot {
  specVersion?: string;
  snapshotId: string;
  capturedAt: string;
  source: Source;
  context?: Context;
  noisePolicy?: NoisePolicy;
  resources: Resource[];
  relationships?: Relationship[];
  digest?: Digest;
}

export interface Relationship {
  specVersion?: string;
  relationshipId: string;
  type: RelationshipType;
  from: string;
  to: string;
  attributes?: Record<string, unknown>;
  observedAt?: string;
  confidence?: number;
  discoverySources?: string[];
  firstSeenAt?: string;
  lastSeenAt?: string;
  confirmation?: "unreviewed" | "confirmed" | "rejected";
  reviewedBy?: string;
  reviewedAt?: string;
  reviewNote?: string;
}

export interface Actor {
  name: string;
  email?: string;
}

export interface ApprovalScope {
  eventIds?: string[];
  resourceIds?: string[];
  validFrom?: string;
  validUntil?: string;
}

export interface Approval {
  specVersion?: string;
  approvalId: string;
  decision: ApprovalDecision;
  decidedAt: string;
  actor: Actor;
  scope?: ApprovalScope;
  reason?: string;
  evidence?: EvidenceItem[];
}

export interface ExpectedChange {
  changeKind: string;
  resourceId: string;
  path?: string;
  summary?: string;
}

export interface Release {
  specVersion?: string;
  releaseId: string;
  version: string;
  status: ReleaseStatus;
  createdAt: string;
  startedAt?: string;
  completedAt?: string;
  source: Source;
  context?: Context;
  subject?: Subject;
  expectedChanges?: ExpectedChange[];
  evidence?: EvidenceItem[];
  notes?: string;
}
