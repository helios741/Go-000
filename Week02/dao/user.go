package dao

import (
	"Week02/model"
	"database/sql"
	"github.com/pkg/errors"
)

func mockSql(user *model.User) error {
	// 处理user，
	return sql.ErrNoRows
}

func GetUserById(id int) (model.User, error) {
	var user model.User
	// 这里简单模拟查询一个sql
	err := mockSql(&user)
	if err != nil {
		return user, errors.Wrap(err, "dao GetUserById fail")
	}
	return user, nil
}