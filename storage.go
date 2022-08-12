package main

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

var localDb db

func initDb() error {
	jsonFile, err := os.Open(baseConfig.dbFile)
	defer jsonFile.Close()
	if err != nil {
		localDb.data = make(map[string]string)
		return nil
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &localDb.data)
	if err != nil {
		return err
	}
	return nil
}

func writeDb() {
	file, err := json.Marshal(localDb.data)
	err = os.WriteFile(baseConfig.dbFile, file, 0666)
	if err != nil {
		fmt.Println("error writing to file")
	}
}

// add a new element
func (d *db) add(k string, v string) error {
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
func (d *db) get(k string) (*string, error) {
	d.m.Lock()
	defer d.m.Unlock()
	_, ok := d.data[k]
	if ok {
		b := d.data[k]
		return &b, nil
	}
	return nil, errors.New(fmt.Sprintf("key \"%s\" not found", k))
}

// del delete an element by key
func (d *db) del(k string) error {
	d.m.Lock()
	defer d.m.Unlock()
	_, ok := d.data[k]
	if ok {
		delete(d.data, k)
		return nil
	}
	return errors.New(fmt.Sprintf("key \"%s\" not found", k))
}

// update an element by key
func (d *db) update(k string, v string) {
	d.m.Lock()
	defer d.m.Unlock()
	d.data[k] = v
}

// findAll prints all elements
func (d *db) findAll() map[string]string {
	d.m.Lock()
	defer d.m.Unlock()
	return d.data
}
