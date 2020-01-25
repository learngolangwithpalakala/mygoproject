package main

import (
	"encoding/json"
	"fmt"
	_roleRepo "github.com/bxcodec/go-clean-arch/admin/role/repository"
	_userHttpDeliver "github.com/bxcodec/go-clean-arch/admin/user/delivery/http"
	_userRepo "github.com/bxcodec/go-clean-arch/admin/user/repository"
	_userUcase "github.com/bxcodec/go-clean-arch/admin/user/usecase"
	_userRoleRepo "github.com/bxcodec/go-clean-arch/admin/user_role/repository"
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/middleware"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	//"github.com/spf13/viper"
	"log"
	"os"

	//"net/url"
	//"os"
	"time"
)
  type Config struct {
	  Debug bool `json:"debug"`
    Server struct {
    	Address string `json:"address"`
    	Host    string `json:"host"`
	}  `json:"server"`
    Context struct {
    	Timeout int64 `json:"timeout"`
	}`json:"context"`
  	Database struct {
      	  Host string `json:"host"`
      	  Port string `json:"port"`
	  }`json:"database"`
  }

func LoadConfiguration(filename string) (Config,error) {
	  var config Config
	  configFile, err := os.Open(filename)
	  defer configFile.Close()
	  if err != nil {
		  return config, err
	  }
	  jsonParser := json.NewDecoder(configFile)
	  err = jsonParser.Decode(&config)
	  return config, err
     }

 func main() {
	 fmt.Println("starting application")
     configuration , _ := LoadConfiguration("config.json")
     fmt.Println(" database port",configuration.Database.Port)
	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	userRep := _userRepo.NewMysqlUserRepository(db)
	roleRep:= _roleRepo.NewMysqlRoleRepository(db)
	userRoleRep:= _userRoleRepo.NewMysqlUserRoleRepository(db)
	fmt.Println(configuration.Context.Timeout)
	timeoutContext := time.Duration(configuration.Context.Timeout) * time.Second
	uuc := _userUcase.NewUserUsecase(userRep, roleRep,userRoleRep,timeoutContext)
	_userHttpDeliver.NewUserHandler(e, uuc)

	log.Fatal(e.Start(configuration.Server.Address))
}
