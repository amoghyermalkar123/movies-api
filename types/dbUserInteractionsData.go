package types

type UserInteractionData struct {
	Name    string `pg:"name"`
	Rating  int64  `pg:"rating"`
	Comment string `pg:"comment"`
}
