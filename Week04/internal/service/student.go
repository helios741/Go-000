package service

import (
	v1 "Week04/api/students/v1"
	"Week04/configs/errcode"
	"Week04/internal/biz"
	"context"
	"errors"
	"fmt"
)

type StudentService struct {
	v1.UnimplementedStudentServer
	suc *biz.StudentUsecase
}

func NewStudentService(suc *biz.StudentUsecase) v1.StudentServer  {
	return &StudentService{suc: suc}
}

func (ss *StudentService) GetById(ctx context.Context, r *v1.StudentRequest) (*v1.StudentReply, error) {
	stu, err := ss.suc.Get(r.Id)
	if err != nil {
		if errors.Is(err, errcode.SqlNotFound) {
			fmt.Println("this is sqlNotFound", err)
			return nil, err
		}
		fmt.Println("error: ", err)
		return nil, err
	}
	return &v1.StudentReply{
		Id: stu.Id,
		Age: stu.Age,
		Sex: stu.Sex,
		Interest: stu.Like,
	}, nil
}