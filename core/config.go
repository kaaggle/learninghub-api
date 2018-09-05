package core

type Config struct {
	BaseURL string
	Secret  string
	Database
}

type Database struct {
	URL string
}

func NewConfig() *Config {
	return &Config{
		BaseURL:  "localhost:6000",
		Database: Database{"mongodb://school-system:school-system1@ds237192.mlab.com:37192/school-system"},
		Secret:   "MySECRET",
	}
}
