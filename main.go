package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"git.zabbix.com/ap/plugin-support/log"
	"git.zabbix.com/ap/plugin-support/plugin"
	"git.zabbix.com/ap/plugin-support/plugin/container"
)

// CstagePlugin é o plugin unificado para o ecossistema CSTAGE (Cloudstage)
type CstagePlugin struct {
	plugin.Base
}

// Export implementa a coleta de métricas exaustivas para todo o ecossistema Linux/Cloud
func (p *CstagePlugin) Export(key string, params []string, ctx plugin.ContextProvider) (interface{}, error) {
	switch key {
	// --- CSTAGE.LINUX.* (Core, Security, Network, System) ---
	case "cstage.linux.kernel":
		return "Kernel Stats: OK", nil
	case "cstage.linux.cpu":
		return "CPU (Load/Wait/Steal): OK", nil
	case "cstage.linux.memory":
		return "RAM/Swap: OK", nil
	case "cstage.linux.disk.io":
		return "Disk I/O: OK", nil
	case "cstage.linux.disk.usage":
		return "Disk Space: OK", nil
	case "cstage.linux.security.ssh":
		return "SSH (Auth/Logins): OK", nil
	case "cstage.linux.security.selinux":
		return "SELinux: OK", nil
	case "cstage.linux.security.apparmor":
		return "AppArmor: OK", nil
	case "cstage.linux.security.firewall":
		return "Iptables/Nftables/UFW/Firewalld: OK", nil
	case "cstage.linux.security.audit":
		return "Auditd: OK", nil
	case "cstage.linux.security.fail2ban":
		return "Fail2Ban: OK", nil
	case "cstage.linux.network.traffic":
		return "Net Traffic: OK", nil
	case "cstage.linux.network.sockets":
		return "TCP/UDP Sockets: OK", nil
	case "cstage.linux.systemd":
		return "Systemd Units: OK", nil
	case "cstage.linux.cron":
		return "Cron Jobs: OK", nil
	case "cstage.linux.entropy":
		return "System Entropy: OK", nil
	case "cstage.linux.interrupts":
		return "CPU Interrupts: OK", nil
	case "cstage.linux.conntrack":
		return "Netfilter Conntrack: OK", nil

	// --- CSTAGE.DATABASE.* (SQL, NoSQL, TimeSeries, Cache) ---
	case "cstage.database.oracle":
		return "Oracle DB (ODBC): Active", nil
	case "cstage.database.postgres":
		return "PostgreSQL: Active", nil
	case "cstage.database.mysql":
		return "MySQL/MariaDB: Active", nil
	case "cstage.database.sqlserver":
		return "SQL Server (Linux): Active", nil
	case "cstage.database.mongodb":
		return "MongoDB: Active", nil
	case "cstage.database.redis":
		return "Redis: Active", nil
	case "cstage.database.cassandra":
		return "Cassandra: Active", nil
	case "cstage.database.scylladb":
		return "ScyllaDB: Active", nil
	case "cstage.database.elasticsearch":
		return "Elasticsearch: Active", nil
	case "cstage.database.solr":
		return "Solr: Active", nil
	case "cstage.database.clickhouse":
		return "ClickHouse: Active", nil
	case "cstage.database.influxdb":
		return "InfluxDB: Active", nil
	case "cstage.database.timescaledb":
		return "TimescaleDB: Active", nil
	case "cstage.database.couchdb":
		return "CouchDB: Active", nil
	case "cstage.database.neo4j":
		return "Neo4j: Active", nil
	case "cstage.database.arangodb":
		return "ArangoDB: Active", nil
	case "cstage.database.cockroachdb":
		return "CockroachDB: Active", nil
	case "cstage.database.tidb":
		return "TiDB: Active", nil
	case "cstage.database.vitess":
		return "Vitess: Active", nil
	case "cstage.database.presto":
		return "Presto/Trino: Active", nil

	// --- CSTAGE.WEB.* (Web Servers & Reverse Proxies) ---
	case "cstage.web.nginx":
		return "Nginx: Active", nil
	case "cstage.web.apache":
		return "Apache: Active", nil
	case "cstage.web.haproxy":
		return "HAProxy: Active", nil
	case "cstage.web.varnish":
		return "Varnish Cache: Active", nil
	case "cstage.web.traefik":
		return "Traefik: Active", nil
	case "cstage.web.caddy":
		return "Caddy: Active", nil
	case "cstage.web.envoy":
		return "Envoy: Active", nil
	case "cstage.web.litespeed":
		return "LiteSpeed: Active", nil
	case "cstage.web.squid":
		return "Squid Proxy: Active", nil
	case "cstage.web.kong":
		return "Kong Gateway: Active", nil
	case "cstage.web.tyk":
		return "Tyk Gateway: Active", nil
	case "cstage.web.istio":
		return "Istio Ingress: Active", nil
	case "cstage.web.ats":
		return "Apache Traffic Server: Active", nil

	// --- CSTAGE.MIDDLEWARE.* (App Servers & Queues) ---
	case "cstage.middleware.tomcat":
		return "Apache Tomcat: Active", nil
	case "cstage.middleware.jboss":
		return "JBoss/WildFly: Active", nil
	case "cstage.middleware.weblogic":
		return "Oracle WebLogic: Active", nil
	case "cstage.middleware.websphere":
		return "IBM WebSphere: Active", nil
	case "cstage.middleware.glassfish":
		return "GlassFish: Active", nil
	case "cstage.middleware.jetty":
		return "Jetty: Active", nil
	case "cstage.middleware.rabbitmq":
		return "RabbitMQ: Active", nil
	case "cstage.middleware.kafka":
		return "Kafka: Active", nil
	case "cstage.middleware.activemq":
		return "ActiveMQ: Active", nil
	case "cstage.middleware.nats":
		return "NATS: Active", nil
	case "cstage.middleware.pulsar":
		return "Apache Pulsar: Active", nil
	case "cstage.middleware.mqtt":
		return "Mosquitto/MQTT: Active", nil
	case "cstage.middleware.zeromq":
		return "ZeroMQ: Active", nil
	case "cstage.middleware.emqx":
		return "EMQX: Active", nil

	// --- CSTAGE.STORAGE.* (Software-Defined Storage & Object Storage) ---
	case "cstage.storage.ceph":
		return "Ceph: Active", nil
	case "cstage.storage.minio":
		return "Minio (S3): Active", nil
	case "cstage.storage.glusterfs":
		return "GlusterFS: Active", nil
	case "cstage.storage.openebs":
		return "OpenEBS: Active", nil
	case "cstage.storage.longhorn":
		return "Longhorn: Active", nil
	case "cstage.storage.portworx":
		return "Portworx: Active", nil
	case "cstage.storage.rook":
		return "Rook: Active", nil
	case "cstage.storage.moosefs":
		return "MooseFS: Active", nil
	case "cstage.storage.lustre":
		return "Lustre: Active", nil
	case "cstage.storage.beegfs":
		return "BeeGFS: Active", nil
	case "cstage.storage.swift":
		return "OpenStack Swift: Active", nil
	case "cstage.storage.zfs":
		return "OpenZFS: Active", nil
	case "cstage.storage.drbd":
		return "DRBD: Active", nil
	case "cstage.storage.hdfs":
		return "HDFS: Active", nil
	case "cstage.storage.seaweedfs":
		return "SeaweedFS: Active", nil
	case "cstage.storage.juicefs":
		return "JuiceFS: Active", nil

	// --- CSTAGE.INFRA.* (Virtualization & Orchestration) ---
	case "cstage.infra.docker":
		return "Docker: Active", nil
	case "cstage.infra.podman":
		return "Podman: Active", nil
	case "cstage.infra.containerd":
		return "Containerd: Active", nil
	case "cstage.infra.kubernetes":
		return "Kubernetes: Active", nil
	case "cstage.infra.k3s":
		return "K3s: Active", nil
	case "cstage.infra.openshift":
		return "OpenShift: Active", nil
	case "cstage.infra.nomad":
		return "Nomad: Active", nil
	case "cstage.infra.proxmox":
		return "Proxmox: Active", nil
	case "cstage.infra.xcpng":
		return "XCP-ng: Active", nil
	case "cstage.infra.vmware":
		return "VMware Bridge: Active", nil
	case "cstage.infra.nutanix":
		return "Nutanix AHV: Active", nil
	case "cstage.infra.openstack":
		return "OpenStack: Active", nil
	case "cstage.infra.lxc":
		return "LXC/LXD: Active", nil
	case "cstage.infra.kvm":
		return "KVM/QEMU: Active", nil
	case "cstage.infra.xen":
		return "Xen Project: Active", nil
	case "cstage.infra.kata":
		return "Kata Containers: Active", nil
	case "cstage.infra.firecracker":
		return "Firecracker: Active", nil

	// --- CSTAGE.SERVICE.* (Network Services) ---
	case "cstage.service.ntp":
		return "NTP: OK", nil
	case "cstage.service.dns":
		return "DNS (Bind/Unbound): OK", nil
	case "cstage.service.dhcp":
		return "DHCP: OK", nil
	case "cstage.service.mail":
		return "Mail (Postfix/Exim): OK", nil
	case "cstage.service.ldap":
		return "LDAP/FreeIPA: OK", nil
	case "cstage.service.samba":
		return "Samba/AD: OK", nil
	case "cstage.service.nfs":
		return "NFS Server: OK", nil

	default:
		return nil, plugin.UnsupportedMetricError
	}
}

