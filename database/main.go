package database

import (
	"database/sql"
	"log"
	"strings"

	// database/sql driver
	_ "github.com/lib/pq"
)

// Up migrate up
func Up() {
	log.Println("Migrate up.")
	tx, err := DB.Begin()
	if err != nil {
		log.Println(err)
	}
	for _, v := range Config.Definition {
		log.Printf("Creating %s table", v.Table.Name)
		if _, err := tx.Exec(v.Table.Create); err != nil {
			log.Println(err)
			tx.Rollback()
		}
	}
	tx.Commit()
}

// Down migrate down
func Down() {
	tx, err := DB.Begin()
	if err != nil {
		log.Println(err)
	}
	for _, v := range Config.Definition {
		if _, err := tx.Exec(v.Table.Delete); err != nil {
			log.Println(err)
			tx.Rollback()
		}
	}
	tx.Commit()
}

type Configuration struct {
	RDBMS      string       `json:"rdbms"`
	Host       string       `json:"host"`
	Port       string       `json:"port"`
	Database   string       `json:"database"`
	User       string       `json:"user"`
	Password   string       `json:"password"`
	Sslmode    string       `json:"sslmode"`
	Definition []Definition `json:"definition"`
}

type Definition struct {
	Table Table `json:"table"`
}

type Table struct {
	Name   string `json:"name"`
	Create string `json:"create"`
	Delete string `json:"delete"`
}

// Config stores database configuration
var (
	Config = Configuration{}
	DB     *sql.DB
	err    error
)

// ConnectPostgresql establishes connection to PostgreSQL rdbms
func ConnectPostgresql() {
	if Config.RDBMS != "postgresql" {
		return
	}
	log.Printf(
		`Connecting to [postgresql] Host[%s] Post[%s] Database[%s] User[%s] Sslmode[%s]`,
		Config.Host,
		Config.Port,
		Config.Database,
		Config.User,
		Config.Sslmode,
	)
	psqlConnectionURL := strings.Join([]string{"postgres://", Config.User, ":", Config.Password,
		"@", Config.Host, ":", Config.Port, "/", Config.Database, "?sslmode=", Config.Sslmode}, "")
	DB, err = sql.Open("postgres", psqlConnectionURL)
	log.Printf("%+v", DB.Stats())
	panic(err)
}
