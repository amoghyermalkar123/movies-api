package db

import "github.com/go-pg/pg"

type Db struct {
	dba *pg.DB
}
