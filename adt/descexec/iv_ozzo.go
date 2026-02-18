package descexec

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/identity"
	"github.com/orglang/go-sdk/adt/revnum"
)

func (dto ExecRef) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.DescID, identity.Required...),
		validation.Field(&dto.DescRN, revnum.Required...),
	)
}
