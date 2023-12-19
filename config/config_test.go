package config

import (
	"testing"
)

func Test_Config_Validate(t *testing.T) {
	config := &Config{
		AppHost:    "localhost",
		AppPort:    "8080",
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:     "testuser",
		DBPassword: "testpassword",
		DBName:     "testdb",
	}
	err := config.Validate()

	if err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}
