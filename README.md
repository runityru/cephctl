# cephctl

[![Go](https://github.com/teran/cephctl/actions/workflows/go.yml/badge.svg)](https://github.com/teran/cephctl/actions/workflows/go.yml)

Small utility to control Ceph cluster configuration just like any other declarative configuration

## Usage

```shell
$ ./cephctl
usage: cephctl [<flags>] <command> [<args> ...]

Small utility to control Ceph cluster configuration just like any other declarative configuration


Flags:
      --[no-]help   Show context-sensitive help (also try --help-long and --help-man).
  -b, --ceph-binary="/usr/bin/ceph"
                    Specify path to ceph binary ($CEPHCTL_CEPH_BINARY)
  -t, --[no-]trace  Enable trace mode ($CEPHCTL_TRACE)

Commands:
help [<command>...]
    Show help.

apply <filename>
    Apply ceph configuration

diff [<flags>] <filename>
    Show difference between running and desired configurations

dump cephconfig
    dump Ceph runtime configuration

healthcheck
    Perform a cluster healthcheck and print report

version
    Print version and exit


```

## Compatibility

All of the changes are tested against Ceph 18.2 (Reef), previous versions are
not tested and not guaranteed to work.
