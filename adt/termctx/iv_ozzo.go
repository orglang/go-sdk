package termctx

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/uniqsym"
)

var Optional = []validation.Rule{
	validation.Length(1, 10),
	validation.Each(validation.Required),
}

func (dto BindClaimME) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.BindPH, uniqsym.Optional...),
		validation.Field(&dto.TypeQN, uniqsym.Required...),
	)
}
