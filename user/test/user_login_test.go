package test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/danangkonang/user/config"
	"github.com/danangkonang/user/controller"
	"github.com/danangkonang/user/service"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' database connection", err)
	}
	return sqlDB, mock
}

func Server(db *config.DB) *mux.Router {
	router := mux.NewRouter()
	rest := controller.NewUserController(
		service.NewServiceUser((*config.DB)(db)),
	)
	v1 := router.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/user/login", rest.Login).Methods("POST")
	v1.HandleFunc("/user/register", rest.Register).Methods("POST")
	return router
}

func Test_Login_Ok(t *testing.T) {
	sqlDB, mock := NewMock()
	con := &config.DB{
		Postgresql: sqlDB,
	}
	defer con.Postgresql.Close()

	rows := sqlmock.NewRows([]string{"user_id", "user_name", "password"}).AddRow("1", "user@email.com", "$2a$10$MUWGo1u6p6Xei6RIPxMFQ.5leV0EW84G.EX2IibNZQftw4RB4mPAi")

	query := "SELECT user_id, user_name, password FROM users WHERE user_name = $1"

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("user@email.com").WillReturnRows(rows)

	var jsonStr = []byte(`{"user_name":"user@email.com","password":"testing123"}`)

	request, _ := http.NewRequest("POST", "/v1/user/login", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	Server(&config.DB{Postgresql: con.Postgresql}).ServeHTTP(response, request)

	responseBody := make(map[string]interface{})
	body, _ := ioutil.ReadAll(response.Body)
	err := json.Unmarshal(body, &responseBody)
	if err != nil {
		t.Errorf("can not conver to json: %v", err)
	}
	assert.Equal(t, 200, int(responseBody["status"].(float64)), "ok")
}

func Test_Login_Invalid_Pass(t *testing.T) {
	sqlDB, mock := NewMock()
	con := &config.DB{
		Postgresql: sqlDB,
	}
	defer con.Postgresql.Close()

	rows := sqlmock.NewRows([]string{"user_id", "user_name", "password"}).AddRow("1", "user@email.com", "$2a$10$MUWGo1u6p6Xei6RIPxMFQ.5leV0EW84G.EX2IibNZQftw4RB4mPAi")

	query := "SELECT user_id, user_name, password FROM users WHERE user_name = $1"

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("user@email.com").WillReturnRows(rows)

	var jsonStr = []byte(`{"user_name":"user@email.com","password":""}`)

	request, _ := http.NewRequest("POST", "/v1/user/login", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	Server(&config.DB{Postgresql: con.Postgresql}).ServeHTTP(response, request)

	responseBody := make(map[string]interface{})
	body, _ := ioutil.ReadAll(response.Body)
	err := json.Unmarshal(body, &responseBody)
	if err != nil {
		t.Errorf("can not conver to json: %v", err)
	}
	assert.Equal(t, 400, int(responseBody["status"].(float64)), "invalid password")
}

func Test_Login_Invalid_Email(t *testing.T) {
	sqlDB, mock := NewMock()
	con := &config.DB{
		Postgresql: sqlDB,
	}
	defer con.Postgresql.Close()

	rows := sqlmock.NewRows([]string{"user_id", "user_name", "password"}).AddRow("1", "user@email.com", "$2a$10$MUWGo1u6p6Xei6RIPxMFQ.5leV0EW84G.EX2IibNZQftw4RB4mPAi")

	query := "SELECT user_id, user_name, password FROM users WHERE user_name = $1"

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("invalid@email.com").WillReturnRows(rows)

	var jsonStr = []byte(`{"user_name":"invalid@email.com","password":"invalid"}`)

	request, _ := http.NewRequest("POST", "/v1/user/login", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	Server(&config.DB{Postgresql: con.Postgresql}).ServeHTTP(response, request)

	responseBody := make(map[string]interface{})
	body, _ := ioutil.ReadAll(response.Body)
	err := json.Unmarshal(body, &responseBody)
	if err != nil {
		t.Errorf("can not conver to json: %v", err)
	}
	assert.Equal(t, 400, int(responseBody["status"].(float64)), "invalid email")
}
