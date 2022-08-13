package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type db struct {
	data map[string]string
	m    sync.Mutex
}

var LocalDb db

func InitDb() error {
	jsonFile, err := os.Open(BaseConfig.DbFile)
	defer jsonFile.Close()
	if err != nil {
		LocalDb.data = make(map[string]string)
		return nil
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &LocalDb.data)
	if err != nil {
		return err
	}
	return nil
}

func WriteDb() {
	file, err := json.Marshal(LocalDb.data)
	err = os.WriteFile(BaseConfig.DbFile, file, 0666)
	if err != nil {
		fmt.Println("error writing to file")
	}
}

// Add a new element
func (d *db) Add(k string, v string) error {
	if k == "" {
		return errors.New("key is empty")
	}
	_, ok := d.data[k]
	if !ok {
		d.m.Lock()
		defer d.m.Unlock()
		d.data[k] = v
		return nil
	}
	return errors.New(fmt.Sprintf("key \"%s\" already exists", k))
}

// find an element by key
func (d *db) Get(k string) (*string, error) {
	d.m.Lock()
	defer d.m.Unlock()
	_, ok := d.data[k]
	if ok {
		b := d.data[k]
		return &b, nil
	}
	return nil, errors.New(fmt.Sprintf("key \"%s\" not found", k))
}

// Del delete an element by key
func (d *db) Del(k string) error {
	d.m.Lock()
	defer d.m.Unlock()
	_, ok := d.data[k]
	if ok {
		delete(d.data, k)
		return nil
	}
	return errors.New(fmt.Sprintf("key \"%s\" not found", k))
}

// Update an element by key
func (d *db) Update(k string, v string) error {
	d.m.Lock()
	defer d.m.Unlock()
	_, ok := d.data[k]
	if ok {
		d.data[k] = v
		return nil
	}
	return errors.New(fmt.Sprintf("key \"%s\" not found", k))

}

// FindAll prints all elements
func (d *db) FindAll() map[string]string {
	d.m.Lock()
	defer d.m.Unlock()
	return d.data
}
