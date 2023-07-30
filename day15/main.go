package main

import (
	connection "personal-web/Connection"
	"personal-web/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	connection.DatabaseConnect()

	e.Static("/Assets", "Assets")

	//List Get
	e.GET("/", handler.Home)
	e.GET("/contact", handler.Contact)
	e.GET("/index", handler.Index)
	e.GET("/add-Project", handler.Add_Project)
	e.GET("/form-project", handler.Form_Project)
	e.GET("/testimoni", handler.Testimoni)
	e.GET("/project-detail/:id", handler.Project_Detail)
	e.GET("/edit-project/:id", handler.Get_Edit_Project)

	//List Post
	e.POST("/post-project", handler.Post_Project)
	e.POST("/delete-project/:id", handler.Delete_Project)
	e.POST("/edit-project/:id", handler.Post_Edit_Project)

	//Server
	e.Logger.Fatal(e.Start("localhost:5000"))
}
