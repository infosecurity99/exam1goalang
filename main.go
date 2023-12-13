package main

import (
	"modul/read"
	"modul/task"
)

func main() {
	info, err := read.OpenJson()
	if err != nil {
		panic(err)
	}

	//1 Task .
	task.Task1(info)

	//2 Task .
	task.Task2(info)

	//3 Task
	task.Task3(info)

	//4 Task
	task.Task4(info)

	//5 Task
	task.Task5(info)

	//6 Task
	task.Task6(info)

	//7 Task
	task.Task7(info)

	//8 Task
	task.Task8(info)

	//9 Task
	task.Task9(info)

	//10 Task
	task.Task10(info)

	//11 Task
	task.Task11(info)

	//12 Task
	task.Task12(info)

	//13 Task
	task.Task13(info)

	//14 Task
	task.Task14(info)

	//15 Task
	task.Task15(info)
}
