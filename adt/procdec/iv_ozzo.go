package procdec

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/uniqsym"
)

const (
	MinContNr = 0
	MaxContNr = 10
)

func (dto DecSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.DescQN, uniqsym.Required...),
		validation.Field(&dto.ProviderVS, validation.Required),
		validation.Field(&dto.ClientVSes,
			validation.Length(MinContNr, MaxContNr),
			validation.Each(validation.Required),
		),
	)
}
