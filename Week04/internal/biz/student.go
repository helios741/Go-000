package biz

type Student struct {
	Id int32
	Age int32
	Sex string
	Name string
	Like string
}

type StudentRepo interface {
	GetStudent(id int32) (*Student, error)
}

func NewStudentUsecase(repo StudentRepo) *StudentUsecase {
	return &StudentUsecase{repo: repo}
}

type StudentUsecase struct {
	repo StudentRepo
}

func (su *StudentUsecase) Get(id int32)(*Student, error) {
	stu, err := su.repo.GetStudent(id)
	if err != nil {
		return nil, err
	}
	// some op
	return stu, nil
}