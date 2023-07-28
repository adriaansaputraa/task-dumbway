package handler

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id           int
	Project_Name string
	Start_Date   string
	End_Date     string
	Duration     string
	Description  string
	Post_Date    string
	Time_Post    string
	Node_js      bool
	Golang       bool
	React_Js     bool
	Java_Script  bool
	Image        string
}

var data_Projects = []Project{
	{
		Project_Name: "Project pertama",
		Start_Date:   "23-07-2023",
		End_Date:     "25-08-2023",
		Duration:     Count_Duration("23-07-2023", "25-08-2023"),
		Description:  "This is the description of project 1",
		Post_Date:    time.Now().Format("02-01-2006"),
		Time_Post:    time.Now().Format("15:04"),
		Node_js:      true,
		Golang:       true,
		React_Js:     false,
		Java_Script:  true,
	},
	{
		Project_Name: "Project kedua",
		Start_Date:   "23-07-2023",
		End_Date:     "03-08-2023",
		Duration:     Count_Duration("23-07-2023", "03-08-2023"),
		Description:  "This is the description of project 2",
		Post_Date:    time.Now().Format("02-01-2006"),
		Time_Post:    time.Now().Format("15:04"),
		Node_js:      false,
		Golang:       true,
		React_Js:     true,
		Java_Script:  true,
	},
	{
		Project_Name: "Project ketiga",
		Start_Date:   "23-07-2023",
		End_Date:     "25-07-2023",
		Duration:     Count_Duration("23-07-2023", "25-07-2023"),
		Description:  "This is the description of project 3",
		Post_Date:    time.Now().Format("02-01-2006"),
		Time_Post:    time.Now().Format("15:04"),
		Node_js:      true,
		Golang:       true,
		React_Js:     false,
		Java_Script:  true,
	},
}

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

	// dataQuery, errQuery := connection.Conn.Query(context.Background(), "SELECT name, description FROM tb_projects")

	// if errQuery != nil {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }

	// var resultProject []Project

	// for dataQuery.Next() {

	// 	var each = Project{}

	// 	err := dataQuery.Scan(&each.Project_Name, &each.Description)
	// 	if err != nil {
	// 		return c.JSON(http.StatusInternalServerError, err.Error())
	// 	}

	// 	resultProject = append(resultProject, each)
	// }

	data := map[string]interface{}{
		"Projects": data_Projects,
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
	new_Start_Date := start_Date.Format("02-01-2006")

	end_Date, err := time.Parse("2006-01-02", add_End_Date)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	new_End_Date := end_Date.Format("02-01-2006")

	var new_Project = Project{
		Project_Name: add_Project_Name,
		Start_Date:   new_Start_Date,
		End_Date:     new_End_Date,
		Duration:     Count_Duration(new_Start_Date, new_End_Date),
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

func Count_Duration(d1, d2 string) string {
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
	new_Start_Date := start_Date.Format("02-01-2006")

	end_Date, err := time.Parse("2006-01-02", add_End_Date)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	new_End_Date := end_Date.Format("02-01-2006")

	var edit_Project = Project{
		Project_Name: add_Project_Name,
		Start_Date:   new_Start_Date,
		End_Date:     new_End_Date,
		Duration:     Count_Duration(new_Start_Date, new_End_Date),
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
