package database

import (
	"database/sql"
	"io/ioutil"
	"testing"

	"github.com/go-yaml/yaml"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func TestDatabaseInit(t *testing.T) {
	DB, err = sql.Open("sqlite3", "test")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestConfigInit(t *testing.T) {
	Config = Configuration{}
	b, err := ioutil.ReadFile("test_config.yaml")
	if err != nil {
		t.Error(err.Error())
	}
	yaml.Unmarshal(b, &Config)
}

func TestConnectPostgresql(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "ConnectPostgresql"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err, ok := recover().(error); ok && err != nil {
					t.Error(err.Error())
				}
			}()
			Connect()
		})
	}
}

func TestUp(t *testing.T) {
	TestDatabaseInit(t)
	TestConfigInit(t)
	tests := []struct {
		name string
	}{
		{name: "Migrate up"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Up()
		})
	}
}

func TestDown(t *testing.T) {
	TestDatabaseInit(t)
	TestConfigInit(t)
	tests := []struct {
		name string
	}{
		{name: "Migrate down"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Down()
		})
	}
}
