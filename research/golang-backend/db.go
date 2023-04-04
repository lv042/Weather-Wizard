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
	// Log a message with the prefix "DatabaseManager: "
	log.Default().Println("DatabaseManager: ", s)
}

func (d *DBManager) LogError(message string) {
	// Create a red color function for log messages
	red := color.New(color.FgRed).SprintFunc()

	// Log an error message with the prefix "DatabaseManager: " and in red color
	log.Fatal("DatabaseManager: ", red(message))
}

func (d *DBManager) GetInfo() {
	// Log the string representation of the DBManager struct
	d.Log(fmt.Sprintf("%+v", d))
}

//sql functions

func (d *DBManager) Query(sql string) *gorm.DB {
	result := d.db.Raw(sql).Scan(&gorm.Model{})
	return result
}

//Additional functions

func (d *DBManager) logAllTables() {
	// Execute a raw SQL query to retrieve the names of all tables in the "ws" schema
	rows, err := d.db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Rows()
	if err != nil {
		// If there was an error executing the query, log the error message
		d.Log(err.Error())
	}

	// Initialize a string to store the log message
	var log = "Tables: "

	// Loop through the query results
	for rows.Next() {
		var table_name string
		rows.Scan(&table_name)
		// Append the table name to the log message
		log += table_name + ", "
	}

	// Log the log message
	d.Log(log)
}

func (d *DBManager) fillDB() {
	// Define a function to be executed when the main function returns (i.e. in case of panic)
	defer func() {
		if r := recover(); r != nil {
			// If a panic occurred, log a message indicating that the database fill failed
			d.Log(fmt.Sprintf("Failed to fill database: %s", r))
		}
	}()

	// Check if the specified SQL file exists
	sqlFilePath := filepath.Join(".", "sql", d.insertSql)
	if _, err := os.Stat(sqlFilePath); os.IsNotExist(err) {
		// If the file does not exist, panic with an error message
		panic(fmt.Sprintf("SQL file '%s' not found", sqlFilePath))
	}

	// Read the contents of the SQL file
	sqlBytes, err := ioutil.ReadFile(sqlFilePath)
	if err != nil {
		// If there was an error reading the file, panic with the error
		panic(err)
	}
	sql := string(sqlBytes)

	// Execute the SQL statements contained in the file
	if err := d.db.Exec(sql).Error; err != nil {
		// If there was an error executing the SQL, panic with the error
		panic(err)
	}

	// Log a message indicating that the SQL was executed successfully
	d.Log("Executed " + d.insertSql)
}

func (d *DBManager) runSqlSetupFiles() {
	// Check if the specified SQL file exists
	sqlFilePath := filepath.Join(".", "sql", d.rebuildSql)
	if _, err := os.Stat(sqlFilePath); os.IsNotExist(err) {
		// If the file does not exist, log an error message
		d.LogError(fmt.Sprintf("SQL file '%s' not found", sqlFilePath))
		return
	}

	// Read the contents of the SQL file
	sqlBytes, err := ioutil.ReadFile(sqlFilePath)
	if err != nil {
		// If there was an error reading the file, log the error message
		d.LogError(err.Error())
		return
	}
	sql := string(sqlBytes)

	// Execute the SQL statements contained in the file
	if err := d.db.Exec(sql).Error; err != nil {
		// If there was an error executing the SQL, log the error message
		d.LogError(err.Error())
		return
	}

	// Log a message indicating that the SQL was executed successfully
	d.Log("Executed " + d.rebuildSql)
}

func (d *DBManager) logWeatherData() {
	// Initialize a slice of maps to store the weather data
	var data []map[string]interface{}

	// Find the weather data from the "weather_data" table
	result := d.db.Table("weather_data").Find(&data)

	// Check if there was an error while finding the data
	if result.Error != nil {
		// Log the error message
		d.Log(result.Error.Error())
		return
	}

	// Check if no rows were affected (i.e. no data was found)
	if result.RowsAffected == 0 {
		// Log a message indicating that no data was found
		d.Log("No weather data found")
		return
	}

	// Loop through the data and log each item as a JSON string
	for _, item := range data {
		// Convert the map to a JSON string
		jsonString, err := json.Marshal(item)

		// Check if there was an error while converting to JSON
		if err != nil {
			// Log the error message
			d.Log(err.Error())
			continue
		}

		// Log the JSON string
		d.Log(string(jsonString))
	}
}

