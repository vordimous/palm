package git

import (
	"fmt"
	"os"
	"os/exec"
)

func Run(args *[]string) string {
	var (
		cmdOut []byte
		err    error
	)
	if cmdOut, err = exec.Command("git", *args...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git rev-parse command: ", err)
		os.Exit(1)
	}
	return string(cmdOut)
}
