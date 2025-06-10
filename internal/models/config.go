package models

import "time"

type Config struct {
	NumFiles   int
	Iterations int
	Workers    int
	Timeout    time.Duration
}
