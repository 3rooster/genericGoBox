package syncMap

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type obj struct {
	data string
}

func TestPointer(t *testing.T) {
	m := Map[string, *obj]{}
	v1Key := "key"
	v1 := &obj{data: "v1"}
	m.Store(v1Key, v1)
	data, ok := m.Load(v1Key)

	require.Equal(t, true, ok)
	require.Equal(t, v1, data)

	v2Key := "nil"
	m.Store(v2Key, nil)
	data, ok = m.Load(v2Key)
	require.Equal(t, true, ok)
	require.Nil(t, data, "expected nil")
}
