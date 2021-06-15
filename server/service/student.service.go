package service

import "gin_server/server/model"

func GetStructFromData(rows [][]string) []model.Student {
	students := []model.Student{}
	for _, columns := range rows {
		var student model.Student
		if columns[0] == "name" {
			continue
		}
		student.Name = columns[0]
		student.Gender = columns[1]
		student.IdentityNumber = columns[2]
		student.Class = columns[3]
		student.Pin = columns[4]
		students = append(students, student)

	}
	return students
}

func Add(a, b int64) int64 {
	return a + b
}
