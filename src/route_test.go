package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)


type MockUrlRepo struct {
	logger *log.Logger
	conn   Conn
}

func (repo MockUrlRepo) getLatestSequence() int                { return 99 }
func (repo MockUrlRepo) insert(url UrlDetail) bool             { return true }
func (repo MockUrlRepo) get(shortUrlId int) (UrlDetail, error) { return UrlDetail{OriginalUrl:"mocked_original_url"}, nil }
var (
	app = &App{
		logger: log.New(os.Stdout, "", log.Ldate|log.Llongfile),
	}
)
var router *httprouter.Router

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
func setup() {
	fmt.Println("Setup")
	app.loadEnv("../resources/.env-dev")
	router = app.route(MockUrlRepo{nil, Conn{}})
}
func shutdown() {
	fmt.Println("Shutdown")
}

func TestCreateShortUrl(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/create", strings.NewReader("{\n    \"original_url\":\"mock url\"\n}"))
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("create api failed")
	}
	response := writer.Result()
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if !strings.Contains(string(data), "\"status\":\"success\"") {
		t.Errorf("Expected status:success %v", string(data))
	}
	if !strings.Contains(string(data), "/99") {
		t.Errorf("Expected http://localhost:5432/99 but got %v", string(data))
	}
}

func TestGetShortUrl(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/get", strings.NewReader("{\n    \"short_url\": \"http://localhost:5432/70\"\n}"))
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("create api failed")
	}
	response := writer.Result()
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if !strings.Contains(string(data), "\"status\":\"success\"") {
		t.Errorf("Expected status:success %v", string(data))
	}
	if !strings.Contains(string(data), "mocked_original_url") {
		t.Errorf("Expected original url mocked_original_url but got %v", string(data))
	}

}
