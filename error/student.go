package error

import "fmt"

type StudentNotFoundError struct {
	Id int64
}

func (e StudentNotFoundError) Error() string {
	return fmt.Sprintf("%d", e.Id) + " Not Found!"
}

type DuplicateStudentError struct {
	Id int64
}

func (e DuplicateStudentError) Error() string {
	return fmt.Sprintf("%d", e.Id) + "Already Exists!"
}

type EmptyStoreError struct {
}

func (e EmptyStoreError) Error() string {
	return "No Records Yet!"
}
