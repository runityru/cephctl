package ceph

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"os/exec"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	cephModels "github.com/runityru/cephctl/ceph/models"
	"github.com/runityru/cephctl/models"
)

type Ceph interface {
	ApplyCephConfigOption(ctx context.Context, section, key, value string) error
	ClusterReport(ctx context.Context) (models.ClusterReport, error)
	ClusterStatus(ctx context.Context) (models.ClusterStatus, error)
	DumpConfig(ctx context.Context) (models.CephConfig, error)
	ListDevices(ctx context.Context) ([]models.Device, error)
	RemoveCephConfigOption(ctx context.Context, section, key string) error
}

type ceph struct {
	binaryPath string
}

func New(binaryPath string) Ceph {
	return &ceph{
		binaryPath: binaryPath,
	}
}

func (c *ceph) ApplyCephConfigOption(ctx context.Context, section, key, value string) error {
	bin, args := mkCommand(c.binaryPath, []string{"config", "set", section, key, value})

	cmd := exec.CommandContext(ctx, bin, args...)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "error applying configuration")
	}
	return nil
}

func (c *ceph) ClusterReport(ctx context.Context) (models.ClusterReport, error) {
	buf := &bytes.Buffer{}
	bin, args := mkCommand(c.binaryPath, []string{"report", "--format=json"})

	cmd := exec.CommandContext(ctx, bin, args...)
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return models.ClusterReport{}, errors.Wrap(err, "error retrieving report")
	}

	log.Tracef("command output: `%s`", buf.String())

	rep := cephModels.Report{}
	if err := json.Unmarshal(buf.Bytes(), &rep); err != nil {
		return models.ClusterReport{}, errors.Wrap(err, "error decoding report")
	}

	return rep.ToSvc()
}

func (c ceph) ClusterStatus(ctx context.Context) (models.ClusterStatus, error) {
	buf := &bytes.Buffer{}
	bin, args := mkCommand(c.binaryPath, []string{"status", "--format=json"})

	cmd := exec.CommandContext(ctx, bin, args...)
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return models.ClusterStatus{}, errors.Wrap(err, "error retrieving cluster status")
	}

	log.Tracef("command output: `%s`", buf.String())

	st := cephModels.Status{}
	if err := json.Unmarshal(buf.Bytes(), &st); err != nil {
		return models.ClusterStatus{}, errors.Wrap(err, "error decoding response")
	}

	return st.ToSvc()
}

func (c *ceph) DumpConfig(ctx context.Context) (models.CephConfig, error) {
	cfg := []cephModels.ConfigOption{}
	buf := &bytes.Buffer{}
	bin, args := mkCommand(c.binaryPath, []string{"config", "dump", "--format=json"})

	cmd := exec.CommandContext(ctx, bin, args...)
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return nil, errors.Wrap(err, "error running command")
	}
	log.Tracef("command output: `%s`", buf.String())

	if err := json.Unmarshal(buf.Bytes(), &cfg); err != nil {
		return nil, errors.Wrap(err, "error decoding response")
	}

	out := make(models.CephConfig)
	for _, v := range cfg {
		if _, ok := out[v.Section]; !ok {
			out[v.Section] = make(map[string]string)
		}

		out[v.Section][v.Name] = v.Value
	}

	return out, nil
}

func (c *ceph) ListDevices(ctx context.Context) ([]models.Device, error) {
	devices := []cephModels.Device{}
	buf := &bytes.Buffer{}
	bin, args := mkCommand(c.binaryPath, []string{"device", "ls", "--format=json"})

	cmd := exec.CommandContext(ctx, bin, args...)
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return nil, errors.Wrap(err, "error listing devices")
	}

	if err := json.Unmarshal(buf.Bytes(), &devices); err != nil {
		return nil, errors.Wrap(err, "error decoding response")
	}

	out := []models.Device{}
	for _, v := range devices {
		out = append(out, models.Device{
			ID:        v.DevID,
			Daemons:   append([]string{}, v.Daemons...),
			WearLevel: v.WearLevel,
		})
	}

	return out, nil
}

func (c *ceph) RemoveCephConfigOption(ctx context.Context, section, key string) error {
	bin, args := mkCommand(c.binaryPath, []string{"config", "rm", section, key})

	cmd := exec.CommandContext(ctx, bin, args...)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "error applying configuration")
	}
	return nil
}
