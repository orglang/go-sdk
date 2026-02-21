package descvar

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/symbol"
	"github.com/orglang/go-sdk/adt/uniqsym"
)

var Optional = []validation.Rule{
	validation.Length(1, 10),
	validation.Each(validation.Required),
}

func (dto VarSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.ChnlPH, symbol.Required...),
		validation.Field(&dto.DescQN, uniqsym.Required...),
	)
}
