package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
}

type SecurityConfig struct {
	Key    string
	Durasi time.Duration
	Issuer string
}

type AppConfig struct {
	AppPort string
}
type MidtransConfig struct {
	ServerKey string
}
type Config struct {
	DbConfig
	AppConfig
	SecurityConfig
	MidtransConfig
}

func (c *Config) readConfig() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c.AppConfig = AppConfig{
		AppPort: os.Getenv("PORT_APP"),
	}
	lifetime, err := strconv.Atoi(os.Getenv("JWT_LIFE_TIME"))
	if err != nil {
		return err
	}
	c.SecurityConfig = SecurityConfig{
		Key:    os.Getenv("KEY"),
		Durasi: time.Duration(lifetime) * time.Second,
		Issuer: os.Getenv("ISSUER"),
	}
	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
	}
	c.MidtransConfig = MidtransConfig{
		ServerKey: os.Getenv("MIDTRANS_SERVER_KEY"),
	}

	if c.DbConfig.Host == "" || c.DbConfig.Port == "" || c.DbConfig.Name == "" || c.DbConfig.User == "" ||
		c.DbConfig.Password == "" || c.DbConfig.Driver == "" ||
		c.SecurityConfig.Key == "" || c.SecurityConfig.Durasi < 0 || c.SecurityConfig.Issuer == "" ||
		c.MidtransConfig.ServerKey == "" {
		return err
	}
	return nil
}

func NewConfig() (*Config, error) {
	c := &Config{}
	err := c.readConfig()
	if err != nil {
		return nil, err
	}
	return c, nil

}
