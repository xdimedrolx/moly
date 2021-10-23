package useCase

import (
	"github.com/asaskevich/govalidator"
	"github.com/xdimedrolx/moly/internal/app/service"
)

type interactor struct {
	services service.Services
}

func (i interactor) Validate(dto interface{}) error {
	if _, err := govalidator.ValidateStruct(dto); err != nil {
		return err
	}
	return nil
}