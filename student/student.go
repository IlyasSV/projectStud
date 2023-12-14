package student

import "fmt"

func CreateStudent(id int, name string, course string) {
	student := Student{ID: id, Name: name, Course: course}
	students = append(students, student)
	fmt.Println("Студент добавлен: ", student)
}

func GetStudentId(id int) {
	for _, student := range students {
		if student.ID == id {
			fmt.Printf("ID: %d, Name: %s, Course: %s\n", student.ID, student.Name, student.Course)
			return
		}
	}
	fmt.Printf("Студент с ID: %d не найден\n", id)
}

func UpdateStudent(id int, name string, course string) {
	for i, student := range students {
		if student.ID == id {
			students[i].Name = name
			students[i].Course = course
			fmt.Println("Данные студента обновлены: ", students[i])
			return
		}
	}
	fmt.Printf("Студент с ID: %d не найден\n", id)
}

func DeleteStudent(id int) {
	for i, student := range students {
		if student.ID == id {
			students = append(students[:i], students[i+1:]...)
			fmt.Println("Студент c удален:", student)
			return
		}
	}
	fmt.Printf("Студент с ID %d не найден.\n", id)
}

func (s *Student) Print() {
	fmt.Printf("ID: %d Name: %s Course: %s\n", s.ID, s.Name, s.Course)
}

func PrintStudents() {
	fmt.Println("\nИнформация о студентах")
	for _, student := range students {
		student.Print()
	}
}
