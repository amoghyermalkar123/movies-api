package types

type User struct {
	tableName struct{} `pg:"users,alias:users"`
	ID        int64    `pg:"id"`
	FullName  string   `pg:"full_name"`
	Email     string   `pg:"email"`
	Password  string   `pg:"password"`
	UserType  string   `pg:"user_type"`
}
