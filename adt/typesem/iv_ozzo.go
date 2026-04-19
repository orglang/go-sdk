package typesem

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/identity"
	"github.com/orglang/go-sdk/adt/revnum"
)

func (dto SemRef) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.TypeID, identity.Required...),
		validation.Field(&dto.TypeRN, revnum.Required...),
	)
}
