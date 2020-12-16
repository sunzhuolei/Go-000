package dao

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"task/Go-000/Week04/internal/model"
)

var ErrRecordNotFound = errors.New("Not Found")

type Dao interface {
	GetUser(context.Context,int) (*model.UserModel,error)
}
var Provider = wire.NewSet(NewDB, NewDao)
type dao struct {
	db *sql.DB
}
func NewDao(db *sql.DB) Dao {
	return &dao{db: db}
}
/**
获取用户信息
 */
func (d *dao) GetUser(ctx context.Context, id int) (*model.UserModel, error) {
	userModel := &model.UserModel{}
	row := d.db.QueryRowContext(ctx, "select uid,name,age from user where uid=?", id)
	err := row.Scan(&userModel.Uid, &userModel.Name, &userModel.Age)
	if err == sql.ErrNoRows {
		return nil, errors.Wrap(ErrRecordNotFound, "No corresponding article")
	}
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get article")
	}
	return userModel, nil
}

func NewDB() (db *sql.DB, cleanup func(), err error) {
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/task")
	fmt.Println(db)
	cleanup = func() {
		if err == nil {
			db.Close()
		}
	}
	return
}
