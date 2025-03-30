package handler

import (
	"log/slog"
	"net"
	"strconv"
	"strings"
	"time"

	. "own-redis/pkg/core"
)

func ProcessSet(t string, addr *net.UDPAddr, conn *net.UDPConn) {
	slog.Info("code 0. SET requested.")
	if RegexSetPX.MatchString(strings.ToUpper(t)) {
		split := strings.Split(t, " ")
		Data.Store(split[1], strings.Join(split[2:len(split)-2], " "))
		t, err := strconv.Atoi(split[len(split)-1])
		if err != nil || t < 0 {
			conn.WriteToUDP([]byte("Wrong time set, not int.\n"), addr)
			slog.Info("code 6. SET not processed, wrong argument for time.")
		}
		go func() {
			time.Sleep(time.Duration(t) * time.Millisecond)
			Data.Delete(split[1])
			slog.Info("code 0. SET Key " + split[1] + " was deleted succesfully.")
		}()
		conn.WriteToUDP([]byte("OK\n"), addr)
		slog.Info("code 0. SET processed succesfully.")

	} else if RegexSet.MatchString(strings.ToUpper(t)) {
		split := strings.Split(t, " ")
		(&Data).Store(split[1], strings.Join(split[2:], " "))
		conn.WriteToUDP([]byte("OK\n"), addr)
		slog.Info("code 0. SET key " + split[1] + " succesfully.")
	} else {
		conn.WriteToUDP([]byte("(error) ERR wrong number of arguments for 'SET' command\n"), addr)
		slog.Error("code 6. SET not processed, wrong arguments.")
	}
}
