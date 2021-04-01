package jsonTypes

type User struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Usype    string `json:"user_type"`
}
