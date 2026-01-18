package uniqref

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/identity"
	"github.com/orglang/go-sdk/adt/revnum"
)

func (dto Msg) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.ID, identity.Required...),
		validation.Field(&dto.RN, revnum.Optional...),
	)
}
