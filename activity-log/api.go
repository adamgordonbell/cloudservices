package api

import "time"

type Activity struct {
	Time        time.Time `json:"time"`
	Description string    `json:"description"`
	ID          int       `json:"id"`
}

type ActivityDocument struct {
	Activity Activity `json:"activity"`
}

type IDDocument struct {
	ID int `json:"id"`
}

type ActivityQueryDocument struct {
	Offset int `json:"offset"`
}
