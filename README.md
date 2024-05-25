# cephctl

[![Verify](https://github.com/teran/cephctl/actions/workflows/verify.yml/badge.svg?branch=master)](https://github.com/teran/cephctl/actions/workflows/verify.yml)

Small utility to control Ceph cluster configuration just like any other declarative
    configuration

## Main features

* Easy-to-use healthcheck which may contain checks against status & configuration
    and indicate some some not trivial issues
* Declarative configuration support which is apply only if needed
* Diff configuration: check what the difference between currently running configuration
    and desired or migrated from other cluster

## Usage

<!-- markdownlint-disable MD013 -->
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
<!-- markdownlint-enable MD013 -->

## How it works

Cephctl uses native Ceph CLIs to work with cluster configuration so it's require
to have Ceph binaries w/ configured `ceph.conf`. Alternatively it's possible
to adjust `ceph` binary path to access ceph in container and/or remote machine.

## Roadmap

* [X] Apply declarative configuration for `ceph config`
* [X] Dump cluster configuration to CephConfig specification
* [X] Diff configuration against running configuration for `ceph config`
* [X] Perform healthcheck based on current cluster status
* [X] Add healthchecks based on current cluster configuration
* [ ] Apply/Dump declarative configuration for `ceph osd set-*` stuff
* [ ] Apply/Dump declarative configuration for Ceph Object Gateway (rgw)
* [ ] Apply/Dump declarative configuration for Pools

## Ceph compatibility

All of the changes are tested against Ceph 18.2 (Reef), previous versions are
not tested and not guaranteed to work.

## Interface compatibility disclaimer

If you gonna use cephctl as a library for your purposes please feel free to
but please note a few things:

1. Cephctl doesn't use internal packages to allow you to do whatever you like.
    Cephctl project doesn't aim to limit your usage.
2. Internal program interfaces are not guaranteed to be stable between releases
    since they're written and serve for internal purposes.
3. CLI interface (until 1.0.x at least) is also not guaranteed to be stable:
    subcommands and options are subjects to change between versions.

## Installation

cephctl is released in following ways to achieve compatibility and provide
an easy way for end users.

### Pre-compiled binary

Pre-compiled binaries are available on per-release basis and provided on
[GitHub Releases page](https://github.com/teran/cephctl/releases). Automatically
generated changelog is available for each release. And binaries are available for:

* FreeBSD (amd64v1, amd64v2, amd64v3, arm64)
* Linux (amd64v1, amd64v2, amd64v3, arm64)
* macOS (amd64v1, amd64v2, amd64v3, arm64)
* Windows (amd64v1, amd64v2, amd64v3, arm64)

Any of them could be used on end-user machine to interact with Ceph
via SSH just like the following way:

```shell
cephctl --ceph-binary='ssh mon01 ceph' healthcheck
```

or by using environment variables to specify ceph binary:

```shell
export CEPHCTL_CEPH_BINARY='ssh mon01 ceph'
cephctl healthcheck
```

### Container image

Since cephctl uses ceph binary to achieve cluster data, container image based
on ceph official release image is also available. This image is designed as
drop-in replacement for official ceph image to use for `cephadm shell` command.

Container image is available at [GitHub Packages](https://github.com/teran/cephctl/pkgs/container/cephctl%2Fceph)

### Build from source

It's possible to build cephctl from source by simply running the following
command:

<!-- markdownlint-disable MD013 -->
```shell
go build -v -ldflags="-X 'main.appVersion=$(git rev-parse --short HEAD) (trunk build)' -X 'main.buildTimestamp=$(date -u +%Y-%m-%dT%H:%m:%SZ)'" -o dist/cephctl ./cmd/cephctl/...
```
<!-- markdownlint-enable MD013 -->
