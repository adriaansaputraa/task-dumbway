package handler

import (
	"context"
	"html/template"
	"net/http"
	"net/mail"
	connection "personal-web/Connection"
	"strconv"
	"time"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

type User struct {
	Id              int
	Username        string
	Email           string
	Hashed_Password string
}

type UserLoginSession struct {
	IsLogin bool
	Name    string
}

var data_Projects = []Project{}
var user_Login_Session = UserLoginSession{}

// ------------------------------------------------------------------------HOME--------------------------------------------------------//
func Home(c echo.Context) error {

	tmpl, err := template.ParseFiles("view/Home.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sess, _ := session.Get("session", c)
	if sess.Values["isLogin"] != true {
		user_Login_Session.IsLogin = false
	} else {
		user_Login_Session.IsLogin = true
		user_Login_Session.Name = sess.Values["name"].(string)
	}

	data := map[string]interface{}{
		"FlashMessage":     sess.Values["message"],
		"UserLoginSession": user_Login_Session,
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), data)
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

	sess, _ := session.Get("session", c)

	dataQuery, errQuery := connection.Conn.Query(context.Background(), "SELECT tb_projects.id, tb_projects.name, tb_projects.start_date, tb_projects.end_date, tb_projects.description, tb_projects.technologies, tb_projects.image, tb_projects.post_date, tb_projects.time_post FROM tb_projects LEFT JOIN tb_user ON tb_projects.user_id = tb_user.id where tb_projects.user_id = $1 ", sess.Values["id"].(int))

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
	add_Image := c.Get("dataFile").(string)

	sess, _ := session.Get("session", c)

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_projects (name, start_date, end_date, description, technologies, image, post_date, time_post, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", add_Project_Name, add_Start_Date, add_End_Date, add_Description, add_techonologies, add_Image, add_post_date, add_time_post, sess.Values["id"].(int))

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
	sess, _ := session.Get("session", c)
	if sess.Values["isLogin"] != true {
		user_Login_Session.IsLogin = false
	} else {
		user_Login_Session.IsLogin = true
		user_Login_Session.Name = sess.Values["name"].(string)
	}

	data := map[string]interface{}{
		"FlashMessage":     sess.Values["message"],
		"UserLoginSession": user_Login_Session,
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), data)
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
	add_Image := c.Get("dataFile").(string)

	StartDate, _ := time.Parse("2006-01-02", add_Start_Date)
	EndDate, _ := time.Parse("2006-01-02", add_End_Date)

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_projects SET name = $1, start_date = $2, end_date = $3, description = $4, technologies = $5, image = $6 WHERE id = $7", add_Project_Name, StartDate, EndDate, add_Description, add_techonologies, add_Image, idToInt)

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

// -----------------------------------------------------FORM REGISTER------------------------------------------------------//
func Form_Register(c echo.Context) error {
	tmpl, err := template.ParseFiles("view/Form-Register.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	flash := map[string]interface{}{
		"FlashMessage": sess.Values["message"],
		"FlashStatus":  sess.Values["status"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), flash)
}

// -----------------------------------------------------REGISTER--------------------------------------------------------//
func Register(c echo.Context) error {

	input_Username := c.FormValue("input-username")
	input_Email := c.FormValue("input-email")
	input_Password := c.FormValue("input-password")

	_, errEmail := mail.ParseAddress(input_Email)

	if errEmail != nil {
		return redirectWithMessage(c, "Invalid email format", false, "/form-register")
	}

	hased_Password, errHase := bcrypt.GenerateFromPassword([]byte(input_Password), 10)

	if errHase != nil {
		return c.JSON(http.StatusInternalServerError, errHase.Error())
	}

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_user (name, email, password) VALUES ($1, $2, $3)", input_Username, input_Email, hased_Password)

	if err != nil {
		return redirectWithMessage(c, "Registration failed", false, "/form-register")
	}

	return redirectWithMessage(c, "Registration Success", true, "/")

}

// -----------------------------------------------------FORM LOGIN------------------------------------------------------//
func Form_Login(c echo.Context) error {
	tmpl, err := template.ParseFiles("view/Form-Login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	flash := map[string]interface{}{
		"FlashMessage": sess.Values["message"],
		"FlashStatus":  sess.Values["status"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), flash)
}

// -----------------------------------------------------LOGIN-------------------------------------------------------//
func Login(c echo.Context) error {
	input_Email := c.FormValue("input-email")
	input_Password := c.FormValue("input-password")

	if input_Email == "" || input_Password == "" {
		return redirectWithMessage(c, "All fields must be filled", false, "/form-login")
	}

	_, errEmail := mail.ParseAddress(input_Email)

	if errEmail != nil {
		return redirectWithMessage(c, "Invalid email format", false, "/form-login")
	}

	user := User{}

	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT id, name, email, password FROM tb_user WHERE email=$1", input_Email).Scan(&user.Id, &user.Username, &user.Email, &user.Hashed_Password)

	if errQuery != nil {
		return redirectWithMessage(c, "Email/password wrong", false, "/form-login")
	}

	errPasword := bcrypt.CompareHashAndPassword([]byte(user.Hashed_Password), []byte(input_Password))

	if errPasword != nil {
		return redirectWithMessage(c, "Email/password wrong", false, "/form-login")
	}

	//Set session login if sucess

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800 // 3 JAM -> berapa lama expired
	sess.Values["message"] = "Login success!"
	sess.Values["status"] = true
	sess.Values["name"] = user.Username
	sess.Values["email"] = user.Email
	sess.Values["id"] = user.Id
	sess.Values["isLogin"] = true
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func Logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

// -----------------------------------------------------REDIRECT WITH MESSAGE-------------------------------------------------------//

func redirectWithMessage(c echo.Context, message string, status bool, redirectPath string) error {
	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, redirectPath)
}
