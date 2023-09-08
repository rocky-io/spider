package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	protodata "github.com/rockyskate/spider/proto/go"
	"google.golang.org/protobuf/proto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Person struct {
	Name    string `form:"name" binding:"required"`
	Address string `form:"address"`
}

// UserShow godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  Person
// @Router       /user/show [get]
func UserShow(c *gin.Context) {
	var person Person
	if err := c.ShouldBindQuery(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	log.Println("====== Only Bind By Query String ======")
	log.Println(person.Name)
	log.Println(person.Address)
	c.JSON(http.StatusOK, person)
}

type User struct {
	ID        uint   `gorm:"primarykey"`
	Username  string `gorm:"size:32"`
	Password  string `gorm:"size:32"`
	Nickname  string `gorm:"size:16"`
	Age       uint8
	Intro     sql.NullString `gorm:"size:512"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func UserAdd(c *gin.Context) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/foodbase?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,  // disable datetime precision support, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // smart configure based on used version
	}), &gorm.Config{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "", "message": err.Error()})
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	db.AutoMigrate(&User{})
	db.Create(&User{Username: "rocky", Password: "123456", Nickname: "管理员", Age: 16})

	// Read
	var user User
	db.First(&user, "username = ?", "rocky")
	c.JSON(http.StatusOK, user)

	// // Update - 将 product 的 price 更新为 200
	// db.Model(&user).Update("Age", 30)
	// // Update - 更新多个字段
	// db.Model(&user).Updates(User{Age: 10, Nickname: "aaa"}) // 仅更新非零值字段
	// //db.Model(&user).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - 删除 product
	// db.Delete(&user, 1)
}

type helloServer struct {
	protodata.UnimplementedGreeterServer
}

func Prototest(c *gin.Context) {
	user := protodata.User{
		Username: "rocky",
		Password: "123456",
		Birtiday: "2020-06-08",
	}
	fmt.Println(&user)
	data, _ := proto.Marshal(&user)
	fmt.Println(data)
	proto.Unmarshal(data, &user)
	fmt.Println(&user)
}
