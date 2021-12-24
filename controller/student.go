package controller

import (
	"databaseExp/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func StudentInsert(ctx *gin.Context) {
	// 接收数据
	var newStu models.Student
	ctx.ShouldBind(&newStu)
	// 存入数据库
	err, rows := models.StudentInsertDB(&newStu)
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

func StudentQuery(ctx *gin.Context) {
	// 接收数据
	var stuQuery models.StudentQuery
	ctx.ShouldBindQuery(&stuQuery)
	// 查询数据库
	var totalRows int64
	students := make([]models.Student, 0)
	err, rows := models.StudentQueryDB(&stuQuery, &students, &totalRows)
	// fmt.Println(students)
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
			"data":         students,
		})
	}
}

func StudentUpdate(ctx *gin.Context) {
	// 接收数据
	var newStu models.Student
	ctx.ShouldBind(&newStu)
	// 存入数据库
	err, rows := models.StudentUpdateDB(&newStu)
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

func StudentDelete(ctx *gin.Context) {
	// 接收数据
	var id int
	id, _ = strconv.Atoi(ctx.Query("id"))
	// 存入数据库
	err, rows := models.StudentDeleteDB(id)
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
