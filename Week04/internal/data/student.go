package data

import (
	"Week04/configs/errcode"
	"Week04/internal/biz"
	"fmt"
	"github.com/pkg/errors"
)

var _ biz.StudentRepo = (*studentRepo)(nil)


type studentRepo struct {

}

func NewStudentRepo() biz.StudentRepo {
	return new(studentRepo)
}

func mockStudent(id int32) (*biz.Student, error) {
	if id == 1001 {
		return nil, errors.Wrap(errcode.SqlNotFound, fmt.Sprintf("sql err for mockStudent"))
	} else if id == 1002 {
		return nil, errors.New("this is sql error")
	}
	stu := biz.Student{
		Id: id,
		Name: "Helios",
		Like: "dsdsd",
		Age: 23,
	}
	return &stu, nil
}

func (sr *studentRepo) GetStudent(id int32) (*biz.Student, error)  {
	// mock student date for db
	return mockStudent(id)
}