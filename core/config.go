package core

import (
	"errors"
	"fmt"
	"os"
)

type Config struct {
	BaseURL        string
	Secret         string
	CasbinConfPath string
	Database
}

type Database struct {
	URL string
}

func NewConfig() (*Config, error) {
	baseURLEnv := os.Getenv("LEARNINGHUB_API_URL")
	dbURL := os.Getenv("KAAGGLE_DB_URL")
	secret := os.Getenv("KAAGGLE_SECRET")
	//casbinConfPath := os.Getenv("CHALK_CASBIN_CONF_PATH")
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	casbinConfPath := dir + "/authorization/conf/"

	if baseURLEnv != "" && dbURL != "" && secret != "" {
		return &Config{
			BaseURL:        baseURLEnv,
			Database:       Database{dbURL},
			CasbinConfPath: casbinConfPath,
			Secret:         secret,
		}, nil
	}

	return nil, errors.New("Please add LEARNINGHUB_API_URL, KAAGGLE_DB_URL and KAAGGLE_SECRET environmental variables")
}

func (c *Config) String() string {
	return fmt.Sprintf("Using config with the following details. URL: %s. DB_URL: %s.", c.BaseURL, c.Database.URL)
}
