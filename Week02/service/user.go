package service

import (
	"Week02/dao"
	"Week02/model"
	"github.com/pkg/errors"
)

func GetUserById(id int) (model.User, error)  {
	user, err := dao.GetUserById(id)
	if err != nil {
		return user, errors.WithMessage(err, "service GetUserById fail")
	}
	// 处理逻辑
	return user, nil
}
