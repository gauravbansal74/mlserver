package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gauravbansal74/mlserver/server/model/user"
)

func TestAllDataSourcesHandler(t *testing.T) {
	setUpMongo()
	ts := httptest.NewServer(http.HandlerFunc(AllDataSourcesHandler))
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	t.Log(string(body))
	res.Body.Close()
	if err != nil {
		t.Errorf("handler returned unexpected error:  %v", err.Error())
	}
	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("handler returned unexpected statuscode:  %v", res.StatusCode)
	}
}

func TestAllDataSourcesHandlerWithToken(t *testing.T) {
	token, err := registerUserAndGetToken()
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(AllDataSourcesHandler))
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
