package SIGINT

import (
	"os"
	"syscall"

	"github.com/songgao/stacktraces"
)

func init() {
	stacktraces.Set(os.Stderr, syscall.SIGINT)
}
