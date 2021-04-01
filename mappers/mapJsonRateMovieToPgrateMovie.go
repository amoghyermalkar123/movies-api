package mappers

import (
	"movieserver/jsonTypes"
	"movieserver/types"
)

func MapJsonRateMovieToPgrateMovie(req *jsonTypes.RateMovie) *types.UserInteraction {
	return &types.UserInteraction{
		UserID:  req.UserID,
		MovieID: req.MovieID,
		Rating:  int64(req.Rating),
	}
}
