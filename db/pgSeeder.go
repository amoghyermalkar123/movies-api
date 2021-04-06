package db

func (d *Db) DeleteTestMovieData(name string) {
	_, _ = d.dba.Exec(`delete from movies where name = ?`, name)
}
