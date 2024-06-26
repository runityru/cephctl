package spec

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewFromDescriptionSingle(t *testing.T) {
	r := require.New(t)

	descs, err := NewFromDescription("testdata/sample_NewFromDescriptionSingle.yaml")
	r.NoError(err)
	r.Len(descs, 1)
	r.Equal("CephConfig", descs[0].Kind)
	r.JSONEq(`{"global":{"rbd_cache":"true"},"osd":{"rocksdb_perf":"true"}}`, string(descs[0].Spec))
}

func TestNewFromDescriptionMulti(t *testing.T) {
	r := require.New(t)

	descs, err := NewFromDescription("testdata/sample_NewFromDescriptionMulti.yaml")
	r.NoError(err)
	r.Len(descs, 2)

	r.Equal("CephConfig", descs[0].Kind)
	r.JSONEq(`{"global":{"rbd_cache":"true"},"osd":{"rocksdb_perf":"true"}}`, string(descs[0].Spec))

	r.Equal("CephOSDConfig", descs[1].Kind)
	r.JSONEq(`{"allow_crimson":true}`, string(descs[1].Spec))
}
