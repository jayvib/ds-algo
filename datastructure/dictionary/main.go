package main

import (
	"errors"
	"fmt"
	"sync"
)

var ErrNotFound = errors.New("item not found")

type DictVal string
type DictKey string

func NewDictionary() *Dictionary {
	d := &Dictionary{}
	d.Reset()
	return d
}

type Dictionary struct {
	elements map[DictKey]DictVal
	rwmu sync.RWMutex
}

func (d *Dictionary) Put(key DictKey, value DictVal) {
	d.rwmu.Lock()
	defer d.rwmu.Unlock()
	d.elements[key] = value // TODO: What if there's an existing data?
}

func (d *Dictionary) Remove(key DictKey) {
	delete(d.elements, key)
}

func (d *Dictionary) IsExists(key DictKey) bool {
	_, exists := d.elements[key]
	return exists
}

func (d *Dictionary) Find(key DictKey) (DictVal, error) {
	d.rwmu.RLock()
	defer d.rwmu.RUnlock()
	val, exists := d.elements[key]
	if !exists {
		return "", fmt.Errorf("cannot find item for %s: %w", key, ErrNotFound)
	}
	return  val, nil
}

func (d *Dictionary) Reset() {
	d.rwmu.Lock()
	defer d.rwmu.Unlock()
	d.elements = make(map[DictKey]DictVal)
}

func main() {

}
