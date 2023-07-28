package handler

import (
	"context"
	"html/template"
	"net/http"
	connection "personal-web/Connection"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id           int
	Project_Name string
	Start_Date   time.Time
	End_Date     time.Time
	Duration     string
	Description  string
	Post_Date    string
	Time_Post    string
	Technologies []string
	Node_js      bool
	Golang       bool
	React_Js     bool
	Java_Script  bool
	Image        string
}

var data_Projects = []Project{}

// ------------------------------------------------------------------------HOME--------------------------------------------------------//
func Home(c echo.Context) error {

	tmpl, err := template.ParseFiles("view/Home.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return tmpl.Execute(c.Response(), nil)
}

//------------------------------------------------------------------------CONTACT--------------------------------------------------------//

func Contact(c echo.Context) error {
	tmpl, err := template.ParseFiles("view/Contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return tmpl.Execute(c.Response(), nil)
}

//------------------------------------------------------------------------INDEX--------------------------------------------------------//

func Index(c echo.Context) error {
	tmpl, err := template.ParseFiles("view/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return tmpl.Execute(c.Response(), nil)
}

//-------------------------------------------------------GET ADD PROJECT FOR FRONT END-------------------------------------------------------//

func Add_Project(c echo.Context) error {
	tmpl, err := template.ParseFiles("view/Add-Project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	dataQuery, errQuery := connection.Conn.Query(context.Background(), "SELECT id, name, start_date, end_date, description, technologies, image FROM tb_projects")

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var resultProject []Project

	for dataQuery.Next() {

		var each = Project{}

		err := dataQuery.Scan(&each.Id, &each.Project_Name, &each.Start_Date, &each.End_Date, &each.Description, &each.Technologies, &each.Image)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		each.Duration = Count_Duration(each.Start_Date, each.End_Date)
		if checkValue(each.Technologies, "nodejs") {
			each.Node_js = true
		}
		if checkValue(each.Technologies, "golang") {
			each.Golang = true
		}
		if checkValue(each.Technologies, "reactjs") {
			each.React_Js = true
		}
		if checkValue(each.Technologies, "javascript") {
			each.Java_Script = true
		}

		each.Post_Date = time.Now().Format("02-01-2006")
		each.Time_Post = time.Now().Format("15:04")

		resultProject = append(resultProject, each)
	}

	data := map[string]interface{}{
		"Projects": resultProject,
	}
	return tmpl.Execute(c.Response(), data)
}

//-------------------------------------------------------POST ADD PROJECT FOR BACK END----------------------------------------------------//

func Post_Project(c echo.Context) error {
	add_Project_Name := c.FormValue("input-projectname")
	add_Start_Date := c.FormValue("input-startdate")
	add_End_Date := c.FormValue("input-endDate")
	add_Description := c.FormValue("input-descripton")
	add_Node_Js := c.FormValue("input-nodejs")
	add_Golang := c.FormValue("input-golang")
	add_React_Js := c.FormValue("input-reactjs")
	add_Javascript := c.FormValue("input-javascript")

	start_Date, err := time.Parse("2006-01-02", add_Start_Date)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	end_Date, err := time.Parse("2006-01-02", add_End_Date)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var new_Project = Project{
		Project_Name: add_Project_Name,
		Start_Date:   start_Date,
		End_Date:     end_Date,
		Duration:     Count_Duration(start_Date, end_Date),
		Description:  add_Description,
		Post_Date:    time.Now().Format("02-01-2006"),
		Time_Post:    time.Now().Format("15:04"),
		Node_js:      (add_Node_Js == "on"),
		Golang:       (add_Golang == "on"),
		React_Js:     (add_React_Js == "on"),
		Java_Script:  (add_Javascript == "on"),
	}

	data_Projects = append(data_Projects, new_Project)

	return c.Redirect(http.StatusMovedPermanently, "/add-Project")

}

//---------------------------------------------------------------FORM PROJECT---------------------------------------------------------//

func Form_Project(c echo.Context) error {
	tmpl, err := template.ParseFiles("view/Form-Project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return tmpl.Execute(c.Response(), nil)
}

//----------------------------------------------------------------TESTIMONI-----------------------------------------------------------//

func Testimoni(c echo.Context) error {
	tmpl, err := template.ParseFiles("view/Testimoni.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return tmpl.Execute(c.Response(), nil)
}

//----------------------------------------------------------------PROJECT DETAIL-----------------------------------------------------------//

func Project_Detail(c echo.Context) error {
	id := c.Param("id")
	idToInt, _ := strconv.Atoi(id)

	tmpl, err := template.ParseFiles("view/Project-Detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var ProjectDetail = Project{}

	for index, data := range data_Projects {
		if idToInt == index {
			ProjectDetail = Project{
				Project_Name: data.Project_Name,
				Start_Date:   data.Start_Date,
				End_Date:     data.End_Date,
				Duration:     data.Duration,
				Description:  data.Description,
				Post_Date:    data.Post_Date,
				Time_Post:    data.Time_Post,
				Node_js:      data.Node_js,
				Golang:       data.Golang,
				React_Js:     data.React_Js,
				Java_Script:  data.Java_Script,
			}
		}
	}

	data := map[string]interface{}{
		"Id":      id,
		"Project": ProjectDetail,
	}
	return tmpl.Execute(c.Response(), data)
}

//-------------------------------------------------------DELETE PROJECT----------------------------------------------------//

func Delete_Project(c echo.Context) error {
	id := c.Param("id")
	idToInt, _ := strconv.Atoi(id)

	data_Projects = append(data_Projects[:idToInt], data_Projects[idToInt+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/add-Project") //execute = respons apa yang mau dipanggil
}

//-----------------------------------------------------COUNT DURATION----------------------------------------------------//

func Count_Duration(d1, d2 time.Time) string {

	diff := d2.Sub(d1)
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

//-----------------------------------------------------POST EDIT PROJECT----------------------------------------------------//

func Post_Edit_Project(c echo.Context) error {
	id := c.Param("id")
	idToInt, _ := strconv.Atoi(id)

	add_Project_Name := c.FormValue("input-projectname")
	add_Start_Date := c.FormValue("input-startdate")
	add_End_Date := c.FormValue("input-endDate")
	add_Description := c.FormValue("input-descripton")
	add_Node_Js := c.FormValue("input-nodejs")
	add_Golang := c.FormValue("input-golang")
	add_React_Js := c.FormValue("input-reactjs")
	add_Javascript := c.FormValue("input-javascript")

	start_Date, err := time.Parse("2006-01-02", add_Start_Date)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	end_Date, err := time.Parse("2006-01-02", add_End_Date)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var edit_Project = Project{
		Project_Name: add_Project_Name,
		Start_Date:   start_Date,
		End_Date:     end_Date,
		Duration:     Count_Duration(start_Date, end_Date),
		Description:  add_Description,
		Post_Date:    time.Now().Format("02-01-2006"),
		Time_Post:    time.Now().Format("15:04"),
		Node_js:      (add_Node_Js == "on"),
		Golang:       (add_Golang == "on"),
		React_Js:     (add_React_Js == "on"),
		Java_Script:  (add_Javascript == "on"),
	}

	data_Projects[idToInt] = edit_Project
	return c.Redirect(http.StatusMovedPermanently, "/add-Project")
}

//-----------------------------------------------------GET EDIT PROJECT----------------------------------------------------//

func Get_Edit_Project(c echo.Context) error {
	id := c.Param("id")
	idToInt, _ := strconv.Atoi(id)

	tmpl, err := template.ParseFiles("view/Edit-Project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var ProjectDetail = Project{}

	for index, data := range data_Projects {
		if idToInt == index {
			ProjectDetail = Project{
				Project_Name: data.Project_Name,
				Start_Date:   data.Start_Date,
				End_Date:     data.End_Date,
				Duration:     data.Duration,
				Description:  data.Description,
				Post_Date:    data.Post_Date,
				Time_Post:    data.Time_Post,
				Node_js:      data.Node_js,
				Golang:       data.Golang,
				React_Js:     data.React_Js,
				Java_Script:  data.Java_Script,
			}
		}
	}

	data := map[string]interface{}{
		"Id":      id,
		"Project": ProjectDetail,
	}
	return tmpl.Execute(c.Response(), data)
}

//-----------------------------------------------------LIST TECHNOLOGIS------------------------------------------------------//

func checkValue(list []string, object string) bool {
	for _, data := range list {
		if data == object {
			return true
		}
	}
	return false
}
