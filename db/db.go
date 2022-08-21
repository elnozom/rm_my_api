package db

import (
	"fmt"
	"mynozom/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var (
	DB *gorm.DB
)

func New(db string) (*gorm.DB, error) {
	// conStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_NAME"))
	conStr := fmt.Sprintf("sqlserver://%s:%s@%s:1433?database=%s", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), db)
	fmt.Println(conStr)
	DB, err := gorm.Open("mssql", conStr)
	if err != nil {
		return nil, err
	}
	DB.LogMode(true)
	return DB, nil
}
