package procstep

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/identity"
)

func (dto StepSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.ExecID, identity.Required...),
		validation.Field(&dto.ProcID, identity.Required...),
		validation.Field(&dto.ProcES, validation.Required),
	)
}
