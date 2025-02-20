package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// InitializeDB 用于初始化数据库连接并自动创建表
func InitializeDB() {
	var err error
	// 数据库连接字符串：用户名:密码@tcp(主机:端口)/数据库名?charset=utf8&parseTime=True&loc=Local
	dsn := "root:12345678@tcp(127.0.0.1:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	// 使用 AutoMigrate 自动创建或更新表结构
	err = DB.AutoMigrate(&Transaction{}).Error
	if err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}

	fmt.Println("Database initialized and tables created if not exist")
}

// Transaction 代表支付记录表
type Transaction struct {
	ID             uint   `gorm:"primary_key"`
	OrderID        string `gorm:"not null"`
	TransactionID  string `gorm:"not null"`
	Amount         float64
	UserID         uint
	CreditCardInfo string
	CreatedAt      string
	UpdatedAt      string
}
