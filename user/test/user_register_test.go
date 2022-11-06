package test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/danangkonang/user/model"
	"github.com/danangkonang/user/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var u = &model.UserRegister{
	UserId:   uuid.New().String(),
	UserName: "momo@mail.com",
	Password: "password",
}

func Test_Register(t *testing.T) {
	db, mock := NewMock()
	repo := &service.Database{Postgresql: db}

	sqlmock.NewRows([]string{"user_id", "user_name", "password"})

	query := "INSERT INTO users (user_id, user_name, password) VALUES ($1, $2, $3)"

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(u.UserId, u.UserName, u.Password).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Register(u)
	assert.NoError(t, err)
}

func Test_register_error(t *testing.T) {
	db, mock := NewMock()
	repo := &service.Database{Postgresql: db}

	sqlmock.NewRows([]string{"user_id", "user_name", "password"})

	query := "INSERT INTO users (user_id, user_name, password) VALUES ($1, $2, $3)"

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(u.UserId, u.UserName, u.Password).WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.Register(u)
	assert.Error(t, err)
}

func Test_register(t *testing.T) {
	db, mock := NewMock()
	repo := &service.Database{Postgresql: db}

	sqlmock.NewRows([]string{"user_id", "user_name", "password"})

	query := "INSERT INTO users (user_id, user_name, password) VALUES ($1, $2, $3)"

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(u.UserId, u.UserName, u.Password).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Register(u)
	assert.NoError(t, err)
}
