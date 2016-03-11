node "ultramaster" {}

node /dockeragent/ {
    class {'echo_server':
        echo_message => "${hostname} : ${ipaddress_eth0}"
    }
}
