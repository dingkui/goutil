package cmdutil

import (
	"bytes"
	"fmt"
	"gitee.com/dk83/goutils/zlog"
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
			zlog.Error("failed to kill: %s, error: %s", cmd.Path, err)
		}
		go func() {
			<-done // allow goroutine to exit
		}()
		zlog.Error("process:%s killed", cmd.Path)
		return err, true
	case err = <-done:
		return err, false
	}
}

func RunTool(exePath string, cmdOut **exec.Cmd, args ...string) (err error, exitCode int, outStr string) {
	var cmd *exec.Cmd
	zlog.Debug("RunTool:%s %s", exePath, strings.Join(args, " "))
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
		zlog.Error(err, -1)
		return err, -1, err.Error()
	}

	var isTimeout bool
	err, isTimeout = cmdRunWithTimeout(cmd, time.Duration(3600*time.Minute))
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

func ConvertVideo(ffempgPath, src, dst string) {
	//ffmpeg  -y -i  DFUD0404.mp4 -vcodec h264 -loglevel quiet    -vf scale='if(gt(a,320/500),320,-1)':'if(gt(a,320/500),-1,320)':force_original_aspect_ratio=decrease,pad='iw+mod(iw\,2)':'ih+mod(ih\,2)'  DFUD04042.mp4
	cmdline := make([]string, 0)
	cmdline = append(cmdline, "-y")
	cmdline = append(cmdline, "-i", src)
	cmdline = append(cmdline, "-vcodec", "h264", "-loglevel", "quiet", "-vf", "scale='if(gt(a,320/500),320,-1)':'if(gt(a,320/500),-1,320)':force_original_aspect_ratio=decrease,pad='iw+mod(iw\\,2)':'ih+mod(ih\\,2)'")
	cmdline = append(cmdline, dst)
	//zlog.Debug(strings.Join(cmdline," "))
	RunTool(ffempgPath, nil, cmdline...)

}

func GenVideoThumb(ffempgPath, src, dst string) {
	//ffmpeg -i C:\Users\DELL\Music\gg.mp4 -y -f image2 -t 1 -frames:v 1 3.jpg
	cmdline := make([]string, 0)
	cmdline = append(cmdline, "-y")
	cmdline = append(cmdline, "-i", src)
	cmdline = append(cmdline, "-f", "image2", "-t", "1", "-frames:v", "1")
	cmdline = append(cmdline, dst)

	//zlog.Debug(strings.Join(cmdline," "))
	RunTool(ffempgPath, nil, cmdline...)
}

//文件格式转换
func ConvertFileType(ffempgPath, src, dst string) {
	cmdline := make([]string, 0)
	cmdline = append(cmdline, "-i", src)
	//cmdline = append(cmdline, "-codec:a")
	//cmdline = append(cmdline, "libmp3lame")
	//cmdline = append(cmdline, "-b:a")
	//cmdline = append(cmdline, "64k")
	cmdline = append(cmdline, dst)
	RunTool(ffempgPath, nil, cmdline...)
}
