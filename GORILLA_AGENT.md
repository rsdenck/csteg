============================================================
GORILLA AGENT (ZABBIX AGENT 3 STYLE)
GO + NATIVE OPENTELEMETRY
============================================================

GOALS

- replace zabbix_agentd with gorilla_agent
- native otel metrics, logs, traces
- passive + active mode
- exporter embedded
- linux first-class support
- template naming pattern compatible (zt-os-linux)
- key naming pattern lowercase snake_case
- no pt-br identifiers
- no uppercase keys

============================================================
FILESYSTEM LAYOUT (LINUX)
============================================================

/etc/gorilla/
└── gorilla_agent/
    ├── bin/
    │   └── gorilla_agent
    │
    ├── conf/
    │   ├── gorilla_agent.yaml
    │   ├── exporters.yaml
    │   ├── otel.yaml
    │   └── modules/
    │       ├── os_linux.yaml
    │       ├── network.yaml
    │       ├── process.yaml
    │       └── disk.yaml
    │
    ├── exporter/
    │   ├── prometheus.sock
    │   ├── otlp.sock
    │   └── http.sock
    │
    ├── tracer/
    │   ├── spans.db
    │   └── buffer/
    │
    └── log/
        └── gorilla_agent.log


============================================================
PROJECT STRUCTURE (GO)
============================================================

gorilla_agent/
├── cmd/
│   └── gorilla_agent/
│       └── main.go
│
├── internal/
│   ├── config/
│   │   └── loader.go
│   │
│   ├── runtime/
│   │   ├── scheduler.go
│   │   ├── worker_pool.go
│   │   └── lifecycle.go
│   │
│   ├── passive/
│   │   ├── listener.go
│   │   └── handler.go
│   │
│   ├── active/
│   │   ├── fetch_config.go
│   │   ├── sender.go
│   │   └── batcher.go
│   │
│   ├── collectors/
│   │   ├── os/
│   │   │   └── linux/
│   │   │       ├── cpu.go
│   │   │       ├── memory.go
│   │   │       ├── disk.go
│   │   │       ├── filesystem.go
│   │   │       ├── network.go
│   │   │       ├── uptime.go
│   │   │       └── load.go
│   │   │
│   │   ├── process/
│   │   │   └── process.go
│   │   │
│   │   ├── system/
│   │   │   └── hostname.go
│   │   │
│   │   └── custom/
│   │       └── user_parameter.go
│   │
│   ├── exporter/
│   │   ├── otlp.go
│   │   ├── prometheus.go
│   │   ├── http.go
│   │   └── registry.go
│   │
│   ├── telemetry/
│   │   ├── metrics.go
│   │   ├── logs.go
│   │   ├── traces.go
│   │   └── resource.go
│   │
│   ├── security/
│   │   ├── tls.go
│   │   ├── psk.go
│   │   └── auth.go
│   │
│   └── protocol/
│       ├── zabbix_compatible.go
│       └── otel_native.go
│
└── go.mod


============================================================
AGENT CAPABILITIES (MATCHING ZABBIX CLASSIC)
============================================================

- passive mode (tcp listener)
- active mode (push to backend)
- scheduler
- user parameters
- prefork worker pool (goroutines)
- tls support
- psk support
- configurable start_workers
- allow/deny key

============================================================
OTEL NATIVE IMPLEMENTATION
============================================================

EXPORTS:

1) metrics
   - otlp grpc
   - otlp http
   - prometheus endpoint

2) logs
   - otlp log exporter

3) traces
   - span generation per collection cycle
   - internal tracing for collector latency

RESOURCE ATTRIBUTES:

service.name = gorilla_agent
service.namespace = gorilla
host.name = <hostname>
os.type = linux
agent.version = 3

============================================================
KEY NAMING STANDARD (ZABBIX STYLE COMPATIBLE)
============================================================

template naming:

zt-os-linux
zt-network-linux
zt-process-linux

metric key pattern:

cpu_usage
cpu_load_1m
memory_total
memory_used
memory_free
disk_used
disk_free
disk_total
fs_inode_used
net_rx_bytes
net_tx_bytes
net_rx_packets
net_tx_packets
process_count
system_uptime
hostname

all lowercase
snake_case only
no spaces
no uppercase
no pt-br

============================================================
PASSIVE MODE FLOW
============================================================

frontend/backend -> tcp 10050
request key
gorilla_agent resolves collector
returns value
also emits otel metric internally

============================================================
ACTIVE MODE FLOW
============================================================

gorilla_agent -> backend
fetch item list
schedule collectors
batch values
push via otlp or custom protocol

============================================================
EXPORTER MODES
============================================================

mode: embedded
mode: otlp_only
mode: prometheus_only
mode: dual_export

============================================================
CONFIG (gorilla_agent.yaml)
============================================================

server: backend.local
mode: active
listen_port: 10050
start_workers: 5
hostname: node01
tls_enabled: true
exporter: otlp
otel_endpoint: http://collector:4317

============================================================
WINDOWS STRUCTURE
============================================================

C:\Program Files\Gorilla Agent\
    bin\
        gorilla_agent.exe
    conf\
        gorilla_agent.yaml
    log\
        gorilla_agent.log
    tracer\
    exporter\

installed as windows service:
"Gorilla Agent"

============================================================
FINAL OBJECTIVE
============================================================

- full zabbix agent parity
- otel native metrics/logs/traces
- exporter embedded
- template compatible naming
- enterprise ready
- multi-platform
- high performance
- minimal footprint
============================================================
END
============================================================
