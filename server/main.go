package server

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"github.com/WolvenSpirit/hayaku/auth"
	"github.com/WolvenSpirit/hayaku/cache"
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
	XML        string
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
	Public      bool   `yaml:"public"`
	SQL         string `yaml:"sql"`
	Resource    []interface{}
	Headers     map[string]string
	ContentType string `yaml:"contentType"`
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
		XML:        "application/xml",
	}
)

func init() {
	S.init()
	S.DefineAPI = ServerAPI{}
}

func assertType(v interface{}, t *sql.ColumnType) interface{} {
	switch t.ScanType().Name() {
	case "int":
		return v.(int)
	case "float32":
		return v.(float32)
	case "float64":
		return v.(float64)
	case "string":
		return v.(string)
	default:
		fmt.Println(t.ScanType().Name())
		return nil
	}
}

func getResult(result *sql.Rows) []map[string]interface{} {
	columnNumber, _ := result.Columns()
	columnTypes, _ := result.ColumnTypes()
	resultSet := make([]map[string]interface{}, 0)
	for result.Next() {
		columns := make([]interface{}, len(columnNumber))
		columnPointers := make([]interface{}, len(columnNumber))
		queryResult := make(map[string]interface{})
		for k := range columnNumber {
			columnPointers[k] = &columns[k]
		}
		err := result.Scan(columnPointers...)
		if err != nil {
			log.Println(err.Error())
		}
		for i, name := range columnNumber {
			v := columnPointers[i].(*interface{})
			value := *v
			var ok bool
			// t := columnTypes[i]
			switch columnTypes[i].ScanType().Name() {
			case "int32": // even if it reports 32bit it fails at casting to that
				//				fmt.Println("int32")
				if queryResult[name], ok = value.(int64); !ok {
					//	log.Println("Error asserting!")
				}
			case "int64":
				//				fmt.Println("int64")
				if queryResult[name], ok = columns[i].(int64); !ok {
					//	log.Println("Error asserting!")
				}
			case "float32":
				fmt.Println("float32")
				if queryResult[name], ok = columns[i].(float32); !ok {
					//	log.Println("Error asserting!")
				}
			case "float64":
				//				fmt.Println("float64")
				if queryResult[name], ok = columns[i].(float64); !ok {
					//	log.Println("Error asserting!")
				}
			case "uint32":
				//				fmt.Println("uint32")
				if queryResult[name], ok = columns[i].(uint32); !ok {
					//	log.Println("Error asserting!")
				}
			case "uint64":
				//				fmt.Println("uint64")
				if queryResult[name], ok = columns[i].(uint64); !ok {
					//	log.Println("Error asserting!")
				}
			case "string":
				//				fmt.Println("string")
				if queryResult[name] = fmt.Sprintf("%s", *v); !ok {
					//	log.Println("Error asserting!")
				}
			case "time.Time":

			case "interface {}":
				//				fmt.Println("float64 forced from interface{}")
				if queryResult[name], ok = columns[i].(float64); !ok {
					//	log.Println("Error asserting!")
				}
			default:
				//				fmt.Println(columnTypes[i].DatabaseTypeName(), columnTypes[i].ScanType())
				queryResult[name] = fmt.Sprintf("%s", *v)
			}
			// queryResult[name] = *v
		}
		resultSet = append(resultSet, queryResult)
	}
	return resultSet
}

func jsonMarshal(result []map[string]interface{}) []byte {
	b, _ := json.Marshal(&result)
	return b
}
func xmlMarshal(result []map[string]interface{}) []byte {
	o := struct {
		Data []map[string]interface{} `xml:"data"`
	}{result}
	b, err := xml.Marshal(&o)
	go func() { panic(err) }()
	return b
}

// Get logic for GET type request
func handleGet(wr http.ResponseWriter, r *http.Request, data method) {
	var err error
	d := data.(Get)
	result, err := database.DB.Query(d.SQL, d.Resource...)
	switch d.ContentType {
	case ContentType.JSON:
		wr.Header().Add("Content-Type", d.ContentType)
		wr.Write(jsonMarshal(getResult(result)))
	/*
		case ContentType.XML:
			wr.Header().Add("Content-Type", d.ContentType)
			wr.Write(xmlMarshal(getResult(result)))
	*/
	default:
		wr.Header().Add("Content-Type", ContentType.JSON)
		wr.Write(jsonMarshal(getResult(result)))
	}
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
		token := r.Header.Get("X-Auth-Token")
		scmd := cache.Client.Get(token)
		storedToken, err := scmd.Result()
		go func() { panic(err) }()
		if !auth.CompareHashAndPassword([]byte(storedToken), []byte(token)) {
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
