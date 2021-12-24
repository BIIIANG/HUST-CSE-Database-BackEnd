package dao

import (
	"databaseExp/settings"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

func InitMySQL(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	DB, err = gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "gorm_",
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("Failed to connect database, err: " + err.Error())
	}

	return err

}

func CloseMySQL() {
	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Println("Close Error.")
		return
	}
	sqlDB.Close()
}
