package store

import (
	error2 "github.com/negaranabestani/students/error"
	"github.com/negaranabestani/students/model"
	"go.uber.org/zap"
)

func NewStudentMemory(logger *zap.Logger) *StudentInMemory {
	return &StudentInMemory{
		Students: make(map[int64]model.Student),
		Logger:   logger,
	}
}

type StudentInMemory struct {
	Students map[int64]model.Student
	Logger   *zap.Logger
}

func (m *StudentInMemory) GetAll() ([]model.Student, error) {
	result := make([]model.Student, 0)
	if len(m.Students) == 0 {
		return nil, error2.EmptyStoreError{}
	}
	for _, v := range m.Students {
		result = append(result, v)
	}
	return result, nil
}
func (m *StudentInMemory) Get(Id int64) (model.Student, error) {
	s, ok := m.Students[Id]
	if !ok {
		return s, error2.StudentNotFoundError{Id: Id}
	}
	return s, nil
}
func (m *StudentInMemory) Save(student model.Student) error {
	s, ok := m.Students[student.Id]
	if ok {
		return error2.DuplicateStudentError{Id: s.Id}
	}
	m.Students[student.Id] = student
	return nil
}
