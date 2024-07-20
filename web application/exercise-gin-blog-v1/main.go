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

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/posts", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"posts": Posts,
		})
	})

	r.GET("/posts/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "ID harus berupa angka",
			})
			return
		}

		for _, post := range Posts {
			if id == post.ID {
				c.JSON(http.StatusOK, gin.H{
					"post": post,
				})
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Postingan tidak ditemukan",
		})
	})

	r.POST("/posts", func(c *gin.Context) {
		var post Post

		err := c.ShouldBindJSON(&post)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		now := time.Now()
		post.ID = len(Posts) + 1
		post.CreatedAt = now
		post.UpdatedAt = now

		Posts = append(Posts, post)
		c.JSON(http.StatusCreated, gin.H{
			"message": "Postingan berhasil ditambahkan",
			"post":    post,
		})
	})

	return r
}

func main() {
	r := SetupRouter()

	r.Run(":8080")
}
