package spec

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v3"
)

type description struct {
	Kind string          `json:"kind"`
	Spec json.RawMessage `json:"spec"`
}

func NewFromDescription(filename string) (string, json.RawMessage, error) {
	desc := description{}

	data, err := os.ReadFile(filename)
	if err != nil {
		return "", nil, errors.Wrap(err, "error reading configuration file")
	}

	switch strings.ToLower(filepath.Ext(filename)) {
	case ".yml", ".yaml":
		type intermediate struct {
			Kind string `yaml:"kind"`
			Spec any    `yaml:"spec"`
		}

		d := intermediate{}
		err := yaml.Unmarshal(data, &d)
		if err != nil {
			return "", nil, errors.Wrap(err, "error unmarshaling intermediate configuration")
		}

		spec, err := json.Marshal(d.Spec)
		if err != nil {
			return "", nil, errors.Wrap(err, "error marshaling intermediate configuration")
		}
		return d.Kind, json.RawMessage(spec), nil
	case ".json":
		// skip since supported natively
	default:
		return "", nil, errors.Errorf("unexpected file format: `%s`", filepath.Ext(filename))
	}

	return desc.Kind, desc.Spec, json.Unmarshal(data, &desc)
}
