package types

import "time"

type Movie struct {
	tableName   struct{}  `pg:"movies"`
	ID          int64     `pg:"id"`
	Name        string    `pg:"name"`
	PublishedAt time.Time `pg:"publishedat"`
}
