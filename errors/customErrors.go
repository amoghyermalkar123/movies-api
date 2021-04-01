package errors

var FailedError = map[string]string{
	"Error":  "could not calculate movies data",
	"Status": "Failed",
}

var BindError = map[string]string{
	"Error":  "bad request payload",
	"Status": "Failed",
}

var EmptyRequestError = map[string]string{
	"Error":  "the search term(s) cannot be empty",
	"Status": "Failed",
}

var NoDataFoundError = map[string]string{
	"Error":  "no data found",
	"Status": "Failed",
}

var AuthError = map[string]string{
	"Error":  "user authentication failed",
	"Status": "Failed",
}

var RateMovieError = map[string]string{
	"Error":  "failed to rate the movie",
	"Status": "Failed",
}

var CommentMovieError = map[string]string{
	"Error":  "failed to comment on the movie",
	"Status": "Failed",
}

var ContentCreationError = map[string]string{
	"Error":  "failed to create new content",
	"Status": "Failed",
}
