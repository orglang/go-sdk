package procdec

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/identity"
	"github.com/orglang/go-sdk/adt/termctx"
	"github.com/orglang/go-sdk/adt/uniqsym"
)

func (dto DecSpecME) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.X, validation.Required),
		validation.Field(&dto.ProcQN, uniqsym.Required...),
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
