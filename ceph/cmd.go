package ceph

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

func mkCommand(cephBinary string, args []string) (string, []string) {
	outCmd := shellCommand
	outArgs := []string{"-c", strings.Join(append([]string{cephBinary}, args...), " ")}

	log.Tracef("preparing command: `%s` `%#v`", outCmd, outArgs)

	return shellCommand, outArgs
}
