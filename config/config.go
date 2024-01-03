package config

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sync"

	"github.com/gemurdock/KeyFinder-GoLang/util"
)

var lock = &sync.Mutex{}
var singleton *Config

type Config struct {
	AppHost    string
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	AppDir     string
}

type ConfigError struct {
	Message string
}

func GetConfigInstance(autoload bool) *Config {
	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleton == nil {
			singleton = &Config{}
			if autoload {
				singleton.Load()
			}
		}
	}
	return singleton
}

func GetAppWorkingDir() (string, error) {
	if singleton != nil && singleton.AppDir != "" {
		return singleton.AppDir, nil
	}
	lock.Lock()
	defer lock.Unlock()
	if singleton == nil {
		singleton = &Config{}
	}

	uniqueAppFiles := []string{
		"variables.env", "variables.env.default", "Makefile", "LICENSE",
	}
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

checkpath:
	filenames, err := util.GetAllFilesInPath(dir)
	if err != nil {
		return "", err
	}

	for _, filename := range filenames {
		for _, uniqueAppFile := range uniqueAppFiles {
			if filename == uniqueAppFile {
				return dir, nil
			}
		}
	}

	if dir != "/" { // Continue to check until root or app working dir is found
		dir = filepath.Dir(dir)
		goto checkpath
	}

	return "", fmt.Errorf("could not find app working directory")
}

func (c *Config) Load() {
	c.AppHost = os.Getenv("APP_HOST")
	c.AppPort = os.Getenv("APP_PORT")
	c.DBHost = os.Getenv("POSTGRES_HOST")
	c.DBPort = os.Getenv("POSTGRES_PORT")
	c.DBUser = os.Getenv("POSTGRES_USER")
	c.DBPassword = os.Getenv("POSTGRES_PASSWORD")
	c.DBName = os.Getenv("POSTGRES_DB")
	c.loadAppDir()

	c.Validate()
}

func (c *Config) LoadTestingValues() {
	c.AppHost = "localhost"
	c.AppPort = "3000"
	c.DBHost = "localhost"
	c.DBPort = "5432"
	c.DBUser = "postgres"
	c.DBPassword = "password"
	c.DBName = "keyfinder_test"
	c.loadAppDir()

	c.Validate()
}

func (c *Config) Validate() error {
	structValue := reflect.ValueOf(c).Elem()
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		if field.String() == "" {
			return fmt.Errorf("missing environment variable in Config.Validate()")
		}
	}
	return nil
}

func (c *Config) loadAppDir() {
	if c.AppDir == "" {
		appDir, err := GetAppWorkingDir()
		if err != nil {
			panic(err)
		}
		c.AppDir = appDir
	}
}
