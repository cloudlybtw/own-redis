package core

import (
	"flag"
	"log/slog"
	"os"
)

func ParseFlags() int {
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
	return *port
}
