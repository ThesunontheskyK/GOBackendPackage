package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thesunonthesky/GoBackendPackage/repository"
	"github.com/thesunonthesky/GoBackendPackage/handler"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


func main(){
	dsn := "root:password@tcp(127.0.0.1:3306)/my_database"
	db,err := sqlx.Connect("mysql",dsn)
	if err != nil{
		log.Fatalln(err)
	}

	userRepo := repository.NewUserRepositoryDB(db)
	userHandler := handler.NewUserHandler(userRepo)

	r := gin.Default()

	r.GET("/users",userHandler.GetUsers)
	r.POST("/users",userHandler.CreateUser)

	r.Run(":8080")
}