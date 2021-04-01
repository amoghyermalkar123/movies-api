package jsonTypes

type Movie struct {
	Name       string        `json:"name"`
	UserRating int           `json:"rating"`
	Comments   []UserComment `json:"comments"`
}
