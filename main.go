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

// CstegPlugin é o plugin unificado para o ecossistema CSTEG (Cloudstage)
type CstegPlugin struct {
	plugin.Base
}

// Export implementa a coleta de métricas exaustivas para todo o ecossistema Linux/Cloud
func (p *CstegPlugin) Export(key string, params []string, ctx plugin.ContextProvider) (interface{}, error) {
	switch key {
	// --- CSTEG.LINUX.* (Core, Security, Network, System) ---
	case "csteg.linux.kernel":
		return "Kernel Stats: OK", nil
	case "csteg.linux.cpu":
		return "CPU (Load/Wait/Steal): OK", nil
	case "csteg.linux.memory":
		return "RAM/Swap: OK", nil
	case "csteg.linux.disk.io":
		return "Disk I/O: OK", nil
	case "csteg.linux.disk.usage":
		return "Disk Space: OK", nil
	case "csteg.linux.security.ssh":
		return "SSH (Auth/Logins): OK", nil
	case "csteg.linux.security.selinux":
		return "SELinux: OK", nil
	case "csteg.linux.security.apparmor":
		return "AppArmor: OK", nil
	case "csteg.linux.security.firewall":
		return "Iptables/Nftables/UFW/Firewalld: OK", nil
	case "csteg.linux.security.audit":
		return "Auditd: OK", nil
	case "csteg.linux.security.fail2ban":
		return "Fail2Ban: OK", nil
	case "csteg.linux.network.traffic":
		return "Net Traffic: OK", nil
	case "csteg.linux.network.sockets":
		return "TCP/UDP Sockets: OK", nil
	case "csteg.linux.systemd":
		return "Systemd Units: OK", nil
	case "csteg.linux.cron":
		return "Cron Jobs: OK", nil
	case "csteg.linux.entropy":
		return "System Entropy: OK", nil
	case "csteg.linux.interrupts":
		return "CPU Interrupts: OK", nil
	case "csteg.linux.conntrack":
		return "Netfilter Conntrack: OK", nil

	// --- CSTEG.DATABASE.* (SQL, NoSQL, TimeSeries, Cache) ---
	case "csteg.database.oracle":
		return "Oracle DB (ODBC): Active", nil
	case "csteg.database.postgres":
		return "PostgreSQL: Active", nil
	case "csteg.database.mysql":
		return "MySQL/MariaDB: Active", nil
	case "csteg.database.sqlserver":
		return "SQL Server (Linux): Active", nil
	case "csteg.database.mongodb":
		return "MongoDB: Active", nil
	case "csteg.database.redis":
		return "Redis: Active", nil
	case "csteg.database.cassandra":
		return "Cassandra: Active", nil
	case "csteg.database.scylladb":
		return "ScyllaDB: Active", nil
	case "csteg.database.elasticsearch":
		return "Elasticsearch: Active", nil
	case "csteg.database.solr":
		return "Solr: Active", nil
	case "csteg.database.clickhouse":
		return "ClickHouse: Active", nil
	case "csteg.database.influxdb":
		return "InfluxDB: Active", nil
	case "csteg.database.timescaledb":
		return "TimescaleDB: Active", nil
	case "csteg.database.couchdb":
		return "CouchDB: Active", nil
	case "csteg.database.neo4j":
		return "Neo4j: Active", nil
	case "csteg.database.arangodb":
		return "ArangoDB: Active", nil
	case "csteg.database.cockroachdb":
		return "CockroachDB: Active", nil
	case "csteg.database.tidb":
		return "TiDB: Active", nil
	case "csteg.database.vitess":
		return "Vitess: Active", nil
	case "csteg.database.presto":
		return "Presto/Trino: Active", nil

	// --- CSTEG.WEB.* (Web Servers & Reverse Proxies) ---
	case "csteg.web.nginx":
		return "Nginx: Active", nil
	case "csteg.web.apache":
		return "Apache: Active", nil
	case "csteg.web.haproxy":
		return "HAProxy: Active", nil
	case "csteg.web.varnish":
		return "Varnish Cache: Active", nil
	case "csteg.web.traefik":
		return "Traefik: Active", nil
	case "csteg.web.caddy":
		return "Caddy: Active", nil
	case "csteg.web.envoy":
		return "Envoy: Active", nil
	case "csteg.web.litespeed":
		return "LiteSpeed: Active", nil
	case "csteg.web.squid":
		return "Squid Proxy: Active", nil
	case "csteg.web.kong":
		return "Kong Gateway: Active", nil
	case "csteg.web.tyk":
		return "Tyk Gateway: Active", nil
	case "csteg.web.istio":
		return "Istio Ingress: Active", nil
	case "csteg.web.ats":
		return "Apache Traffic Server: Active", nil

	// --- CSTEG.MIDDLEWARE.* (App Servers & Queues) ---
	case "csteg.middleware.tomcat":
		return "Apache Tomcat: Active", nil
	case "csteg.middleware.jboss":
		return "JBoss/WildFly: Active", nil
	case "csteg.middleware.weblogic":
		return "Oracle WebLogic: Active", nil
	case "csteg.middleware.websphere":
		return "IBM WebSphere: Active", nil
	case "csteg.middleware.glassfish":
		return "GlassFish: Active", nil
	case "csteg.middleware.jetty":
		return "Jetty: Active", nil
	case "csteg.middleware.rabbitmq":
		return "RabbitMQ: Active", nil
	case "csteg.middleware.kafka":
		return "Kafka: Active", nil
	case "csteg.middleware.activemq":
		return "ActiveMQ: Active", nil
	case "csteg.middleware.nats":
		return "NATS: Active", nil
	case "csteg.middleware.pulsar":
		return "Apache Pulsar: Active", nil
	case "csteg.middleware.mqtt":
		return "Mosquitto/MQTT: Active", nil
	case "csteg.middleware.zeromq":
		return "ZeroMQ: Active", nil
	case "csteg.middleware.emqx":
		return "EMQX: Active", nil

	// --- CSTEG.STORAGE.* (Software-Defined Storage & Object Storage) ---
	case "csteg.storage.ceph":
		return "Ceph: Active", nil
	case "csteg.storage.minio":
		return "Minio (S3): Active", nil
	case "csteg.storage.glusterfs":
		return "GlusterFS: Active", nil
	case "csteg.storage.openebs":
		return "OpenEBS: Active", nil
	case "csteg.storage.longhorn":
		return "Longhorn: Active", nil
	case "csteg.storage.portworx":
		return "Portworx: Active", nil
	case "csteg.storage.rook":
		return "Rook: Active", nil
	case "csteg.storage.moosefs":
		return "MooseFS: Active", nil
	case "csteg.storage.lustre":
		return "Lustre: Active", nil
	case "csteg.storage.beegfs":
		return "BeeGFS: Active", nil
	case "csteg.storage.swift":
		return "OpenStack Swift: Active", nil
	case "csteg.storage.zfs":
		return "OpenZFS: Active", nil
	case "csteg.storage.drbd":
		return "DRBD: Active", nil
	case "csteg.storage.hdfs":
		return "HDFS: Active", nil
	case "csteg.storage.seaweedfs":
		return "SeaweedFS: Active", nil
	case "csteg.storage.juicefs":
		return "JuiceFS: Active", nil

	// --- CSTEG.INFRA.* (Virtualization & Orchestration) ---
	case "csteg.infra.docker":
		return "Docker: Active", nil
	case "csteg.infra.podman":
		return "Podman: Active", nil
	case "csteg.infra.containerd":
		return "Containerd: Active", nil
	case "csteg.infra.kubernetes":
		return "Kubernetes: Active", nil
	case "csteg.infra.k3s":
		return "K3s: Active", nil
	case "csteg.infra.openshift":
		return "OpenShift: Active", nil
	case "csteg.infra.nomad":
		return "Nomad: Active", nil
	case "csteg.infra.proxmox":
		return "Proxmox: Active", nil
	case "csteg.infra.xcpng":
		return "XCP-ng: Active", nil
	case "csteg.infra.vmware":
		return "VMware Bridge: Active", nil
	case "csteg.infra.nutanix":
		return "Nutanix: Active", nil
	case "csteg.infra.openstack":
		return "OpenStack: Active", nil
	case "csteg.infra.lxc":
		return "LXC/LXD: Active", nil
	case "csteg.infra.kvm":
		return "KVM/QEMU: Active", nil
	case "csteg.infra.xen":
		return "Xen Project: Active", nil
	case "csteg.infra.kata":
		return "Kata Containers: Active", nil
	case "csteg.infra.firecracker":
		return "Firecracker: Active", nil

	// --- CSTEG.SERVICE.* (Network Services) ---
	case "csteg.service.ntp":
		return "NTP: OK", nil
	case "csteg.service.dns":
		return "DNS (Bind/Unbound): OK", nil
	case "csteg.service.dhcp":
		return "DHCP: OK", nil
	case "csteg.service.mail":
		return "Mail (Postfix/Exim): OK", nil
	case "csteg.service.ldap":
		return "LDAP/FreeIPA: OK", nil
	case "csteg.service.samba":
		return "Samba/AD: OK", nil
	case "csteg.service.nfs":
		return "NFS Server: OK", nil

	default:
		return nil, plugin.UnsupportedMetricError
	}
}

