package db

import (
	e "movieserver/errors"
	"movieserver/types"

	"github.com/go-pg/pg"
)

func CheckAdminDetails(adminID int64, email, password string) error {
	db, err := getDbSession()

	if err != nil {
		return e.ErrDatabaseErrored
	}

	dets := &types.User{}
	err = db.Model(dets).
		Where("id = ?", adminID).
		Where("email = ?", email).
		Where("password = ?", password).
		Where("user_type = ?", "admin").
		Select(dets)

	if err != nil {
		if err == pg.ErrNoRows {
			return e.ErrNoDataFound
		}
		return err
	}
	return nil
}

func CreateContent(data *types.Movie) error {
	db, err := getDbSession()

	if err != nil {
		return e.ErrDatabaseErrored
	}

	err = db.Insert(data)

	if err != nil {
		return err
	}
	return nil
}
