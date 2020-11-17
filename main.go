package main

import (
	"os"
	"time"

	"github.com/pannoi/wiel/package/pannoi"
)

const version = "v0.0.1"

var (
	config    pannoi.Config
	gitCommit string
	startTime int64
)

func main() {
	startTime = time.Now().Unix()
}

// ReadEnvOrPanic is a helper that returns an env var or panics if not set
func ReadEnvOrPanic(name string) string {
	val := os.Getenv(name)
	if val == "" {
		panic("Env var " + name + " must be set")
	}
	return val
}
