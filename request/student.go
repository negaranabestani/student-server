package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Student struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
}

func (req Student) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.FirstName, validation.Required, is.UTFLetter),
		validation.Field(&req.LastName, validation.Required, is.UTFLetter),
	)
}
