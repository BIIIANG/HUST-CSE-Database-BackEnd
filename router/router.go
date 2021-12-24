package router

import (
	"github.com/gin-gonic/gin"

	"databaseExp/controller"
	"databaseExp/middlewares"
	"databaseExp/settings"
)

func SetupRouter() *gin.Engine {
	if settings.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Use(middlewares.Cors())

	r.Static("/js", "./static/js")
	r.Static("/css", "./static/css")
	r.Static("/img", "./static/img")
	r.Static("/fonts", "./static/fonts")

	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	r.GET("/ping", controller.Ping)
	studentGroup := r.Group("api/student")
	{
		studentGroup.GET("/query", controller.StudentQuery)
		studentGroup.POST("/insert", controller.StudentInsert)
		studentGroup.PUT("/update", controller.StudentUpdate)
		studentGroup.DELETE("/delete", controller.StudentDelete)
	}
	courseGroup := r.Group("api/course")
	{
		courseGroup.GET("/query", controller.CourseQuery)
		courseGroup.POST("/insert", controller.CourseInsert)
		courseGroup.PUT("/update", controller.CourseUpdate)
		courseGroup.DELETE("/delete", controller.CourseDelete)
	}
	scGroup := r.Group("api/sc")
	{
		scGroup.GET("/query", controller.SCQuery)
		scGroup.POST("/insert", controller.SCInsert)
		scGroup.PUT("/update", controller.SCUpdate)
		scGroup.DELETE("/delete", controller.SCDelete)
	}
	appGroup := r.Group("api/app")
	{
		appGroup.GET("/infograde", controller.InfoGradeQuery)
		appGroup.GET("/deptinfo", controller.DeptInfoQuery)
		appGroup.GET("/deptrank", controller.DeptRankQuery)
	}

	return r
}
