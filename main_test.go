package main

import (
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func Test_populateStructFromYAML(t *testing.T) {
	b := readFile(ConfigFile)
	type args struct {
		data []byte
		obj  interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test populate config struct",
			args: args{
				data: b, obj: &Configuration},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			populateStructFromYAML(tt.args.data, tt.args.obj)
		})
	}
}

func Test_readFile(t *testing.T) {
	expect, _ := ioutil.ReadFile(APIfile)
	type args struct {
		fileName string
	}
	tests := []struct {
		name            string
		args            args
		wantFilebytearr []byte
	}{
		{
			name: "read a file",
			args: args{
				fileName: APIfile,
			},
			wantFilebytearr: expect,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFilebytearr := readFile(tt.args.fileName); !reflect.DeepEqual(gotFilebytearr, tt.wantFilebytearr) {
				t.Errorf("readFile() = %v, want %v", gotFilebytearr, tt.wantFilebytearr)
			}
		})
	}
}

func Test_dontPanicBeHappy(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Panic!!!",
		},
		{
			name: "Don't panic!!!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer dontPanicBeHappy()
			if tt.name == "Panic!!!" {
				panic(errors.New("Panicking!!!"))
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "call main",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testtingMain = true
			main()
		})
	}
}

func Test_parseFlags(t *testing.T) {
	os.Args = []string{"--cfg", "", "--db", "", "--api", "", "--dev"}
	tests := []struct {
		name string
	}{
		{
			"parse flags",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseFlags()
		})
	}
}
