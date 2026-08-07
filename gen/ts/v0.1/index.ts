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
