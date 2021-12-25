package model

// 学生结构体
type Student struct {
	Id uint8
	Name string
	Age uint8
	Grade SubjectGrade
}

func NewStudent(id uint8, name string, age uint8, grade *SubjectGrade) *Student {
	return &Student{
		Id: id,
		Name: name,
		Age: age,
		Grade: *grade,
	}
}

