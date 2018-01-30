package transactions

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/boltdb/bolt"
)

const (
	defaultDatabase    = "transactions.moon"
	bucketTransactions = "transactions"
)

// NewBoltStore creates a new store object
func NewBoltStore(path string) (*BoltStore, error) {
	store := &BoltStore{
		DatabaseFile: path,
	}

	err := store.Start()
	if err != nil {
		return nil, err
	}

	return store, nil
}

// BoltStore implements the IStore and uses BoltDB for storage of transactions
type BoltStore struct {
	DatabaseFile string
	db           *bolt.DB
}

// Start loads the database of exist if not it will create a database
func (s *BoltStore) Start() error {
	if len(s.DatabaseFile) == 0 {
		s.DatabaseFile = defaultDatabase
	}

	extension := filepath.Ext(s.DatabaseFile)
	if len(extension) == 0 {
		return fmt.Errorf("invalid database file: %v", s.DatabaseFile)
	}

	var err error

	// Create the store directory if doesnt exist
	storeDir := filepath.Dir(s.DatabaseFile)
	if _, err = os.Stat(storeDir); os.IsNotExist(err) {
		err = os.MkdirAll(storeDir, 0755)
		if err != nil {
			return err
		}
	}

	s.db, err = bolt.Open(s.DatabaseFile, 0600, nil)
	if err != nil {
		return err
	}

	err = s.createTransactionBucketIfNotExists()
	if err != nil {
		return err
	}

	return nil
}

// Stop stops the database
func (s *BoltStore) Stop() {
	s.db.Close()
}

func (s *BoltStore) createTransactionBucketIfNotExists() error {
	return s.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketTransactions))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}

// GetTransactions returns all transactions from the bolt database
func (s *BoltStore) GetTransactions() ([]*Transaction, error) {
	/*transactions := make([]*Transaction, 0)

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketTransactions))
		b.ForEach(func(k, v []byte) error {
			transaction := &Transaction{}
			if err := json.Unmarshal(v, transaction); err != nil {
				return err
			}

			transactions = append(transactions, transaction)
			return nil
		})
		return nil
	})

	if err != nil {
		return nil, err
	}

	return transactions, nil*/

	return mockTransactions(), nil
}

// AddTransaction adds a transaction to the bolt database
func (s *BoltStore) AddTransaction(t *Transaction) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketTransactions))
		id, _ := b.NextSequence()
		t.ID = int(id)

		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}

		return b.Put(itob(t.ID), buf)
	})
}

// UpdateTransaction updates a transaction in the database
func (s *BoltStore) UpdateTransaction(t *Transaction) error {
	return nil
}

// DeleteTransaction deletes a transaction from the bolt database
func (s *BoltStore) DeleteTransaction(t *Transaction) error {
	return nil
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
