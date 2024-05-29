package cephconfig

import (
	"context"

	yaml "gopkg.in/yaml.v3"

	"github.com/teran/cephctl/models"
	"github.com/teran/cephctl/printer"
	"github.com/teran/cephctl/service"
)

type DumpCephConfigConfig struct {
	Printer printer.Printer
	Service service.Service
}

func DumpCephConfig(ctx context.Context, dc DumpCephConfigConfig) error {
	type outputSpec struct {
		Kind string            `yaml:"kind"`
		Spec models.CephConfig `yaml:"spec"`
	}

	cfg, err := dc.Service.DumpConfig(ctx)
	if err != nil {
		return err
	}

	spec := outputSpec{
		Kind: "CephConfig",
		Spec: cfg,
	}

	data, err := yaml.Marshal(spec)
	if err != nil {
		return err
	}

	dc.Printer.Println(string(data))
	return nil
}
