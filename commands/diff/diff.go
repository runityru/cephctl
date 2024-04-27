package commands

import (
	"context"
	"strings"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/teran/cephctl/ceph/config/spec"
	"github.com/teran/cephctl/ceph/config/spec/cephconfig"
	"github.com/teran/cephctl/models"
	"github.com/teran/cephctl/service"
)

type DiffConfig struct {
	Service  service.Service
	Colorize bool
	SpecFile string
}

func Diff(ctx context.Context, ac DiffConfig) error {
	kind, specData, err := spec.NewFromDescription(ac.SpecFile)
	if err != nil {
		return err
	}

	color.NoColor = !ac.Colorize

	switch strings.ToLower(kind) {
	case "cephconfig":
		cfg, err := cephconfig.New(specData)
		if err != nil {
			return err
		}

		changes, err := ac.Service.DiffCephConfig(ctx, cfg)
		if err != nil {
			return err
		}

		for _, change := range changes {
			log.Tracef("change: %#v", change)

			switch change.Kind {
			case models.CephConfigDifferenceKindAdd:
				color.Green("+ %s %s %s", change.Section, change.Key, *change.Value)
			case models.CephConfigDifferenceKindChange:
				color.Yellow("~ %s %s %s -> %s", change.Section, change.Key, *change.OldValue, *change.Value)
			case models.CephConfigDifferenceKindRemove:
				color.Red("- %s %s", change.Section, change.Key)
			}
		}

	default:
		return errors.Errorf("unexpected specification kind: `%s`", kind)
	}

	return nil
}
