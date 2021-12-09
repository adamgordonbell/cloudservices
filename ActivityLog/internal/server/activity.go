package server

import (
	"fmt"
	"sync"
	"time"
)

type Activity struct {
	Time        time.Time `json:"time"`
	Description string    `json:"description"`
	Id          uint64    `json:"id"`
}

type Activities struct {
	mu         sync.Mutex
	activities []Activity
}

func (c *Activities) Insert(activity Activity) uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	activity.Id = uint64(len(c.activities))
	c.activities = append(c.activities, activity)
	return activity.Id
}

var ErrIdNotFound = fmt.Errorf("Id not found")

func (c *Activities) Retrieve(id uint64) (Activity, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if id >= uint64(len(c.activities)) {
		return Activity{}, ErrIdNotFound
	}
	return c.activities[id], nil
}
