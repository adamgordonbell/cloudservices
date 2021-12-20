package client

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

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

type Activities struct {
	URL string
}

func (c *Activities) Insert(activity Activity) (int, error) {
	var document IDDocument
	activityDoc := ActivityDocument{Activity: activity}
	bytes, err := json.Marshal(activityDoc)
	if err != nil {
		return -1, err
	}
	req, err := http.NewRequest(http.MethodGet, c.URL, strings.NewReader(string(bytes)))
	if err != nil {
		return -1, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return -1, err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return -1, err
	}
	err = json.Unmarshal(body, &document)
	if err != nil {
		return -1, err
	}
	return document.ID, nil
}

var ErrIDNotFound = errors.New("Id not found")

func (c *Activities) Retrieve(id int) (Activity, error) {
	var document ActivityDocument
	idDoc := IDDocument{ID: id}
	bytes, err := json.Marshal(idDoc)
	if err != nil {
		return document.Activity, err
	}
	req, err := http.NewRequest(http.MethodGet, c.URL, strings.NewReader(string(bytes)))
	if err != nil {
		return document.Activity, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return document.Activity, err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return document.Activity, err
	}
	err = json.Unmarshal(body, &document)
	if err != nil {
		return document.Activity, err
	}
	return document.Activity, nil
}
