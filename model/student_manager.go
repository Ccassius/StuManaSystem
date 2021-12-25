package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// 文件路径
const FILE_NAME = "D:\\StuInfo.txt"

// 学生管理员结构体
type StudentManager struct {
	Name string
	Age uint8
	studentMap map[uint8]*Student
}

// 学生管理员的构造方法
func NewStudentManager(name string, age uint8) StuManaSystem{
	return &StudentManager{
		Name: name,
		Age: age,
		studentMap: make(map[uint8]*Student),
	}
}

// 接口
type StuManaSystem interface {
	GetStuInfo() error
	SaveStuInfo() error
	AddStu(id uint8, name string, age uint8, chinese uint8, math uint8, english uint8, total uint8, average float32)  error
	DeleteStu(id uint8) error
	ChangeStu(stu *Student, newId uint8, newName string, newAge, newChinese, newMath, newEnglish uint8) error
	MenuShow()
	ShowAllStu() error
	SelectStu(id uint8) error
	IsExistStu(id uint8) (bool, error)
	Start()
}

// 从文件中读取数据
func (s *StudentManager) GetStuInfo() error {
	// 从文件中读取数据
	content, err := ioutil.ReadFile(FILE_NAME)
	if err != nil {
		return err
	}

	// 反序列化
	json.Unmarshal(content, &s.studentMap)
	return nil
}


// 将数据序列化存储到文件中
func (s *StudentManager) SaveStuInfo() error {
	// 序列化
	content, err :=  json.Marshal(&s.studentMap)
	if err != nil {
		return err
	}

	// 存储到文件中, 文件权限为可读写
	err = ioutil.WriteFile(FILE_NAME, content, 0666)
	if err != nil {
		return err
	}
	return nil
}


// 增加一个新学生
func (s *StudentManager) AddStu(id uint8, name string, age uint8, chinese uint8, math uint8, english uint8, total uint8, average float32)  error {
	grade := NewSubjectGrade(chinese, math, english, total, average)
	stu := NewStudent(id, name, age, grade)
	s.studentMap[stu.Id] = stu
	return nil
}


// 删除一个学生（学生map）
func (s *StudentManager) DeleteStu(id uint8) error {
	if _, ok := s.studentMap[id]; ok {
		delete(s.studentMap, id)
		return nil
	} else {
		return errors.New(fmt.Sprintf("删除失败，不存在id为 %d 的学生\n", id))
	}
}


// 修改一个学生
func (s *StudentManager) ChangeStu(stu *Student, newId uint8, newName string, newAge, newChinese, newMath, newEnglish uint8) error{
	if stu == nil {
		return errors.New("修改失败，参数 stu 为 nil")
	}
	stu.Id = newId
	stu.Name = newName
	stu.Age = newAge
	stu.Grade.Chinese = newChinese
	stu.Grade.Math = newMath
	stu.Grade.English = newEnglish
	stu.Grade.Total = newChinese + newEnglish + newMath
	stu.Grade.Average = float32(stu.Grade.Total) / 3.0
	return nil
}


//菜单
func (s *StudentManager) MenuShow() {
	fmt.Println("欢迎使用学生管理系统")
	fmt.Printf("%d........................增加学生\n", ADD)
	fmt.Printf("%d........................删除学生\n", DELETE)
	fmt.Printf("%d........................改变学生信息\n", CHANGE)
	fmt.Printf("%d........................查询学生信息\n", SELECT)
	fmt.Printf("%d........................显示所有学生\n", SHOW)
	fmt.Printf("%d........................退出\n", EXIT)
	fmt.Print("请输入选择:\n")
}


//显示所有学生
func (s *StudentManager) ShowAllStu() error {
	if s.studentMap == nil {
		return errors.New("studentMap 为 nil")
	}
	if len(s.studentMap) == 0 {
		fmt.Println("没有学生信息")
	} else {
		fmt.Println("Id\t\t姓名\t\t年龄\t\t语文\t\t数学\t\t英语\t\t总分\t\t平均分")
		for _, stu := range s.studentMap {
			fmt.Printf("%v\t\t%v\t\t%v\t\t%v\t\t%v\t\t%v\t\t%v\t\t%v\n", stu.Id, stu.Name, stu.Age, stu.Grade.Chinese,
				stu.Grade.Math, stu.Grade.English, stu.Grade.Total, stu.Grade.Average)
		}
	}
	return nil
}


