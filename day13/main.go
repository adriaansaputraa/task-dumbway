package main

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/Assets", "Assets")

	e.GET("/home", home)

	e.GET("/contact", contact)

	e.GET("/index", index)

	e.GET("/myProject", myProject)

	e.GET("/testimonial", testimonial)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

// handle

func home(c echo.Context) error {
	tmpl, err := template.ParseFiles("view/Home.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return tmpl.Execute(c.Response(), nil) //execute = respons apa yang mau dipanggil
}

func contact(c echo.Context) error {
	tmpl, err := template.ParseFiles("view/Contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return tmpl.Execute(c.Response(), nil) //execute = respons apa yang mau dipanggil
}

func index(c echo.Context) error {
	tmpl, err := template.ParseFiles("view/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return tmpl.Execute(c.Response(), nil) //execute = respons apa yang mau dipanggil
}

func myProject(c echo.Context) error {
	tmpl, err := template.ParseFiles("view/MyProject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return tmpl.Execute(c.Response(), nil) //execute = respons apa yang mau dipanggil
}

func testimonial(c echo.Context) error {
	tmpl, err := template.ParseFiles("view/Testimoni.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return tmpl.Execute(c.Response(), nil) //execute = respons apa yang mau dipanggil
}
