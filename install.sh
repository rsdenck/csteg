#!/bin/bash
# CSTEG Agent Installation Script (Pi-hole Style)
# Target: /opt/csteg/
# Theme: Cloudstage Green

# Colors
GREEN='\033[0;32m'
NC='\033[0m' # No Color
BOLD='\033[1m'

set -e

INSTALL_DIR="/opt/csteg"
BINARY_NAME="csteg-agent"

# ASCII ART (Cloudstage)
show_banner() {
    echo -e "${GREEN}${BOLD}"
    echo "          .::."
    echo "        .::::::."
    echo "   ... :::::::::: "
    echo "  :::::::::::::::::      _                 _     _"
    echo " :::::::::::::::::::  ____ _                 _     _                      "
    echo " :::::::::::::::::: / ___| | ___  _   _  __| |___| |_ __ _  __ _  ___  "
    echo "  ::::::::::::::::: | |   | |/ _ \| | | |/ _\` / __| __/ _\` |/ _\` |/ _ \\ "
    echo "   ''' ::::::::::   | |___| | (_) | |_| | (_| \__ \ || (_| | (_| |  __/ "
    echo "        '::::::'     \____|_|\___/ \__,_|\__,_|___/\__\__,_|\__, |\___| "
    echo "          '::'                                              |___/      "
    echo -e "${NC}"
}

clear
show_banner

echo -e "[${GREEN}✓${NC}] Iniciando Instalação do Agente CSTEG..."

# 1. Root Check
if [ "$EUID" -ne 0 ]; then 
  echo -e "[${GREEN}✗${NC}] Erro: Por favor, execute como root (sudo)"
  exit 1
fi
echo -e "[${GREEN}✓${NC}] Root user check"

# 2. Distro Detection
if [ -f /etc/os-release ]; then
    . /etc/os-release
    DISTRO=$ID
    echo -e "[${GREEN}✓${NC}] Distribuição detectada: ${BOLD}${NAME}${NC}"
else
    DISTRO="unknown"
    echo -e "[${GREEN}i${NC}] Distribuição não identificada claramente."
fi

# 3. Zabbix Server/Proxy IP Configuration
if [ -z "$ZABBIX_SERVER_IP" ]; then
    echo -ne "[${GREEN}?${NC}] Digite o IP do Zabbix Server ou Proxy para liberar no firewall: "
    read ZABBIX_SERVER_IP
fi

if [ -z "$ZABBIX_SERVER_IP" ]; then
    echo -e "[${GREEN}✗${NC}] Erro: IP do Zabbix Server/Proxy é obrigatório para as regras de firewall."
    exit 1
fi

# 4. Dependency Check (Go)
if ! command -v go &> /dev/null; then
    echo -e "[${GREEN}✗${NC}] Erro: Go (Golang) não encontrado. Necessário para compilação."
    exit 1
fi
echo -e "[${GREEN}✓${NC}] Go dependency check"

# 5. Preparar Diretórios
mkdir -p $INSTALL_DIR
mkdir -p $INSTALL_DIR/conf
mkdir -p $INSTALL_DIR/logs
echo -e "[${GREEN}✓${NC}] Diretórios criados em $INSTALL_DIR"

# 6. Compilar Agente
if [ -f "main.go" ]; then
    echo -e "[${GREEN}i${NC}] Compilando binário customizado com Zabbix SDK..."
    go build -o $BINARY_NAME main.go
    mv $BINARY_NAME $INSTALL_DIR/
    echo -e "[${GREEN}✓${NC}] Binário instalado em $INSTALL_DIR/$BINARY_NAME"
else
    echo -e "[${GREEN}✗${NC}] Erro: main.go não encontrado para compilação."
    exit 1
fi

# 7. Firewall Configuration (Exclusive Access)
echo -e "[${GREEN}i${NC}] Configurando firewall para porta 10050 (Acesso exclusivo: ${ZABBIX_SERVER_IP})..."

AGENT_PORT=10050

if command -v ufw &> /dev/null && ufw status | grep -q "active"; then
    echo -e "[${GREEN}i${NC}] Detectado UFW ativo. Aplicando regras..."
    ufw allow from $ZABBIX_SERVER_IP to any port $AGENT_PORT comment 'CSTEG Agent Access'
    ufw deny $AGENT_PORT
    echo -e "[${GREEN}✓${NC}] Regras UFW aplicadas."
