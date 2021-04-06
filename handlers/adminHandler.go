package handlers

import (
	e "movieserver/errors"
	"movieserver/jsonTypes"
	"movieserver/mappers"
	m "movieserver/messages"
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) AddNewContent(c echo.Context) error {
	movieData := &jsonTypes.MovieData{}

	if err := c.Bind(movieData); err != nil {
		return c.JSON(http.StatusInternalServerError, e.BindError)
	}

	isAuthenticated := h.Auth.AdminCheck(movieData.AdminID)

	if !isAuthenticated {
		return c.JSON(http.StatusInternalServerError, e.AuthError)
	}

	dbMovie := mappers.MapJMovieToDbMovieData(movieData)
	err := h.Db.CreateContent(dbMovie)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, e.ContentCreationError)
	}

	return c.JSON(http.StatusOK, m.Success("Content Created Successfully"))
}
