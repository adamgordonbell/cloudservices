package server

import (
	"fmt"
	"sync"
	"time"
)

type Record struct {
	Time     time.Time `json:"time"`
	Activity []byte    `json:"activity"`
	Id       uint64    `json:"id"`
}

type Activities struct {
	mu      sync.Mutex
	records []Record
}

func (c *Activities) Insert(record Record) uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	record.Id = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Id
}

var ErrIdNotFound = fmt.Errorf("Id not found")

func (c *Activities) Retrieve(id uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if id >= uint64(len(c.records)) {
		return Record{}, ErrIdNotFound
	}
	return c.records[id], nil
}
