package ceph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMkCommand(t *testing.T) {
	r := require.New(t)

	bin, args := mkCommand("testcmd", []string{"arg1", "arg2"})
	r.Equal(shellCommand, bin)
	r.Equal([]string{
		shellArg, "testcmd \"arg1\" \"arg2\"",
	}, args)
}
