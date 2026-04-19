package termdec

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
		validation.Field(&dto.TermQN, uniqsym.Required...),
		validation.Field(&dto.LiabVar, validation.Required),
		validation.Field(&dto.AssetVars,
			validation.Length(MinContNr, MaxContNr),
			validation.Each(validation.Required),
		),
	)
}
