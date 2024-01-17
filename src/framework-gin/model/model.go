package model

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string `json:"name" validate:"min=3,max=40"`
	CPF  string `json:"cpf" validate:"len=11,regexp=^[0-9]*$"`
	RG   string `json:"rg" validate:"len=9,regexp=^[0-9]*$"`
}

func ValidateStudents(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}
	return nil
}
