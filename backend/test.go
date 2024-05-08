package main

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
)

func testingGin() {
	router := gin.Default()

	// Enable CORS
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

    // Run the server on port 8080
    router.Run(":8080")
}