package ceph

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func mkCommand(cephBinary string, args []string) (string, []string) {
	outCmd := shellCommand

	escapedArgs := []string{}
	for _, arg := range args {
		escapedArgs = append(escapedArgs, strconv.Quote(arg))
	}

	outArgs := []string{"-c", strings.Join(append([]string{cephBinary}, escapedArgs...), " ")}

	log.Tracef("preparing command: `%s` `%#v`", outCmd, outArgs)

	return shellCommand, outArgs
}
