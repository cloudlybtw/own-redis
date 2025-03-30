package core

import (
	"log/slog"
	"net"
	"os"
	"strconv"
)

func Initialize(port int) *net.UDPConn {
	udpAddr, err := net.ResolveUDPAddr("udp", ":"+strconv.Itoa(port))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("code 0. Succesfull connection to port " + strconv.Itoa(port))
	return conn
}
