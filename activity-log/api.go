package api

import "time"

type Activity struct {
	Time        time.Time `json:"time"`
	Description string    `json:"description"`
	ID          uint64    `json:"id"`
}

type ActivityDocument struct {
	Activity Activity `json:"activity"`
}

type IDDocument struct {
	ID uint64 `json:"id"`
}
