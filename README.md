# cephctl

[![Verify](https://github.com/runityru/cephctl/actions/workflows/verify.yml/badge.svg?branch=master)](https://github.com/runityru/cephctl/actions/workflows/verify.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/runityru/cephctl)](https://goreportcard.com/report/github.com/runityru/cephctl)
[![Go Reference](https://pkg.go.dev/badge/github.com/runityru/cephctl.svg)](https://pkg.go.dev/github.com/runityru/cephctl)

Small utility to control Ceph cluster configuration just like any other declarative
configuration

## Main features

- Easy-to-use healthcheck which may contain checks against status & configuration
  and indicate some some not trivial issues
- Declarative configuration support which is apply only if needed
- Diff configuration: check what the difference between currently running configuration
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
  -d, --[no-]debug  Enable debug mode ($CEPHCTL_DEBUG)
  -t, --[no-]trace  Enable trace mode (debug mode on steroids) ($CEPHCTL_TRACE)
  -c, --[no-]color  Colorize diff output ($CEPHCTL_COLOR)

Commands:
help [<command>...]
    Show help.

apply <filename>
    Apply ceph configuration

diff <filename>
    Show difference between running and desired configurations

dump cephconfig
    dump Ceph runtime configuration

dump cephosdconfig
    dump Ceph OSD configuration

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

## Declarative configuration format

cephctl uses a declarative YAML specification to describe the desired Ceph cluster
configuration. The file can contain one or more configuration documents separated
by `---`.

Each document has two mandatory fields:

- `kind` — defines the type of configuration section (see below)
- `spec` — the actual configuration payload

### kind: CephConfig

Maps to `ceph config set` commands. The spec is a nested map of configuration
sections and their key-value pairs:

```yaml
kind: CephConfig
spec:
  <section>:
    <key>: "<value>"
```

- **section** — any valid Ceph configuration section (e.g., `global`, `mon`,
  `osd`, `client.radosgw`, `mgr`, etc.)
- **key** — any valid Ceph configuration parameter within that section
- **value** — the value as a string (YAML strings, quoted or unquoted)

Example:

```yaml
kind: CephConfig
spec:
  global:
    rbd_cache: "true"
    osd_pool_default_size: 3
  client.radosgw:
    rgw_cache_lru_size: 100
    rgw_crypt_require_ssl: "true"
  osd:
    rocksdb_perf: "true"
```

### kind: CephOSDConfig

Maps to `ceph osd set-*` commands. The spec defines OSD-level operational
parameters:

```yaml
kind: CephOSDConfig
spec:
  allow_crimson: <bool>
  backfillfull_ratio: <float>
  full_ratio: <float>
  nearfull_ratio: <float>
  require_min_compat_client: <string>
```

Available fields and their defaults:

| Field                       | Type   | Default | Description                                                                                              |
| --------------------------- | ------ | ------- | -------------------------------------------------------------------------------------------------------- |
| `allow_crimson`             | bool   | `false` | Enable experimental Crimson OSD backend (risky for production)                                           |
| `backfillfull_ratio`        | float  | `0.9`   | OSD considered full for backfill purposes at this ratio                                                  |
| `full_ratio`                | float  | `0.95`  | OSD considered full and blocks writes at this ratio                                                      |
| `nearfull_ratio`            | float  | `0.85`  | OSD considered near-full at this ratio                                                                   |
| `require_min_compat_client` | string | `reef`  | Minimum allowed client version (`luminous`, `nautilus`, `octopus`, `pacific`, `quincy`, `reef`, `squid`) |

### Multi-document example

Both kinds can be combined in a single file using the `---` separator:

```yaml
---
kind: CephConfig
spec:
  global:
    rbd_cache: "true"
  client.radosgw:
    rgw_cache_lru_size: 100
---
kind: CephOSDConfig
spec:
  allow_crimson: false
  backfillfull_ratio: 0.9
  full_ratio: 0.95
  nearfull_ratio: 0.85
  require_min_compat_client: reef
```

## Health checks

`cephctl healthcheck` runs the following checks and reports the status for each
indicator. Each check is reported as **GOOD**, **AT_RISK**, **DANGEROUS**, or
**UNKNOWN**.

| Indicator                 | Type             | GOOD               | AT_RISK           | DANGEROUS        | Description                                                  |
| ------------------------- | ---------------- | ------------------ | ----------------- | ---------------- | ------------------------------------------------------------ |
| `CLUSTER_STATUS`          | overall          | `HEALTH_OK`        | `HEALTH_WARN`     | `HEALTH_ERR`     | Overall cluster health status from `ceph status`             |
| `QUORUM`                  | monitors         | all mons in quorum | some mons missing | —                | Whether all monitor nodes are participating in quorum        |
| `MON_DOWN`                | monitors         | 0                  | >0                | —                | Count of monitor nodes that are not up                       |
| `OSD_DOWN`                | OSD              | 0                  | >0                | —                | Count of OSDs in down state                                  |
| `OSD_OUT`                 | OSD              | 0                  | >0                | —                | Count of OSDs in out state                                   |
| `DOWN_PGS`                | placement groups | 0                  | —                 | >0               | PGs stored on down OSDs with no available copy               |
| `UNCLEAN_PGS`             | placement groups | 0                  | >0                | —                | PGs not in clean state (e.g., recovering, backfilling)       |
| `INACTIVE_PGS`            | placement groups | 0                  | —                 | >0               | PGs that cannot perform IO (inactive)                        |
| `IP_COLLISION`            | networking       | no collisions      | —                 | collisions found | Duplicate front or back IP addresses across OSD hosts        |
| `MUTES_AMOUNT`            | health           | 0                  | >0                | —                | Muted health checks that could mask real issues              |
| `OSD_METADATA_SIZE`       | storage          | ≤7%                | >15%              | >20%             | OSD metadata (block.db) size as percentage of total capacity |
| `OSD_NUM_DAEMON_VERSIONS` | versions         | 1 version          | 2 versions        | >2 versions      | Number of distinct OSD daemon versions running               |
| `ALLOW_CRIMSON`           | OSD              | disabled           | enabled           | —                | Whether experimental Crimson OSD is allowed                  |
| `DEVICE_HEALTH_WEAROUT`   | hardware         | no worn devices    | wear >50%         | wear >75%        | SSD/NVMe devices with high wear level                        |

### Using health checks for monitoring

You can run health checks against a remote cluster via SSH:

```shell
cephctl --ceph-binary='ssh user@mon01 ceph' healthcheck
```

Or by setting the environment variable:

```shell
export CEPHCTL_CEPH_BINARY='ssh user@mon01 ceph'
cephctl healthcheck
```

## Roadmap

- [x] v0.0.0
  - [x] Apply declarative configuration for `ceph config`
  - [x] Dump cluster configuration to CephConfig specification
  - [x] Diff configuration against running configuration for `ceph config`
  - [x] Perform healthcheck based on current cluster status
  - [x] Add healthchecks based on current cluster configuration
- [x] v0.1.0
  - [x] Additional healthchecks based on hardware status
  - [x] FreeBSD support in builds
  - [x] Remote Ceph cluster access via SSH
- [x] v0.2.0
  - [x] Apply/Dump declarative configuration for `ceph osd set-*` stuff
- [ ] v0.3.0
  - [ ] Apply/Dump declarative configuration for Ceph Object Gateway (rgw)
- [ ] v0.4.0
  - [ ] Apply/Dump declarative configuration for Pools
- [ ] v0.5.0
  - [ ] Live balancing PGs across OSDs

## Ceph compatibility

All of the changes are tested against Ceph 18.2 (Reef) and 19.2 (Squid), previous
versions are not tested and not guaranteed to work.

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
[GitHub Releases page](https://github.com/runityru/cephctl/releases). Automatically
generated changelog is available for each release. And binaries are available for:

- FreeBSD (amd64v1, amd64v2, amd64v3, arm64)
- Linux (amd64v1, amd64v2, amd64v3, arm64)
- macOS (amd64v1, amd64v2, amd64v3, arm64)
- Windows (amd64v1, amd64v2, amd64v3, arm64)

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

Container image is available at [GitHub Packages](https://github.com/runityru/cephctl/pkgs/container/cephctl%2Fceph)

To replace official Ceph image with the one containing cephctl in cephadm
clusters just do:

```shell
# Set container image as a global parameter to all components
ceph config set mgr mgr/cephadm/container_image_base ghcr.io/runityru/cephctl/ceph

# Run upgrade procedure
ceph orch upgrade start --ceph_version=18.2.2
```

Please note Ceph orch will automatically replace `container_image` parameter
for each component with specific sha256 image ID instead of tag we defined
manually. It's OK and it's a guarantee the image won't be changed in your
cluster.

### Build from source

It's possible to build cephctl from source by simply running the following
command:

```shell
goreleaser build --snapshot --clean
```

or manually via Go compiler

<!-- markdownlint-disable MD013 -->

```shell
go build -v -ldflags="-X 'main.appVersion=$(git rev-parse --short HEAD) (trunk build)' -X 'main.buildTimestamp=$(date -u +%Y-%m-%dT%H:%m:%SZ)'" -o dist/cephctl ./cmd/cephctl/...
```

<!-- markdownlint-enable MD013 -->

## Contribution

cephctl is an open source project so you have the following ways to contribute:

- Documentation
- Fill issues
- Fix bugs
- Suggest/implement new features
- Or any other way, if you have any doubts please fill free to [open discussion](https://github.com/runityru/cephctl/discussions)

### Something about guidelines for the code

There's no actually a particular guidelines for the code but there are some
common rules about it based on TDD, DDD and SOLID, the list could be extended:

#### Split layers of abstraction

Each abstraction layer should be isolated: transport layer from data layer,
DTO models from business logic models and so on.

#### Write unit tests

Since we gonna have amount of packages we can easily write tests, please don't
ignore such ability.

#### Isolate tests

In many projects I've seen how test are using system-wide configuration files,
system-wide binaries and so on - I think this is a bad, insecure, unsafe and
irreproducible practice not allowing to be sure the tests are passed in new
environment (on a new developer machine for instance).

So all the tests in cephctl are isolated:

- code running any commands runs scripts in tests emulating the expected behavior
- command output payload is gathered from real installations
- the only thing you need to run tests is go compiler
