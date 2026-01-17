package typeexp

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/identity"
	"github.com/orglang/go-sdk/adt/qualsym"
)

func (dto ExpSpecME) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.K, kindRequired...),
		validation.Field(&dto.Link, validation.Required.When(dto.K == LinkExp), validation.Skip.When(dto.K != LinkExp)),
		validation.Field(&dto.Tensor, validation.Required.When(dto.K == TensorExp), validation.Skip.When(dto.K != TensorExp)),
		validation.Field(&dto.Lolli, validation.Required.When(dto.K == LolliExp), validation.Skip.When(dto.K != LolliExp)),
		validation.Field(&dto.Plus, validation.Required.When(dto.K == PlusExp), validation.Skip.When(dto.K != PlusExp)),
		validation.Field(&dto.With, validation.Required.When(dto.K == WithExp), validation.Skip.When(dto.K != WithExp)),
	)
}

func (dto LinkSpecME) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.TypeQN, qualsym.Required...),
	)
}

func (dto ProdSpecME) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.ValES, validation.Required),
		validation.Field(&dto.ContES, validation.Required),
	)
}

func (dto SumSpecME) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Choices,
			validation.Required,
			validation.Length(1, 10),
			validation.Each(validation.Required),
		),
	)
}

func (dto ChoiceSpecME) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Label, qualsym.Required...),
		validation.Field(&dto.ContES, validation.Required),
	)
}

func (dto ExpRefME) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.ExpID, identity.Required...),
		validation.Field(&dto.K, kindRequired...),
	)
}

var kindRequired = []validation.Rule{
	validation.Required,
	validation.In(OneExp, LinkExp, TensorExp, LolliExp, PlusExp, WithExp),
}
