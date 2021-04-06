package db

import (
	e "movieserver/errors"
	"movieserver/types"

	"github.com/go-pg/pg"
)

func (d *Db) CheckAdminDetails(adminID int64, email, password string) error {

	dets := &types.User{}
	err := d.dba.Model(dets).
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

func (d *Db) CreateContent(data *types.Movie) error {
	movieDoesExist := &types.Movie{}

	_, _ = d.dba.Model(data).Query(movieDoesExist, `
	select name from movies where name = ?
	`, data.Name)

	if movieDoesExist.Name != "" {
		return e.ErrDataAlreadyExists
	}

	err := d.dba.Insert(data)

	if err != nil {
		return err
	}
	return nil
}
