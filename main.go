package main

import (
	"flag"
	"fmt"
	"os"
)

func printHelp() {
	fmt.Println("Own Redis\n\nUsage:\n  own-redis [--port <N>]\n  own-redis --help\n\nOptions:\n  --help       Show this screen.\n  --port N     Port number.")
}

func main() {
	port := flag.Int("port", 8080, "The port on which UDP protocol are transferred.")
	help := flag.Bool("help", false, "Displays help message.")
	flag.Parse()
	if *help {
		printHelp()
		os.Exit(0)
	}
	fmt.Println(*port)
}
