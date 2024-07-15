package apputil

import (
	"bytes"
	"fmt"
	"gitee.com/dk83/goutils/dlog"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

func cmdRunWithTimeout(cmd *exec.Cmd, timeout time.Duration) (error, bool) {
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	var err error
	select {
	case <-time.After(timeout):
		// timeout
		if err = cmd.Process.Kill(); err != nil {
			dlog.Error("failed to kill: %s, error: %s", cmd.Path, err)
		}
		go func() {
			<-done // allow goroutine to exit
		}()
		dlog.Error("process:%s killed", cmd.Path)
		return err, true
	case err = <-done:
		return err, false
	}
}

func RunCmd(exePath string, cmdOut **exec.Cmd, args ...string) (err error, exitCode int, outStr string) {
	var cmd *exec.Cmd
	dlog.Debug("RunTool:%s %s", exePath, strings.Join(args, " "))
	if cmdOut != nil {
		(*cmdOut).Path = exePath
		(*cmdOut).Args = append([]string{exePath}, args...)

		cmd = *cmdOut
	} else {
		cmd = exec.Command(exePath, args...)
	}

	var out bytes.Buffer
	var errout bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errout
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := cmd.Start(); err != nil {
		dlog.Error(err, -1)
		return err, -1, err.Error()
	}

	var isTimeout bool
	err, isTimeout = cmdRunWithTimeout(cmd, 3600*time.Minute)
	if isTimeout {
		exitCode = -1
		return fmt.Errorf("time out"), exitCode, "time out"
	}
	if err != nil {
		//glog.Error(err)
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		} else {
			// This will happen (in OSX) if `name` is not available in $PATH,
			// in this situation, exit code could not be get, and stderr will be
			// empty string very likely, so we use the default fail code, and format err
			// to string and set to stderr
			//log.Printf("Could not get exit code for failed program: %v, %v", name, args)
			exitCode = -1
		}
		outStr = out.String()
		return err, exitCode, errout.String()
	} else {
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()
	}

	outStr = out.String()
	return nil, exitCode, outStr
}
