package differ

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	diff "github.com/r3labs/diff/v3"
	log "github.com/sirupsen/logrus"
	"github.com/teran/go-ptr"

	"github.com/teran/cephctl/models"
)

const flattenMapSeparator = ":::"

type Differ interface {
	DiffCephConfig(ctx context.Context, from, to models.CephConfig) ([]models.CephConfigDifference, error)
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

	log.Tracef("diff generated: %#v", changelog)

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
		key := pathParts[1]

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

func flattenMap(in map[string]map[string]string) map[string]string {
	out := make(map[string]string)
	for k, v := range in {
		for j, m := range v {
			out[k+flattenMapSeparator+j] = m
		}
	}
	return out
}
