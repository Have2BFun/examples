package examples

import (
	"github.com/gin-gonic/gin"
)

// SimpleAPI 简单的 API 示例
func SimpleAPI() *gin.Engine {
	r := gin.Default()

	// 获取所有用户
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"users": []string{"Alice", "Bob", "Charlie"},
		})
	})

	// 获取单个用户
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"id":   id,
			"name": "User " + id,
		})
	})

	// 创建用户
	r.POST("/users", func(c *gin.Context) {
		var json map[string]interface{}
		if err := c.ShouldBindJSON(&json); err == nil {
			c.JSON(201, gin.H{"message": "User created", "data": json})
		} else {
			c.JSON(400, gin.H{"error": err.Error()})
		}
	})

	return r
}
