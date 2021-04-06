package db

import (
	e "movieserver/errors"
	"movieserver/types"

	"github.com/go-pg/pg"
)

func (d *Db) GetUserDetails(userID int64, userType string) (*types.User, error) {
	dets := &types.User{}
	err := d.dba.Model(dets).
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

func (d *Db) CheckUserDetails(ID int64, email, password string) error {

	dets := &types.User{}
	err := d.dba.Model(dets).
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
