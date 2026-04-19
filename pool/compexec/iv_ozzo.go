package compexec

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/compvar"
	"github.com/orglang/go-sdk/adt/uniqsym"
)

func (dto ExecSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.TermQN, uniqsym.Required...),
		validation.Field(&dto.LiabVar, validation.Required),
		validation.Field(&dto.AssetVars, compvar.Required...),
	)
}
