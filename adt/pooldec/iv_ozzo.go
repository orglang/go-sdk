package pooldec

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/poolbind"
	"github.com/orglang/go-sdk/adt/uniqsym"
)

func (dto DecSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.PoolQN, uniqsym.Required...),
		validation.Field(&dto.ProviderBS, validation.Required),
		validation.Field(&dto.ClientBSes, poolbind.Optional...),
	)
}
