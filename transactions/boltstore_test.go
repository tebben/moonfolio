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
		assert.Fail(t, "database file should exist")
	}

	os.Remove(defaultDatabase)
}

func TestCreateBoltStoreError(t *testing.T) {
	store, err := NewBoltStore("blabla")

	assert.Error(t, err)
	assert.Nil(t, store)
}

func TestAddTransaction(t *testing.T) {
	store, err := NewBoltStore(testdatabase)
	transaction := createTestTransaction()

	err = store.AddTransaction(transaction)

	assert.NoError(t, err)
	os.Remove(testdatabase)
}

func TestGetTransactions(t *testing.T) {
	store, err := NewBoltStore(testdatabase)

	err = store.AddTransaction(createTestTransaction())
	err = store.AddTransaction(createTestTransaction2())
	transactions, err := store.GetTransactionsTemp()

	assert.NoError(t, err)
	assert.Len(t, transactions, 2)

	assert.Equal(t, transactions[0].ID, 1)
	assert.Equal(t, transactions[0].CoinID, "BTC")
	assert.Equal(t, transactions[0].CoinAmount, 13.37)
	assert.Equal(t, transactions[0].DateTime, int64(1515940982305))
	assert.Equal(t, transactions[0].PriceUSD, 99.99)

	assert.Equal(t, transactions[1].ID, 2)
	assert.Equal(t, transactions[1].CoinID, "VSX")
	assert.Equal(t, transactions[1].CoinAmount, 14.48)
	assert.Equal(t, transactions[1].DateTime, int64(1515940982316))
	assert.Equal(t, transactions[1].PriceUSD, 88.88)

	os.Remove(testdatabase)
}

func createTestTransaction() *Transaction {
	return &Transaction{Type: TransactionBuy, CoinID: "BTC", CoinAmount: 13.37, DateTime: 1515940982305, PriceUSD: 99.99}
}

func createTestTransaction2() *Transaction {
	return &Transaction{Type: TransactionBuy, CoinID: "VSX", CoinAmount: 14.48, DateTime: 1515940982316, PriceUSD: 88.88}
}
