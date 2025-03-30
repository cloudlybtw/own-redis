package handler

import (
	"log/slog"
	"net"
	"strings"

	. "own-redis/pkg/core"
)

func ProcessGet(t string, addr *net.UDPAddr, conn *net.UDPConn) {
	slog.Info("code 0. GET requested.")
	if RegexGet.MatchString(strings.ToUpper(t)) {
		split := strings.Split(t, " ")
		value, _ := Data.Load(split[1])
		str, exists := value.(string)
		if exists {
			conn.WriteToUDP([]byte(str+"\n"), addr)
		} else {
			conn.WriteToUDP([]byte("(nil)\n"), addr)
		}
		slog.Info("code 0. GET key " + split[1] + " processed succesfully.")

	} else {
		conn.WriteToUDP([]byte("(error) ERR wrong number of arguments for 'GET' command\n"), addr)
		slog.Error("code 6. GET not processed, wrong arguments.")
	}
}
