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
func NewDBManager(name string) *dbManager {
	var d = dbManager{db: nil, name: name, insertSql: "insert.sql", rebuildSql: "rebuild.sql"}
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	var err error
	d.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		d.Log("Failed to connect database")
	}

	d.Log("Connected to database")

	return &d
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

func (d *dbManager) ToString() string {
	return fmt.Sprintf("Running %s ", d.name)
}

//setup function

func (d *dbManager) setupDb() {
	//run sql file
	d.runSqlSetupFiles()

	//show tables again
	d.logAllTables()

	//fill db
	d.fillDB()

	//show tables
	d.logAllTables()
}

//Misc functions

func (d *dbManager) Close() {
	db, err := d.db.DB()
	if err != nil {
		d.Log(err.Error())
	}
	db.Close()
}

func (d *dbManager) Log(s string) {
	log.Default().Println("DatabaseManager: ", s)
}

func (d *dbManager) GetInfo() {
	d.Log(fmt.Sprintf("%+v", d))
}

//sql functions

func (d *dbManager) Query(sql string) *gorm.DB {
	result := d.db.Raw(sql).Scan(&gorm.Model{})
	return result
}

//Additional functions

func (d *dbManager) logAllTables() {
	rows, err := d.db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema='ws'").Rows()
	if err != nil {
		d.Log(err.Error())
	}
	var log = "Tables: "

	for rows.Next() {
		var table_name string
		rows.Scan(&table_name)
		log += table_name + ", "
	}
	d.Log(log)
}

func (d *dbManager) fillDB() {
	defer func() {
		if r := recover(); r != nil {
			d.Log(fmt.Sprintf("Failed to fill database: %s", r))
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
	d.Log("Executed " + d.insertSql)
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
	d.Log("Executed " + d.rebuildSql)

}

func (d *dbManager) logWeatherData() {
	//log the most recent weather data entries from the table weather_data
	rows, err := d.db.Raw("SELECT * FROM weather_data").Rows()
	if err != nil {
		d.Log(err.Error())
	}
	//return rows and  log them
	//print rows
	for rows.Next() {
		var id int
		var temperature float64
		var humidity float64
		var pressure float64
		var date string
		rows.Scan(&id, &temperature, &humidity, &pressure, &date)
		d.Log(fmt.Sprintf("ID: %d, Temperature: %f, Humidity: %f, Pressure: %f, Date: %s", id, temperature, humidity, pressure, date))
	}
}

func (d *dbManager) addWeatherData(data string) {
	//add weather data to the table we	ather_data
	d.db.Create(data)
}