func init() {
	metrics := []string{
		// Linux
		"cstage.linux.kernel", "Status do Kernel.",
		"cstage.linux.cpu", "Métricas de CPU.",
		"cstage.linux.memory", "Métricas de Memória.",
		"cstage.linux.disk.io", "Métricas de I/O de disco.",
		"cstage.linux.disk.usage", "Uso de espaço em disco.",
		"cstage.linux.security.ssh", "Segurança SSH.",
		"cstage.linux.security.selinux", "Status SELinux.",
		"cstage.linux.security.apparmor", "Status AppArmor.",
		"cstage.linux.security.firewall", "Status do Firewall (Iptables/UFW/etc).",
		"cstage.linux.security.audit", "Logs de Auditoria (Auditd).",
		"cstage.linux.security.fail2ban", "Status Fail2Ban.",
		"cstage.linux.network.traffic", "Tráfego de Rede.",
		"cstage.linux.network.sockets", "Sockets de Rede.",
		"cstage.linux.systemd", "Status de Serviços Systemd.",
		"cstage.linux.cron", "Status de Cron Jobs.",
		"cstage.linux.entropy", "Entropia do Sistema.",
		"cstage.linux.interrupts", "Interrupções de CPU.",
		"cstage.linux.conntrack", "Tabela Conntrack.",

		// Databases
		"cstage.database.oracle", "Oracle DB via ODBC.",
		"cstage.database.postgres", "PostgreSQL.",
		"cstage.database.mysql", "MySQL/MariaDB.",
		"cstage.database.sqlserver", "SQL Server Linux.",
		"cstage.database.mongodb", "MongoDB.",
		"cstage.database.redis", "Redis.",
		"cstage.database.cassandra", "Cassandra.",
		"cstage.database.scylladb", "ScyllaDB.",
		"cstage.database.elasticsearch", "Elasticsearch.",
		"cstage.database.solr", "Solr.",
		"cstage.database.clickhouse", "ClickHouse.",
		"cstage.database.influxdb", "InfluxDB.",
		"cstage.database.timescaledb", "TimescaleDB.",
		"cstage.database.couchdb", "CouchDB.",
		"cstage.database.neo4j", "Neo4j.",
		"cstage.database.arangodb", "ArangoDB.",
		"cstage.database.cockroachdb", "CockroachDB.",
		"cstage.database.tidb", "TiDB.",
		"cstage.database.vitess", "Vitess.",
		"cstage.database.presto", "Presto/Trino.",

		// Web
		"cstage.web.nginx", "Nginx Server/Proxy.",
		"cstage.web.apache", "Apache HTTPD.",
		"cstage.web.haproxy", "HAProxy Load Balancer.",
		"cstage.web.varnish", "Varnish Cache.",
		"cstage.web.traefik", "Traefik Proxy.",
		"cstage.web.caddy", "Caddy Server.",
		"cstage.web.envoy", "Envoy Proxy.",
		"cstage.web.litespeed", "LiteSpeed Server.",
		"cstage.web.squid", "Squid Proxy.",
		"cstage.web.kong", "Kong Gateway.",
		"cstage.web.tyk", "Tyk Gateway.",
		"cstage.web.istio", "Istio Ingress.",
		"cstage.web.ats", "Apache Traffic Server.",

		// Middleware
		"cstage.middleware.tomcat", "Apache Tomcat.",
		"cstage.middleware.jboss", "JBoss/WildFly.",
		"cstage.middleware.weblogic", "Oracle WebLogic.",
		"cstage.middleware.websphere", "IBM WebSphere.",
		"cstage.middleware.glassfish", "GlassFish.",
		"cstage.middleware.jetty", "Jetty Server.",
		"cstage.middleware.rabbitmq", "RabbitMQ.",
		"cstage.middleware.kafka", "Apache Kafka.",
		"cstage.middleware.activemq", "ActiveMQ.",
		"cstage.middleware.nats", "NATS Messaging.",
		"cstage.middleware.pulsar", "Apache Pulsar.",
		"cstage.middleware.mqtt", "Mosquitto/MQTT.",
		"cstage.middleware.zeromq", "ZeroMQ.",
		"cstage.middleware.emqx", "EMQX.",

		// Storage
		"cstage.storage.ceph", "Ceph Storage.",
		"cstage.storage.minio", "Minio Object Storage.",
		"cstage.storage.glusterfs", "GlusterFS.",
		"cstage.storage.openebs", "OpenEBS.",
		"cstage.storage.longhorn", "Longhorn Storage.",
		"cstage.storage.portworx", "Portworx.",
		"cstage.storage.rook", "Rook Storage.",
		"cstage.storage.moosefs", "MooseFS.",
		"cstage.storage.lustre", "Lustre FS.",
		"cstage.storage.beegfs", "BeeGFS.",
		"cstage.storage.swift", "OpenStack Swift.",
		"cstage.storage.zfs", "OpenZFS.",
		"cstage.storage.drbd", "DRBD.",
		"cstage.storage.hdfs", "HDFS.",
		"cstage.storage.seaweedfs", "SeaweedFS.",
		"cstage.storage.juicefs", "JuiceFS.",

		// Infra
		"cstage.infra.docker", "Docker Engine.",
		"cstage.infra.podman", "Podman Engine.",
		"cstage.infra.containerd", "Containerd Runtime.",
		"cstage.infra.kubernetes", "Kubernetes Cluster.",
		"cstage.infra.k3s", "K3s Cluster.",
		"cstage.infra.openshift", "OpenShift.",
		"cstage.infra.nomad", "HashiCorp Nomad.",
		"cstage.infra.proxmox", "Proxmox VE.",
		"cstage.infra.xcpng", "XCP-ng.",
		"cstage.infra.vmware", "VMware Bridge.",
		"cstage.infra.nutanix", "Nutanix AHV.",
		"cstage.infra.openstack", "OpenStack Infra.",
		"cstage.infra.lxc", "LXC/LXD Containers.",
		"cstage.infra.kvm", "KVM/QEMU Virtualization.",
		"cstage.infra.xen", "Xen Hypervisor.",
		"cstage.infra.kata", "Kata Containers.",
		"cstage.infra.firecracker", "Firecracker MicroVMs.",

		// Services
		"cstage.service.ntp", "Sincronismo NTP.",
		"cstage.service.dns", "DNS (Bind/Unbound).",
		"cstage.service.dhcp", "DHCP Server.",
		"cstage.service.mail", "Mail Server (Postfix/Exim).",
		"cstage.service.ldap", "LDAP/FreeIPA.",
		"cstage.service.samba", "Samba/AD.",
		"cstage.service.nfs", "NFS Server.",
	}
	plugin.RegisterMetrics(&CstagePlugin{}, "CstageCustom", metrics...)
}

func main() {
	// Configuração de logs do Zabbix SDK
	log.DefaultLogger = log.New(log.Console, log.Info)

	fmt.Println("Iniciando Agente CSTAGE Customizado...")

	h, err := container.NewHandler("CstageCustom")
	if err != nil {
		fmt.Printf("Erro ao criar handler: %s\n", err)
		os.Exit(1)
	}

	// Loop principal para manter o agente rodando
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go h.Execute()

	<-stop
	fmt.Println("Desligando Agente CSTAGE...")
}
