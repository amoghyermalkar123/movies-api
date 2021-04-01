package jsonTypes

import "time"

type MovieData struct {
	AdminID     int64     `json:"admin_id"`
	Name        string    `json:"name"`
	PublishedAt time.Time `json:"publishedat"`
}
