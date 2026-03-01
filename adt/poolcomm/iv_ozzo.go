package poolcomm

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (dto CommSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.ImplRef, validation.Required),
		validation.Field(&dto.PoolES, validation.Required),
	)
}
