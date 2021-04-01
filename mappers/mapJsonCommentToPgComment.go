package mappers

import (
	"movieserver/jsonTypes"
	"movieserver/types"
)

func MapJsonCommentToPgComment(req *jsonTypes.Comment) *types.UserInteraction {

	return &types.UserInteraction{
		UserID:  req.UserID,
		MovieID: req.MovieID,
		Comment: req.Comment,
	}
}
