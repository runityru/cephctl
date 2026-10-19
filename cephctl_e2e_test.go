//go:build e2e

// Package cephctl_e2e contains end-to-end tests that run the cephctl binary
// against a real single-node Ceph demo container (squid / tentacle) using
// github.com/teran/go-docker-testsuite. These tests are gated behind the
// `e2e` build tag so they are excluded from `go build ./...`, `go vet ./...`
// and plain `go test ./...` — they require a running Docker daemon and a
// Ceph container image that takes several minutes to bootstrap.
//
// Run with:
//
//	go test -tags e2e -v ./...
//
// Target a single Ceph image (default: v19.2.6 and v20.2.4) via CEPH_IMAGE:
//
//	CEPH_IMAGE=ghcr.io/teran/ceph-container/ceph:v20.2.4 go test -tags e2e -v ./...
//
// go-docker-testsuite is consumed through a local `replace` directive (see
// go.mod) until a release containing docker.Container.Exec and
// ceph.WithBinds is published.
package cephctl_e2e

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/docker/docker/client"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	ceph "github.com/teran/go-docker-testsuite/applications/ceph"
)

// defaultImages are the Ceph release trains in scope for this regression test.
// They match the demo images published from github.com/teran/ceph-container.
var defaultImages = []string{
	"ghcr.io/teran/ceph-container/ceph:v19.2.6", // squid
	"ghcr.io/teran/ceph-container/ceph:v20.2.4", // tentacle
}

// imagesUnderTest returns the images to test. When CEPH_IMAGE is set only that
// single image is exercised, mirroring go-docker-testsuite's override convention.
func imagesUnderTest() []string {
	if v := os.Getenv("CEPH_IMAGE"); v != "" {
		return []string{v}
	}
	return defaultImages
}

// requireDocker skips the test (rather than failing) when running with -short
// or when the Docker daemon is unreachable, so the suite degrades gracefully
// on machines without Docker. Mirrors go-docker-testsuite's convention.
func requireDocker(t *testing.T) {
	t.Helper()

	if testing.Short() {
		t.Skip("skipping e2e test in -short mode")
	}

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		t.Skipf("skipping e2e test: unable to create docker client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := cli.Ping(ctx); err != nil {
		t.Skipf("skipping e2e test: docker daemon unavailable: %v", err)
	}
}

// buildCephctl compiles the cephctl binary into a temp dir so it can be
// bind-mounted into the container. It returns the path to the binary. If the
// build fails the test is skipped rather than failed.
func buildCephctl(t *testing.T) string {
	t.Helper()

	// `go test` runs with the working directory set to the package source
	// directory, which for this file is the cephctl repository root.
	repoRoot, err := os.Getwd()
	require.NoError(t, err)

	binPath := filepath.Join(t.TempDir(), "cephctl")

	cmd := exec.Command("go", "build", "-o", binPath, "github.com/runityru/cephctl/cmd/cephctl")
	cmd.Dir = repoRoot
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Skipf("skipping e2e test: failed to build cephctl binary: %v\n%s", err, out)
	}

	return binPath
}

func TestCephctlAgainstRealCeph(t *testing.T) {
	requireDocker(t)

	binPath := buildCephctl(t)

	for _, image := range imagesUnderTest() {
		image := image
		t.Run(image, func(t *testing.T) {
			runCephctlE2E(t, image, binPath)
		})
	}
}

// runCephctlE2E starts a Ceph demo container with the cephctl binary
// bind-mounted and exercises the real `cephctl dump cephosdconfig` and
// `cephctl healthcheck` commands against it. This is the exact command from
// issue #115 — a successful exit without a panic proves the report decoding
// regression is fixed.
func runCephctlE2E(t *testing.T, image, binPath string) {
	t.Helper()

	// Ceph bootstrap is slow (mon/mgr/osd/rgw bring-up); give it a generous budget.
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Minute)
	defer cancel()

	app, err := ceph.NewWithImageT(t, ctx, image, ceph.WithBinds(binPath+":/usr/local/bin/cephctl:ro"))
	require.NoError(t, err)
	// NewWithImageT ties the container lifecycle to the test via t.Cleanup.

	t.Run("dump cephosdconfig", func(t *testing.T) {
		res, err := app.Container().Exec(ctx, []string{
			"cephctl", "dump", "cephosdconfig",
			"--ceph-binary", "/usr/bin/ceph",
		})
		require.NoError(t, err, "exec cephosdconfig failed: %s", res.Combined())
		require.Equal(t, 0, res.ExitCode, "stderr: %s", string(res.Stderr))

		var doc struct {
			Kind string `yaml:"kind"`
		}
		require.NoError(t, yaml.Unmarshal(res.Combined(), &doc), "output is not valid YAML:\n%s", res.Combined())
		require.Equal(t, "CephOSDConfig", doc.Kind)
	})

	t.Run("healthcheck", func(t *testing.T) {
		res, err := app.Container().Exec(ctx, []string{
			"cephctl", "healthcheck",
			"--ceph-binary", "/usr/bin/ceph",
		})
		require.NoError(t, err, "exec healthcheck failed: %s", res.Combined())
		require.Equal(t, 0, res.ExitCode, "stderr: %s", string(res.Stderr))
	})
}
