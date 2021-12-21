package client

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (a Activity) String() string {
	return fmt.Sprintf("ID:%d\t\"%s\"\t%d-%d-%d",
		a.ID, a.Description, a.Time.Year(), a.Time.Month(), a.Time.Day())
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
	activityDoc := ActivityDocument{Activity: activity}
	bytes, err := json.Marshal(activityDoc)
	if err != nil {
		return -1, err
	}
	req, err := http.NewRequest(http.MethodPost, c.URL, strings.NewReader(string(bytes)))
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
	var document IDDocument
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
	if res.StatusCode == 404 {
		return document.Activity, errors.New("Not Found")
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
