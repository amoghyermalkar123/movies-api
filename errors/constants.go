package errors

type constError string

func (err constError) Error() string {
	return string(err)
}

const (
	ErrDatabaseErrored = constError("database errored out")
	ErrNoDataFound     = constError("no data found")
)
