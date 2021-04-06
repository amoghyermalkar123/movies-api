package handlers

import (
	e "movieserver/errors"
	"movieserver/mappers"
	m "movieserver/messages"
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) SearchForMovies(c echo.Context) error {
	name := c.Param("name")

	if name == "" {
		return c.JSON(http.StatusInternalServerError, e.EmptyRequestError)
	}

	response, err := h.Db.GetMovieSearchDetails(name)

	if err != nil {
		if err == e.ErrNoDataFound {
			return c.JSON(http.StatusInternalServerError, e.NoDataFoundError)
		}
		return err
	}

	retVal, err := mappers.MapDBMovieToJsonMovieResponse(response)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, e.FailedError)
	}
	return c.JSON(http.StatusOK, m.Success(retVal))
}
