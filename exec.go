package utils

import (
	"os/exec"
	"syscall"
)

func ExecWithOutput(dir string, name string, hideWindow bool, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	if dir != "" {
		cmd.Dir = dir
	}

	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: hideWindow}
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func ExecStart(dir string, name string, hideWindow bool, arg ...string) error {
	cmd := exec.Command(name, arg...)
	if dir != "" {
		cmd.Dir = dir
	}

	//
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: hideWindow}
	err := cmd.Start()
	return err
}
