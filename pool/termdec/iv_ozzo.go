package termdec

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/termvar"
	"github.com/orglang/go-sdk/adt/uniqsym"
)

func (dto DecSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.TermQN, uniqsym.Required...),
		validation.Field(&dto.LiabVar, validation.Required),
		validation.Field(&dto.AssetVars, termvar.Required...),
	)
}
