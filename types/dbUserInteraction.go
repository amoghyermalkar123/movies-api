package types

type UserInteraction struct {
	tableName struct{} `pg:"user_interactions"`
	ID        int64    `pg:"id"`
	UserID    int64    `pg:"user_id"`
	MovieID   int64    `pg:"movie_id"`
	Rating    int64    `pg:"rating"`
	Comment   string   `pg:"comment"`
}
