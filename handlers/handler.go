package handlers

import (
	"movieserver/auth"
	"movieserver/db"
)

type Handler struct {
	Auth auth.Auth
	Db   *db.Db
}
