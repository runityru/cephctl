package cephosdconfig

import (
	"github.com/creasty/defaults"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/runityru/cephctl/models"
)

func New(in []byte) (models.CephOSDConfig, error) {
	spec := models.CephOSDConfig{}
	if err := defaults.Set(&spec); err != nil {
		return models.CephOSDConfig{}, errors.Wrap(err, "error setting default values")
	}

	if err := yaml.Unmarshal(in, &spec); err != nil {
		return models.CephOSDConfig{}, errors.Wrap(err, "error decoding spec file")
	}

	return spec, nil
}
