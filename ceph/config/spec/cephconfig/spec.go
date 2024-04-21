package cephconfig

import (
	"github.com/pkg/errors"
	"github.com/teran/cephctl/models"
	yaml "gopkg.in/yaml.v3"
)

func New(in []byte) (models.CephConfig, error) {
	spec := models.CephConfig{}
	if err := yaml.Unmarshal(in, &spec); err != nil {
		return nil, errors.Wrap(err, "error decoding spec file")
	}

	return spec, nil
}
