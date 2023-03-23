package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"path/filepath"
)

var db *gorm.DB

func main() {
	setup_db()
	app := fiber.New()

	// Define routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}

func init_() {

}

func setup_db() {
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Default().Println("Connected to database")

	//show tables
	log_all_tables()

	//run sql file
	run_sql_setup_files()

}

func run_sql_setup_files() {
	// Run SQL file
	sqlFilePath := filepath.Join(".", "schema.sql")
	sqlBytes, err := ioutil.ReadFile(sqlFilePath)
	if err != nil {
		panic(err)
	}
	sql := string(sqlBytes)
	if err := db.Exec(sql).Error; err != nil {
		panic(err)
	}
	log.Default().Println("Executed schema.sql")

	// Show tables again
	log_all_tables()
}

func log_all_tables() {
	rows, err := db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema='public'").Rows()
	if err != nil {
		log.Default().Println(err)
	}

	for rows.Next() {
		var table_name string
		rows.Scan(&table_name)
		log.Default().Println(table_name)
	}
}

func print_all_routes() {

}
