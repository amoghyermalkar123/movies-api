package db

import (
	e "movieserver/errors"
	"movieserver/types"

	"github.com/go-pg/pg"
)

func GetUserDetails(userID int64, userType string) (*types.User, error) {
	db, err := getDbSession()

	if err != nil {
		return nil, e.ErrDatabaseErrored
	}

	dets := &types.User{}
	err = db.Model(dets).
		Where("id = ?", userID).
		Where("user_type = ?", userType).
		Select(dets)

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, e.ErrNoDataFound
		}
		return nil, err
	}
	return dets, nil
}

func CheckUserDetails(ID int64, email, password string) error {
	db, err := getDbSession()

	if err != nil {
		return e.ErrDatabaseErrored
	}

	dets := &types.User{}
	err = db.Model(dets).
		Where("id = ?", ID).
		Where("email = ?", email).
		Where("password = ?", password).
		Select(dets)

	if err != nil {
		if err == pg.ErrNoRows {
			return err
		}
		return err
	}
	return nil
}
