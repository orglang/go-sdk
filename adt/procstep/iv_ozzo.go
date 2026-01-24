package procstep

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (dto StepSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.ExecRef, validation.Required),
		validation.Field(&dto.ProcES, validation.Required),
	)
}
