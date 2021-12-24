package controller

import (
	"databaseExp/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SCInsert(ctx *gin.Context) {
	// 接收数据
	var newSC models.SC
	ctx.ShouldBind(&newSC)
	// 存入数据库
	err, rows := models.SCInsertDB(&newSC)
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

func SCQuery(ctx *gin.Context) {
	// 接收数据
	var scQuery models.SCQuery
	ctx.ShouldBindQuery(&scQuery)
	// 查询数据库
	var totalRows int64
	scs := make([]models.SC, 0)
	err, rows := models.SCQueryDB(&scQuery, &scs, &totalRows)
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
			"data":         scs,
		})
	}
}

func SCUpdate(ctx *gin.Context) {
	// 接收数据
	var newSC models.SC
	ctx.ShouldBind(&newSC)
	// 存入数据库
	err, rows := models.SCUpdateDB(&newSC)
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

func SCDelete(ctx *gin.Context) {
	// 接收数据
	var id int
	id, _ = strconv.Atoi(ctx.Query("id"))
	// 存入数据库
	err, rows := models.SCDeleteDB(id)
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
