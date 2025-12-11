package main

// import (
// 	"log"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func main() {
// 	dsn := "host=localhost user=postgres dbname=test password=pass sslmode=disable"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// вызываем сиды
// 	SeedUsers(db)
// 	SeedProducts(db)

// 	log.Println("Seeding done!")
// }
