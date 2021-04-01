package mappers

import (
	"errors"
	"math"
	"movieserver/jsonTypes"
	"movieserver/types"
)

func MapDBMovieToJsonMovieResponse(movieData []*types.MovieSearchDetails) (*jsonTypes.MovieSearchResponse, error) {
	if len(movieData) == 0 {
		return nil, errors.New("failed mapping")
	}
	finalResponse := &jsonTypes.MovieSearchResponse{}
	userComments := make([]*jsonTypes.UserComment, 0)

	finalResponse.MovieName = movieData[0].MovieName
	finalResponse.PublishedAt = movieData[0].PublishedAt

	num := float64(len(movieData))
	count := 0.0

	for _, item := range movieData {
		userComments = append(userComments, &jsonTypes.UserComment{
			UserName: item.Commenter,
			Comment:  item.Comment,
		})

		count = count + float64(item.Rating)
	}

	averageRating := count / num
	roundedAvgRating := math.Round(averageRating*100) / 100

	finalResponse.AverageRating = roundedAvgRating
	finalResponse.TotalRatingsCount = int(num)
	finalResponse.Comments = userComments

	return finalResponse, nil
}
