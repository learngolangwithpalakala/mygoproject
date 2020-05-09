package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//"github.com/learngolangwithpalakala/mygoproject/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetDB(dbDriver string, dbName string, dbUser string, dbPassword string, dbServerIp string, dbPort string) (*gorm.DB, error) {
	db, err := gorm.Open(dbDriver, dbUser+":"+dbPassword+"@tcp("+dbServerIp+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Connection Failed to Open")
	}
	fmt.Println("DB Connection Established succesfully")
	// Migrate the schema
	//db.AutoMigrate(&models.User{}, &models.Role{}, &models.UserRole{})
	return db, nil

}
