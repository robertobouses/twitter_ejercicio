package internal

import "fmt"

var db *gorm.DB
var err error

func initialMigration() {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()
	db.AutoMigrate(&User{}, &Tweet{})
}
