package compvar

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/uniqsym"
)

var Required = []validation.Rule{
	validation.Length(1, 10),
	validation.Each(validation.Required),
}

func (dto VarSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.ChnlPH, uniqsym.Optional...),
		validation.Field(&dto.TermQN, uniqsym.Required...),
	)
}
