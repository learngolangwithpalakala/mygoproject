package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_roleRepo "github.com/learngolangwithpalakala/mygoproject/admin/role/repository"
	_userHttpDeliver "github.com/learngolangwithpalakala/mygoproject/admin/user/delivery/http"
	_userRepo "github.com/learngolangwithpalakala/mygoproject/admin/user/repository"
	_userUcase "github.com/learngolangwithpalakala/mygoproject/admin/user/usecase"
	_userRoleRepo "github.com/learngolangwithpalakala/mygoproject/admin/user_role/repository"
	"github.com/learngolangwithpalakala/mygoproject/config"
	_storageHttpDeliver "github.com/learngolangwithpalakala/mygoproject/storage/delivery/http"
	_storageRepo "github.com/learngolangwithpalakala/mygoproject/storage/repository"
	_storageUcase "github.com/learngolangwithpalakala/mygoproject/storage/usecase"
	"net/http"
	//"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type Config struct {
	Debug  bool `json:"debug"`
	Server struct {
		Address      string `json:"address"`
		Host         string `json:"host"`
		FileLocation string `json:"filelocation"`
	} `json:"server"`
	Context struct {
		Timeout int64 `json:"timeout"`
	} `json:"context"`
	Database struct {
		DBDriver   string `json:"dbDriver"`
		DBName     string `json:"dbName"`
		DBUserName string `json:"dbUser"`
		DBPassword string `json:"dbPassword"`
		DBServerIP string `json:"dbServerIp"`
		DBPort     string `json:"dbPort"`
	} `json:"database"`
}

func LoadConfiguration(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)

	if err != nil {
		return config, err
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return config, err
	}
	return config, err
}


func main() {
	fmt.Println("starting application")
	configuration , _ := LoadConfiguration("config.json")
	fmt.Println("Server File System location", configuration.Server.FileLocation)
	fmt.Println("Server Port:", configuration.Server.Address)
	fmt.Println("Server timeout:",configuration.Context.Timeout)
	fmt.Println(" DB Driver:", configuration.Database.DBDriver)
	fmt.Println(" DB Name:", configuration.Database.DBName)
	fmt.Println(" DB Server IP:", configuration.Database.DBServerIP)
	fmt.Println(" DB Port:", configuration.Database.DBPort)
	fmt.Println(" DB UserName:", configuration.Database.DBUserName)
	fmt.Println(" DB Password:", configuration.Database.DBPassword)
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}))


	db, err := config.GetDB(configuration.Database.DBDriver,
		configuration.Database.DBName,
		configuration.Database.DBUserName,
		configuration.Database.DBPassword,
		configuration.Database.DBServerIP,
		configuration.Database.DBPort,
	)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}


	userRep := _userRepo.NewMysqlUserRepository(db)
	roleRep := _roleRepo.NewMysqlRoleRepository(db)
	userRoleRep := _userRoleRepo.NewMysqlUserRoleRepository(db)
	fmt.Println(configuration.Context.Timeout)
	timeoutContext := time.Duration(configuration.Context.Timeout) * time.Second
	fileLocation := configuration.Server.FileLocation
	uuc := _userUcase.NewUserUsecase(userRep, roleRep, userRoleRep, timeoutContext)
	_userHttpDeliver.NewUserHandler(e, uuc)

	_storageRepo := _storageRepo.NewMysqlStorageFileRepository(db)
	stguc := _storageUcase.NewProductUsecase(_storageRepo, timeoutContext)
	_storageHttpDeliver.NewStorageHandler(e, stguc, fileLocation)


	log.Fatal(e.Start(configuration.Server.Address))
}
