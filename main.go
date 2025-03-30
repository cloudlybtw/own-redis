package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

func PrintHelp() {
	fmt.Println("Own Redis\n\nUsage:\n  own-redis [--port <N>]\n  own-redis --help\n\nOptions:\n  --help       Show this screen.\n  --port N     Port number (1-65535).")
}

var (
	regexPing   = regexp.MustCompile(`^PING(\s+.+)?$`)
	preregexSet = regexp.MustCompile(`^SET(\s+.+)?$`)
	regexSet    = regexp.MustCompile(`^SET [\w]* +[\w\s]*$`)
	preregexGet = regexp.MustCompile(`^GET(\s+.+)?$`)
	regexGet    = regexp.MustCompile(`^GET [\w]*$`)
	regexSetPX  = regexp.MustCompile(`^SET [\w]* [\w\s]* PX \d+$`)
	Data        sync.Map
)

func main() {
	port := flag.Int("port", 8080, "The port on which UDP protocol are transferred.")
	help := flag.Bool("help", false, "Displays help message.")

	flag.Parse()
	if *help {
		PrintHelp()
		os.Exit(0)
	}

	if !(*port >= 1 && *port <= 65535) {
		slog.Error("code 6. Port number is not in range (1-65535).")
		os.Exit(1)
	}

	udpAddr, err := net.ResolveUDPAddr("udp", ":"+strconv.Itoa(*port))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("code 0. Succesfull connection to port " + strconv.Itoa(*port))

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
		case regexPing.MatchString(strings.ToUpper(text)):
			slog.Info("code 0. PING requested.")
			conn.WriteToUDP([]byte("PONG\n"), addr)
			slog.Info("code 0. Connection check is succesful.")

		case preregexSet.MatchString(strings.ToUpper(text)):
			ProcessSet(text, addr, conn)
		case preregexGet.MatchString(strings.ToUpper(text)):
			ProcessGet(text, addr, conn)
		default:
			conn.WriteToUDP([]byte("(error) ERR unknown command\n"), addr)
			slog.Error("code 8. Unsupported command requested.")
		}
	}
}

func ProcessSet(t string, addr *net.UDPAddr, conn *net.UDPConn) {
	slog.Info("code 0. SET requested.")
	if regexSetPX.MatchString(strings.ToUpper(t)) {
		split := strings.Split(t, " ")
		Data.Store(split[1], strings.Join(split[2:len(split)-2], " "))
		t, err := strconv.Atoi(split[len(split)-1])
		if err != nil {
			conn.WriteToUDP([]byte("Wrong time set, not int.\n"), addr)
			slog.Info("code 6. SET not processed, wrong argument for time.")
		}
		go func() {
			time.Sleep(time.Duration(t) * time.Millisecond)
			Data.Delete(split[1])
			slog.Info("code 0. Key " + split[1] + " was deleted succesfully.")
		}()
		conn.WriteToUDP([]byte("OK\n"), addr)
		slog.Info("code 0. SET processed succesfully.")

	} else if regexSet.MatchString(strings.ToUpper(t)) {
		split := strings.Split(t, " ")
		(&Data).Store(split[1], strings.Join(split[2:], " "))
		conn.WriteToUDP([]byte("OK\n"), addr)
		slog.Info("code 0. SET key " + split[1] + " succesfully.")
	} else {
		conn.WriteToUDP([]byte("(error) ERR wrong number of arguments for 'SET' command\n"), addr)
		slog.Error("code 6. SET not processed, wrong arguments.")
	}
}

func ProcessGet(t string, addr *net.UDPAddr, conn *net.UDPConn) {
	slog.Info("code 0. GET requested.")
	if regexGet.MatchString(strings.ToUpper(t)) {
		split := strings.Split(t, " ")
		value, _ := (&Data).Load(split[1])
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
