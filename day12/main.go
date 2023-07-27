package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id          int
	NameProject string
	StartDate   string
	EndDate     string
	Duration    string
	Description string
	PostDate    string
	TimePost    string
	Nodejs      bool
	Golang      bool
	ReactJs     bool
	JavaScript  bool
	Image       string
}

var dataProjects = []Project{
	{
		NameProject: "Project pertama",
		StartDate:   "23-07-2023",
		EndDate:     "25-08-2023",
		Duration:    countDuration("23-07-2023", "25-08-2023"),
		Description: "This is the description of project 1",
		PostDate:    time.Now().Format("02-01-2006"),
		TimePost:    time.Now().Format("15:04"),
		Nodejs:      true,
		Golang:      true,
		ReactJs:     false,
		JavaScript:  true,
	},
	{
		NameProject: "Project kedua",
		StartDate:   "23-07-2023",
		EndDate:     "25-07-2023",
		Duration:    countDuration("23-07-2023", "25-07-2023"),
		Description: "This is the description of project 2",
		PostDate:    time.Now().Format("02-01-2006"),
		TimePost:    time.Now().Format("15:04"),
		Nodejs:      false,
		Golang:      true,
		ReactJs:     true,
		JavaScript:  true,
	},
	{
		NameProject: "Project ketiga",
		StartDate:   "23-07-2023",
		EndDate:     "25-07-2024",
		Duration:    countDuration("23-07-2023", "25-07-2024"),
		Description: "This is the description of project 3",
		PostDate:    time.Now().Format("02-01-2006"),
		TimePost:    time.Now().Format("15:04"),
		Nodejs:      true,
		Golang:      true,
		ReactJs:     true,
		JavaScript:  true,
	},
}

func main() {
	e := echo.New()

	e.Static("/Assets", "Assets")

	//List Get
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/index", index)
	e.GET("/myProject", myProject)
	e.GET("/form-project", formProject)
	e.GET("/testimoni", testimonial)
	e.GET("/project-detail/:id", projectDetail)
	e.GET("/edit-project/:id", editProject)

	//List Post
	e.POST("/add-project", addProject)
	e.POST("/edit-project/:id", postEditProject)
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

	data := map[string]interface{}{
		"Projects": dataProjects,
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
				Duration:    data.Duration,
				Description: data.Description,
				PostDate:    data.PostDate,
				TimePost:    data.TimePost,
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

func postEditProject(c echo.Context) error {
	id := c.Param("id")
	idToInt, _ := strconv.Atoi(id)

	addProjectName := c.FormValue("input-projectname")
	addStartDate := c.FormValue("input-startdate")
	addEndDate := c.FormValue("input-endDate")
	addDescription := c.FormValue("input-descripton")
	addNodeJs := c.FormValue("input-nodejs")
	addGolang := c.FormValue("input-golang")
	addReactJs := c.FormValue("input-reactjs")
	addJavascript := c.FormValue("input-javascript")

	startDate, err := time.Parse("2006-01-02", addStartDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	newStartDate := startDate.Format("02-01-2006")

	endtDate, err := time.Parse("2006-01-02", addEndDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	newEndDate := endtDate.Format("02-01-2006")

	var editProject = Project{
		NameProject: addProjectName,
		StartDate:   newStartDate,
		EndDate:     newEndDate,
		Duration:    countDuration(newStartDate, newEndDate),
		Description: addDescription,
		PostDate:    time.Now().Format("02-01-2006"),
		TimePost:    time.Now().Format("15:04"),
		Nodejs:      (addNodeJs == "on"),
		Golang:      (addGolang == "on"),
		ReactJs:     (addReactJs == "on"),
		JavaScript:  (addJavascript == "on"),
	}

	dataProjects[idToInt] = editProject
	return c.Redirect(http.StatusMovedPermanently, "/myProject")
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

	return c.Redirect(http.StatusMovedPermanently, "/myProject")
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

	startDate, err := time.Parse("2006-01-02", addStartDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	newStartDate := startDate.Format("02-01-2006")

	endtDate, err := time.Parse("2006-01-02", addEndDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	newEndDate := endtDate.Format("02-01-2006")

	var newProject = Project{
		NameProject: addProjectName,
		StartDate:   newStartDate,
		EndDate:     newEndDate,
		Duration:    countDuration(newStartDate, newEndDate),
		Description: addDescription,
		PostDate:    time.Now().Format("02-01-2006"),
		TimePost:    time.Now().Format("15:04"),
		Nodejs:      (addNodeJs == "on"),
		Golang:      (addGolang == "on"),
		ReactJs:     (addReactJs == "on"),
		JavaScript:  (addJavascript == "on"),
	}

	dataProjects = append(dataProjects, newProject)

	fmt.Println("input StartDate", newStartDate)
	fmt.Println("input EndDate", newEndDate)
	fmt.Println("input Node Js", addNodeJs)
	fmt.Println("input Golang", addGolang)
	fmt.Println("input ReactJs", addReactJs)
	fmt.Println("input JavaScript", addJavascript)

	return c.Redirect(http.StatusMovedPermanently, "/myProject")

}

func countDuration(d1, d2 string) string {
	date_1, _ := time.Parse("02-01-2006", d1)
	date_2, _ := time.Parse("02-01-2006", d2)

	diff := date_2.Sub(date_1)
	days := int(diff.Hours() / 24)
	weeks := days / 7
	months := days / 30

	if months >= 12 {
		return strconv.Itoa(months/12) + " Tahun"
	}
	if months > 0 {
		return strconv.Itoa(months) + " Bulan"
	}
	if weeks > 0 {
		return strconv.Itoa(weeks) + " Minggu"
	}

	return strconv.Itoa(days) + " Hari"
}

func editProject(c echo.Context) error {
	id := c.Param("id")
	idToInt, _ := strconv.Atoi(id)

	var ProjectDetail = Project{}

	for index, data := range dataProjects {
		if idToInt == index {
			ProjectDetail = Project{
				NameProject: data.NameProject,
				StartDate:   data.StartDate,
				EndDate:     data.EndDate,
				Duration:    data.Duration,
				Description: data.Description,
				PostDate:    data.PostDate,
				TimePost:    data.TimePost,
				Nodejs:      data.Nodejs,
				Golang:      data.Golang,
				ReactJs:     data.ReactJs,
				JavaScript:  data.JavaScript,
			}
		}
	}
	fmt.Println("Received data:", ProjectDetail.NameProject, ProjectDetail.StartDate)

	data := map[string]interface{}{
		"Id":      id,
		"Project": ProjectDetail,
	}
	tmpl, err := template.ParseFiles("view/Edit-MyProject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), data)

}
