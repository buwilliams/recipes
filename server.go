package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//db, err := sql.Open("mysql", "root@/twitter_db")
	//if err != nil {
	//panic(err)
	//}

	router := gin.Default()

	//router.GET("/user", getAllUsers(db))
	//router.GET("/tweet", getAllTweets(db))
	router.GET("/ping", getPing())

	router.Run()
}

func getPing() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

func getAllUsers(db *sql.DB) func(c *gin.Context) {
	u := User{}
	users, err := u.GetAllUsers(db)

	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		c.JSON(200, users)
	}
}

func getAllTweets(db *sql.DB) func(c *gin.Context) {
	u := Tweet{}
	tweets, err := u.GetAllTweets(db)

	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		c.JSON(200, tweets)
	}
}
