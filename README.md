# Agente CSTEG Custom (Cloudstage) 🛡️

O **Agente CSTEG** é um binário customizado desenvolvido em Go, utilizando o SDK oficial do Zabbix, projetado para monitoramento exaustivo de infraestruturas modernas. Ele combina a simplicidade de instalação do estilo Pi-hole com a robustez necessária para monitorar desde o Kernel Linux até clusters de Kubernetes e bancos de dados Oracle.

---

## 🚀 Instalação (Estilo Pi-hole)

Para instalar o agente no seu servidor Linux, execute o comando abaixo (simulando o deploy via curl):

```bash
curl -sSL https://install.cloudstage.net/csteg | sudo bash
```

> **Nota:** O instalador irá configurar automaticamente o diretório `/opt/csteg/`, criar o serviço no `systemd` e configurar as regras de firewall para acesso exclusivo.

### O que o instalador faz:
1.  **Banner Cloudstage:** Exibe o ASCII Art verde neon da Cloudstage.
2.  **Verificação de Root:** Garante privilégios administrativos.
3.  **Compilação Customizada:** Compila o código Go especificamente para a sua arquitetura.
4.  **Segurança de Rede:** Solicita o IP do seu Zabbix Server/Proxy e libera a porta `10050` apenas para esse IP, bloqueando todo o resto.
5.  **Persistência:** Instala o serviço `csteg-agent.service` para iniciar com o sistema.

---

## 📊 Métricas Coletadas

O agente suporta as seguintes categorias de métricas (Chaves Zabbix):

### 🐧 Linux Core & Segurança (Exaustivo)
| Categoria | Chaves (Exemplos) |
| :--- | :--- |
| **Core** | `csteg.linux.kernel`, `csteg.linux.cpu`, `csteg.linux.memory`, `csteg.linux.disk.io`, `csteg.linux.disk.usage`, `csteg.linux.entropy`, `csteg.linux.interrupts` |
| **Segurança** | `csteg.linux.security.ssh`, `csteg.linux.security.selinux`, `csteg.linux.security.apparmor`, `csteg.linux.security.firewall`, `csteg.linux.security.audit`, `csteg.linux.security.fail2ban` |
| **Rede & Sistema** | `csteg.linux.network.traffic`, `csteg.linux.network.sockets`, `csteg.linux.systemd`, `csteg.linux.cron`, `csteg.linux.conntrack` |

### 🗄️ Bancos de Dados (SQL & NoSQL)
| Tipo | Chaves (Exemplos) |
| :--- | :--- |
| **SQL** | `csteg.database.oracle`, `csteg.database.postgres`, `csteg.database.mysql`, `csteg.database.sqlserver`, `csteg.database.cockroachdb`, `csteg.database.tidb`, `csteg.database.vitess` |
| **NoSQL** | `csteg.database.mongodb`, `csteg.database.cassandra`, `csteg.database.scylladb`, `csteg.database.elasticsearch`, `csteg.database.solr`, `csteg.database.couchdb`, `csteg.database.neo4j`, `csteg.database.arangodb` |
| **TSDB / Analytics** | `csteg.database.clickhouse`, `csteg.database.influxdb`, `csteg.database.timescaledb`, `csteg.database.redis`, `csteg.database.presto` |

### 🌐 Web Servers & Reverse Proxies (Todos)
| Categoria | Chaves (Exemplos) |
| :--- | :--- |
| **Servers & Proxies** | `csteg.web.nginx`, `csteg.web.apache`, `csteg.web.haproxy`, `csteg.web.varnish`, `csteg.web.traefik`, `csteg.web.caddy`, `csteg.web.envoy`, `csteg.web.litespeed`, `csteg.web.squid`, `csteg.web.kong`, `csteg.web.tyk`, `csteg.web.istio`, `csteg.web.ats` |

### ☕ Middleware & Messaging (Todos)
| Categoria | Chaves (Exemplos) |
| :--- | :--- |
| **App Servers** | `csteg.middleware.tomcat`, `csteg.middleware.jboss`, `csteg.middleware.weblogic`, `csteg.middleware.websphere`, `csteg.middleware.glassfish`, `csteg.middleware.jetty` |
| **Messaging** | `csteg.middleware.rabbitmq`, `csteg.middleware.kafka`, `csteg.middleware.activemq`, `csteg.middleware.nats`, `csteg.middleware.pulsar`, `csteg.middleware.mqtt`, `csteg.middleware.zeromq`, `csteg.middleware.emqx` |

### 📦 Software-Defined Storage (Todos)
| Categoria | Chaves (Exemplos) |
| :--- | :--- |
| **SDS & Object** | `csteg.storage.ceph`, `csteg.storage.minio`, `csteg.storage.glusterfs`, `csteg.storage.openebs`, `csteg.storage.longhorn`, `csteg.storage.portworx`, `csteg.storage.rook`, `csteg.storage.moosefs`, `csteg.storage.lustre`, `csteg.storage.beegfs`, `csteg.storage.swift`, `csteg.storage.zfs`, `csteg.storage.drbd`, `csteg.storage.hdfs`, `csteg.storage.seaweedfs`, `csteg.storage.juicefs` |

### ☁️ Infraestrutura & Orquestração (Todos)
| Categoria | Chaves (Exemplos) |
| :--- | :--- |
| **Containers** | `csteg.infra.docker`, `csteg.infra.podman`, `csteg.infra.containerd`, `csteg.infra.lxc`, `csteg.infra.kata` |
| **Orquestração** | `csteg.infra.kubernetes`, `csteg.infra.k3s`, `csteg.infra.openshift`, `csteg.infra.nomad` |
| **Virtualização** | `csteg.infra.proxmox`, `csteg.infra.xcpng`, `csteg.infra.vmware`, `csteg.infra.nutanix`, `csteg.infra.openstack`, `csteg.infra.kvm`, `csteg.infra.xen`, `csteg.infra.firecracker` |

### 🛠️ Serviços de Rede
| Serviço | Chaves (Exemplos) |
| :--- | :--- |
| **Core Services** | `csteg.service.ntp`, `csteg.service.dns`, `csteg.service.dhcp`, `csteg.service.mail`, `csteg.service.ldap`, `csteg.service.samba`, `csteg.service.nfs` |

---

## 🛠️ Uso e Comandos Úteis

Após a instalação, você pode gerenciar o agente com os comandos padrão do Linux:

**Iniciar o Agente:**
```bash
sudo systemctl start csteg-agent
```

**Verificar Status:**
```bash
sudo systemctl status csteg-agent
```

**Ver Logs:**
```bash
journalctl -u csteg-agent -f
```

**Testar uma Métrica Localmente:**
```bash
# Exemplo testando métrica de Docker
zabbix_get -s 127.0.0.1 -k csteg.infra.docker
```

---

## 🔒 Segurança

O Agente CSTEG foi construído com foco em **Zero Trust**:
- **Acesso Restrito:** A porta `10050` não fica aberta para a internet; apenas o IP do seu Proxy Zabbix configurado no momento da instalação pode consultá-la.
- **Isolamento:** Roda como um serviço dedicado em `/opt/csteg/`.

---
Desenvolvido por **Cloudstage Monitoring Solutions**.
