package controller

import (
	"databaseExp/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CourseInsert(ctx *gin.Context) {
	// 接收数据
	var newCou models.Course
	ctx.ShouldBind(&newCou)
	// 存入数据库
	err, rows := models.CourseInsertDB(&newCou)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  1,
			"msg":   "fail",
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":         0,
			"msg":          "success",
			"rowsAffected": rows,
		})
	}
}

func CourseQuery(ctx *gin.Context) {
	// 接收数据
	var couQuery models.CourseQuery
	ctx.ShouldBindQuery(&couQuery)
	// 查询数据库
	var totalRows int64
	courses := make([]models.Course, 0)
	err, rows := models.CourseQueryDB(&couQuery, &courses, &totalRows)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  1,
			"msg":   "fail",
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":         0,
			"msg":          "success",
			"rows":         totalRows,
			"rowsAffected": rows,
			"data":         courses,
		})
	}
}

func CourseUpdate(ctx *gin.Context) {
	// 接收数据
	var newCou models.Course
	ctx.ShouldBind(&newCou)
	// 存入数据库
	err, rows := models.CourseUpdateDB(&newCou)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  1,
			"msg":   "fail",
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":         0,
			"msg":          "success",
			"rowsAffected": rows,
		})
	}
}

func CourseDelete(ctx *gin.Context) {
	// 接收数据
	var id int
	id, _ = strconv.Atoi(ctx.Query("id"))
	// 存入数据库
	err, rows := models.CourseDeleteDB(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  1,
			"msg":   "fail",
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":         0,
			"msg":          "success",
			"rowsAffected": rows,
		})
	}
}
