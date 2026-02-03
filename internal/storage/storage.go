package storage

import "github.com/Dipumane1318/PROJECT_STUDENTS_REST_API/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)

	GetStudentById(id int64) (types.Student,error)
	GetStudents() ([]types.Student, error)
}