elif command -v firewall-cmd &> /dev/null && systemctl is-active --quiet firewalld; then
    echo -e "[${GREEN}i${NC}] Detectado Firewalld ativo. Aplicando regras..."
    firewall-cmd --permanent --new-zone=csteg-zone || true
    firewall-cmd --permanent --zone=csteg-zone --add-source=$ZABBIX_SERVER_IP
    firewall-cmd --permanent --zone=csteg-zone --add-port=$AGENT_PORT/tcp
    firewall-cmd --reload
    echo -e "[${GREEN}✓${NC}] Regras Firewalld aplicadas (Zona csteg-zone)."
else
    echo -e "[${GREEN}i${NC}] Usando Iptables para configuração direta..."
    # Limpa regras anteriores para a porta
    iptables -D INPUT -p tcp --dport $AGENT_PORT -j ACCEPT 2>/dev/null || true
    iptables -D INPUT -p tcp --dport $AGENT_PORT -s $ZABBIX_SERVER_IP -j ACCEPT 2>/dev/null || true
    
    # Permite apenas do IP do Proxy/Server
    iptables -A INPUT -p tcp -s $ZABBIX_SERVER_IP --dport $AGENT_PORT -m comment --comment "CSTEG Agent Access" -j ACCEPT
    # Bloqueia o resto para essa porta
    iptables -A INPUT -p tcp --dport $AGENT_PORT -j DROP
    echo -e "[${GREEN}✓${NC}] Regras Iptables aplicadas."
fi

# 8. Configuração do Systemd
echo -e "[${GREEN}i${NC}] Configurando serviço Systemd (csteg-agent.service)..."
cat <<EOF > /etc/systemd/system/csteg-agent.service
[Unit]
Description=CSTEG Zabbix Agent (Cloudstage)
After=network.target

[Service]
Type=simple
ExecStart=$INSTALL_DIR/$BINARY_NAME
Restart=always
User=root
WorkingDirectory=$INSTALL_DIR

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
# systemctl enable csteg-agent &> /dev/null
# systemctl start csteg-agent &> /dev/null

echo ""
echo -e "${GREEN}${BOLD}--- Instalação Concluída com Sucesso! ---${NC}"
echo -e "O agente CSTEG agora monitora EXAUSTIVAMENTE todo o ecossistema:"
echo -e " - ${GREEN}Linux Core:${NC} Kernel, CPU, RAM, I/O, Entropy, Interrupts, Security (SSH/SELinux/Audit)"
echo -e " - ${GREEN}Databases:${NC} Oracle, Postgres, MySQL, MongoDB, Redis, Cassandra, ScyllaDB, Cockroach, Presto"
echo -e " - ${GREEN}Web/Proxies:${NC} Nginx, Apache, HAProxy, Traefik, Caddy, Envoy, Kong, Tyk, Istio"
echo -e " - ${GREEN}Middleware:${NC} Tomcat, JBoss, WebLogic, RabbitMQ, Kafka, NATS, Pulsar, ZeroMQ"
echo -e " - ${GREEN}Storage (SDS):${NC} Ceph, Minio, GlusterFS, OpenEBS, Longhorn, ZFS, DRBD, HDFS, SeaweedFS"
echo -e " - ${GREEN}Infra/Cloud:${NC} Docker, Podman, K8S, Nomad, Proxmox, VMware, OpenStack, KVM, Xen, Kata"
echo -e " - ${GREEN}Serviços:${NC} NTP, DNS, DHCP, LDAP, Samba, NFS, Mail (Postfix/Exim)"
echo ""
echo -e "[${GREEN}🔒${NC}] ${BOLD}Segurança de Rede:${NC} Porta 10050 restrita exclusivamente para o IP ${BOLD}${ZABBIX_SERVER_IP}${NC}."
echo ""
echo -e "Para iniciar o serviço: ${BOLD}systemctl start csteg-agent${NC}"
echo -e "Para verificar o status: ${BOLD}systemctl status csteg-agent${NC}"
