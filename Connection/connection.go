package connection

import (
	"fmt"
	models "order-api/Models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBInitialize(dbUsername, dbPass, dbPort, dbName, dbHost string) *gorm.DB {
	// Create Conn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUsername, dbPass, dbHost, dbPort, dbName)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(models.Order{}, models.Item{})
	return db

}
