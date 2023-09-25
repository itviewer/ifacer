package utils

import (
    "fmt"
    "net"
    "strings"
)

// IpMaskToCIDR 格式类似 192.168.1.10/255.255.255.255
func IpMaskToCIDR(ipMask string) string {
    ips := strings.Split(ipMask, "/")
    length, _ := net.IPMask(net.ParseIP(ips[1]).To4()).Size()
    return fmt.Sprintf("%s/%v", ips[0], length)
}
