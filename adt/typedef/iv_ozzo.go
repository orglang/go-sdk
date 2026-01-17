package typedef

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/identity"
	"github.com/orglang/go-sdk/adt/qualsym"
	"github.com/orglang/go-sdk/adt/revnum"
)

func (dto DefSpecME) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.TypeQN, qualsym.Required...),
		validation.Field(&dto.TypeES, validation.Required),
	)
}

func (dto IdentME) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.DefID, identity.Required...),
	)
}

func (dto DefRefME) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.DefID, identity.Required...),
		validation.Field(&dto.DefRN, revnum.Optional...),
	)
}

func (dto DefSnapME) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.DefID, identity.Required...),
		validation.Field(&dto.DefRN, revnum.Optional...),
		validation.Field(&dto.TypeES, validation.Required),
	)
}
