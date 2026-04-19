package compsem

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/identity"
	"github.com/orglang/go-sdk/adt/revnum"
)

func (dto SemRef) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CompID, identity.Required...),
		validation.Field(&dto.CompRN, revnum.Required...),
	)
}
