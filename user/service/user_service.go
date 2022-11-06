package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/danangkonang/user/config"
	"github.com/danangkonang/user/model"
)

type ServiceUser interface {
	Login(email string) (*model.UserLogin, error)
	Register(user *model.UserRegister) error
}

func NewServiceUser(Con *config.DB) ServiceUser {
	return &Database{
		Postgresql: Con.Postgresql,
	}
}

func (r *Database) Login(email string) (*model.UserLogin, error) {
	user := new(model.UserLogin)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := r.Postgresql.QueryRowContext(ctx, "SELECT user_id, user_name, password FROM users WHERE user_name = $1", email).Scan(&user.UserId, &user.UserName, &user.Password)

	switch {
	case err == sql.ErrNoRows:
		return nil, errors.New("user name not found")
	case err != nil:
		return nil, err
	}
	return user, nil
}

func (r *Database) Register(user *model.UserRegister) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var us string
	err := r.Postgresql.QueryRowContext(ctx, "SELECT user_name FROM users WHERE user_name = $1", user.UserName).Scan(&us)
	// if err != nil {
	// }
	if err == sql.ErrNoRows {
		fmt.Println("oke")
	}
	// fmt.Println(us)
	return err
	// switch {
	// case err == sql.ErrNoRows:
	// 	query := "INSERT INTO users (user_id, user_name, password) VALUES ($1, $2, $3)"
	// 	_, err := r.Postgresql.ExecContext(ctx, query, user.UserId, user.UserName, user.Password)
	// 	if err != nil {
	// 		return err
	// 	}
	// case err != nil:
	// 	return err
	// }
	// return nil
}
