package core

import "fmt"

func PrintHelp() {
	fmt.Println("Own Redis\n\nUsage:\n  own-redis [--port <N>]\n  own-redis --help\n\nOptions:\n  --help       Show this screen.\n  --port N     Port number (1-65535).")
}
