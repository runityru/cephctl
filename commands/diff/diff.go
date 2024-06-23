package diff

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/runityru/cephctl/ceph/config/spec"
	"github.com/runityru/cephctl/ceph/config/spec/cephconfig"
	"github.com/runityru/cephctl/ceph/config/spec/cephosdconfig"
	"github.com/runityru/cephctl/models"
	"github.com/runityru/cephctl/printer"
	"github.com/runityru/cephctl/service"
)

type DiffConfig struct {
	Service  service.Service
	Printer  printer.Printer
	SpecFile string
}

func Diff(ctx context.Context, ac DiffConfig) error {
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

			changes, err := ac.Service.DiffCephConfig(ctx, cfg)
			if err != nil {
				return err
			}

			for _, change := range changes {
				log.WithFields(log.Fields{
					"component": "command",
				}).Tracef("change: %#v", change)

				switch change.Kind {
				case models.CephConfigDifferenceKindAdd:
					ac.Printer.Green("+ %s %s %s", change.Section, change.Key, *change.Value)
				case models.CephConfigDifferenceKindChange:
					ac.Printer.Yellow("~ %s %s %s -> %s", change.Section, change.Key, *change.OldValue, *change.Value)
				case models.CephConfigDifferenceKindRemove:
					ac.Printer.Red("- %s %s", change.Section, change.Key)
				}
			}

		case "cephosdconfig":
			cfg, err := cephosdconfig.New(desc.Spec)
			if err != nil {
				return err
			}

			changes, err := ac.Service.DiffCephOSDConfig(ctx, cfg)
			if err != nil {
				return err
			}

			for _, change := range changes {
				log.WithFields(log.Fields{
					"component": "command",
				}).Tracef("change: %#v", change)

				ac.Printer.Yellow("~ %s %s -> %s", change.Key, change.OldValue, change.Value)
			}

		default:
			return errors.Errorf("unexpected specification kind: `%s`", desc.Kind)
		}
	}

	return nil
}
