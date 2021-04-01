package mappers

import (
	"movieserver/jsonTypes"
	"movieserver/types"
)

func MapJMovieToDbMovieData(data *jsonTypes.MovieData) *types.Movie {
	return &types.Movie{
		Name:        data.Name,
		PublishedAt: data.PublishedAt,
	}
}
