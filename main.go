package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
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

	udpAddr, err := net.ResolveUDPAddr("udp", ":"+strconv.Itoa(*port))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Succesfull connection to port", *port)

	for {
		var buf [512]byte
		_, addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print("> ", string(buf[0:]))

		// Write back the message over UPD
		conn.WriteToUDP(buf[0:], addr)
	}
}
