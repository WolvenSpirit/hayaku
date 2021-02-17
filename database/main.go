package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	// database/sql drivers
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const (
	// POSTGRESQL driver
	POSTGRESQL = "postgresql"
	// MYSQL driver
	MYSQL = "mysql"
	// SQLITE driver
	SQLITE          = "sqlite"
	defaultDatabase = "hayaku_default_database"
)

// Up migrate up
func Up() {
	log.Println("Migrate up.")
	if DB == nil {
		os.Exit(3)
	}
	tx, err := DB.Begin()
	if err != nil {
		log.Println(err)
	}
	for _, v := range Config.Definition {
		if v.Table.Flavor == Config.RDBMS {
			log.Printf("Creating %s table", v.Table.Name)
			if _, err := tx.Exec(v.Table.Create); err != nil {
				log.Println(err)
				tx.Rollback()
			}
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
		if v.Table.Flavor == Config.RDBMS {
			if _, err := tx.Exec(v.Table.Delete); err != nil {
				log.Println(err)
				tx.Rollback()
			}
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
	Flavor string `json:"driver"`
}

// Config stores database configuration
var (
	Config = Configuration{}
	DB     *sql.DB
	err    error
)

// Connect establishes connection using the chosen driver
func Connect() {
	var dsn string
	log.Printf(
		`Connecting to [%s] Host[%s] Post[%s] Database[%s] User[%s] Sslmode[%s] \n DSN[%s]`,
		Config.RDBMS,
		Config.Host,
		Config.Port,
		Config.Database,
		Config.User,
		Config.Sslmode,
		dsn,
	)
	switch Config.RDBMS {
	case POSTGRESQL:

		dsn = strings.Join([]string{"postgres://", Config.User, ":", Config.Password,
			"@", Config.Host, ":", Config.Port, "/", Config.Database, "?sslmode=", Config.Sslmode}, "")
		DB, err = sql.Open("postgres", dsn)
		log.Printf("%+v", DB.Stats())
		panic(err)
	case MYSQL:
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s", Config.User, Config.Password, Config.Host, Config.Database)
		DB, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Println(err.Error())
			os.Exit(3)
		}
		log.Printf("%+v", DB.Stats())
		panic(err)
	case SQLITE:
		DB, err = sql.Open("sqlite3", Config.Database)
		panic(err)
	default:
		DB, err = sql.Open("sqlite3", defaultDatabase)
		panic(err)
	}
}
