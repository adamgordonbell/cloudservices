package server

import (
	"database/sql"
	"errors"
	"sync"

	api "github.com/adamgordonbell/cloudservices/activity-log"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

const create string = `
		CREATE TABLE [activities] (
		id INTEGER NOT NULL PRIMARY KEY,
		time TEXT,
		description TEXT
		);
`
const file string = "activities.db"

type Activities struct {
	db *sql.DB
	mu sync.Mutex
}

func NewActivities() (*Activities, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(create); err != nil {
		return nil, err
	}
	return &Activities{
		db: db,
	}, nil
}

func (c *Activities) Insert(activity api.Activity) (int, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	//do we need a prepare or just db.Exec works?
	insStmt, err := c.db.Prepare("INSERT INTO activities VALUES(NULL,?,?);")
	if err != nil {
		return 0, err
	}
	defer insStmt.Close()

	res, err := insStmt.Exec(activity.Time, activity.Description)
	if err != nil {
		return 0, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}

	log.Printf("Added %v as %s", activity, id)
	return int(id), nil
}

var ErrIDNotFound = errors.New("Id not found")

func (c *Activities) Retrieve(id int) (api.Activity, error) {
	log.Printf("Getting %d", id)
	c.mu.Lock()
	defer c.mu.Unlock()

	// Query DB row based on ID
	row := c.db.QueryRow("SELECT * FROM interval WHERE id=?", id)

	// Parse row into Interval struct
	activity := api.Activity{}
	var err error
	if err := row.Scan(&activity.ID, &activity.Time, &activity.Description); err == sql.ErrNoRows {
		log.Printf("Id not found")
		return api.Activity{}, ErrIDNotFound
	}
	return activity, err
}
