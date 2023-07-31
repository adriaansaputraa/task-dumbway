package main

import (
	connection "personal-web/Connection"
	middleware "personal-web/Middleware"
	"personal-web/handler"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	connection.DatabaseConnect()

	e.Static("/Assets", "Assets")
	e.Static("/uploads", "uploads")

	//List Get
	e.GET("/", handler.Home)
	e.GET("/contact", handler.Contact)
	e.GET("/index", handler.Index)
	e.GET("/add-Project", handler.Add_Project)
	e.GET("/form-project", handler.Form_Project)
	e.GET("/form-register", handler.Form_Register)
	e.GET("/form-login", handler.Form_Login)
	e.GET("/testimoni", handler.Testimoni)
	e.GET("/project-detail/:id", handler.Project_Detail)
	e.GET("/edit-project/:id", handler.Get_Edit_Project)

	//List Post
	e.POST("/post-project", middleware.Upload_File(handler.Post_Project))
	e.POST("/delete-project/:id", handler.Delete_Project)
	e.POST("/edit-project/:id", middleware.Upload_File(handler.Post_Edit_Project))
	e.POST("/register", handler.Register)
	e.POST("/login", handler.Login)
	e.POST("/logout", handler.Logout)

	//Server
	e.Logger.Fatal(e.Start("localhost:5000"))

}
