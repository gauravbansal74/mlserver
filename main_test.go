package main

import (
	"reflect"
	"testing"

	"github.com/gauravbansal74/mlserver/cmd"
)

func TestMain(t *testing.T) {
	if !isFunc(cmd.Execute) {
		t.Error("cmd.Execute should be function")
	}
}

func isFunc(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Func
}
