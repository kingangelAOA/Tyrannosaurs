package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
	"github.com/jinzhu/gorm"
	"fmt"
	"tyrannosaurs/config"
)

var once sync.Once
const connectFormat = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
var DB *gorm.DB
var Ctx map[string]string

func init()  {
	once.Do(func() {
		config, err := config.Env.GetConfig(*config.E)
		if err != nil {
			panic(err)
		}
		mysql := config.Mysql
		db, err := gorm.Open("mysql", fmt.Sprintf(connectFormat, mysql.Username, mysql.Password, mysql.Host, mysql.Port, mysql.DB))
		if err != nil {
			panic(err)
		}
		db.DB().SetMaxIdleConns(mysql.MaxIdleConns)
		db.DB().SetMaxOpenConns(mysql.MaxOpenConns)
		DB = db
		createTable()
	})
}

func createTable()  {
	if !DB.HasTable(&Plan{}) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Plan{})
	}
	if !DB.HasTable(&ThreadGroup{}) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&ThreadGroup{})
	}
	if !DB.HasTable(&Scenario{}) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Scenario{})
	}
	if !DB.HasTable(&Sampler{}) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Sampler{})
	}
	if !DB.HasTable(&HttpRequest{}) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&HttpRequest{})
	}
	if !DB.HasTable(&Cookie{}) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Cookie{})
	}
	if !DB.HasTable(&TestData{}) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&TestData{})
	}
	if !DB.HasTable(&Script{}) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Script{})
	}
}