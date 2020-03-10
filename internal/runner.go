package internal

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

// CommandRunner is a struct that will handle calling exec.Cmd or syscall.Exec using a boolean
// it keeps all of the logic in a single interface to keep testing easier
type CommandRunner struct {
	path      string
	isSyscall bool
	reader    *strings.Reader
}

func (c CommandRunner) SetReader(rdr *strings.Reader) {
	c.reader = rdr
}

func (c CommandRunner) Run(args []string) error {
	// if this is a syscall, hand it off
	if c.isSyscall {
		// this allows us to only add args, not the binary path to
		// keep everything consistent
		if args[0] != "multipass" {
			args = append([]string{"multipass"}, args...)
		}

		return syscall.Exec(c.path, args, os.Environ())
	}

	cmd := exec.Command(c.path, args...)

	// if the reader is set, we need tp
	if c.reader != nil {
		cmd.Stdin = c.reader
	}

	if cmd.Stdout == nil {
		cmd.Stdout = os.Stdout
	}
	if cmd.Stderr == nil {
		cmd.Stderr = os.Stderr
	}

	return cmd.Run()
}

func (c *CommandRunner) UseSyscall(t bool) {
	c.isSyscall = t
}

func NewRunner(file string) Runner {
	path, err := exec.LookPath(file)
	if err != nil {
		log.Fatal("unable to find multipass")
	}

	return &CommandRunner{
		path: path,
	}
}
