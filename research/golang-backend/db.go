package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

//DBManager struct

type DBManager struct {
	db         *gorm.DB
	name       string
	insertSql  string
	rebuildSql string
}

// NewDBManager constructor for DBManager
func NewDBManager(name string, dsn string) *DBManager {
	var d = DBManager{db: nil, name: name, insertSql: "insert.sql", rebuildSql: "rebuild.sql"}

	var err error
	d.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		d.LogError("Failed to connect database")
	}

	d.Log("Connected to database")

	return &d
}

//basic functions

func (d *DBManager) SetDBManager(db *gorm.DB, name string) {
	d.SetName(name)
	d.SetDB(db)
}

func (d *DBManager) SetDB(db *gorm.DB) {
	d.db = db
}

func (d *DBManager) SetName(name string) {
	d.name = name
}

func (d *DBManager) GetDB() *gorm.DB {
	return d.db
}

func (d *DBManager) GetName() string {
	return d.name
}

func (d *DBManager) ToString() string {
	return fmt.Sprintf("Running %s ", d.name)
}

//setup function

func (d *DBManager) setupDb() {
	//run sql files
	d.runSqlSetupFiles()

	//show tables
	d.logAllTables()

	//fill db
	d.fillDB()

	//show tables again
	d.logAllTables()
}

//Misc functions

func (d *DBManager) Close() {
	db, err := d.db.DB()
	if err != nil {
		d.Log(err.Error())
	}
	db.Close()
}

func (d *DBManager) Log(s string) {
	log.Default().Println("DatabaseManager: ", s)
}

func (d *DBManager) LogError(message string) {
	red := color.New(color.FgRed).SprintFunc()
	log.Fatal("DatabaseManager: ", red(message))
}

func (d *DBManager) GetInfo() {
	d.Log(fmt.Sprintf("%+v", d))
}

//sql functions

func (d *DBManager) Query(sql string) *gorm.DB {
	result := d.db.Raw(sql).Scan(&gorm.Model{})
	return result
}

//Additional functions

func (d *DBManager) logAllTables() {
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

func (d *DBManager) fillDB() {
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

func (d *DBManager) runSqlSetupFiles() {
	// Check if SQL file exists
	sqlFilePath := filepath.Join(".", "sql", d.rebuildSql)
	if _, err := os.Stat(sqlFilePath); os.IsNotExist(err) {
		d.LogError(fmt.Sprintf("SQL file '%s' not found", sqlFilePath))
		return
	}

	// Read SQL file
	sqlBytes, err := ioutil.ReadFile(sqlFilePath)
	if err != nil {
		d.LogError(err.Error())
		return
	}
	sql := string(sqlBytes)

	// Execute SQL
	if err := d.db.Exec(sql).Error; err != nil {
		d.LogError(err.Error())
		return
	}

	d.Log("Executed " + d.rebuildSql)
}

func (d *DBManager) logWeatherData() {
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

// WeatherData ORM with GORM
type WeatherData struct {
	Timestamp        time.Time `gorm:"column:timestamp"`
	Temperature      float64   `gorm:"column:temperature"`
	Humidity         float64   `gorm:"column:humidity"`
	Pressure         float64   `gorm:"column:pressure"`
	ObstacleDetected bool      `gorm:"column:obstacle_detected"`
	LightIntensity   float64   `gorm:"column:light_intensity"`
}

// Crud operations

// GetWeatherDataByTimestampJSON Get weather data by timestamp
func (d *DBManager) GetWeatherDataByTimestampJSON(timestamp string) (string, error) {
	// Validate the timestamp format using a regular expression
	validTimestamp := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z$`)
	if !validTimestamp.MatchString(timestamp) {
		return "", fmt.Errorf("Invalid timestamp format: %s", timestamp)
	}

	// Parse the timestamp
	timestampParsed, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return "", err
	}

	var weatherData WeatherData
	result := d.db.Table("weather_data").Where("timestamp = ?", timestampParsed).First(&weatherData)
	if result.Error != nil {
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		return "No weather data found for the specified timestamp", nil
	}

	weatherDataJSON, err := json.Marshal(weatherData)
	if err != nil {
		return "", err
	}

	return string(weatherDataJSON), nil
}

// GetAllWeatherDataJSON Get all weather data as JSON
func (d *DBManager) GetAllWeatherDataJSON() ([]byte, error) {
	var data []WeatherData
	result := d.db.Table("weather_data").Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("no weather data found")
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

// DeleteWeatherDataJSON Delete weather data by timestamp
func (d *DBManager) DeleteWeatherDataJSON(jsonStr string) (string, error) {
	var data struct {
		Timestamp string `json:"timestamp"`
	}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", err
	}

	timestamp, err := time.Parse(time.RFC3339, data.Timestamp)
	if err != nil {
		return "", err
	}

	result := d.db.Table("weather_data").Where("timestamp = ?", timestamp).Delete(&WeatherData{})
	if result.Error != nil {
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		return "No weather data found for the specified timestamp", nil
	}

	return "Weather data deleted", nil
}

// UpdateWeatherDataJSON Update weather data by timestamp
// UpdateWeatherDataJSON Update weather data by timestamp
func (d *DBManager) UpdateWeatherDataJSON(jsonStr string) (string, error) {
	var data WeatherData
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", err
	}

	// Check if data with the same timestamp already exists
	existingData := WeatherData{}
	result := d.db.Table("weather_data").Where("timestamp = ?", data.Timestamp).First(&existingData)
	if result.Error != nil {
		return "", result.Error
	}
	if result.RowsAffected == 0 {
		return "No weather data found for the specified timestamp", nil
	}

	// Update the weather data
	result = d.db.Table("weather_data").
		Where("timestamp = ?", data.Timestamp).
		Updates(WeatherData{
			Temperature:      data.Temperature,
			Humidity:         data.Humidity,
			Pressure:         data.Pressure,
			ObstacleDetected: data.ObstacleDetected,
			LightIntensity:   data.LightIntensity,
		})
	if result.Error != nil {
		return "", result.Error
	}

	return "Weather data updated", nil
}

// CreateWeatherDataJSON Create weather data
func (d *DBManager) CreateWeatherDataJSON(jsonStr string) (string, error) {
	var data WeatherData
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", err
	}

	// Check if data with the same timestamp already exists
	existingData := WeatherData{}
	result := d.db.Table("weather_data").Where("timestamp = ?", data.Timestamp).First(&existingData)
	if result.RowsAffected > 0 {
		return "Weather data already exists for the specified timestamp", nil
	}

	result = d.db.Table("weather_data").Create(&data)
	if result.Error != nil {
		return "", result.Error
	}

	return "Weather data created", nil
}
