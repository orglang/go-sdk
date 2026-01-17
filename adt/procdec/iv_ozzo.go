package procdec

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/identity"
	"github.com/orglang/go-sdk/adt/qualsym"
	"github.com/orglang/go-sdk/adt/termctx"
)

func (dto DecSpecME) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.X, validation.Required),
		validation.Field(&dto.ProcQN, qualsym.Required...),
		validation.Field(&dto.Ys, termctx.Optional...),
	)
}

func (dto IdentME) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.DecID, identity.Required...),
	)
}

func (dto DecRefME) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.DecID, identity.Required...),
	)
}
