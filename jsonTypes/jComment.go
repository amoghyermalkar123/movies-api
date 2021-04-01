package jsonTypes

type Comment struct {
	UserID  int64  `json:"user_id"`
	MovieID int64  `json:"movie_id"`
	Comment string `json:"comment"`
}
