package main

import (
	"encoding/json"
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
	var data []map[string]interface{}
	result := d.db.Table("weather_data").Find(&data)
	if result.Error != nil {
		d.Log(result.Error.Error())
		return
	}
	if result.RowsAffected == 0 {
		d.Log("No weather data found")
		return
	}
	for _, item := range data {
		jsonString, err := json.Marshal(item)
		if err != nil {
			d.Log(err.Error())
			continue
		}
		d.Log(string(jsonString))
	}
}

func (d *dbManager) getWeatherData() (string, error) {
	var data []map[string]interface{}
	result := d.db.Table("weather_data").Find(&data)
	if result.Error != nil {
		return "", result.Error
	}
	if result.RowsAffected == 0 {
		return "", nil
	}
	jsonString, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonString), nil
}

//Crud operations

type WeatherData struct {
	Timestamp int64                  `json:"timestamp" gorm:"primaryKey"`
	Data      map[string]interface{} `json:"data"`
}

// Read a weather data record by timestamp
func (d *dbManager) readWeatherData(timestamp string) ([]map[string]interface{}, error) {
	var data []map[string]interface{}
	result := d.db.Table("weather_data").Where("timestamp = ?", timestamp).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

// Create a new weather data entry with the given timestamp and data
func (d *dbManager) createWeatherData(timestamp string, data string) error {
	result := d.db.Table("weather_data").Create(map[string]interface{}{
		"timestamp": timestamp,
		"data":      data,
	})
	return result.Error
}

// Update the weather data entry with the given timestamp
func (d *dbManager) updateWeatherData(timestamp string, data string) error {
	result := d.db.Table("weather_data").Where("timestamp = ?", timestamp).Update("data", data)
	return result.Error
}

// Delete the weather data entry with the given timestamp
func (d *dbManager) deleteWeatherData(timestamp string) error {
	result := d.db.Table("weather_data").Where("timestamp = ?", timestamp).Delete(&map[string]interface{}{})
	return result.Error
}
