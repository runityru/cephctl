package spec

import (
	"encoding/json"
	"io"
	"os"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v3"
)

type Description struct {
	Kind string          `json:"kind"`
	Spec json.RawMessage `json:"spec"`
}

type yamlIntermediate struct {
	Kind string `yaml:"kind"`
	Spec any    `yaml:"spec"`
}

func NewFromDescription(filename string) ([]Description, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "error opening spec file")
	}
	defer fp.Close()

	docs := []Description{}
	dec := yaml.NewDecoder(fp)
	for {
		v := yamlIntermediate{}
		err := dec.Decode(&v)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, errors.Wrap(err, "error unmarshaling document")
		}
		spec, err := json.Marshal(v.Spec)
		if err != nil {
			return nil, errors.Wrap(err, "error marshaling intermediate data structure")
		}

		docs = append(docs, Description{
			Kind: v.Kind,
			Spec: json.RawMessage(spec),
		})
	}

	return docs, nil
}
