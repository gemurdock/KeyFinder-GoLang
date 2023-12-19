package db

import (
	"fmt"
	"sync"
	"time"

	"github.com/gemurdock/KeyFinder-GoLang/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var lock = &sync.Mutex{}
var singleton *DatabaseConnection

type DatabaseConnection struct {
	config *config.Config
	conn   *gorm.DB
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

func (dc *DatabaseConnection) LoadConfig(config *config.Config) {
	dc.config = config
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
	var maxRetries = 3
	var retryCount = 0

	if singleton.config == nil {
		return fmt.Errorf("config is nil")
	}

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dc.config.DBHost, dc.config.DBPort, dc.config.DBUser, dc.config.DBPassword, dc.config.DBName)

	fmt.Printf("Connecting to %s:%s as %s\n", dc.config.DBHost, dc.config.DBPort, dc.config.DBUser)

	var db *gorm.DB
	var err error

	for retryCount < maxRetries {
		db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
		if err != nil {
			retryCount++
			fmt.Printf("(%d of %d) Failed to connect to database. Retrying in 5 seconds. Error: %s\n", retryCount, maxRetries, err.Error())
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}

	if retryCount == maxRetries || err != nil {
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
