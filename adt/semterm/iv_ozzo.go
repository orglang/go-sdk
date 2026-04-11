package semterm

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/identity"
	"github.com/orglang/go-sdk/adt/revnum"
)

func (dto TermRef) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.TermID, identity.Required...),
		validation.Field(&dto.TermRN, revnum.Required...),
	)
}
