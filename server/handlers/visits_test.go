package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gauravbansal74/mlserver/server/model/user"
)

func Test_visitOverAllNoToken(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(visitOverAll))
	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("content-type", "application/json")
	_, err = (&http.Client{}).Do(req)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_visitOverAll(t *testing.T) {
	token, err := registerUserAndGetToken()
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(visitOverAll))
	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", token)
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		t.Fatal(err)
	}
	_, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Errorf("handler returned unexpected error:  %v", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned unexpected statuscode:  %v", res.StatusCode)
	}
	userLogin := user.Entity{
		Email:    email,
		Password: password,
	}
	err = userLogin.Delete()
	if err != nil {
		t.Errorf("handler returned unexpected error:  %v", err.Error())
	}
}

func Test_visitsByDateNoToken(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(visitsByDate))
	req, err := http.NewRequest("GET", ts.URL+"/1519008141000", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("content-type", "application/json")
	_, err = (&http.Client{}).Do(req)
	if err != nil {
		t.Fatal(err)
	}
}
