package xactexp

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

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
		validation.Field(&dto.XactQN, uniqsym.Required...),
	)
}

func (dto ResSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.ValES, validation.Required),
		validation.Field(&dto.ContES, validation.Required),
	)
}

func (dto LaborSpec) Validate() error {
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
		validation.Field(&dto.ProcQN, uniqsym.Required...),
		validation.Field(&dto.ContES, validation.Required),
	)
}

var kindRequired = []validation.Rule{
	validation.Required,
	validation.In(One, Link, Tensor, Lolli, Plus, With),
}
