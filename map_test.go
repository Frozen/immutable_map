package immutable_map

import "testing"
import "github.com/stretchr/testify/require"

func TestMap(t *testing.T) {
	m := New()
	require.False(t, m.Contains([]byte("tx")))
	m2 := m.Insert([]byte("tx"), 5)
	require.True(t, m2.Contains([]byte("tx")))
	require.False(t, m.Contains([]byte("tx")))

	rs1, ok1 := m.Get([]byte("tx"))
	require.Equal(t, nil, rs1)
	require.Equal(t, false, ok1)

	rs2, ok2 := m2.Get([]byte("tx"))
	require.Equal(t, 5, rs2)
	require.Equal(t, true, ok2)
}

func TestMap_InsertGet(t *testing.T) {
	t.Run("find equal value", func(t *testing.T) {
		m := New()
		m2 := m.Insert([]byte("tx"), 5)
		value, ok := m2.Get([]byte("tx"))
		require.Equal(t, 5, value)
		require.True(t, ok)
	})
	t.Run("find less value", func(t *testing.T) {
		m := New()
		m2 := m.Insert([]byte("c"), 5)
		value, ok := m2.Get([]byte("0"))
		require.False(t, ok)
		require.Nil(t, value)
	})
	t.Run("find gt value", func(t *testing.T) {
		m := New()
		m2 := m.Insert([]byte("a"), 5)
		value, ok := m2.Get([]byte("b"))
		require.False(t, ok)
		require.Nil(t, value)
	})

	t.Run("insert gt values", func(t *testing.T) {
		m := New()
		m2 := m.Insert([]byte("items"), 5).Insert([]byte("call"), 6)
		require.Equal(t, 6, m2.Get1([]byte("call")))
		require.Equal(t, 5, m2.Get1([]byte("items")))
	})

	t.Run("test right override", func(t *testing.T) {
		key1 := []byte("ab")
		key2 := []byte("a")

		m := New().Insert(key1, 10)
		require.True(t, m.Contains(key1))
		require.Equal(t, 10, m.Get1(key1))

		m2 := m.Insert(key2, 5)
		// contains prev data
		require.True(t, m2.Contains(key1))
		require.Equal(t, 10, m2.Get1(key1))
		// contains new data
		require.True(t, m2.Contains(key2))
		require.Equal(t, 5, m2.Get1(key2))
	})

}

func TestValuesByKeys(t *testing.T) {
	m := New().Insert([]byte("t"), 10)
	m2 := m.Insert([]byte("tx"), 5)

	require.Equal(t, nil, m.Get1([]byte("tx")))
	require.Equal(t, 5, m2.Get1([]byte("tx")))
	require.Equal(t, 10, m2.Get1([]byte("t")))
}

func TestOverride(t *testing.T) {
	path := []byte("t")
	m := New().Insert(path, 10)
	m2 := m.Insert(path, 1)

	require.Equal(t, 10, m.Get1(path))
	require.Equal(t, 1, m2.Get1(path))
}
