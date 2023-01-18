package database

import (
	"fmt"
	"log"
	"sully/todo-app/config"
	"sully/todo-app/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialise() {
	var err error
	user := config.GetConfig("DB_USERNAME")
	pass := config.GetConfig("DB_PASSWORD")
	host := config.GetConfig("DB_HOST")
	port := config.GetConfig("DB_PORT")
	name := config.GetConfig("DB_NAME")

	// [user[:password]@][net[(addr)]]/dbname[?param1=value1&paramN=valueN]
	dts := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)
	DB, err = gorm.Open(mysql.Open(dts))

	if err != nil {
		log.Fatal("Error opening database: " + err.Error())
	}

	fmt.Println("Database connected")

	if err = DB.AutoMigrate(&model.TodoItem{}); err != nil {
		log.Fatal("Error with database automigration")
	}
	fmt.Println("Automigration complete")
}
