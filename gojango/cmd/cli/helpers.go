package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func setup() {
	err := godotenv.Load()
	if err != nil {
		exitGracefully(err)
	}

	path, err := os.Getwd()
	if err != nil {
		exitGracefully(err)
	}

	gg.RootPath = path
	gg.DB.DataType = os.Getenv("DATABASE_TYPE")
}

func getDSN() string {
	dbType := gg.DB.DataType

	if dbType == "pgx" {
		dbType = "postgres"
	}

	var dsn string
	if dbType == "postgres" {
		if os.Getenv("DATABASE_PASS") != "" {
			dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_PASS"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"),
			)
		} else {
			dsn = fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"),
			)
		}
	} else {
		dsn = "mysql://" + gg.BuildDSN()
	}
	return dsn
}

func showHelp() {
	color.Yellow(`Available commands:
    help                  - show the help commands
    version               - print application version
    migrate               - runs all up migrations that have not been run
    migrate down          - reverses the most recent migration
    migrate reset         - runs all down migrations and then up
    make migration <name> - creates two new up and down migrations
    make auth             - creates and runs migrations for auth tables & middleware
    make handler <name>   - creates a stub handler in the handlers directory
    make model <name>     - creates a new model in the data directory
    make Session          - creates a table in the database as a session store
    make mail <name>      - creates two starter mail templates in the mail directory
    `)
}
