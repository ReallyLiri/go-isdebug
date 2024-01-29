package isdebug

import (
	"github.com/mitchellh/go-ps"
	"os"
)

func IsDebugging() bool {
	pid := os.Getppid()
	for pid != 0 {
		switch p, err := ps.FindProcess(pid); {
		case err != nil:
			return false
		case p.Executable() == "dlv", p.Executable() == "debugserver":
			return true
		default:
			pid = p.PPid()
		}
	}
	return false
}
