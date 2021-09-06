package ttype

import "time"

type PingResult struct {
	Domain    string
	PingDelay time.Duration
}
