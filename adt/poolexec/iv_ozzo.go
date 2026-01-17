package poolexec

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/identity"
)

func (dto ExecSpecME) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.SupID, identity.Optional...),
	)
}
