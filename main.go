package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	"github.com/WolvenSpirit/hayaku/cache"
	"github.com/WolvenSpirit/hayaku/database"
	"github.com/WolvenSpirit/hayaku/server"
	"github.com/go-yaml/yaml"
)

// Config struct
type Config struct {
	Host  string      `yaml:"host"`
	Port  string      `yaml:"port"`
	Redis cache.Redis `yaml:"redis"`
}

var (
	// ConfigFile yaml
	ConfigFile string = "config.yaml"
	// DatabaseFile yaml
	DatabaseFile string = "database.yaml"
	// Configuration structure
	Configuration = Config{}
	// APIfile API definition file
	APIfile string = "api.yaml"
	// Development run
	Development  = false
	testtingMain = false
)

func dontPanicBeHappy() {
	if err, ok := recover().(error); ok {
		log.Println(err)
	}
}

func parseFlags() {
	flags := struct {
		cfg string
		db  string
		api string
		dev string
	}{
		cfg: "--cfg",
		db:  "--db",
		api: "--api",
		dev: "--dev",
	}
	for k, arg := range os.Args {
		switch arg {
		case flags.api:
			APIfile = os.Args[k+1]
		case flags.cfg:
			ConfigFile = os.Args[k+1]
		case flags.db:
			DatabaseFile = os.Args[k+1]
		case flags.dev:
			Development = true
		}
	}
}

func readFile(fileName string) (filebytearr []byte) {
	filebytearr, err := ioutil.ReadFile(fileName)
	go func() {
		defer dontPanicBeHappy()
		panic(err)
	}()
	return
}

func populateStructFromYAML(data []byte, obj interface{}) {
	yaml.Unmarshal(data, obj)
}

func init() {
	parseFlags()
	defer dontPanicBeHappy()
	log.Println("Parsed configuration file arguments:", os.Args[1:])
	populateStructFromYAML(readFile(ConfigFile), &Configuration)
	populateStructFromYAML(readFile(DatabaseFile), &database.Config)
	populateStructFromYAML(readFile(APIfile), &server.S.DefineAPI)
	server.S.HTTPserver.Addr = fmt.Sprintf("%s:%s", Configuration.Host, Configuration.Port)
	server.S.RegisterAPI()
	if !testtingMain {
		database.Connect()
	}
	cache.ConnectRedis(fmt.Sprintf(Configuration.Redis.Host), Configuration.Redis.Password, Configuration.Redis.DBI)
}

func main() {
	if !testtingMain {
		database.Up()
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	log.Printf("Listening on [%s] Development run [%t]", server.S.HTTPserver.Addr, Development)
	go func() {
		log.Println(server.S.HTTPserver.ListenAndServe())
	}()
	if !testtingMain {
		<-sig
	}
	log.Println("Shutting down server.")
	database.DB.Close()
	server.S.HTTPserver.Shutdown(context.Background())
}
