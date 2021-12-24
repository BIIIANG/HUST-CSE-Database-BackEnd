package main

import (
	"databaseExp/dao"
	"databaseExp/models"
	"databaseExp/router"
	"databaseExp/settings"
	"fmt"
)

func main() {
	var err error

	// 加载配置文件
	if err = settings.Init("config/config.ini"); err != nil {
		fmt.Printf("Load config from file failed, err: %v\n", err)
		return
	}

	// 连接数据库
	err = dao.InitMySQL(settings.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("Init MySQL failed, err: %v\n", err)
		return
	}
	// 模型绑定
	err = dao.DB.AutoMigrate(&models.Student{})
	if err != nil {
		panic("Migrate database failed, err: " + err.Error())
	}
	err = dao.DB.AutoMigrate(&models.Course{})
	if err != nil {
		panic("Migrate database failed, err: " + err.Error())
	}
	err = dao.DB.AutoMigrate(&models.SC{})
	if err != nil {
		panic("Migrate database failed, err: " + err.Error())
	}
	// 程序退出关闭数据库连接
	defer dao.CloseMySQL()

	// 注册路由
	r := router.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port)); err != nil {
		fmt.Printf("Server startup failed, err: %v\n", err)
	}

}
