package immutable_map_test

import (
	"testing"

	. "github.com/frozen/immutable_map"
	"github.com/stretchr/testify/require"
)

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

	t.Run("test override 1", func(t *testing.T) {
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

	t.Run("test override 2", func(t *testing.T) {
		path := []byte("t")
		m := New().Insert(path, 10)
		m2 := m.Insert(path, 1)

		require.Equal(t, 10, m.Get1(path))
		require.Equal(t, 1, m2.Get1(path))
	})

	t.Run("test absence of value on same path", func(t *testing.T) {
		path := []byte("aa")
		m := New().Insert(path, 10)

		inf, ok := m.Get([]byte("a"))
		require.False(t, ok)
		require.Nil(t, inf)
	})

}

func TestValuesByKeys(t *testing.T) {
	m := New().Insert([]byte("t"), 10)
	m2 := m.Insert([]byte("tx"), 5)

	require.Equal(t, nil, m.Get1([]byte("tx")))
	require.Equal(t, 5, m2.Get1([]byte("tx")))
	require.Equal(t, 10, m2.Get1([]byte("t")))
}

func TestContains(t *testing.T) {
	t.Run("empty search bytes", func(t *testing.T) {
		require.False(t, New().Insert(nil, 5).Contains(nil))
	})
}

func TestCount(t *testing.T) {
	t.Run("check empty", func(t *testing.T) {
		require.Equal(t, 0, New().Count())
	})
	t.Run("", func(t *testing.T) {
		require.Equal(t, 1, New().Insert([]byte{5}, 5).Count())
	})
}

func TestMap_ToStringMap(t *testing.T) {
	t.Run("test multiple values", func(t *testing.T) {
		v := New().
			Insert([]byte("abc"), 5).
			Insert([]byte("cba"), 6).
			Insert([]byte("ab"), 7).
			Insert([]byte("a"), 8)

		m := v.ToStringMap()

		require.Equal(t, map[string]interface{}{
			"abc": 5,
			"cba": 6,
			"ab":  7,
			"a":   8,
		}, m)
	})
	t.Run("test empty", func(t *testing.T) {
		v := New()
		m := v.ToStringMap()
		require.Equal(t, map[string]interface{}{}, m)
	})
}

func TestMap_ToSlice(t *testing.T) {
	t.Run("test multiple values", func(t *testing.T) {
		v := New().
			Insert([]byte("abc"), 5).
			Insert([]byte("cba"), 6).
			Insert([]byte("ab"), 7).
			Insert([]byte("a"), 8)

		m := v.ToSlice()

		require.Equal(t, []KeyValue{
			{[]byte("a"), 8},
			{[]byte("ab"), 7},
			{[]byte("abc"), 5},
			{[]byte("cba"), 6},
		}, m)
	})
	t.Run("test empty", func(t *testing.T) {
		v := New()
		m := v.ToSlice()
		require.Len(t, m, 0)
	})
}
