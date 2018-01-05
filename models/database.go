package models

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	DB_USER     = "root"
	DB_NAME     = "expense"
	DB_PASSWORD = "Behrooz1993"
)

type Datastore interface {
	AddUser(user *User) error
	SelectByCellphone(user *User) error
}

type DB struct {
	database *gorm.DB
}

var Mgr *DB

func init() {
	connectionString := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		log.Panic(err)
	}
	log.Print("Database initialiized successfully!")

	db.AutoMigrate(&User{})

	Mgr = &DB{db}
}
