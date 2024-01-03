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

	if err == nil {
		t.Errorf("Validation failed to catch missing AppDir")
	}

	config.AppDir = "/"

	err = config.Validate()
	if err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func Test_Get_App_Working_Dir(t *testing.T) {
	workingDir, err := GetAppWorkingDir()
	if err != nil {
		t.Errorf("Could not get app working dir: %v", err)
	}
	if workingDir == "" {
		t.Errorf("App working dir should not be empty")
	}
	if workingDir == "/" {
		t.Errorf("App working dir should not be root")
	}
}
