package spec

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewFromDescriptionYAML(t *testing.T) {
	r := require.New(t)

	kind, spec, err := NewFromDescription("testdata/sample_NewFromDescriptionYAML.yaml")
	r.NoError(err)
	r.Equal("CephConfig", kind)
	r.JSONEq(`{"global":{"rbd_cache":"true"},"osd":{"rocksdb_perf":"true"}}`, string(spec))
}

func TestNewFromDescriptionJSON(t *testing.T) {
	r := require.New(t)

	kind, spec, err := NewFromDescription("testdata/sample_NewFromDescriptionJSON.json")
	r.NoError(err)
	r.Equal("CephConfig", kind)
	r.JSONEq(`{"global":{"rbd_cache":"true"},"osd":{"rocksdb_perf":"true"}}`, string(spec))
}
