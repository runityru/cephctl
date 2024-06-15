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
		shellArg, "testcmd 'arg1' 'arg2'",
	}, args)
}

func TestHandleArg(t *testing.T) {
	type testCase struct {
		name   string
		in     string
		expOut string
	}

	tcs := []testCase{
		{
			name:   "simple string",
			in:     "simple string",
			expOut: "'simple string'",
		},
		{
			name:   "string with special characters",
			in:     `!@#$%^&*()_-+=\/:'"`,
			expOut: `'!@#$%^&*()_-+=\\/:\'"'`,
		},
		{
			name:   "single-quoted string",
			in:     `'quoted string'`,
			expOut: `'\'quoted string\''`,
		},
		{
			name:   "double-quoted string",
			in:     `"quoted string"`,
			expOut: `'"quoted string"'`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			v := handleArg(tc.in)
			r.Equal(tc.expOut, v)
		})
	}
}
