package db

import (
	"fmt"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var lock = &sync.Mutex{}
var singleton *DatabaseConnection

type DatabaseConnection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
	conn     *gorm.DB
}

func GetDatabaseInstance() *DatabaseConnection {
	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleton == nil {
			singleton = &DatabaseConnection{}
		}
	}
	return singleton
}

func (dc *DatabaseConnection) GetConnection() *gorm.DB {
	return dc.conn
}

func (dc *DatabaseConnection) Ping() (success bool, err error) {
	if dc.conn == nil {
		return false, fmt.Errorf("database connection is nil")
	}
	sqlDB, err := dc.conn.DB()
	if err != nil {
		return false, fmt.Errorf("failed to get database connection: %w", err)
	}
	err = sqlDB.Ping()
	if err != nil {
		return false, fmt.Errorf("failed to ping database: %w", err)
	}
	return true, nil
}

func (dc *DatabaseConnection) ConnectToDatabase() error {
	var maxRetries = 5
	var retryCount = 0

	// Retrieve environment variables
	dc.Host = os.Getenv("POSTGRES_HOST")
	dc.Port = os.Getenv("POSTGRES_PORT")
	dc.User = os.Getenv("POSTGRES_USER")
	dc.Password = os.Getenv("POSTGRES_PASSWORD")
	dc.DBname = os.Getenv("POSTGRES_DB")

	// Check if any of the required environment variables are missing
	if dc.Host == "" || dc.Port == "" || dc.User == "" || dc.Password == "" || dc.DBname == "" {
		panic("Error: One or more PostgreSQL environment variables are missing.")
	}

	// Construct the connection string
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dc.Host, dc.Port, dc.User, dc.Password, dc.DBname)

	fmt.Printf("Connecting to %s:%s as %s to %s\n", dc.Host, dc.Port, dc.User, dc.DBname)

	var db *gorm.DB
	var err error

	for retryCount < maxRetries {
		db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
		if err != nil {
			retryCount++
			fmt.Printf("(%d of %d) Failed to connect to database. Retrying in 5 seconds. Error: %s\n", retryCount, maxRetries, err.Error())
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	dc.conn = db
	fmt.Print("Successfully connected to database\n")

	return nil
}

func (dc *DatabaseConnection) CloseConnection() error {
	if dc.conn == nil {
		return fmt.Errorf("database connection is nil")
	}
	sqlDB, err := dc.conn.DB()
	if err != nil {
		return fmt.Errorf("failed to get database connection: %w", err)
	}
	err = sqlDB.Close()
	if err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}
	return nil
}
