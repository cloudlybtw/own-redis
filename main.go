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
)

func main() {
	port := flag.Int("port", 8080, "The port on which UDP protocol are transferred.")
	help := flag.Bool("help", false, "Displays help message.")
	// var m map[string]string
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
			conn.WriteToUDP([]byte("PONG\n"), addr)
			slog.Info("code 0. Connection check is succesful.")

		case preregexSet.MatchString(strings.ToUpper(text)):
			if regexSetPX.MatchString(strings.ToUpper(text)) {
				conn.WriteToUDP([]byte("OKPX\n"), addr)
				slog.Info("code 0. SET processed succesfully.")

			} else if regexSet.MatchString(strings.ToUpper(text)) {
				conn.WriteToUDP([]byte("OK\n"), addr)
				slog.Info("code 0. SET processed succesfully.")

			} else {
				conn.WriteToUDP([]byte("Wrong arguments: \"SET [key] [val]\"\n"), addr)
				slog.Info("code 6. SET not processed, wrong arguments.")
			}

		case preregexGet.MatchString(strings.ToUpper(text)):
			if regexGet.MatchString(strings.ToUpper(text)) {
				conn.WriteToUDP([]byte("GOT\n"), addr)
				slog.Info("code 0. GET processed succesfully.")

			} else {
				conn.WriteToUDP([]byte("Wrong arguments: \"GET [key]\"\n"), addr)
				slog.Info("code 6. GET not processed, wrong arguments.")
			}

		default:
			conn.WriteToUDP([]byte("Unknown command\n"), addr)
			slog.Info("code 8. Unsupported command")
		}
	}
}
