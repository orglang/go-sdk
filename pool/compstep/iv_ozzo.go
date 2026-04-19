package compstep

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (dto StepSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CompRef, validation.Required),
		validation.Field(&dto.PoolExp, validation.Required),
	)
}
