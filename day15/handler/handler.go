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

	dataQuery, errQuery := connection.Conn.Query(context.Background(), "SELECT id, name, start_date, end_date, description, technologies, image, post_date, time_post FROM tb_projects")

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var resultProject []Project

	for dataQuery.Next() {

		var each = Project{}

		err := dataQuery.Scan(&each.Id, &each.Project_Name, &each.Start_Date, &each.End_Date, &each.Description, &each.Technologies, &each.Image, &each.Post_Date, &each.Time_Post)
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
	add_techonologies := []string{add_Node_Js, add_Golang, add_React_Js, add_Javascript}
	add_post_date := time.Now().Format("02-01-2006")
	add_time_post := time.Now().Format("15:04")

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_projects (name, start_date, end_date, description, technologies, image, post_date, time_post) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", add_Project_Name, add_Start_Date, add_End_Date, add_Description, add_techonologies, "default.jgp", add_post_date, add_time_post)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

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

	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT id, name, start_date, end_date, description, technologies, image, post_date, time_post FROM tb_projects WHERE id=$1", idToInt).Scan(&ProjectDetail.Id, &ProjectDetail.Project_Name, &ProjectDetail.Start_Date, &ProjectDetail.End_Date, &ProjectDetail.Description, &ProjectDetail.Technologies, &ProjectDetail.Image, &ProjectDetail.Post_Date, &ProjectDetail.Time_Post)

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	ProjectDetail.Duration = Count_Duration(ProjectDetail.Start_Date, ProjectDetail.End_Date)
	if checkValue(ProjectDetail.Technologies, "nodejs") {
		ProjectDetail.Node_js = true
	}
	if checkValue(ProjectDetail.Technologies, "golang") {
		ProjectDetail.Golang = true
	}
	if checkValue(ProjectDetail.Technologies, "reactjs") {
		ProjectDetail.React_Js = true
	}
	if checkValue(ProjectDetail.Technologies, "javascript") {
		ProjectDetail.Java_Script = true
	}

	data := map[string]interface{}{
		"Project":   ProjectDetail,
		"StartDate": ProjectDetail.Start_Date.Format("02-01-2006"),
		"EndDate":   ProjectDetail.End_Date.Format("02-01-2006"),
	}
	return tmpl.Execute(c.Response(), data)
}

//-------------------------------------------------------DELETE PROJECT----------------------------------------------------//

func Delete_Project(c echo.Context) error {
	id := c.Param("id")
	idToInt, _ := strconv.Atoi(id)

	connection.Conn.Exec(context.Background(), "DELETE FROM tb_projects WHERE id=$1", idToInt)

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
	add_techonologies := []string{add_Node_Js, add_Golang, add_React_Js, add_Javascript}

	StartDate, _ := time.Parse("2006-01-02", add_Start_Date)
	EndDate, _ := time.Parse("2006-01-02", add_End_Date)

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_projects SET name = $1, start_date = $2, end_date = $3, description = $4, technologies = $5 WHERE id = $6", add_Project_Name, StartDate, EndDate, add_Description, add_techonologies, idToInt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

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

	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT id, name, start_date, end_date, description, technologies, image, post_date, time_post FROM tb_projects WHERE id=$1", idToInt).Scan(&ProjectDetail.Id, &ProjectDetail.Project_Name, &ProjectDetail.Start_Date, &ProjectDetail.End_Date, &ProjectDetail.Description, &ProjectDetail.Technologies, &ProjectDetail.Image, &ProjectDetail.Post_Date, &ProjectDetail.Time_Post)

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if checkValue(ProjectDetail.Technologies, "nodejs") {
		ProjectDetail.Node_js = true
	}
	if checkValue(ProjectDetail.Technologies, "golang") {
		ProjectDetail.Golang = true
	}
	if checkValue(ProjectDetail.Technologies, "reactjs") {
		ProjectDetail.React_Js = true
	}
	if checkValue(ProjectDetail.Technologies, "javascript") {
		ProjectDetail.Java_Script = true
	}

	start_date := ProjectDetail.Start_Date.Format("2006-01-02")
	end_date := ProjectDetail.End_Date.Format("2006-01-02")

	data := map[string]interface{}{
		"Project":   ProjectDetail,
		"StartDate": start_date,
		"EndDate":   end_date,
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
