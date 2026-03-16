package poolexec

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/implvar"
	"github.com/orglang/go-sdk/adt/uniqsym"
)

func (dto ExecSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.DescQN, uniqsym.Required...),
		validation.Field(&dto.LiabVar, validation.Required),
		validation.Field(&dto.AssetVars, implvar.Required...),
	)
}
