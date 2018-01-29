package transactions

import (
	"path/filepath"

	"github.com/boltdb/bolt"
)

const (
	defaultDatabase = "transactions.moon"
)

// NewBoltStore creates a new store object
func NewBoltStore(path string) (*BoltStore, error) {
	store := &BoltStore{
		DatabaseFile: path,
	}

	err := store.CreateOrTryLoadDatabase()
	if err != nil {
		return nil, err
	}

	return store, nil
}

// BoltStore implements the IStore and uses BoltDB for storage of transactions
type BoltStore struct {
	DatabaseFile string
}

// CreateOrTryLoadDatabase loads the database of exist if not it will create a database
func (s *BoltStore) CreateOrTryLoadDatabase() error {
	extension := filepath.Ext(s.DatabaseFile)
	if len(extension) == 0 {
		s.DatabaseFile = defaultDatabase
	}

	db, err := bolt.Open(s.DatabaseFile, 0600, nil)
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		return err
	}

	return nil
}

// GetTransactions returns all transactions from the bolt database
func (s *BoltStore) GetTransactions() ([]*Transaction, error) {
	return mockTransactions(), nil
}

// AddTransaction adds a transaction to the bolt database
func (s *BoltStore) AddTransaction(t *Transaction) error {
	return nil
}

// UpdateTransaction updates a transaction in the database
func (s *BoltStore) UpdateTransaction(t *Transaction) error {
	return nil
}

// DeleteTransaction deletes a transaction from the bolt database
func (s *BoltStore) DeleteTransaction(t *Transaction) error {
	return nil
}
