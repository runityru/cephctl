package apply

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"github.com/runityru/cephctl/ceph/config/spec"
	"github.com/runityru/cephctl/ceph/config/spec/cephconfig"
	"github.com/runityru/cephctl/ceph/config/spec/cephosdconfig"
	"github.com/runityru/cephctl/service"
)

type ApplyConfig struct {
	Service  service.Service
	SpecFile string
}

func Apply(ctx context.Context, ac ApplyConfig) error {
	descs, err := spec.NewFromDescription(ac.SpecFile)
	if err != nil {
		return err
	}

	for _, desc := range descs {
		switch strings.ToLower(desc.Kind) {
		case "cephconfig":
			cfg, err := cephconfig.New(desc.Spec)
			if err != nil {
				return err
			}

			if err := ac.Service.ApplyCephConfig(ctx, cfg); err != nil {
				return err
			}

		case "cephosdconfig":
			cfg, err := cephosdconfig.New(desc.Spec)
			if err != nil {
				return err
			}

			if err := ac.Service.ApplyCephOSDConfig(ctx, cfg); err != nil {
				return err
			}

		default:
			return errors.Errorf("unexpected specification kind: `%s`", desc.Kind)
		}
	}

	return nil
}
