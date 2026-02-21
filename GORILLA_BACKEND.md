============================================================
GORILLA AGENT ↔ GORILLA BACKEND
COMMUNICATION ARCHITECTURE
============================================================

OBJECTIVE

Define a secure, scalable, OTEL-native communication model
between gorilla_agent and gorilla_backend.

============================================================
COMMUNICATION MODES
============================================================

MODE 1: OTEL NATIVE (PRIMARY)
MODE 2: GORILLA CONTROL CHANNEL (CONFIG + MGMT)
MODE 3: COMPATIBILITY PASSIVE MODE (OPTIONAL)

============================================================
GLOBAL DESIGN PRINCIPLE
============================================================

- telemetry always via otlp
- configuration always via control api
- no telemetry via custom json protocol
- no simulated transport
- tls mandatory in production
- multi-tenant header required

============================================================
1) TELEMETRY CHANNEL (PRIMARY DATA FLOW)
============================================================

PROTOCOL:
- otlp grpc (preferred)
- otlp http (fallback)

PORTS:
- 4317 (grpc)
- 4318 (http)

FLOW:

gorilla_agent
    -> collect metrics/logs/traces
    -> create otel resource attributes
    -> batch
    -> export via otlp
    -> gorilla_backend otel receiver
    -> internal pipeline
    -> storage (metrics/logs/traces db)

RESOURCE ATTRIBUTES REQUIRED:

service.name=gorilla_agent
service.namespace=gorilla
agent.version=3
host.name=<hostname>
tenant.id=<uuid>
environment=<prod|staging|dev>

NO custom telemetry format allowed.

============================================================
2) CONTROL CHANNEL (CONFIGURATION + MANAGEMENT)
============================================================

PROTOCOL:
- https rest api
- optional grpc control api

PORT:
- 8443 (default)

FLOW (ACTIVE MODE):

gorilla_agent
    -> authenticate
    -> fetch assigned templates
    -> fetch item definitions
    -> fetch schedule
    -> receive policy updates

ENDPOINTS (BACKEND):

GET  /api/v1/agent/bootstrap
GET  /api/v1/agent/templates
GET  /api/v1/agent/schedule
POST /api/v1/agent/heartbeat
POST /api/v1/agent/status

AUTHENTICATION:

- mTLS (recommended)
or
- agent token (signed jwt)
or
- api key (scoped per tenant)

============================================================
3) PASSIVE MODE (OPTIONAL - ZABBIX STYLE)
============================================================

PORT:
10050

FLOW:

backend -> tcp request key
gorilla_agent resolves collector
returns raw value
also emits otel metric internally

passive mode disabled by default.

============================================================
AGENT REGISTRATION FLOW
============================================================

FIRST START:

1) agent reads gorilla_agent.yaml
2) generates agent_id (uuid)
3) connects to backend bootstrap endpoint
4) registers host metadata
5) receives:
   - tenant_id
   - template assignment
   - policy rules
6) starts scheduler

============================================================
HEARTBEAT MODEL
============================================================

INTERVAL: configurable (default 30s)

POST /api/v1/agent/heartbeat

payload:

- agent_id
- hostname
- version
- last_export_timestamp
- health_status

NO telemetry data in heartbeat.

============================================================
BACKEND INGEST PIPELINE
============================================================

[ otlp receiver ]
        ↓
[ validation layer ]
        ↓
[ tenant isolation filter ]
        ↓
[ metrics processor ]
[ logs processor ]
[ traces processor ]
        ↓
[ storage layer ]

============================================================
SECURITY MODEL
============================================================

REQUIRED:

- tls 1.2+
- cert validation
- hostname verification
- tenant scoping

OPTIONAL:

- certificate pinning
- ip allowlist
- rate limiting

FORBIDDEN:

- plaintext telemetry
- shared global tokens
- anonymous agent registration

============================================================
MULTI-TENANT ENFORCEMENT
============================================================

every telemetry batch must contain:

tenant.id

backend rejects batch if:

- tenant missing
- agent not registered
- agent disabled
- invalid signature

============================================================
FAILOVER STRATEGY
============================================================

agent side:

- local buffer queue
- retry with exponential backoff
- disk fallback if memory full
- max retention configurable

backend side:

- horizontal scalable otlp receivers
- load balancer in front

============================================================
SCALABILITY DESIGN
============================================================

- stateless backend receivers
- sharded storage
- async ingestion
- batching mandatory
- compression enabled (gzip)

============================================================
COMPATIBILITY WITH TEMPLATE MODEL
============================================================

template example:

zt-os-linux

backend defines:

- metric keys
- expected collection interval
- threshold rules

agent receives template
maps template keys to collectors
exports via otel metric name:

gorilla.os.linux.disk_used
gorilla.os.linux.cpu_usage

frontend still displays key:

disk_used
cpu_usage

============================================================
FINAL ARCHITECTURE SUMMARY
============================================================

control plane  -> https api
data plane     -> otlp grpc/http
optional mode  -> passive tcp
security       -> mTLS
multi-tenant   -> mandatory attribute
storage        -> separated pipelines
no simulated telemetry
no custom fake protocol

============================================================
END
============================================================
