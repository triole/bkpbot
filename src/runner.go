package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

// Exitcode contains the return of runcmdexitcode
type tExitcode struct {
	Cmd  string
	Out  string
	Err  string
	Code int
}

func runCmd(cmdstring string) (ex tExitcode) {

	// parse function input and create command set vars
	basecmd, args := makeCommandSet(cmdstring)
	cmd := exec.Command(basecmd, args...)

	// fetch environment, variables and everything
	env := os.Environ()
	cmd.Env = env

	// prepare returning the output
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// cmd.Start does not wait for command to complete, cmd.Wait does that later
	if err := cmd.Start(); err != nil {
		lg.Logf("Error at starting %q, %q", cmdstring, err.Error())
		ex.Code = 1
	}
	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			// The program has exited with an exit code != 0
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				ex.Code = status.ExitStatus()
			}
		} else {
			lg.Logf("Error at waiting for %q, %q", cmdstring, err.Error())
		}
	}

	// assign what was caught
	ex.Cmd = cmdstring
	ex.Out = stdout.String()
	ex.Err = stderr.String()

	return
}

func makeCommandSet(cmdstring string) (basecmd string, args []string) {
	arr := splitter(cmdstring)
	basecmd = arr[0]
	args = arr[1:]
	return
}

func splitter(s string) (r []string) {
	if strings.Contains(s, " ") == true {
		r = strings.Split(s, " ")
	} else {
		r = []string{s}
	}
	return
}
