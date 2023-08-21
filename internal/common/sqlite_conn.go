package common

import (
	"database/sql"
	"fmt"

	// _ "github.com/lib/pq"

	// _ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"

	log "github.com/sirupsen/logrus"
)

// const DBFile = "flotta.db"

// const ConnectionInfo = "user=root password=Password12 host=localhost port=3306 dbname=flotta sslmode=disable" //for postgres
const ConnectionInfo = "root:Password12!@tcp(localhost:3306)/flotta"

// connStr := "user:password@tcp(host:port)/dbname"
//sqlite
/*
func SetupSqliteDB1() {

	log.Info("Setup end nodes sqlite local database")
	// Check if the database file already exists
	if _, err := os.Stat(DBFile); err == nil {
		log.Error("Database file already exists, skipping creation...\n")
	} else if os.IsNotExist(err) {
		// Create the SQLite database file if it doesn't exist
		file, err := os.Create(DBFile)
		if err != nil {
			log.Errorf("Error creating database file: %s \n", err.Error())
			return
		}
		file.Close()
		log.Info("Database file created successfully.")
	} else {
		log.Errorf("Error checking database file: %s \n", err.Error())
		return
	}

	// Open a connection to the SQLite database

	db, err := SQLiteConnect(DBFile)
	if err != nil {
		log.Errorf("Error openning sqlite database file: %s\n", err.Error())
	}
	defer db.Close()

	// Create a table if it doesn't exist
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS wireless_device (
		wireless_device_id INTEGER PRIMARY KEY AUTOINCREMENT,
		wireless_device_name TEXT NOT NULL,
		wireless_device_manufacturer TEXT NULL,
		wireless_device_model TEXT NULL,
		wireless_device_sw_version TEXT NULL,
		wireless_device_identifier TEXT NOT NULL,
		wireless_device_protocol TEXT NULL,
		wireless_device_connection TEXT NULL,
		wireless_device_battery TEXT NULL,
		wireless_device_availability TEXT NULL,
		wireless_device_description TEXT NULL,
		wireless_device_last_seen TEXT NOT NULL
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Errorf("Error creating table: %s \n", err.Error())
		return
	}

	// Create a table if it doesn't exist
	createTableSQL = `
		CREATE TABLE IF NOT EXISTS device_property (
			device_property_id INTEGER PRIMARY KEY AUTOINCREMENT,
			wireless_device_identifier TEXT NOT NULL,
			property_identifier TEXT NOT NULL,

			property_service_uuid TEXT NULL,
			property_name TEXT NOT NULL,
			property_access_mode TEXT NOT NULL,
			property_reading TEXT NULL,
			property_state TEXT NULL,
			property_unit TEXT NULL,
			property_description TEXT NULL,
			property_last_seen TEXT NULL
		);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Errorf("Error creating table: %s \n", err.Error())
		return
	}

	// Create a table if it doesn't exist
	createTableSQL = `
		CREATE TABLE IF NOT EXISTS known_device (
			known_device_id INTEGER PRIMARY KEY AUTOINCREMENT,
			wireless_device_identifier TEXT NOT NULL,
			wireless_device_name TEXT NULL
		);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Errorf("Error creating table: %s \n", err.Error())
		return
	}

	log.Info("Table created successfully or already exists!")
}

func SQLiteConnect1(dbFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SQLite database: %w", err)
	}

	// Set connection properties (optional)
	db.SetMaxOpenConns(10) // Set the maximum number of open connections
	db.SetMaxIdleConns(5)  // Set the maximum number of idle connections

	// Perform a simple query to ensure the connection is valid
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping the database: %w", err)
	}

	// log.Info("Connected to SQLite database successfully")
	return db, nil
}
*/

//general database
func SetupSqliteDB() {

	// Open a connection to the SQLite database

	db, err := DbConnect(ConnectionInfo)
	if err != nil {
		log.Errorf("Error openning connecting to database: %s\n", err.Error())
	}
	defer db.Close()

	// Create a table if it doesn't exist
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS wireless_device (
		wireless_device_id INTEGER PRIMARY KEY AUTO_INCREMENT,
		wireless_device_name TEXT NOT NULL,
		wireless_device_manufacturer TEXT NULL,
		wireless_device_model TEXT NULL,
		wireless_device_sw_version TEXT NULL,
		wireless_device_identifier TEXT NOT NULL,
		wireless_device_protocol TEXT NULL,
		wireless_device_connection TEXT NULL,
		wireless_device_battery TEXT NULL,
		wireless_device_availability TEXT NULL,
		wireless_device_description TEXT NULL,
		wireless_device_last_seen TEXT NOT NULL
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Errorf("Error creating table: %s \n", err.Error())
		return
	}

	// Create a table if it doesn't exist
	createTableSQL = `
		CREATE TABLE IF NOT EXISTS device_property (
			device_property_id INTEGER PRIMARY KEY AUTO_INCREMENT,
			wireless_device_identifier TEXT NOT NULL,
			property_identifier TEXT NOT NULL,
			
			property_service_uuid TEXT NULL,
			property_name TEXT NOT NULL,
			property_access_mode TEXT NOT NULL,
			property_reading TEXT NULL,
			property_state TEXT NULL,
			property_unit TEXT NULL,
			property_description TEXT NULL,
			property_last_seen TEXT NULL
		);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Errorf("Error creating table: %s \n", err.Error())
		return
	}

	// Create a table if it doesn't exist
	createTableSQL = `
		CREATE TABLE IF NOT EXISTS known_device (
			known_device_id INTEGER PRIMARY KEY AUTO_INCREMENT,
			wireless_device_identifier TEXT NOT NULL,
			wireless_device_name TEXT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP , 
			updated_at DATETIME on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Errorf("Error creating table: %s \n", err.Error())
		return
	}

	createTableSQL = `
	CREATE TABLE IF NOT EXISTS last_sync (
		last_sync_id INTEGER PRIMARY KEY NOT NULL AUTO_INCREMENT , 
		last_sync_date_time DATETIME NOT NULL , 
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP , 
		updated_at DATETIME on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Errorf("Error creating table: %s \n", err.Error())
		return
	}
	log.Info("Table created successfully or already exists!")
}

func DbConnect(connStr string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to local database: %w", err)
	}

	// Set connection properties (optional)
	db.SetMaxOpenConns(50) // Set the maximum number of open connections
	db.SetMaxIdleConns(5)  // Set the maximum number of idle connections

	// Perform a simple query to ensure the connection is valid
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping the database: %w", err)
	}

	log.Info("Connected to local database successfully")
	return db, nil
}
