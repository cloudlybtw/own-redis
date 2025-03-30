package handler

import (
	"log/slog"
	"net"
	. "own-redis/pkg/core"
	"strings"
)

func HandleRequests(conn *net.UDPConn) {
	for {
		var buf [512]byte
		_, addr, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			slog.Error(err.Error())
			return
		}
		text := string(buf[0:])
		text = strings.Split(text, "\n")[0]
		text = strings.TrimFunc(text, func(c rune) bool {
			return c == '\n' || c == ' '
		})

		switch {
		case RegexPing.MatchString(strings.ToUpper(text)):
			go ProcessPing(addr, conn)
		case PreregexSet.MatchString(strings.ToUpper(text)):
			go ProcessSet(text, addr, conn)
		case PreregexGet.MatchString(strings.ToUpper(text)):
			go ProcessGet(text, addr, conn)
		default:
			conn.WriteToUDP([]byte("(error) ERR unknown command\n"), addr)
			slog.Error("code 8. Unsupported command requested.")
		}
	}
}
