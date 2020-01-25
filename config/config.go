package config

import (
	"fmt"
	"github.com/bxcodec/go-clean-arch/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetDB() (*gorm.DB, error){
	db, err := gorm.Open( "mysql","admin:password#123@tcp(10.0.75.1:3306)/thools?charset=utf8&parseTime=True&loc=Local")
	//defer db.Close()
	if err!=nil{
		fmt.Println("Connection Failed to Open")
	}
	fmt.Println("Connection Established")
	// Migrate the schema
	db.AutoMigrate(&models.User{},&models.Role{},&models.UserRole{})
	return db, nil
}
