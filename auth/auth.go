package auth

import (
	"movieserver/db"
	"movieserver/jsonTypes"
)

type Auth struct {
	Db *db.Db
}

type isAuthenticated bool

func (a *Auth) AuthenticateUserDetails(userDetails *jsonTypes.User) isAuthenticated {
	err := a.Db.CheckUserDetails(userDetails.ID, userDetails.Email, userDetails.Password)
	return err == nil
}

func (a *Auth) UserCheck(ID int64) bool {
	_, err := a.Db.GetUserDetails(ID, "user")
	return err == nil
}

func (a *Auth) AuthenticateAdmin(adminDetails *jsonTypes.User) isAuthenticated {
	err := a.Db.CheckAdminDetails(adminDetails.ID, adminDetails.Email, adminDetails.Password)
	return err == nil
}

func (a *Auth) AdminCheck(ID int64) bool {
	_, err := a.Db.GetUserDetails(ID, "admin")
	return err == nil
}
