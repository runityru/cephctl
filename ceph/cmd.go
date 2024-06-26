package ceph

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

func mkCommand(cephBinary string, args []string) (string, []string) {
	outCmd := shellCommand

	escapedArgs := []string{}
	for _, arg := range args {
		escapedArgs = append(escapedArgs, handleArg(arg))
	}

	outArgs := []string{shellArg, strings.Join(append([]string{cephBinary}, escapedArgs...), " ")}

	log.Debugf("preparing command: `%s` `%#v`", outCmd, outArgs)

	return shellCommand, outArgs
}

func handleArg(arg string) string {
	arg = strings.ReplaceAll(arg, `\`, `\\`)
	arg = strings.ReplaceAll(arg, "'", `\'`)

	return "'" + arg + "'"
}
