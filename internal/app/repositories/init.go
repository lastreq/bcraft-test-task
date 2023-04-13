package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // comment for linter
)

type PgDatabase struct {
	DB *sql.DB
}

func New() *PgDatabase {
	dbUser, dbPassword, dbName, dbHost, dbPort := os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Println(err.Error())
	}

	fillTable(db)

	return &PgDatabase{
		DB: db,
	}
}

func fillTable(db *sql.DB) {
	sqlStatement := `CREATE TABLE IF NOT EXISTS recipes(
    recipe_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    ingredients TEXT,
    cooking_steps TEXT
    );`

	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Println(err.Error())
	}

	for i := 1; i <= 100; i++ {
		sqlStatement = `
INSERT INTO "recipes" ("name","description", "ingredients", "cooking_steps")
VALUES ($1, $2, $3, $4)`

		_, err = db.Exec(sqlStatement, fmt.Sprintf("name%v", i), fmt.Sprintf("description%v", i), fmt.Sprintf("ingredients%v", i), fmt.Sprintf("cooking_steps%v", i))
		if err != nil {
			log.Println(err.Error())
		}
	}
}
