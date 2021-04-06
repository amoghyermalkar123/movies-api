package db

import (
	"errors"
	"log"
	e "movieserver/errors"
	"movieserver/types"

	"github.com/go-pg/pg"
)

func (d *Db) SearchMovieByName(name string) (*types.Movie, error) {
	result := &types.Movie{}

	_, err := d.dba.Model(result).Query(result, `select * from movies where name ~* ?`, name)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d *Db) GetMovieSearchDetails(name string) ([]*types.MovieSearchDetails, error) {
	result := &types.Movie{}

	response := make([]*types.MovieSearchDetails, 0)

	_, err := d.dba.Model(result).Query(&response, `
		select name as movie_name, published_at, full_name as commenter,rating, comment  from movies m join user_interactions ui on
		m.id = ui.movie_id
		and name = ?
		join users
		on users.id = ui.user_id;
	`, name)

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, e.ErrNoDataFound
		}
		return nil, err
	}

	return response, nil
}

func (d *Db) GetUserInteractedMovies(userID int64) ([]*types.UserInteractionData, error) {
	if userID == 0 {
		return nil, errors.New("user id must not be nil")
	}

	result := make([]*types.UserInteractionData, 0)

	_, err := d.dba.Query(&result, `
		select name, rating, comment from movies m 
		join user_interactions ui 
		on m.id = ui.movie_id
		and ui.user_id = ?;
	`, userID)

	if err != nil {
		if err == pg.ErrNoRows {
			return nil, e.ErrNoDataFound
		}
		return nil, err
	}

	return result, nil
}

func (d *Db) RateMovie(req *types.UserInteraction) error {

	if checkUserInteraction(d.dba, req.UserID, req.MovieID) {
		_, err := d.dba.Exec(`
			update user_interactions
			set rating = ?
			where user_id = ?
			and movie_id = ?
		`, req.Rating, req.UserID, req.MovieID)

		if err != nil {
			return e.ErrDatabaseErrored
		}
	} else {
		_, err := d.dba.Exec(`
			insert into user_interactions(user_id, movie_id, rating)
			values(?, ?, ?)
		`, req.UserID, req.MovieID, req.Rating)

		if err != nil {
			return e.ErrDatabaseErrored
		}
	}

	return nil
}

func checkUserInteraction(db *pg.DB, userID, movieID int64) bool {
	model := &types.UserInteraction{}
	err := db.Model(model).
		Where("user_id = ?", userID).
		Where("movie_id = ?", movieID).
		Select(model)

	if err != nil {
		return false
	}

	if model.ID == 0 {
		return false
	}
	return true
}

func (d *Db) CommentOnMovie(req *types.UserInteraction) error {

	if checkUserInteraction(d.dba, req.UserID, req.MovieID) {
		log.Println("true")
		_, err := d.dba.Exec(`
			update user_interactions
			set comment = ?
			where user_id = ?
			and movie_id = ?
		`, req.Comment, req.UserID, req.MovieID)

		if err != nil {
			return e.ErrDatabaseErrored
		}
	} else {
		_, err := d.dba.Exec(`
			insert into user_interactions(user_id, movie_id, comment)
			values(?, ?, ?)
		`, req.UserID, req.MovieID, req.Comment)

		if err != nil {
			return e.ErrDatabaseErrored
		}
	}

	return nil
}
