package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type application struct {
	app *echo.Echo
}

func newApp() *application {
	app := echo.New()

	app.GET("/user/:id", getUser)
	app.GET("/add", add)

	return &application{
		app: app,
	}
}

func (a *application) start() {
	a.app.Logger.Fatal(a.app.Start("localhost:8888"))
}

func add(c echo.Context) error {
	return c.JSON(http.StatusOK, 2)
}

func getUser(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:  "TomsFavouriteDrink",
		Value: "Beer",
		Path:  "/",
	})

	id := c.Param("id")

	if id == "1234" {
		return c.JSON(200, &user{ID: id, Name: "Andy"})
	}
	return c.NoContent(404)
}

func main() {
	newApp().start()
}
