package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/twitter_ejercicio/app"
	"github.com/robertobouses/twitter_ejercicio/http"
	"github.com/robertobouses/twitter_ejercicio/internal"
	"github.com/robertobouses/twitter_ejercicio/repository"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	fmt.Println("DB_USER:", os.Getenv("DB_USER"))
	fmt.Println("DB_PASS:", os.Getenv("DB_PASS"))
	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
	fmt.Println("DB_DATABASE:", os.Getenv("DB_DATABASE"))

	db, err := internal.NewPostgres(internal.DBConfig{
		User:     os.Getenv("DB_USER"),
		Pass:     os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
	})

	if err != nil {
		panic(err)
	}
	defer db.Close()

	var repo repository.REPOSITORY
	repo, err = repository.NewRepository(db)
	if err != nil {
		panic(err)
	}
	var appController app.APP
	var httpController http.HTTP

	appController = app.NewAPP(repo)
	httpController = http.NewHTTP(appController)

	server := gin.Default()

	db.AutoMigrate(&entity.User{}, &entity.Tweet)

	server.POST("/createAccount", httpController.CreateAccount)
	server.PUT("/configurePassword/:id", httpController.ConfigurePassword)
	server.POST("/publishTweet", httpController.PublishTweet)
	server.GET("/getOwnTweets", httpController.GetOwnTweets)
	server.GET("/getFollowingTweets", httpController.GetFollowingTweets)

	server.Run(":8080")
}