// 根据编号查询学生信息
func (s *StudentManager) SelectStu(id uint8) error {
	if s.studentMap == nil {
		return errors.New("studentMap 为 nil")
	}
	// 判断对应学生是否存在
	if stu, ok := s.studentMap[id]; ok {
		fmt.Println("Id\t\t姓名\t\t年龄\t\t语文\t\t数学\t\t英语\t\t总分\t\t平均分")
		fmt.Printf("%v\t\t%v\t\t%v\t\t%v\t\t%v\t\t%v\t\t%v\t\t%v\n", stu.Id, stu.Name, stu.Age, stu.Grade.Chinese,
			stu.Grade.Math, stu.Grade.English, stu.Grade.Total, stu.Grade.Average)
		return nil
	} else {
		return errors.New(fmt.Sprintf("不存在id为 %d 的学生\n", id))
	}

}


// 判断编号为id的学生是否存在
func (s *StudentManager) IsExistStu(id uint8) (bool, error) {
	if s.studentMap == nil {
		return false, errors.New("studentMap 为 nil,学生不存在")
	}
	_, exist := s.studentMap[id]
	return exist, nil
}


// 整合学生管理系统的功能
func (s *StudentManager) Start() {
	// 从文件中读取数据存储到 s.studentMap 中
	err := s.GetStuInfo()
	if err != nil {
		fmt.Println("读取文件数据出错：", err)
		return
	}

	for {
		s.MenuShow()
		var inputNum uint8
		fmt.Scan(&inputNum)

		switch inputNum {
		case ADD:
			fmt.Println("请输入学生id、姓名、年龄、语文成绩、数学成绩、英语成绩")
			var (
				id uint8
				name string
				age uint8
				chinese uint8
				math uint8
				english uint8
			)
			for {
				fmt.Scan(&id, &name, &age, &chinese, &math, &english)
				exist, err := s.IsExistStu(id)
				if exist {
					fmt.Printf("id为%d的学生已存在\n", id)

					fmt.Println("继续添加请按1，返回上一级请按任意键")
					var flag int
					fmt.Scan(&flag)
					if flag == 1 {
						fmt.Println("请输入学生id、姓名、年龄、语文成绩、数学成绩、英语成绩")
						continue
					} else {
						break
					}

					continue
				}
				err = s.AddStu(id, name, age, chinese,math, english, (chinese + math + english), (float32)(chinese + math + english) / 3)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("添加学生成功")
				}
				fmt.Println("继续添加请按1，返回上一级请按任意键")
				var flag int
				fmt.Scan(&flag)
				if flag == 1 {
					fmt.Println("请输入学生id、姓名、年龄、语文成绩、数学成绩、英语成绩")
					continue
				} else {
					break
				}
			}
		case DELETE:
			var id uint8
			fmt.Println("请输入要删除的学生的id")
			fmt.Scan(&id)
			err := s.DeleteStu(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("删除成功")
			}
		case CHANGE:
			var id uint8
			fmt.Println("请输入需要修改的学生id")
			fmt.Scan(&id)
			exist, err := s.IsExistStu(id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			// 如果编号不存在，则返回提示
			if !exist {
				fmt.Printf("不存在id为 %d 的学生\n", id)
			} else {
				fmt.Println("请输入学生姓名、年龄、语文成绩、数学成绩、英语成绩")
				var (
					name string
					age uint8
					chinese uint8
					math uint8
					english uint8
				)
				fmt.Scan(&name, &age, &chinese, &math, &english)
				stu := s.studentMap[id]
				err := s.ChangeStu(stu, id, name, age, chinese, math, english)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("修改成功")
				}
			}
		case SELECT:
			var id uint8
			fmt.Println("请输入要查询的学生的id")
			fmt.Scan(&id)
			err := s.SelectStu(id)
			if err != nil {
				fmt.Println(err)
			}
		case SHOW:
			err := s.ShowAllStu()
			if err != nil {
				fmt.Println(err)
			}
		case EXIT:
			err := s.SaveStuInfo()
			if err != nil {
				fmt.Println("保存文件出错：", err)
			} else {
				fmt.Println("已退出")
				return
			}
		default:
			fmt.Println("无效的输入!")
		}
	}
}