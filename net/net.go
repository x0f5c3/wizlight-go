package net

import (
	"fmt"
	"net"
	"net/netip"
	"syscall"
)

const PORT = 38999

func CreateUDP() (*net.UDPConn, error) {
	laddr := net.UDPAddrFromAddrPort(netip.AddrPortFrom(netip.IPv4Unspecified(), PORT))
	sock, err := net.ListenUDP("udp", laddr)
	if err != nil {
		return nil, err
	}
	sysSock, err := sock.SyscallConn()
	if err != nil {
		return nil, err
	}
	err = sysSock.Control(func(fd uintptr) {
		err := syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
		if err != nil {
			fmt.Println("Failed to setsockopt due to " + err.Error())
			return
		}
		err = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_BROADCAST, 1)
		if err != nil {
			fmt.Println("Failed to setsockopt due to " + err.Error())
			return
		}
	})
	return sock, err
}
