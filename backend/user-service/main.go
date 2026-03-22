package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
)

// 用户信息结构体
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

// 注册请求结构体
type RegisterRequest struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// 登录请求结构体
type LoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

var (
	rdb = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	ctx = context.Background()
)

func main() {
	router := gin.Default()

	// 注册
	router.POST("/register", func(c *gin.Context) {
		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "参数格式错误"})
			return
		}
		// 密码加密
		hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})
			return
		}
		user := User{
			ID:       time.Now().UnixNano(), // 简化, 应用生产需用数据库自增
			Username: req.Username,
			Phone:    req.Phone,
			Password: string(hashed),
		}
		bs, _ := json.Marshal(user)
		rdb.Set(ctx, "user:phone:"+req.Phone, bs, 0)
		c.JSON(http.StatusOK, gin.H{"msg": "注册成功"})
	})

	// 登录
	router.POST("/login", func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "参数格式错误"})
			return
		}
		val, err := rdb.Get(ctx, "user:phone:"+req.Phone).Result()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
			return
		}
		var user User
		json.Unmarshal([]byte(val), &user)
		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
			return
		}
		// 生成简单token并存Redis以实现登录态
		token := "token_" + req.Phone
		rdb.Set(ctx, "im:token:"+req.Phone, token, time.Hour*24)
		c.JSON(http.StatusOK, gin.H{"msg": "登录成功", "token": token, "user": user})
	})

	// 忘记密码（伪逻辑, 仅示例）
	router.POST("/reset", func(c *gin.Context) {
		type ResetRequest struct{ Phone, NewPassword string }
		var req ResetRequest
		c.ShouldBindJSON(&req)
		val, err := rdb.Get(ctx, "user:phone:"+req.Phone).Result()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
			return
		}
		var user User
		json.Unmarshal([]byte(val), &user)
		hashed, _ := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
		user.Password = string(hashed)
		bs, _ := json.Marshal(user)
		rdb.Set(ctx, "user:phone:"+req.Phone, bs, 0)
		c.JSON(http.StatusOK, gin.H{"msg": "密码重置成功"})
	})

	// 获取用户信息
	router.GET("/profile", func(c *gin.Context) {
		phone := c.Query("phone")
		val, err := rdb.Get(ctx, "user:phone:"+phone).Result()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "未找到用户"})
			return
		}
		var user User
		json.Unmarshal([]byte(val), &user)
		user.Password = ""
		c.JSON(http.StatusOK, user)
	})

	router.Run(":8081")
}