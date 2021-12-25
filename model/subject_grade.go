package model

// 科目成绩结构体，总分和平均分属性使用公开标识符，为了后续序列化到文件时，可以被访问到
type SubjectGrade struct {
	Chinese uint8
	Math uint8
	English uint8
	Total uint8
	Average float32
}

func NewSubjectGrade(chinese, math, english, total uint8, average float32) *SubjectGrade{
	return &SubjectGrade{
		Chinese: chinese,
		Math: math,
		English: english,
		Total: total,
		Average: average,
	}
}