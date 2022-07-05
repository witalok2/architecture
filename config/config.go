package config

import (
	"os"

	logger "github.com/sirupsen/logrus"
)

type DB struct {
	Database string
	User     string
	Password string
	Host     string
	Port     string
	SSL      string
}

type Env struct {
	Postgres DB
}

func LoadConfig() (*Env, error) {
	logger.Debug("Loading environment variables")

	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_DATABASE", "postgres")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_SSL_MODE", "disable")

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSL_MODE")


	return &Env{DB{
		Host:     host,
		User:     user,
		Password: password,
		Database: database,
		Port:     port,
		SSL:      sslmode,
	}}, nil
}

// func LoadConfig() (*Env, error) {
// 	logger.Debug("Loading environment variables")

// 	host, ok := os.LookupEnv("DB_HOST")
// 	if !ok {
// 		return nil, errors.New("env var isn't set: DB_HOST")
// 	}

// 	user, ok := os.LookupEnv("DB_USER")
// 	if !ok {
// 		return nil, errors.New("env var isn't set: DB_USER")
// 	}

// 	password, ok := os.LookupEnv("DB_PASSWORD")
// 	if !ok {
// 		return nil, errors.New("env var isn't set: DB_PASSWORD")
// 	}

// 	database, ok := os.LookupEnv("DB_DATABASE")
// 	if !ok {
// 		return nil, errors.New("env var isn't set: DB_DATABASE")
// 	}

// 	port, ok := os.LookupEnv("DB_PORT")
// 	if !ok {
// 		return nil, errors.New("env var isn't set: DB_PORT")
// 	}

// 	sslmode, ok := os.LookupEnv("DB_SSL_MODE")
// 	if !ok {
// 		return nil, errors.New("env var isn't set: DB_SSL_MODE")
// 	}

// 	return &Env{DB{
// 		Host: host,
// 		User: user,
// 		Password: password,
// 		Database: database,
// 		Port: port,
// 		SSL: sslmode,
// 	}}, nil
// }
