package typeexp

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/identity"
	"github.com/orglang/go-sdk/adt/uniqsym"
)

func (dto ExpSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.K, kindRequired...),
		validation.Field(&dto.Link, validation.Required.When(dto.K == Link), validation.Skip.When(dto.K != Link)),
		validation.Field(&dto.Tensor, validation.Required.When(dto.K == Tensor), validation.Skip.When(dto.K != Tensor)),
		validation.Field(&dto.Lolli, validation.Required.When(dto.K == Lolli), validation.Skip.When(dto.K != Lolli)),
		validation.Field(&dto.Plus, validation.Required.When(dto.K == Plus), validation.Skip.When(dto.K != Plus)),
		validation.Field(&dto.With, validation.Required.When(dto.K == With), validation.Skip.When(dto.K != With)),
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

const (
	MinChoiceNr = 1
	MaxChoiceNr = 10
)

func (dto SumSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Choices,
			validation.Required,
			validation.Length(MinChoiceNr, MaxChoiceNr),
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
	validation.In(One, Link, Tensor, Lolli, Plus, With),
}
