package openai

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func InitDB(dataSourceName string) error {
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	// 在这里检查是否已存在超级管理员账户
	adminUser := User{
		Username: "admin",
		Password: "admin",
	}

	if !IsValidLogin(adminUser) {
		// 如果不存在，创建超级管理员账户
		err := CreateUser(adminUser)
		if err != nil {
			return err
		}
	}

	return nil
}

func CloseDB() {
	db.Close()
}

func IsValidLogin(user User) bool {
	var storedPassword string
	query := "SELECT password FROM users WHERE username = ?"
	err := db.QueryRow(query, user.Username).Scan(&storedPassword)
	if err != nil {
		log.Println("Error querying database:", err)
		return false
	}

	log.Printf("Query: %s, Username: %s, StoredPassword: %s\n", query, user.Username, storedPassword)

	return user.Password == storedPassword
}

func CreateUser(user User) error {
	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	return err
}

// UserExists 函数用于检查用户是否存在
func UserExists(username string) bool {
	var count int
	query := "SELECT COUNT(*) FROM users WHERE username = ?"
	err := db.QueryRow(query, username).Scan(&count)
	if err != nil {
		log.Println("Error querying database:", err)
		return false
	}

	return count > 0
}

/* HandleLogin 处理登录请求的函数 */
func HandleLogin(c *gin.Context) {
	// 从请求中获取用户名和密码
	var user User
	if err := c.BindJSON(&user); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request"})
		return
	}

	log.Println("Received login request for username:", user.Username)

	// 在数据库中查询用户是否存在的逻辑
	if !UserExists(user.Username) {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "User does not exist"})
		return
	}

	// 用户存在，进行密码验证
	if IsValidLogin(user) {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Invalid username or password"})
	}
}

/* HandleRegister 处理注册请求的函数 */
func HandleRegister(c *gin.Context) {
	// 处理注册请求，使用 CreateUser 方法
	// 可以根据具体需求进行实现
}
