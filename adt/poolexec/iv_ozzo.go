package poolexec

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/implsubst"
	"github.com/orglang/go-sdk/adt/uniqsym"
)

func (dto ExecSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.DescQN, uniqsym.Required...),
		validation.Field(&dto.ProviderSS, validation.Required),
		validation.Field(&dto.ClientSSes, implsubst.Required...),
	)
}
