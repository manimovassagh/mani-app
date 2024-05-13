package main

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name    string
	IsAdmin bool
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=mani port=5432 sslmode=disable "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Student{})

	// Create
	db.Create(&Student{Name: "Mani", IsAdmin: false})

	// Read
	var student Student
	db.First(&student, 1) // find product with integer primary key

}
