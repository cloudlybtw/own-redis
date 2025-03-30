package core

import (
	"regexp"
	"sync"
)

var (
	RegexPing   = regexp.MustCompile(`^PING(\s+.+)?$`)
	PreregexSet = regexp.MustCompile(`^SET(\s+.+)?$`)
	RegexSet    = regexp.MustCompile(`^SET [\w]* +[\w\s]*$`)
	PreregexGet = regexp.MustCompile(`^GET(\s+.+)?$`)
	RegexGet    = regexp.MustCompile(`^GET [\w]*$`)
	RegexSetPX  = regexp.MustCompile(`^SET [\w]* [\w\s]* PX \d+$`)
	Data        sync.Map
)
