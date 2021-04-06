package handlers

import (
	e "movieserver/errors"
	"movieserver/jsonTypes"
	"movieserver/mappers"
	m "movieserver/messages"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (h *Handler) Authenticate(c echo.Context) error {
	userDetails := &jsonTypes.User{}

	if err := c.Bind(userDetails); err != nil {
		return c.JSON(http.StatusInternalServerError, e.BindError)
	}

	isAuthenticated := h.Auth.AuthenticateUserDetails(userDetails)

	if isAuthenticated {
		return c.JSON(http.StatusOK, m.Success("Authenticated"))
	}
	return c.JSON(http.StatusInternalServerError, e.AuthError)
}

func (h *Handler) SearchUserInteractedMovies(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	isAuthenticated := h.Auth.UserCheck(int64(id))

	if !isAuthenticated {
		return c.JSON(http.StatusInternalServerError, e.AuthError)
	}

	response, err := h.Db.GetUserInteractedMovies(int64(id))

	if err != nil {
		if err == e.ErrNoDataFound {
			return c.JSON(http.StatusInternalServerError, e.NoDataFoundError)
		}
		return err
	}

	return c.JSON(http.StatusOK, m.Success(response))

}

func (h *Handler) RateMovie(c echo.Context) error {
	ratingData := &jsonTypes.RateMovie{}

	if err := c.Bind(ratingData); err != nil {
		return c.JSON(http.StatusInternalServerError, e.BindError)
	}

	isAuthenticated := h.Auth.UserCheck(ratingData.UserID)

	if !isAuthenticated {
		return c.JSON(http.StatusInternalServerError, e.AuthError)
	}

	if ratingData.MovieID == 0 || ratingData.UserID == 0 || ratingData.Rating == 0 {
		return c.JSON(http.StatusInternalServerError, e.EmptyRequestError)
	}

	dbRateMovieData := mappers.MapJsonRateMovieToPgrateMovie(ratingData)

	err := h.Db.RateMovie(dbRateMovieData)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, e.RateMovieError)
	}

	return c.JSON(http.StatusOK, m.Success("Movie Rated Successfully"))
}

func (h *Handler) CommentOnMovie(c echo.Context) error {
	commentData := &jsonTypes.Comment{}

	if err := c.Bind(commentData); err != nil {
		return c.JSON(http.StatusInternalServerError, e.BindError)
	}

	isAuthenticated := h.Auth.UserCheck(commentData.UserID)

	if !isAuthenticated {
		return c.JSON(http.StatusInternalServerError, e.AuthError)
	}

	if commentData.MovieID == 0 || commentData.UserID == 0 || commentData.Comment == "" {
		return c.JSON(http.StatusInternalServerError, e.EmptyRequestError)
	}

	dbComment := mappers.MapJsonCommentToPgComment(commentData)

	err := h.Db.CommentOnMovie(dbComment)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, e.CommentMovieError)
	}

	return c.JSON(http.StatusOK, m.Success("Commented On The Movie Successfully"))

}
