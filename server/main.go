package server

import (
	"log"
	"net/http"

	"github.com/WolvenSpirit/hayaku/database"
)

// APIMethod type
type APIMethod func(wr http.ResponseWriter, r *http.Request, mData method)

// Method defines stadard request method types
type Method struct {
	GET    string
	POST   string
	PUT    string
	DELETE string
	PATCH  string
}

type contentType struct {
	URLEncoded string
	Multipart  string
	JSON       string
}

type method interface {
	isPublic() bool
}

func (m Get) isPublic() bool {
	return m.Public
}
func (m Post) isPublic() bool {
	return m.Public
}
func (m Patch) isPublic() bool {
	return m.Public
}
func (m Put) isPublic() bool {
	return m.Public
}
func (m Delete) isPublic() bool {
	return m.Public
}

type ServerAPI struct {
	API []API `yaml:"api"`
}
type Get struct {
	Public   bool   `yaml:"public"`
	SQL      string `yaml:"sql"`
	Resource []interface{}
	Headers  map[string]string
}
type Post struct {
	Public   bool   `yaml:"public"`
	SQL      string `yaml:"sql"`
	Resource []interface{}
	Headers  map[string]string
}
type Put struct {
	Public   bool   `yaml:"public"`
	SQL      string `yaml:"sql"`
	Resource []interface{}
	Headers  map[string]string
}
type Delete struct {
	Public   bool   `yaml:"public"`
	SQL      string `yaml:"sql"`
	Resource []interface{}
	Headers  map[string]string
}
type Patch struct {
	Public   bool   `yaml:"public"`
	SQL      string `yaml:"sql"`
	Resource []interface{}
	Headers  map[string]string
}
type Methods struct {
	Get    Get    `yaml:"get"`
	Post   Post   `yaml:"post"`
	Put    Put    `yaml:"put"`
	Delete Delete `yaml:"delete"`
	Patch  Patch  `yaml:"patch"`
}
type Handler struct {
	Path    string  `yaml:"path"`
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
		PATCH:  "PATCH",
	}
	// ContentType accepted
	ContentType = contentType{
		URLEncoded: "application/x-www-form-urlencoded",
		Multipart:  "multipart/form-data",
		JSON:       "application/json",
	}
)

func init() {
	S.init()
	S.DefineAPI = ServerAPI{}
}

// Get logic for GET type request
func handleGet(wr http.ResponseWriter, r *http.Request, data method) {
	var err error
	d := data.(Get)
	result, err := database.DB.Query(d.SQL, d.Resource...)
	for result.Next() {
		// ...
	}
	wr.Write([]byte("GET"))
	panic(err)
}

// Post logic for POST type request
func handlePost(wr http.ResponseWriter, r *http.Request, data method) {
	wr.Write([]byte("POST"))
}

// Put logic for PUT type request
func handlePut(wr http.ResponseWriter, r *http.Request, data method) {
	wr.Write([]byte("PUT"))
}

// Delete logic for DELETE type request
func handleDelete(wr http.ResponseWriter, r *http.Request, data method) {
	wr.Write([]byte("DELETE"))
}

// Patch logic for DELETE type request
func handlePatch(wr http.ResponseWriter, r *http.Request, data method) {
	wr.Write([]byte("DELETE"))
}

// Auth middleware
func (s *Server) Auth(fn APIMethod) APIMethod {
	return func(wr http.ResponseWriter, r *http.Request, mData method) {
		if !mData.isPublic() {
			wr.WriteHeader(http.StatusForbidden)
			wr.Write([]byte("This resource is not public"))
			return
		}
		/*
			t := reflect.TypeOf(mData)
			switch t.Name() {
			case "Get":
			case "Post":
			case "Put":
			case "Patch":
			case "Delete":
			default:
				return
			}
		*/
		fn(wr, r, mData)
	}
}

// Logger middleware
func (s *Server) Logger(fn http.HandlerFunc) http.HandlerFunc {
	return func(wr http.ResponseWriter, r *http.Request) {
		log.Println("Path:", r.RequestURI)
		defer func() {
			if err, ok := recover().(error); ok && err != nil {
				log.Println(err.Error())
			}
		}()
		fn(wr, r)
	}
}

// RegisterAPI register defined api structure
func (s *Server) RegisterAPI() {
	for _, v := range s.DefineAPI.API {

		if v.Handler.Path == "" {
			continue
		}
		log.Println("Registering API path:", v.Handler.Path)
		Mux.HandleFunc(v.Handler.Path, s.Logger(func(wr http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case methods.GET:
				s.Auth(handleGet)(wr, r, v.Handler.Methods.Get)
				break
			case methods.POST:
				s.Auth(handlePost)(wr, r, v.Handler.Methods.Post)
				break
			case methods.PUT:
				s.Auth(handlePut)(wr, r, v.Handler.Methods.Put)
				break
			case methods.DELETE:
				s.Auth(handleDelete)(wr, r, v.Handler.Methods.Delete)
				break
			case methods.PATCH:
				s.Auth(handlePatch)(wr, r, v.Handler.Methods.Patch)
			default:
				wr.Write([]byte("..."))
			}
		}))

	}
	S.HTTPserver.Handler = Mux
}

func (s *Server) init() {
	log.Println("Server object allocated")
	s.HTTPserver = &http.Server{}
}
