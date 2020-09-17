package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Request is Responce
type Request struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Input  string `json:"input"`
}

// Responce is Responce
type Responce struct {
	Status string `json:"status"`
	Output string `json:"output"`
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong\n")
	})
	e.POST("/translation", postTranslationHandler)
	e.Start(":4000")
}

func postTranslationHandler(c echo.Context) error {
	req := Request{}
	c.Bind(&req)
	return c.JSON(http.StatusOK, Responce{Status: "success", Output: "This is a English sentence."})
}
