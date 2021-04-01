package types

import "time"

type MovieSearchDetails struct {
	MovieName   string    `pg:"name"`
	Rating      int       `pg:"rating"`
	Comment     string    `pg:"comment"`
	Commenter   string    `pg:"commenter"`
	PublishedAt time.Time `pg:"published_at"`
}
