package server

import (
	"net/http"
	"testing"
)

func TestServer_RegisterAPI(t *testing.T) {
	s := &Server{}
	s.init()
	type fields struct {
		HTTPserver *http.Server
		mux        *http.ServeMux
		DefineAPI  ServerAPI
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "RegisterAPI",
			fields: fields{
				HTTPserver: s.HTTPserver,
				mux:        http.NewServeMux(),
				DefineAPI:  ServerAPI{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				HTTPserver: tt.fields.HTTPserver,
				mux:        tt.fields.mux,
				DefineAPI:  tt.fields.DefineAPI,
			}
			s.RegisterAPI()
		})
	}
}
