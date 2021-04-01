package main

import (
	"log"
	"movieserver/handlers"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	h := &handlers.Handler{}

	// Routes

	e.POST("/auth/", h.Authenticate)
	e.GET("/search_movie/:name", h.SearchForMovies)
	e.GET("/search_user_movies/:id", h.SearchUserInteractedMovies)
	e.POST("/rate/", h.RateMovie)
	e.POST("/comment/", h.CommentOnMovie)
	e.POST("/create_content/", h.AddNewContent)

	go setEnv()
	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}

func setEnv() {
	workingdir, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	os.Setenv("configPath", workingdir+"/config/")
}
