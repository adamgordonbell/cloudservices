package server

import (
	"database/sql"
	"errors"
	"log"
	"sync"

	api "github.com/adamgordonbell/cloudservices/activity-log"
	// need to get sqlite working
	_ "github.com/mattn/go-sqlite3"
)

const create string = `
		CREATE TABLE IF NOT EXISTS activities (
		id INTEGER NOT NULL PRIMARY KEY,
		time DATETIME NOT NULL,
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
	res, err := c.db.Exec("INSERT INTO activities VALUES(NULL,?,?);", activity.Time, activity.Description)
	if err != nil {
		return 0, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	log.Printf("Added %v as %d", activity, id)
	return int(id), nil
}

var ErrIDNotFound = errors.New("Id not found")

func (c *Activities) Retrieve(id int) (api.Activity, error) {
	log.Printf("Getting %d", id)

	// Query DB row based on ID
	row := c.db.QueryRow("SELECT * FROM activities WHERE id=?", id)

	// Parse row into Interval struct
	activity := api.Activity{}
	var err error
	if err = row.Scan(&activity.ID, &activity.Time, &activity.Description); err == sql.ErrNoRows {
		log.Printf("Id not found")
		return api.Activity{}, ErrIDNotFound
	}
	return activity, err
}

func (c *Activities) List(offset int) ([]api.Activity, error) {
	log.Printf("Getting list from offset %d\n", offset)

	// Query DB row based on ID
	rows, err := c.db.Query("SELECT * FROM activities WHERE ID > ? ORDER BY id DESC LIMIT 100", offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := []api.Activity{}
	for rows.Next() {
		i := api.Activity{}
		err = rows.Scan(&i.ID, &i.Time, &i.Description)
		if err != nil {
			return nil, err
		}
		data = append(data, i)
	}
	return data, nil
}
