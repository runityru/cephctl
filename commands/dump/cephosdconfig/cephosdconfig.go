package cephosdconfig

import (
	"context"

	"gopkg.in/yaml.v3"

	"github.com/runityru/cephctl/models"
	"github.com/runityru/cephctl/printer"
	"github.com/runityru/cephctl/service"
)

type DumpCephOSDConfigConfig struct {
	Printer printer.Printer
	Service service.Service
}

func DumpCephOSDConfig(ctx context.Context, doc DumpCephOSDConfigConfig) error {
	type outputSpec struct {
		Kind string               `yaml:"kind"`
		Spec models.CephOSDConfig `yaml:"spec"`
	}

	cfg, err := doc.Service.DumpOSDConfig(ctx)
	if err != nil {
		return err
	}

	spec := outputSpec{
		Kind: "CephOSDConfig",
		Spec: cfg,
	}

	data, err := yaml.Marshal(spec)
	if err != nil {
		return err
	}

	doc.Printer.Println(string(data))
	return nil
}
