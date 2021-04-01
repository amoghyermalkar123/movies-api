package db

import "log"

func DeleteTestMovieData(name string) {
	pg, err := getDbSession()
	if err != nil {
		log.Println(err)
	}

	_, _ = pg.Exec(`delete from movies where name = ?`, name)
}
