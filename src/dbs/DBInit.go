package dbs

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

var Orm *gorm.DB

func init() {
	err := fmt.Errorf("")
	Orm, err = gorm.Open("mysql", "root:123456@tcp(101.132.138.205:3306)/test?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB := Orm.DB()
	Orm.LogMode(true)
	logger := logrus.New()
	Orm.SetLogger(logger)
	mysqlDB.SetMaxIdleConns(5)
	mysqlDB.SetMaxOpenConns(10)
	mysqlDB.SetConnMaxLifetime(30 * time.Second)
}
