package procdec

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/procbind"
	"github.com/orglang/go-sdk/adt/uniqsym"
)

func (dto DecSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.X, validation.Required),
		validation.Field(&dto.ProcQN, uniqsym.Required...),
		validation.Field(&dto.Ys, procbind.Optional...),
	)
}
