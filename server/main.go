package server

import (
	"log"
	"net/http"
)

// Method defines stadard request method types
type Method struct {
	GET    string
	POST   string
	PUT    string
	DELETE string
}

type ServerAPI struct {
	API []API `yaml:"api"`
}
type Get struct {
	SQL string `yaml:"sql"`
}
type Post struct {
	SQL string `yaml:"sql"`
}
type Put struct {
	SQL string `yaml:"sql"`
}
type Delete struct {
	SQL string `yaml:"sql"`
}
type Methods struct {
	Get    Get    `yaml:"get"`
	Post   Post   `yaml:"post"`
	Put    Put    `yaml:"put"`
	Delete Delete `yaml:"delete"`
}
type Handler struct {
	Path    string  `yaml:"path"`
	Public  bool    `yaml:"public"`
	Methods Methods `yaml:"methods"`
}
type API struct {
	Handler Handler `yaml:"handler"`
}

// Server - http server
type Server struct {
	HTTPserver *http.Server
	mux        *http.ServeMux
	DefineAPI  ServerAPI
}

var (
	// Mux http server multiplexer
	Mux = http.NewServeMux()
	// S http server
	S       = Server{}
	methods = &Method{
		GET:    "GET",
		POST:   "POST",
		PUT:    "PUT",
		DELETE: "DELETE",
	}
)

func init() {
	S.init()
	S.DefineAPI = ServerAPI{}
}

// Get logic for GET type request
func handleGet(wr http.ResponseWriter, r *http.Request, sql string) {
	wr.Write([]byte("GET"))
}

// Post logic for POST type request
func handlePost(wr http.ResponseWriter, r *http.Request, sql string) {
	wr.Write([]byte("POST"))
}

// Put logic for PUT type request
func handlePut(wr http.ResponseWriter, r *http.Request, sql string) {
	wr.Write([]byte("PUT"))
}

// Delete logic for DELETE type request
func handleDelete(wr http.ResponseWriter, r *http.Request, sql string) {
	wr.Write([]byte("DELETE"))
}

// RegisterAPI register defined api structure
func (s *Server) RegisterAPI() {
	for _, v := range s.DefineAPI.API {

		if v.Handler.Path == "" {
			continue
		}
		log.Println("Registering API path:", v.Handler.Path)
		Mux.HandleFunc(v.Handler.Path, func(wr http.ResponseWriter, r *http.Request) {
			log.Println("Path:", v.Handler.Path, r.RequestURI)
			switch r.Method {
			case methods.GET:
				handleGet(wr, r, v.Handler.Methods.Get.SQL)
				break
			case methods.POST:
				handlePost(wr, r, v.Handler.Methods.Post.SQL)
				break
			case methods.PUT:
				handlePut(wr, r, v.Handler.Methods.Put.SQL)
				break
			case methods.DELETE:
				handleDelete(wr, r, v.Handler.Methods.Delete.SQL)
				break
			default:
				wr.Write([]byte("..."))
			}
		})

	}
	S.HTTPserver.Handler = Mux
}

func (s *Server) init() {
	log.Println("Server object allocated")
	s.HTTPserver = &http.Server{}
}
