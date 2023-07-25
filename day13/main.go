package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	connection "personal-web/Connection"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id          int
	NameProject string
	StartDate   string
	EndDate     string
	Duration    string
	Description string
	Nodejs      bool
	Golang      bool
	ReactJs     bool
	JavaScript  bool
	Image       string
}

var dataProjects = []Project{
	{
		NameProject: "Project pertama",
		StartDate:   "23/07/2023",
		EndDate:     "25/08/2023",
		Description: "This is the description of project 1",
		Nodejs:      true,
		Golang:      true,
		ReactJs:     false,
		JavaScript:  true,
	},
	{
		NameProject: "Project kedua",
		StartDate:   "23/07/2023",
		EndDate:     "25/08/2023",
		Description: "This is the description of project 2",
		Nodejs:      false,
		Golang:      true,
		ReactJs:     true,
		JavaScript:  true,
	},
}

func main() {
	e := echo.New()
	connection.DatabaseConnect()

	e.Static("/Assets", "Assets")

	//List Get
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/index", index)
	e.GET("/myProject", myProject)
	e.GET("/form-project", formProject)
	e.GET("/testimoni", testimonial)
	e.GET("/project-detail/:id", projectDetail)

	//List Post
	e.POST("/add-project", addProject)
	e.POST("/delete-project/:id", deleteProject)

	//Server
	e.Logger.Fatal(e.Start("localhost:5000"))
}

// handler

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

	dataQuery, errQuery := connection.Conn.Query(context.Background(), "SELECT name, description FROM tb_projects")

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var resultProject []Project

	for dataQuery.Next() {
		var each = Project{}

		err := dataQuery.Scan(&each.NameProject, &each.Description)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		resultProject = append(resultProject, each)
	}

	data := map[string]interface{}{
		"Projects": resultProject,
		// "db_Project": resultProject,
	}
	return tmpl.Execute(c.Response(), data) //execute = respons apa yang mau dipanggil
}

func projectDetail(c echo.Context) error {
	id := c.Param("id")

	tmpl, err := template.ParseFiles("view/Project-Detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	idToInt, _ := strconv.Atoi(id)

	var ProjectDetail = Project{}

	for index, data := range dataProjects {
		if idToInt == index {
			ProjectDetail = Project{
				NameProject: data.NameProject,
				StartDate:   data.StartDate,
				EndDate:     data.EndDate,
				Description: data.Description,
				Nodejs:      data.Nodejs,
				Golang:      data.Golang,
				ReactJs:     data.ReactJs,
				JavaScript:  data.JavaScript,
			}
		}
	}

	data := map[string]interface{}{
		"Id":      id,
		"Project": ProjectDetail,
	}

	return tmpl.Execute(c.Response(), data) //execute = respons apa yang mau dipanggil
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

func deleteProject(c echo.Context) error {
	id := c.Param("id")
	idToInt, _ := strconv.Atoi(id)

	dataProjects = append(dataProjects[:idToInt], dataProjects[idToInt+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/myProject") //execute = respons apa yang mau dipanggil
}

func addProject(c echo.Context) error {
	addProjectName := c.FormValue("input-projectname")
	addStartDate := c.FormValue("input-startdate")
	addEndDate := c.FormValue("input-endDate")
	addDescription := c.FormValue("input-descripton")
	addNodeJs := c.FormValue("input-nodejs")
	addGolang := c.FormValue("input-golang")
	addReactJs := c.FormValue("input-reactjs")
	addJavascript := c.FormValue("input-javascript")

	var newProject = Project{
		NameProject: addProjectName,
		StartDate:   addStartDate,
		EndDate:     addEndDate,
		Description: addDescription,
		Nodejs:      (addNodeJs == "on"),
		Golang:      (addGolang == "on"),
		ReactJs:     (addReactJs == "on"),
		JavaScript:  (addJavascript == "on"),
	}

	dataProjects = append(dataProjects, newProject)

	fmt.Println("input Node Js", addNodeJs)
	fmt.Println("input Golang", addGolang)
	fmt.Println("input ReactJs", addReactJs)
	fmt.Println("input JavaScript", addJavascript)

	return c.Redirect(http.StatusMovedPermanently, "/myProject")

}
