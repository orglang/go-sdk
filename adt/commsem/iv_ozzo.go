package commsem

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/identity"
	"github.com/orglang/go-sdk/adt/revnum"
)

func (dto SemRef) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommID, identity.Required...),
		validation.Field(&dto.CommRN, revnum.Required...),
	)
}
