package typedef

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/identity"
	"github.com/orglang/go-sdk/adt/uniqsym"
)

func (dto DefSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.TypeQN, uniqsym.Required...),
		validation.Field(&dto.TypeES, validation.Required),
	)
}

func (dto DefSnap) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.DefRef, identity.Required...),
		validation.Field(&dto.TypeES, validation.Required),
	)
}
