package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/Assets", "Assets")

	e.GET("/", handlers.home)

	e.GET("/contact", contact)

	e.GET("/index", index)

	e.GET("/myProject", myProject)

	e.GET("/form-project", formProject)

	e.GET("/testimoni", testimonial)

	e.POST("add-project", addProject)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

// handle

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

func formProject(c echo.Context) error {
	tmpl, err := template.ParseFiles("view/Form-MyProject.html")

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

func addProject(c echo.Context) error {
	ProjectName := c.FormValue("input-projectname")
	StartDate := c.FormValue("input-startdate")
	EndDate := c.FormValue("input-endDate")
	Description := c.FormValue("input-descripton")

	fmt.Println("Project Name :", ProjectName)
	fmt.Println("Start Date :", StartDate)
	fmt.Println("End Date :", EndDate)
	fmt.Println("Description :", Description)

	return c.Redirect(http.StatusMovedPermanently, "/myProject")

}
