package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var Posts = []Post{
	{ID: 1, Title: "Judul Postingan Pertama", Content: "Ini adalah postingan pertama di blog ini.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 2, Title: "Judul Postingan Kedua", Content: "Ini adalah postingan kedua di blog ini.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []User{
	{Username: "user1", Password: "pass1"},
	{Username: "user2", Password: "pass2"},
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		for _, user := range users {
			if user.Username == username && user.Password == password {
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(authMiddleware())

	r.GET("/posts", func(c *gin.Context) {
		idStr := c.Query("id")
		if idStr == "" {
		c.JSON(http.StatusOK, gin.H{"posts": Posts})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	for _, post := range Posts {
		if post.ID == id {
			c.JSON(http.StatusOK, gin.H{"post": post})
			return
		}
	}

	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Postingan tidak ditemukan"})
	})

	r.POST("/posts", func(c *gin.Context) {
		var newPost Post
		err := c.BindJSON(&newPost)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		newPost.ID = len(Posts) + 1
		newPost.CreatedAt = time.Now()
		newPost.UpdatedAt = time.Now()

		Posts = append(Posts, newPost)
		c.JSON(http.StatusCreated, gin.H{"post": newPost})	
	})

	return r
}

func main() {
	r := SetupRouter()

	r.Run(":8080")
}
