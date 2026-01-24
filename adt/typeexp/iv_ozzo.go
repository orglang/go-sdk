package typeexp

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/identity"
	"github.com/orglang/go-sdk/adt/uniqsym"
)

func (dto ExpSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.K, kindRequired...),
		validation.Field(&dto.Link, validation.Required.When(dto.K == LinkExp), validation.Skip.When(dto.K != LinkExp)),
		validation.Field(&dto.Tensor, validation.Required.When(dto.K == TensorExp), validation.Skip.When(dto.K != TensorExp)),
		validation.Field(&dto.Lolli, validation.Required.When(dto.K == LolliExp), validation.Skip.When(dto.K != LolliExp)),
		validation.Field(&dto.Plus, validation.Required.When(dto.K == PlusExp), validation.Skip.When(dto.K != PlusExp)),
		validation.Field(&dto.With, validation.Required.When(dto.K == WithExp), validation.Skip.When(dto.K != WithExp)),
	)
}

func (dto LinkSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.TypeQN, uniqsym.Required...),
	)
}

func (dto ProdSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.ValES, validation.Required),
		validation.Field(&dto.ContES, validation.Required),
	)
}

func (dto SumSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Choices,
			validation.Required,
			validation.Length(1, 10),
			validation.Each(validation.Required),
		),
	)
}

func (dto ChoiceSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.LabQN, uniqsym.Required...),
		validation.Field(&dto.ContES, validation.Required),
	)
}

func (dto ExpRef) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.ExpID, identity.Required...),
		validation.Field(&dto.K, kindRequired...),
	)
}

var kindRequired = []validation.Rule{
	validation.Required,
	validation.In(OneExp, LinkExp, TensorExp, LolliExp, PlusExp, WithExp),
}
