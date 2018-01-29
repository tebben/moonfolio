package transactions

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testdatabase = "test.moon"
)

func TestCreateBoltStoreOk(t *testing.T) {
	store, err := NewBoltStore(testdatabase)

	assert.Nil(t, err)
	assert.NotNil(t, store)

	if _, err := os.Stat(testdatabase); err != nil {
		assert.Error(t, errors.New("database file should exist"))
	}

	os.Remove(testdatabase)
}

func TestCreateDefaultStore(t *testing.T) {
	store, err := NewBoltStore("")

	assert.Nil(t, err)
	assert.NotNil(t, store)

	if _, err := os.Stat(defaultDatabase); err != nil {
		assert.Error(t, errors.New("database file should exist"))
	}

	os.Remove(defaultDatabase)
}

func TestCreateBoltStoreError(t *testing.T) {
	store, err := NewBoltStore("blabla")

	assert.NotNil(t, err)
	assert.Nil(t, store)
}
