package differ

import (
	"context"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	diff "github.com/r3labs/diff/v3"
	log "github.com/sirupsen/logrus"
	"github.com/teran/go-ptr"

	"github.com/runityru/cephctl/models"
)

const flattenMapSeparator = ":::"

type Differ interface {
	DiffCephConfig(ctx context.Context, from, to models.CephConfig) ([]models.CephConfigDifference, error)
	DiffCephOSDConfig(ctx context.Context, from, to models.CephOSDConfig) ([]models.CephOSDConfigDifference, error)
}

type differ struct{}

func New() Differ {
	return &differ{}
}

func (d *differ) DiffCephConfig(ctx context.Context, from, to models.CephConfig) ([]models.CephConfigDifference, error) {
	srcf := flattenMap(from)
	cfgf := flattenMap(to)

	changelog, err := diff.Diff(srcf, cfgf)
	if err != nil {
		return nil, errors.Wrap(err, "error comparing current and desired configuration")
	}

	log.WithFields(log.Fields{
		"component": "differ",
	}).Tracef("diff generated: %#v", changelog)

	changes := []models.CephConfigDifference{}
	for _, change := range changelog {
		if len(change.Path) != 1 {
			return nil, errors.Errorf("unexpected path structure in diff (got %d parts): mostly possible programmer error", len(change.Path))
		}

		pathParts := strings.SplitN(change.Path[0], flattenMapSeparator, 2)
		if len(pathParts) != 2 {
			return nil, errors.Errorf("unexpected path received: no flattened parts found (%d received)", len(pathParts))
		}

		section := pathParts[0]
		if len(section) == 0 {
			return nil, errors.Errorf("section name cannot be empty")
		}

		key := pathParts[1]
		if len(key) == 0 {
			return nil, errors.Errorf("key name cannot be empty")
		}

		switch change.Type {
		case diff.CREATE:
			v, ok := change.To.(string)
			if !ok {
				log.Warnf("unexpected value type: expected string, got %T", v)
				break
			}

			changes = append(changes, models.CephConfigDifference{
				Kind:    models.CephConfigDifferenceKindAdd,
				Section: section,
				Key:     key,
				Value:   ptr.String(v),
			})

		case diff.UPDATE:
			oldV, ok := change.From.(string)
			if !ok {
				log.Warnf("unexpected old value type: expected string, got %T", oldV)
				break
			}

			v, ok := change.To.(string)
			if !ok {
				log.Warnf("unexpected new value type: expected string, got %T", v)
				break
			}

			changes = append(changes, models.CephConfigDifference{
				Kind:     models.CephConfigDifferenceKindChange,
				Section:  section,
				Key:      key,
				OldValue: ptr.String(oldV),
				Value:    ptr.String(v),
			})

		case diff.DELETE:
			changes = append(changes, models.CephConfigDifference{
				Kind:    models.CephConfigDifferenceKindRemove,
				Section: section,
				Key:     key,
			})

		}
	}

	return changes, nil
}

func (d *differ) DiffCephOSDConfig(ctx context.Context, from, to models.CephOSDConfig) ([]models.CephOSDConfigDifference, error) {
	changelog, err := diff.Diff(from, to)
	if err != nil {
		return nil, errors.Wrap(err, "error comparing current and desired configuration")
	}

	changes := []models.CephOSDConfigDifference{}
	for _, change := range changelog {
		log.Printf("single changes: %#v", change)

		var (
			oldValue string
			newValue string
		)
		switch change.Type {
		case diff.UPDATE:
			switch v := change.From.(type) {
			case bool:
				oldValue = strconv.FormatBool(v)
			case float32:
				oldValue = strconv.FormatFloat(float64(v), 'f', 2, 32)
			case string:
				oldValue = v
			default:
				log.Warnf("unexpected old value type: got %T", v)
			}

			switch v := change.To.(type) {
			case bool:
				newValue = strconv.FormatBool(v)
			case float32:
				newValue = strconv.FormatFloat(float64(v), 'f', 2, 32)
			case string:
				newValue = v
			default:
				log.Warnf("unexpected new value type: got %T", v)
			}

			changes = append(changes, models.CephOSDConfigDifference{
				Key:      change.Path[0],
				OldValue: oldValue,
				Value:    newValue,
			})

		default:
			log.Warnf("unexpected change type: `%s`", change.Type)
			break
		}
	}

	log.Printf("changes: %#v", changes)

	return changes, nil
}
