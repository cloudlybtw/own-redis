package main

import (
	. "own-redis/pkg/core"
	. "own-redis/pkg/handler"
)

func main() {
	port := ParseFlags()
	conn := Initialize(port)
	HandleRequests(conn)
}
