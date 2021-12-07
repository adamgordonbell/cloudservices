package server

import (
	"fmt"
	"sync"
)

type Record struct {
	Value  []byte `json:"value"` //wtf?
	Offset uint64 `json:"offset"`
}

type Log struct {
	mu      sync.Mutex
	records []Record
}

func NewLog() *Log { //why?
	return &Log{}
}

func (c *Log) Append(record Record) (uint64, error) { //why error if no chance of it?
	c.mu.Lock()
	defer c.mu.Unlock()
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

var ErrOffsetNotFound = fmt.Errorf("offset not found")

func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if offset >= uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}
