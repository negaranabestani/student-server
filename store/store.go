package store

import "github.com/negaranabestani/students/model"

type Student interface {
	GetAll() ([]model.Student, error)
	Get(Id int64) (model.Student, error)
	Save(student model.Student) error
}
