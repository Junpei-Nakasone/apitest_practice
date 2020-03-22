package main

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

func TestGetUser_CookieMatching(t *testing.T) {
	apitest.New().
		Handler(newApp().app).
		Get("/user/1234").
		Expect(t).
		Cookies(apitest.NewCookie("TomsFavouriteDrink").
			Value("Beer").
			Path("/")).
		Status(http.StatusOK).
		End()
}

func TestGetUser_Success(t *testing.T) {
	apitest.New().
		Handler(newApp().app).
		Get("/user/1234").
		Expect(t).
		Body(`{"id":"1234", "name":"Andy"}`).
		Status(http.StatusOK).
		End()
}
