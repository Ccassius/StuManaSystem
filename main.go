package main

import "StuManaSystem/model"

func main() {
	stuManger := model.NewStudentManager("jack", 20)
	stuManger.Start()
}