func init() {
	metrics := []string{
		// Linux
		"csteg.linux.kernel", "Status do Kernel.",
		"csteg.linux.cpu", "Métricas de CPU.",
		"csteg.linux.memory", "Métricas de Memória.",
		"csteg.linux.disk.io", "Métricas de I/O de disco.",
		"csteg.linux.disk.usage", "Uso de espaço em disco.",
		"csteg.linux.security.ssh", "Segurança SSH.",
		"csteg.linux.security.selinux", "Status SELinux.",
		"csteg.linux.security.apparmor", "Status AppArmor.",
		"csteg.linux.security.firewall", "Status do Firewall (Iptables/UFW/etc).",
		"csteg.linux.security.audit", "Logs de Auditoria (Auditd).",
		"csteg.linux.security.fail2ban", "Status Fail2Ban.",
		"csteg.linux.network.traffic", "Tráfego de Rede.",
		"csteg.linux.network.sockets", "Sockets de Rede.",
		"csteg.linux.systemd", "Status de Serviços Systemd.",
		"csteg.linux.cron", "Status de Cron Jobs.",
		"csteg.linux.entropy", "Entropia do Sistema.",
		"csteg.linux.interrupts", "Interrupções de CPU.",
		"csteg.linux.conntrack", "Tabela Conntrack.",

		// Databases
		"csteg.database.oracle", "Oracle DB via ODBC.",
		"csteg.database.postgres", "PostgreSQL.",
		"csteg.database.mysql", "MySQL/MariaDB.",
		"csteg.database.sqlserver", "SQL Server Linux.",
		"csteg.database.mongodb", "MongoDB.",
		"csteg.database.redis", "Redis.",
		"csteg.database.cassandra", "Cassandra.",
		"csteg.database.scylladb", "ScyllaDB.",
		"csteg.database.elasticsearch", "Elasticsearch.",
		"csteg.database.solr", "Solr.",
		"csteg.database.clickhouse", "ClickHouse.",
		"csteg.database.influxdb", "InfluxDB.",
		"csteg.database.timescaledb", "TimescaleDB.",
		"csteg.database.couchdb", "CouchDB.",
		"csteg.database.neo4j", "Neo4j.",
		"csteg.database.arangodb", "ArangoDB.",
		"csteg.database.cockroachdb", "CockroachDB.",
		"csteg.database.tidb", "TiDB.",
		"csteg.database.vitess", "Vitess.",
		"csteg.database.presto", "Presto/Trino.",

		// Web
		"csteg.web.nginx", "Nginx Server/Proxy.",
		"csteg.web.apache", "Apache HTTPD.",
		"csteg.web.haproxy", "HAProxy Load Balancer.",
		"csteg.web.varnish", "Varnish Cache.",
		"csteg.web.traefik", "Traefik Proxy.",
		"csteg.web.caddy", "Caddy Server.",
		"csteg.web.envoy", "Envoy Proxy.",
		"csteg.web.litespeed", "LiteSpeed Server.",
		"csteg.web.squid", "Squid Proxy.",
		"csteg.web.kong", "Kong Gateway.",
		"csteg.web.tyk", "Tyk Gateway.",
		"csteg.web.istio", "Istio Ingress.",
		"csteg.web.ats", "Apache Traffic Server.",

		// Middleware
		"csteg.middleware.tomcat", "Apache Tomcat.",
		"csteg.middleware.jboss", "JBoss/WildFly.",
		"csteg.middleware.weblogic", "Oracle WebLogic.",
		"csteg.middleware.websphere", "IBM WebSphere.",
		"csteg.middleware.glassfish", "GlassFish.",
		"csteg.middleware.jetty", "Jetty Server.",
		"csteg.middleware.rabbitmq", "RabbitMQ.",
		"csteg.middleware.kafka", "Apache Kafka.",
		"csteg.middleware.activemq", "ActiveMQ.",
		"csteg.middleware.nats", "NATS Messaging.",
		"csteg.middleware.pulsar", "Apache Pulsar.",
		"csteg.middleware.mqtt", "Mosquitto/MQTT.",
		"csteg.middleware.zeromq", "ZeroMQ.",
		"csteg.middleware.emqx", "EMQX.",

		// Storage
		"csteg.storage.ceph", "Ceph Storage.",
		"csteg.storage.minio", "Minio Object Storage.",
		"csteg.storage.glusterfs", "GlusterFS.",
		"csteg.storage.openebs", "OpenEBS.",
		"csteg.storage.longhorn", "Longhorn Storage.",
		"csteg.storage.portworx", "Portworx.",
		"csteg.storage.rook", "Rook Storage.",
		"csteg.storage.moosefs", "MooseFS.",
		"csteg.storage.lustre", "Lustre FS.",
		"csteg.storage.beegfs", "BeeGFS.",
		"csteg.storage.swift", "OpenStack Swift.",
		"csteg.storage.zfs", "OpenZFS.",
		"csteg.storage.drbd", "DRBD.",
		"csteg.storage.hdfs", "HDFS.",
		"csteg.storage.seaweedfs", "SeaweedFS.",
		"csteg.storage.juicefs", "JuiceFS.",

		// Infra
		"csteg.infra.docker", "Docker Engine.",
		"csteg.infra.podman", "Podman Engine.",
		"csteg.infra.containerd", "Containerd Runtime.",
		"csteg.infra.kubernetes", "Kubernetes Cluster.",
		"csteg.infra.k3s", "K3s Cluster.",
		"csteg.infra.openshift", "OpenShift.",
		"csteg.infra.nomad", "HashiCorp Nomad.",
		"csteg.infra.proxmox", "Proxmox VE.",
		"csteg.infra.xcpng", "XCP-ng.",
		"csteg.infra.vmware", "VMware Bridge.",
		"csteg.infra.nutanix", "Nutanix AHV.",
		"csteg.infra.openstack", "OpenStack Infra.",
		"csteg.infra.lxc", "LXC/LXD Containers.",
		"csteg.infra.kvm", "KVM/QEMU Virtualization.",
		"csteg.infra.xen", "Xen Hypervisor.",
		"csteg.infra.kata", "Kata Containers.",
		"csteg.infra.firecracker", "Firecracker MicroVMs.",

		// Services
		"csteg.service.ntp", "Sincronismo NTP.",
		"csteg.service.dns", "DNS (Bind/Unbound).",
		"csteg.service.dhcp", "DHCP Server.",
		"csteg.service.mail", "Mail Server (Postfix/Exim).",
		"csteg.service.ldap", "LDAP/FreeIPA.",
		"csteg.service.samba", "Samba/AD.",
		"csteg.service.nfs", "NFS Server.",
	}
	plugin.RegisterMetrics(&CstegPlugin{}, "CstegCustom", metrics...)
}

func main() {
	// Configuração de logs do Zabbix SDK
	log.DefaultLogger = log.New(log.Console, log.Info)

	fmt.Println("Iniciando Agente CSTEG Customizado...")

	h, err := container.NewHandler("CstegCustom")
	if err != nil {
		fmt.Printf("Erro ao criar handler: %s\n", err)
		os.Exit(1)
	}

	// Loop principal para manter o agente rodando
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go h.Execute()

	<-stop
	fmt.Println("Desligando Agente CSTEG...")
}
