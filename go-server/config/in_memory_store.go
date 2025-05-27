package main

// import (
// 	"encoding/json"
// 	"log"

// 	badger "github.com/dgraph-io/badger/v4"
// )

// type BadgerStore struct {
// 	db *badger.DB
// }
// type StoredVector struct {
// 	ID     string    `json:"id"`
// 	Vector []float64 `json:"vector"`
// }

// func NewBadgerStore() (*BadgerStore, error) {
// 	opts := badger.DefaultOptions("/tmp/badger").WithLogger(nil)
// 	db, err := badger.Open(opts)
// 	if err != nil {
// 		log.Fatalf("Failed to open BadgerDB: %v", err)
// 	}
// 	return &BadgerStore{db: db}, nil
// }

// func (bs *BadgerStore) Insert(id string, vector []float64) error {
// 	sv := StoredVector{
// 		ID:     id,
// 		Vector: vector}

// 	data, err := json.Marshal(sv)
// 	if err != nil {
// 		log.Printf("Error marshalling vector: %v", err)
// 		return err
// 	}
// 	err = bs.db.Update(func(txn *badger.Txn) error {
// 		return txn.Set([]byte(id), data)
// 	})
// 	return err
// }

// func (bs *BadgerStore) Get(id string) (StoredVector, error) {
// 	var sv StoredVector
// 	err := bs.db.View(func(txn *badger.Txn) error {
// 		item, err := txn.Get([]byte(id))
// 		if err != nil {
// 			return err
// 		}
// 		return item.Value(func(val []byte) error {
// 			return json.Unmarshal(val, &sv)
// 		})
// 	})
// 	return sv, err
// }

// func (bs *BadgerStore) Close() error {
// 	return bs.db.Close()
// }
