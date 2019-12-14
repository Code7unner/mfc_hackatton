package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func Connect() *sql.DB {
	//config, err := loadConfig()
	config, err := loadHerokuConfig()
	if err != nil {
		log.Fatalf("Error loading config.env file: %v", err)
	}

	//db, err := initDatabase(config)
	db, err := initDatabaseWithHeroku(config)
	if err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	return db
}

// Config represents structure of the config.env
type Config struct {
	dbUser string
	dbPass string
	dbName string
	dbHost string
	dbPort string
}

func loadConfig() (config *Config, err error) {
	err = godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading config.env file")
	}

	config = &Config {
		dbUser : os.Getenv("db_user"),
		dbPass : os.Getenv("db_pass"),
		dbName : os.Getenv("db_name"),
		dbHost : os.Getenv("db_host"),
		dbPort : os.Getenv("db_port"),
	}

	return config, err
}

func initDatabase(c *Config) (db *sql.DB, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.dbHost, c.dbPort, c.dbUser, c.dbPass, c.dbName)

	log.Println(psqlInfo)

	db, err = sql.Open("postgres", psqlInfo)
	return db, err
}

func initDatabaseWithHeroku(c *HerokuConfig) (db *sql.DB, err error) {
	db, err = sql.Open("postgres", c.dbUrl)
	if err != nil {
		return nil, err
	}

	return db, nil
}

type HerokuConfig struct {
	dbUrl string
}

func loadHerokuConfig() (config *HerokuConfig, err error) {
	err = godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading heroku.env file")
	}

	config = &HerokuConfig{dbUrl: os.Getenv("DATABASE_URL")}

	return config, err
}