// WeatherData ORM
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
	// Validate the format of the input timestamp using a regular expression
	validTimestamp := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z$`)
	if !validTimestamp.MatchString(timestamp) {
		// Return an error if the format is invalid
		return "", fmt.Errorf("Invalid timestamp format: %s", timestamp)
	}

	// Parse the timestamp string into a time.Time value
	timestampParsed, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		// Return the error if there was a problem parsing the timestamp
		return "", err
	}

	// Retrieve the weather data with the specified timestamp
	var weatherData WeatherData
	result := d.db.Table("weather_data").Where("timestamp = ?", timestampParsed).First(&weatherData)
	if result.Error != nil {
		// Return the error if there was a problem retrieving the data
		return "", result.Error
	}

	// Check if no rows were affected (i.e. no data was found)
	if result.RowsAffected == 0 {
		// Return a message indicating that no data was found
		return "No weather data found for the specified timestamp", nil
	}

	// Marshal the weather data into a JSON string
	weatherDataJSON, err := json.Marshal(weatherData)
	if err != nil {
		// Return the error if there was a problem marshaling the data
		return "", err
	}

	// Return the JSON string
	return string(weatherDataJSON), nil
}

// GetAllWeatherDataJSON Get all weather data as JSON
func (d *DBManager) GetAllWeatherDataJSON() ([]byte, error) {
	// Retrieve all weather data records
	var data []WeatherData
	result := d.db.Table("weather_data").Find(&data)
	if result.Error != nil {
		// Return the error if there was a problem retrieving the data
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		// Return an error if no data was found
		return nil, errors.New("no weather data found")
	}

	// Marshal the weather data into a JSON string
	jsonData, err := json.Marshal(data)
	if err != nil {
		// Return the error if there was a problem marshaling the data
		return nil, err
	}

	// Return the JSON string
	return jsonData, nil
}

// DeleteWeatherDataJSON Delete weather data by timestamp
func (d *DBManager) DeleteWeatherDataJSON(jsonStr string) (string, error) {
	// Unmarshal the JSON string into a struct
	var data struct {
		Timestamp string `json:"timestamp"`
	}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		// Return the error if there was a problem unmarshaling the JSON
		return "", err
	}

	// Parse the "timestamp" field into a time.Time value
	timestamp, err := time.Parse(time.RFC3339, data.Timestamp)
	if err != nil {
		// Return the error if there was a problem parsing the timestamp
		return "", err
	}

	// Delete the weather data with the specified timestamp
	result := d.db.Table("weather_data").Where("timestamp = ?", timestamp).Delete(&WeatherData{})
	if result.Error != nil {
		// Return the error if there was a problem deleting the data
		return "", result.Error
	}

	// Check if no rows were affected (i.e. no data was found)
	if result.RowsAffected == 0 {
		// Return a message indicating that no data was found
		return "No weather data found for the specified timestamp", nil
	}

	// Return a success message
	return "Weather data deleted", nil
}

// UpdateWeatherDataJSON Update weather data by timestamp
func (d *DBManager) UpdateWeatherDataJSON(jsonStr string) (string, error) {
	// Unmarshal the JSON string into a `WeatherData` struct
	var data WeatherData
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		// Return the error if there was a problem unmarshaling the JSON
		return "", err
	}

	// Retrieve the existing weather data with the same timestamp
	existingData := WeatherData{}
	result := d.db.Table("weather_data").Where("timestamp = ?", data.Timestamp).First(&existingData)
	if result.Error != nil {
		// Return the error if there was a problem retrieving the existing data
		return "", result.Error
	}
	if result.RowsAffected == 0 {
		// Return a message indicating that no data was found
		return "No weather data found for the specified timestamp", nil
	}

	// Update the existing weather data
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
		// Return the error if there was a problem updating the data
		return "", result.Error
	}

	// Return a success message
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
