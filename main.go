package main

import (
	"fmt"
	"mynozom/config"
	"mynozom/db"
	"mynozom/handler"
	"mynozom/repo"
	"mynozom/router"

	"github.com/jinzhu/gorm"
)

func main() {
	r := router.New()
	v1 := r.Group("/api")
	db1, err := db.New(config.Config("DB_NAME1"))
	if err != nil {
		panic(err)
	}
	db2, err := db.New(config.Config("DB_NAME2"))
	if err != nil {
		panic(err)
	}
	db3, err := db.New(config.Config("DB_NAME3"))
	if err != nil {
		panic(err)
	}
	db4, err := db.New(config.Config("DB_NAME4"))

	dbs := []*gorm.DB{db1, db2, db3, db4}
	if err != nil {
		panic(err)
	}
	userRepo := repo.NewUserRepo(dbs)
	accountRepo := repo.NewAccountRepo(dbs)
	h := handler.NewHandler(dbs, userRepo, accountRepo)
	h.Register(v1)
	port := fmt.Sprintf(":%s", config.Config("PORT"))
	r.Logger.Fatal(r.Start(port))

}
