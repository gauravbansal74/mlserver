package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gauravbansal74/mlserver/server/model/user"
)

func TestRegister(t *testing.T) {
	setUpMongo()
	ts := httptest.NewServer(http.HandlerFunc(registerHandler))
	userRegistration := user.Entity{
		Email:    email,
		Password: password,
	}
	userBytes, err := json.Marshal(userRegistration)
	if err != nil {
		t.Fatal(err)
	}
	res, err := http.Post(ts.URL, "application/json", bytes.NewBuffer(userBytes))
	if err != nil {
		t.Fatal(err)
	}
	bodyData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Errorf("handler returned unexpected error:  %v", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned unexpected statuscode:  %v", res.StatusCode)
	}

	if string(bodyData) != expectedSuccessRegister {
		t.Errorf("handler returned unexpected response:  %v", string(bodyData))
	}
	err = userRegistration.Delete()
	if err != nil {
		t.Errorf("handler returned unexpected error:  %v", err.Error())
	}
}

func TestRegisterEmpty(t *testing.T) {
	setUpMongo()
	ts := httptest.NewServer(http.HandlerFunc(registerHandler))
	res, err := http.Post(ts.URL, "application/json", nil)
	if err != nil {
		t.Fatal(err)
	}
	bodyData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Errorf("handler returned unexpected error:  %v", err.Error())
	}
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("handler returned unexpected statuscode:  %v", res.StatusCode)
	}

	if string(bodyData) != expectedInteralError {
		t.Errorf("handler returned unexpected response:  %v", string(bodyData))
	}
}

func TestLogin(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(registerHandler))
	userLogin := user.Entity{
		Email:    email,
		Password: password,
	}
	userBytes, err := json.Marshal(userLogin)
	if err != nil {
		t.Fatal(err)
	}
	res, err := http.Post(ts.URL, "application/json", bytes.NewBuffer(userBytes))
	if err != nil {
		t.Fatal(err)
	}
	bodyData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Errorf("handler returned unexpected error:  %v", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned unexpected statuscode:  %v", res.StatusCode)
	}

	if string(bodyData) != expectedSuccessRegister {
		t.Errorf("handler returned unexpected response:  %v", string(bodyData))
	}

	tsl := httptest.NewServer(http.HandlerFunc(loginHandler))

	resp, err := http.Post(tsl.URL, "application/json", bytes.NewBuffer(userBytes))
	if err != nil {
		t.Fatal(err)
	}
	bodyDat, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Errorf("handler returned unexpected error:  %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("handler returned unexpected statuscode:  %v", resp.StatusCode)
	}

	if string(bodyDat) != expectedSuccess {
		t.Errorf("handler returned unexpected response:  %v", string(bodyDat))
	}
	err = userLogin.Delete()
	if err != nil {
		t.Errorf("handler returned unexpected error:  %v", err.Error())
	}
}

func TestLoginWrong(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(registerHandler))
	userLogin := user.Entity{
		Email:    email,
		Password: password,
	}
	userBytes, err := json.Marshal(userLogin)
	if err != nil {
		t.Fatal(err)
	}
	res, err := http.Post(ts.URL, "application/json", bytes.NewBuffer(userBytes))
	if err != nil {
		t.Fatal(err)
	}
	bodyData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Errorf("handler returned unexpected error:  %v", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned unexpected statuscode:  %v", res.StatusCode)
	}

	if string(bodyData) != expectedSuccessRegister {
		t.Errorf("handler returned unexpected response:  %v", string(bodyData))
	}

	tsl := httptest.NewServer(http.HandlerFunc(loginHandler))
	userLogin.Password = "1232112321"
	userBytesLogin, err := json.Marshal(userLogin)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Post(tsl.URL, "application/json", bytes.NewBuffer(userBytesLogin))
	if err != nil {
		t.Fatal(err)
	}
	bodyDat, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Errorf("handler returned unexpected error:  %v", err.Error())
	}
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("handler returned unexpected statuscode:  %v", resp.StatusCode)
	}

	if string(bodyDat) != expectedWrongPassword {
		t.Errorf("handler returned unexpected response:  %v", string(bodyDat))
	}
	err = userLogin.Delete()
	if err != nil {
		t.Errorf("handler returned unexpected error:  %v", err.Error())
	}
}

func TestLoginEmpty(t *testing.T) {
	setUpMongo()
	ts := httptest.NewServer(http.HandlerFunc(registerHandler))
	res, err := http.Post(ts.URL, "application/json", nil)
	if err != nil {
		t.Fatal(err)
	}
	bodyData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Errorf("handler returned unexpected error:  %v", err.Error())
	}
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("handler returned unexpected statuscode:  %v", res.StatusCode)
	}

	if string(bodyData) != expectedInteralError {
		t.Errorf("handler returned unexpected response:  %v", string(bodyData))
	}
}

func Test_profileHandler(t *testing.T) {
	token, err := registerUserAndGetToken()
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(profileHandler))
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
