package handlers

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func home(c echo.Context) error {
	tmpl, err := template.ParseFiles("view/Home.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return tmpl.Execute(c.Response(), nil) //execute = respons apa yang mau dipanggil
}
