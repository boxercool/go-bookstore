package config

import (
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {

	uri := fmt.Sprintln("root" + ":" + "root" + "@tcp(" + "127.0.0.1" + ":" + "3306" + ")/" + "poorvika")
	//	d,err:=sql.Open("mysql", "root:root@tcp(localhost:3306)/poorvika")
	fmt.Println("Database Connection", uri)
	d, err := gorm.Open("mysql", strings.TrimSpace(uri)+"?parseTime=True")
	d.SingularTable(true)
	// db.LogMode(true)
	if err != nil {
		log.Fatal(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
