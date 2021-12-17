package client

import (
	"errors"
	"time"
)

type Activity struct {
	Time        time.Time `json:"time"`
	Description string    `json:"description"`
	ID          int       `json:"id"`
}

type Activities struct {
	activities []Activity
}

func (c *Activities) Insert(activity Activity) (int, error) {
	activity.ID = len(c.activities)
	c.activities = append(c.activities, activity)
	return activity.ID, nil
}

var ErrIDNotFound = errors.New("Id not found")

func (c *Activities) Retrieve(id int) (Activity, error) {
	a := Activity{Time: time.Now(), Description: "fake"}
	c.Insert(a)

	if id >= len(c.activities) {
		return Activity{}, ErrIDNotFound
	}
	return c.activities[id], nil
}
