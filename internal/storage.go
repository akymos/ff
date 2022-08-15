package internal

import (
	"errors"
	"fmt"
	bolt "go.etcd.io/bbolt"
)

func InitDb() (*bolt.DB, error) {
	db, err := bolt.Open(BaseConfig.DbFile, 0600, nil)
	if err != nil {
		return nil, err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("ff"))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Add a new element
func Add(k string, v string) error {
	val, err := Get(k)
	if val != nil {
		return errors.New(fmt.Sprintf("key \"%s\" already exists", k))
	}
	err = BaseConfig.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("ff"))
		return b.Put([]byte(k), []byte(v))
	})
	if err != nil {
		return err
	}
	return nil
}

// find an element by key
func Get(k string) (*string, error) {
	var val string
	err := BaseConfig.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("ff"))
		val = string(b.Get([]byte(k)))
		return nil
	})
	if err != nil {
		return nil, err
	}
	if val == "" {
		return nil, errors.New(fmt.Sprintf("key \"%s\" not found", k))
	}
	return &val, nil

}

// Del delete an element by key
func Del(k string) error {
	_, err := Get(k)
	if err != nil {
		return err
	}
	err = BaseConfig.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("ff"))
		return b.Delete([]byte(k))
	})
	return err
}

// Update an element by key
func Update(k string, v string) error {
	val, err := Get(k)
	if val == nil {
		return errors.New(fmt.Sprintf("key \"%s\" not found", k))
	}
	err = BaseConfig.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("ff"))
		return b.Put([]byte(k), []byte(v))
	})
	if err != nil {
		return err
	}
	return nil
}

// FindAll prints all elements
func FindAll() map[string]string {
	data := make(map[string]string)
	err := BaseConfig.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("ff"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			data[string(k)] = string(v)
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return data
}
