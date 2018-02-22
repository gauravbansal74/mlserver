package logger

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	type args struct {
		debug bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "debug:False",
			args: args{
				debug: false,
			},
		},
		{
			name: "debug:True",
			args: args{
				debug: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init(tt.args.debug)
		})
	}
}

func TestFatal(t *testing.T) {
	type args struct {
		err     error
		message string
		other   []map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "fatal:Error",
			args: args{
				err:     fmt.Errorf("fatal error"),
				message: "fatal error message",
			},
		},
		{
			name: "fatal:Error",
			args: args{
				err:     fmt.Errorf("fatal error"),
				message: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Fatal(tt.args.err, tt.args.message, tt.args.other...)
		})
	}
}

func TestDebug(t *testing.T) {
	otherObject := make([]map[string]interface{}, 1)
	otherObject[0] = make(map[string]interface{})
	otherObject[0]["test"] = "test"
	type args struct {
		message string
		other   []map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Debug messsage",
			args: args{
				message: "debug message",
			},
		},
		{
			name: "Debug all params",
			args: args{
				message: "debug message second",
				other:   nil,
			},
		},
		{
			name: "Debug all params",
			args: args{
				message: "debug message second",
				other:   otherObject,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debug(tt.args.message, tt.args.other...)
		})
	}
}

func TestInfo(t *testing.T) {
	otherObject := make([]map[string]interface{}, 1)
	otherObject[0] = make(map[string]interface{})
	otherObject[0]["test"] = "test"
	type args struct {
		message string
		other   []map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Info messsage",
			args: args{
				message: "info message",
			},
		},
		{
			name: "Info all params",
			args: args{
				message: "info message second",
				other:   nil,
			},
		},
		{
			name: "Info all params",
			args: args{
				message: "info message second",
				other:   otherObject,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.message, tt.args.other...)
		})
	}
}

func TestError(t *testing.T) {
	otherObject := make([]map[string]interface{}, 1)
	otherObject[0] = make(map[string]interface{})
	otherObject[0]["test"] = "test"
	type args struct {
		err     error
		message string
		other   []map[string]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "error messsage",
			args: args{
				err:     fmt.Errorf("error message"),
				message: "error message",
			},
		},
		{
			name: "error all params",
			args: args{
				err:     fmt.Errorf("error message second"),
				message: "error message second",
				other:   nil,
			},
		},
		{
			name: "error all params",
			args: args{
				err:     fmt.Errorf("error message all params"),
				message: "error message second",
				other:   otherObject,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(tt.args.err, tt.args.message, tt.args.other...)
		})
	}
}
