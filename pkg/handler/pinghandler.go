package handler

import (
	"log/slog"
	"net"
)

func ProcessPing(addr *net.UDPAddr, conn *net.UDPConn) {
	slog.Info("code 0. PING requested.")
	conn.WriteToUDP([]byte("PONG\n"), addr)
	slog.Info("code 0. Connection check is succesful.")
}
