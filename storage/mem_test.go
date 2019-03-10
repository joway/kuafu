package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemStorage(t *testing.T) {
	stg := NewMemStorage()
	k := "test"
	v := []byte("xxx")
	err := stg.Add(k, v)
	assert.NoError(t, err)

	got, err := stg.Get(k)
	assert.NoError(t, err)
	assert.Equal(t, v, got)

	got, err = stg.Get("not_found")
	assert.Error(t, err)
	assert.Nil(t, got)
}
