package typedef

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/uniqsym"
)

func (dto DefSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.TypeQN, uniqsym.Required...),
		validation.Field(&dto.TypeExp, validation.Required),
	)
}

func (dto DefSnap) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.TypeRef, validation.Required),
		validation.Field(&dto.DefSpec, validation.Required),
	)
}
