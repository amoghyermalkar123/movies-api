package jsonTypes

type RateMovie struct {
	UserID  int64 `json:"user_id"`
	MovieID int64 `json:"movie_id"`
	Rating  int   `json:"rating"`
}
