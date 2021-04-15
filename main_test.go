package main

import (
	"reflect"
	"testing"
)

func Test_populateStructFromYAML(t *testing.T) {
	type args struct {
		data []byte
		obj  interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			populateStructFromYAML(tt.args.data, tt.args.obj)
		})
	}
}

func Test_readFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name            string
		args            args
		wantFilebytearr []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFilebytearr := readFile(tt.args.fileName); !reflect.DeepEqual(gotFilebytearr, tt.wantFilebytearr) {
				t.Errorf("readFile() = %v, want %v", gotFilebytearr, tt.wantFilebytearr)
			}
		})
	}
}
