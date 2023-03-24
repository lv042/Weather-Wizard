package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//dbManager struct

type dbManager struct {
	db         *gorm.DB
	name       string
	insertSql  string
	rebuildSql string
}

// NewDBManager constructor for dbManager
func NewDBManager(db *gorm.DB, name string) *dbManager {
	return &dbManager{db: db, name: name, insertSql: "insert.sql", rebuildSql: "rebuild.sql"}
}

//basic functions

func (d *dbManager) SetDBManager(db *gorm.DB, name string) {
	d.SetName(name)
	d.SetDB(db)
}

func (d *dbManager) SetDB(db *gorm.DB) {
	d.db = db
}

func (d *dbManager) SetName(name string) {
	d.name = name
}

func (d *dbManager) GetDB() *gorm.DB {
	return d.db
}

func (d *dbManager) GetName() string {
	return d.name
}

//setup function

func (d *dbManager) setupDb() {
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	var err error
	d.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Default().Println("Connected to database")

	//show tables
	d.logAllTables()

	//run sql file
	d.runSqlSetupFiles()

	//show tables again
	d.logAllTables()

	//fill db
	d.fillDB()

	//show all tables content

}

//Additional functions

func (d *dbManager) logAllTables() {
	rows, err := d.db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema='ws'").Rows()
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

func (d *dbManager) fillDB() {
	defer func() {
		if r := recover(); r != nil {
			log.Default().Println("fill_db() Error: ", r)
		}
	}()

	// Check if SQL file exists
	sqlFilePath := filepath.Join(".", "sql", d.insertSql)
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
	if err := d.db.Exec(sql).Error; err != nil {
		panic(err)
	}
	log.Default().Println("Executed ", d.insertSql)
}

func (d *dbManager) runSqlSetupFiles() {
	// Check if SQL file exists
	sqlFilePath := filepath.Join(".", "sql", d.rebuildSql)
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
	if err := d.db.Exec(sql).Error; err != nil {
		panic(err)
	}
	log.Default().Println("Executed ", d.rebuildSql)

}
