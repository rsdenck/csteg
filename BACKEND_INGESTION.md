============================================================
GORILLA OBSERVABILITY PLATFORM
BACKEND INGESTION ARCHITECTURE
============================================================

OBJECTIVE

Design a scalable, multi-tenant, OTEL-native ingestion
architecture for:

- metrics
- logs
- traces

============================================================
HIGH LEVEL ARCHITECTURE
============================================================

               ┌────────────────────┐
               │   gorilla_agent    │
               └─────────┬──────────┘
                         │ OTLP
                         ▼
               ┌────────────────────┐
               │  ingress gateway   │
               │  (load balancer)   │
               └─────────┬──────────┘
                         ▼
               ┌────────────────────┐
               │ otlp receiver tier │
               └─────────┬──────────┘
                         ▼
               ┌────────────────────┐
               │ validation layer   │
               └─────────┬──────────┘
                         ▼
               ┌────────────────────┐
               │ tenant isolation   │
               └─────────┬──────────┘
                         ▼
      ┌──────────────┬───────────────┬───────────────┐
      ▼              ▼               ▼
 metrics pipeline   logs pipeline    traces pipeline
      ▼              ▼               ▼
 metrics storage    logs storage     traces storage

============================================================
INGRESS LAYER
============================================================

COMPONENT:
gorilla_ingress

RESPONSIBILITIES:
- tls termination
- rate limiting
- request validation
- compression handling
- routing to otlp receivers

PORTS:
- 4317 (grpc)
- 4318 (http)

REQUIREMENTS:
- tls 1.2+
- multi-tenant header validation
- reject anonymous traffic

============================================================
OTLP RECEIVER TIER
============================================================

COMPONENT:
gorilla_otel_receiver

FUNCTION:
- receive otlp metrics/logs/traces
- decode protobuf
- forward to processing layer

SCALE:
- horizontally scalable
- stateless
- auto-discoverable

============================================================
VALIDATION LAYER
============================================================

VALIDATIONS:

- agent_id exists
- tenant_id valid
- signature valid
- schema valid
- payload size limit
- timestamp sanity check

REJECTION CONDITIONS:

- missing resource attributes
- invalid tenant
- malformed batch
- expired certificate

============================================================
TENANT ISOLATION LAYER
============================================================

FUNCTION:

- enforce strict tenant boundary
- attach tenant_id internally
- prevent cross-tenant contamination

IMPLEMENTATION:

- middleware enrichment
- context-based tenant injection
- partition key tagging

============================================================
METRICS PIPELINE
============================================================

PROCESSORS:

- batch processor
- attribute normalizer
- aggregation processor
- retention policy router

STORAGE OPTIONS:

PRIMARY:
- timeseries optimized database

REQUIREMENTS:

- high write throughput
- label-based indexing
- shard by tenant
- retention configurable per tenant
- compression enabled

DATA MODEL:

metric_name
timestamp
value
labels:
  host.name
  tenant.id
  service.name
  environment
  template_id

============================================================
LOGS PIPELINE
============================================================

PROCESSORS:

- structured log parser
- severity normalizer
- enrichment processor
- retention router

STORAGE REQUIREMENTS:

- full-text search
- time-based partitioning
- tenant isolation
- fast filtering by host/service

DATA MODEL:

timestamp
tenant_id
host.name
service.name
severity
body
attributes (jsonb)

============================================================
TRACES PIPELINE
============================================================

PROCESSORS:

- span validator
- trace assembler
- dependency mapper
- service graph builder

STORAGE REQUIREMENTS:

- trace_id index
- span_id index
- parent-child mapping
- service dependency storage
- time-based retention

DATA MODEL:

trace_id
span_id
parent_span_id
tenant_id
service.name
operation
start_time
duration
attributes

============================================================
STORAGE LAYER DESIGN
============================================================

METRICS DB:
- append-only
- partitioned by time + tenant
- retention per tenant
- compression required

LOGS DB:
- columnar or inverted index
- partition by time
- tenant filter mandatory

TRACES DB:
- optimized for trace reconstruction
- secondary index on trace_id

============================================================
RETENTION MANAGEMENT
============================================================

PER TENANT CONFIG:

metrics_retention_days
logs_retention_days
traces_retention_days

automatic lifecycle policy:
- hot storage
- warm storage
- delete

============================================================
BACKPRESSURE STRATEGY
============================================================

AGENT SIDE:
- local buffer
- exponential retry
- disk spillover

BACKEND SIDE:
- ingestion queue
- rate limiter
- overload protection

============================================================
SCALABILITY MODEL
============================================================

- stateless ingestion tier
- sharded storage
- load balancer in front
- autoscaling based on:
    ingestion_rate
    cpu
    memory
    queue depth

============================================================
SECURITY MODEL
============================================================

REQUIRED:

- mTLS preferred
- certificate rotation
- agent identity validation
- per-tenant rate limiting
- audit logging

FORBIDDEN:

- shared tenant credentials
- plaintext ingestion
- telemetry without tenant.id

============================================================
OBSERVABILITY OF THE OBSERVABILITY
============================================================

internal metrics exposed:

gorilla.ingestion.rate
gorilla.ingestion.errors
gorilla.pipeline.latency
gorilla.storage.write_latency
gorilla.queue.depth

exported via internal otel

============================================================
FAILURE ISOLATION
============================================================

- metrics failure does not block logs
- logs failure does not block traces
- independent pipelines
- circuit breaker per pipeline

============================================================
FINAL ARCHITECTURE PRINCIPLES
============================================================

- otel native
- multi-tenant first
- horizontally scalable
- storage separated by signal type
- zero simulated data
- strict validation
- enterprise security
- template compatible

============================================================
END
============================================================
