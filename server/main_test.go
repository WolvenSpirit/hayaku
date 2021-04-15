package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"reflect"
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

func Test_getResult(t *testing.T) {
	type args struct {
		result *sql.Rows
	}
	tests := []struct {
		name string
		args args
		want []map[string]interface{}
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getResult(tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonMarshal(t *testing.T) {
	var st []map[string]interface{}
	expect, _ := json.Marshal(&st)
	type args struct {
		result []map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "",
			args: args{result: st},
			want: expect,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jsonMarshal(tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jsonMarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
