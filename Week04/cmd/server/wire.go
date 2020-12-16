// +build wireinject

package main

import (
	pb "Week04/api/students/v1"
	"Week04/internal/biz"
	"Week04/internal/data"
	"Week04/internal/service"
	"github.com/google/wire"
)
var SuperSet = wire.NewSet(data.NewStudentRepo, biz.NewStudentUsecase, service.NewStudentService)
func initServer() pb.StudentServer {
	panic(wire.Build(SuperSet))
}