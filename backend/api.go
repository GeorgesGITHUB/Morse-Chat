package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// NOTE: Ordering of the functions represents the run order

func enableCORS(router *gin.Engine) {
	router.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    })
}

func registerAPItoEndpoint(router *gin.Engine) {
//	http://localhost:8080/api/users/id?username=Bob&password=p1
    router.GET("/api/users/id", func(c *gin.Context) {
		username := c.Query("username")
        password := c.Query("password")
		
		var db Database
		db.OpenConnection()
		defer db.CloseConnection()
		
		user_id, err := db.GetUser(username,password)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
		}

		c.JSON(http.StatusOK, gin.H{"user_id": strconv.Itoa(user_id)})
	})

//	http://localhost:8080/api/users
    router.POST("/api/users", func(c *gin.Context) {
        var newUser User
		if err := c.BindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

        username, password := newUser.Username, newUser.Password

		var db Database
		db.OpenConnection()
		defer db.CloseConnection()
		
		err := db.PostUser(username,password)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Failed posting User"})
            return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Successfully posted User"})
	})
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}