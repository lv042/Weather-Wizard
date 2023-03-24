package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var insertSql = "insert.sql"
var rebuildSql = "rebuild.sql"

type dbManager struct {
	db *gorm.DB
}

func main() {
	setup_db()
	fiberApp = fiber.New()

	// Define routes
	fiberApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	fiberApp.Listen(":3000")

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
	runSqlSetupFiles()

	//show tables again
	log_all_tables()

	//fill db
	fill_db()

	//show all tables content

}

func runSqlSetupFiles() {
	// Check if SQL file exists
	sqlFilePath := filepath.Join(".", "sql", rebuildSql)
	if _, err := os.Stat(sqlFilePath); os.IsNotExist(err) {
		panic(fmt.Sprintf("SQL file '%s' not found", sqlFilePath))
	}

	// Read SQL file
	sqlBytes, err := ioutil.ReadFile(sqlFilePath)
	if err != nil {
		panic(err)
	}
	sql := string(sqlBytes)

	// Execute SQL
	if err := db.Exec(sql).Error; err != nil {
		panic(err)
	}
	log.Default().Println("Executed ", rebuildSql)

}

func fill_db() {
	defer func() {
		if r := recover(); r != nil {
			log.Default().Println("fill_db() Error: ", r)
		}
	}()

	// Check if SQL file exists
	sqlFilePath := filepath.Join(".", "sql", insertSql)
	if _, err := os.Stat(sqlFilePath); os.IsNotExist(err) {
		panic(fmt.Sprintf("SQL file '%s' not found", sqlFilePath))
	}

	// Read SQL file
	sqlBytes, err := ioutil.ReadFile(sqlFilePath)
	if err != nil {
		panic(err)
	}
	sql := string(sqlBytes)

	// Execute SQL
	if err := db.Exec(sql).Error; err != nil {
		panic(err)
	}
	log.Default().Println("Executed ", insertSql)
}

func log_all_tables() {
	rows, err := db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema='ws'").Rows()
	if err != nil {
		log.Default().Println(err)
	}
	log.Default().Println("Tables: ")
	for rows.Next() {
		var table_name string
		rows.Scan(&table_name)
		log.Default().Println(table_name)
	}
	//line break
	log.Default().Println()
}
