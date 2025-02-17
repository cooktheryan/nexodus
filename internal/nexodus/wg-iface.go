package nexodus

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net"
)

var interfaceErr = errors.New("interface setup error")

// ifaceExists returns true if the input matches a net interface
func ifaceExists(logger *zap.SugaredLogger, iface string) bool {
	_, err := net.InterfaceByName(iface)
	if err != nil {
		logger.Debugf("existing link not found: %s", iface)
		return false
	}

	return true
}

// getIPv4Iface get the IP of the specified net interface
func getIPv4Iface(ifname string) net.IP {
	interfaces, _ := net.Interfaces()
	for _, inter := range interfaces {
		if inter.Name != ifname {
			continue
		}
		addrs, err := inter.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			switch ip := addr.(type) {
			case *net.IPNet:
				if ip.IP.DefaultMask() != nil {
					return ip.IP
				}
			}
		}
	}

	return nil
}

// enableForwardingIPv4 for linux nodes that are hub bouncers
func enableForwardingIPv4(logger *zap.SugaredLogger) error {
	cmdOut, err := RunCommand("sysctl", "-w", "net.ipv4.ip_forward=1")
	if err != nil {
		return fmt.Errorf("failed to enable IP Forwarding for this hub-router: %w", err)
	}
	logger.Debugf("%v", cmdOut)
	return nil
}
