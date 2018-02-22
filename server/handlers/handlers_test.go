package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/gauravbansal74/mlserver/pkg/database"

	"github.com/gauravbansal74/mlserver/server/model/user"
)

var (
	email                   = time.Now().String() + "@example.com"
	password                = "123456789"
	ID                      = "12321321312312"
	expectedSuccess         = `{"message":"success"}`
	expectedSuccessRegister = `{"message":"Successfully registered."}`
	expectedInteralError    = `{"message":"internal server error"}`
	expectedWrongPassword   = `{"message":"email or password invalid"}`
)

func registerUserAndGetToken() (string, error) {
	ts := httptest.NewServer(http.HandlerFunc(registerHandler))
	userLogin := user.Entity{
		Email:    email,
		Password: password,
	}
	userBytes, err := json.Marshal(userLogin)
	if err != nil {
		return "", err
	}
	res, err := http.Post(ts.URL, "application/json", bytes.NewBuffer(userBytes))
	if err != nil {
		return "", err
	}
	bodyData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Status code wrong")
	}

	if string(bodyData) != expectedSuccessRegister {
		return "", fmt.Errorf("response message is different")
	}

	tsl := httptest.NewServer(http.HandlerFunc(loginHandler))

	resp, err := http.Post(tsl.URL, "application/json", bytes.NewBuffer(userBytes))
	if err != nil {
		return "", err
	}
	bodyDat, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Status code wrong")
	}

	if string(bodyDat) != expectedSuccess {
		return "", fmt.Errorf("response message is different")
	}
	token := resp.Header.Get("Authorization")
	return token, nil
}

func setUpMongo() {
	dbInfo := database.LoadConfig("127.0.0.1", "mlserver")
	dbInfo.Init()
}
