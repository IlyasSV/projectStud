package main

import (
	"awesomeProject2/student"
	"fmt"
)

func main() {
	fmt.Println("\nДобавляем нового студента в базу данных Студенты")
	student.CreateStudent(1, "Петя", "2 курс")
	student.CreateStudent(2, "Dona", "3 course")
	student.CreateStudent(3, "Jon", "3 course")
	student.CreateStudent(4, "Jorden", "4 course")

	fmt.Println("\nПолучаем информацию из базы данных о студенте по id")
	student.GetStudentId(2)
	student.GetStudentId(4)
	student.GetStudentId(6)

	fmt.Println("\nОбновляем информацию в базе данных о студенте")
	student.UpdateStudent(1, "Petr", "2 course")
	student.UpdateStudent(4, "Jorden", "5 course")

	fmt.Println("\nУдаляем информацию из базы данных о студенте")
	student.DeleteStudent(3)

	fmt.Println("\nВыводим актуальную информацию из базы данных о студенте")
	student.PrintStudents()

}
