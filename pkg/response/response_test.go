package response

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendMessage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SendMessage(w, http.StatusOK, "success")
	}))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err.Error())
	}
	bodyData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err.Error())
	}
	// Check the response body is what we expect.
	expected := `{"message":"success"}`
	if string(bodyData) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(bodyData), expected)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned unexpected status code: got %v want %v",
			res.StatusCode, http.StatusOK)
	}
}

func TestSendMessageEmpty(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SendMessage(w, http.StatusOK, "")
	}))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err.Error())
	}
	bodyData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err.Error())
	}
	// Check the response body is what we expect.
	expected := `{"message":""}`
	if string(bodyData) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(bodyData), expected)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned unexpected status code: got %v want %v",
			res.StatusCode, http.StatusOK)
	}
}

func TestSendError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SendError(w, http.StatusBadRequest, "failed")
	}))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err.Error())
	}
	bodyData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err.Error())
	}
	// Check the response body is what we expect.
	expected := `{"message":"failed"}`
	if string(bodyData) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(bodyData), expected)
	}
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("handler returned unexpected status code: got %v want %v",
			res.StatusCode, http.StatusBadRequest)
	}
}

func TestSendJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var dataResult interface{}
		SendJSON(w, http.StatusOK, dataResult)
	}))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err.Error())
	}
	bodyData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err.Error())
	}
	// Check the response body is what we expect.
	expected := ``
	if string(bodyData) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(bodyData), expected)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned unexpected status code: got %v want %v",
			res.StatusCode, http.StatusOK)
	}
}

func TestSendJSONWithResult(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dataResult := []string{"one", "two", "three"}
		SendJSON(w, http.StatusOK, dataResult)
	}))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err.Error())
	}
	bodyData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err.Error())
	}
	// Check the response body is what we expect.
	expected := `["one","two","three"]`
	if string(bodyData) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(bodyData), expected)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned unexpected status code: got %v want %v",
			res.StatusCode, http.StatusOK)
	}
}
