package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/WolvenSpirit/hayaku/database"
	"github.com/WolvenSpirit/hayaku/server"
	"github.com/go-yaml/yaml"
)

// Config struct
type Config struct {
	Host string
	Port string
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
	}{
		cfg: "--cfg",
		db:  "--db",
		api: "--api",
	}
	for k, arg := range os.Args {
		switch arg {
		case flags.api:
			APIfile = os.Args[k+1]
		case flags.cfg:
			ConfigFile = os.Args[k+1]
		case flags.db:
			DatabaseFile = os.Args[k+1]
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
	log.Println("Parsed configuration files:", os.Args[1:])
	populateStructFromYAML(readFile(ConfigFile), &Configuration)
	populateStructFromYAML(readFile(DatabaseFile), &database.Config)
	populateStructFromYAML(readFile(APIfile), &server.S.DefineAPI)
}

func main() {
	log.Printf("Marshalled: \n%+v \n%+v \n%+v", Configuration, database.Config, server.S.DefineAPI)
	server.S.HTTPserver.Addr = fmt.Sprintf("%s:%s", Configuration.Host, Configuration.Port)
	server.S.RegisterAPI()
	log.Println("Listening on ", server.S.HTTPserver.Addr)
	err := server.S.HTTPserver.ListenAndServe()
	panic(err)
}
