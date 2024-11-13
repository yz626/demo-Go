package main

import (
	"Gin-jwt/utils"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func main() {
	r := gin.Default()

	r.POST("/login", Login)

	v1 := r.Group("/vip", JWTAuth())
	{
		v1.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "hello",
				"data": nil,
			})
		})

		v1.GET("/sayHello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "sayHello",
				"data": nil,
			})
		})
	}

	r.Run(":8080")
}

func Login(c *gin.Context) {
	// 身份认证
	name := c.PostForm("name")
	_ = c.PostForm("password")

	// 生成token
	token, err := utils.CreateToken(name)
	if err != nil {
		c.JSON(500, gin.H{
			"code":  500,
			"msg":   "token生成失败",
			"error": err.Error(),
			"data":  nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{
				"code": 401,
				"msg":  "token为空",
				"data": nil,
			})
			c.Abort()
			return
		}

		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 {
			c.JSON(401, gin.H{
				"code": 401,
				"msg":  "token格式不正确",
				"data": nil,
			})
			c.Abort()
			return
		}

		tokenString = parts[1]

		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{
				"code":  401,
				"msg":   "token认证失败",
				"error": err.Error(),
				"data":  nil,
			})
			c.Abort()
			return
		}

		if claims.ExpiresAt.Time.Unix() < time.Now().Unix() {
			c.JSON(401, gin.H{
				"code": 401,
				"msg":  "token已过期",
				"data": nil,
			})
			c.Abort()
			return
		}

		if claims.Name != "admin" {
			c.JSON(401, gin.H{
				"code": 401,
				"msg":  "权限不足",
				"data": nil,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
