package config

type Config struct {
	Port     int
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     int
	Name     string
	Username string
	Password string
}
