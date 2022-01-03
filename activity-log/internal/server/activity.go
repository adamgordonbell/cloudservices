package server

import (
	"errors"
	"sync"

	api "github.com/adamgordonbell/cloudservices/activity-log"
	log "github.com/sirupsen/logrus"
)

type Activities struct {
	mu         sync.Mutex
	activities []api.Activity
}

func (c *Activities) Insert(activity api.Activity) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	activity.ID = len(c.activities) + 1
	c.activities = append(c.activities, activity)
	log.Printf("Added %v", activity)
	return activity.ID
}

var ErrIDNotFound = errors.New("Id not found")

func (c *Activities) Retrieve(id int) (api.Activity, error) {
	log.Printf("Getting %d", id)
	c.mu.Lock()
	defer c.mu.Unlock()
	if id > len(c.activities) {
		log.Printf("Id not found")
		return api.Activity{}, ErrIDNotFound
	}
	return c.activities[id-1], nil
}
