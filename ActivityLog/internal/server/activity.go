package server

import (
	"errors"
	"sync"
	"time"
)

type Activity struct {
	Time        time.Time `json:"time"`
	Description string    `json:"description"`
	ID          uint64    `json:"id"`
}

type Activities struct {
	mu         sync.Mutex
	activities []Activity
}

func (c *Activities) Insert(activity Activity) uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	activity.ID = uint64(len(c.activities))
	c.activities = append(c.activities, activity)
	return activity.ID
}

var ErrIDNotFound = errors.New("Id not found")

func (c *Activities) Retrieve(id uint64) (Activity, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if id >= uint64(len(c.activities)) {
		return Activity{}, ErrIDNotFound
	}
	return c.activities[id], nil
}
