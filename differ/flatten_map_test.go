package differ

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFlattenMap(t *testing.T) {
	r := require.New(t)

	out := flattenMap(map[string]map[string]string{
		"key1": {
			"key2": "value3",
		},
		"key3": {
			"key4": "value5",
		},
	})
	r.Equal(map[string]string{
		"key1:::key2": "value3",
		"key3:::key4": "value5",
	}, out)
}
