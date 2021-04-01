package jsonTypes

import "time"

type UserComment struct {
	UserName string `json:"userName"`
	Comment  string `json:"comment"`
}

type MovieSearchResponse struct {
	MovieName         string         `json:"movie_name"`
	AverageRating     float64        `json:"avg_rating"`
	Comments          []*UserComment `json:"comments"`
	TotalRatingsCount int            `json:"allRatingsCount"`
	PublishedAt       time.Time      `json:"publishedAt"`
}
