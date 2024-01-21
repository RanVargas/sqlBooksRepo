package migrations

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func RunMigrationZero() {
	log.Println("Running migration")
	err := godotenv.Load("D:\\Practice\\ran-vargas-naithergrand-2024-1-16-dev-backend-coding-challenge-library-api\\local.env")
	if err != nil {
		panic(err)
	}
	//
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sqlFilePath := os.Getenv("PATH_TO_SQL_FILE_MIGRATION")

	if sqlFilePath == "" {
		panic("No file for migration found")
	}
	//host := "localhost"
	//port := 5432
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, dbUser, dbPassword, dbName)
	//postgres://postgres:password@localhost:5432/example?sslmode=disable&search_path=public

	psqlAnother := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable&search_path=public", dbUser, dbPassword, dbName)
	fmt.Println(psqlAnother)
	m, err := migrate.New(
		"file://internal//migrations",
		psqlAnother)
	if err != nil {

		log.Fatal(err)
	}

	if err := m.Steps(2); err != nil {
		log.Fatal(err)
	} // Run first 2 migrations

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
