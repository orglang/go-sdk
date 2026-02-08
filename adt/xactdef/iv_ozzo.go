package xactdef

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/uniqsym"
)

func (dto DefSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.XactQN, uniqsym.Required...),
		validation.Field(&dto.XactES, validation.Required),
	)
}
