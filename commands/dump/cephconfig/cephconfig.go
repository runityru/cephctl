package cephconfig

import (
	"context"
	"fmt"

	"github.com/teran/cephctl/models"
	"github.com/teran/cephctl/service"
	yaml "gopkg.in/yaml.v3"
)

type DumpCephConfigConfig struct {
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

	fmt.Println(string(data))
	return nil
}
