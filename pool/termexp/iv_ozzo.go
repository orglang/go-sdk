package termexp

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/symbol"
	"github.com/orglang/go-sdk/adt/uniqsym"
)

var expKindRequired = []validation.Rule{
	validation.Required,
	validation.In(Hire, Fire, Apply, Quit, Acquire, Release, Accept, Detach, Spawn),
}

func (dto ExpSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.K, expKindRequired...),
		validation.Field(&dto.Hire, validation.Required.When(dto.K == Hire)),
		validation.Field(&dto.Fire, validation.Required.When(dto.K == Fire)),
		validation.Field(&dto.Apply, validation.Required.When(dto.K == Apply)),
		validation.Field(&dto.Quit, validation.Required.When(dto.K == Quit)),
		validation.Field(&dto.Acquire, validation.Required.When(dto.K == Acquire)),
		validation.Field(&dto.Release, validation.Required.When(dto.K == Release)),
		validation.Field(&dto.Accept, validation.Required.When(dto.K == Accept)),
		validation.Field(&dto.Detach, validation.Required.When(dto.K == Detach)),
		validation.Field(&dto.Spawn, validation.Required.When(dto.K == Spawn)),
	)
}

func (dto HireSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommChnlPH, symbol.Required...),
		validation.Field(&dto.ProcTermQN, uniqsym.Required...),
		validation.Field(&dto.ContExp, validation.Required),
	)
}

func (dto FireSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommChnlPH, uniqsym.Required...),
	)
}

func (dto ApplySpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommChnlPH, symbol.Required...),
		validation.Field(&dto.ProcTermQN, uniqsym.Required...),
		validation.Field(&dto.ContExp, validation.Required),
	)
}

func (dto QuitSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.ProcDescQN, uniqsym.Required...),
	)
}

func (dto AcquireSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommChnlPH, symbol.Required...),
		validation.Field(&dto.ContExp, validation.Required),
	)
}

func (dto ReleaseSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommChnlPH, symbol.Required...),
	)
}

func (dto AcceptSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommChnlPH, symbol.Required...),
		validation.Field(&dto.ContExp, validation.Required),
	)
}

func (dto DetachSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommChnlPH, symbol.Required...),
	)
}

const (
	MinValNr = 0
	MaxValNr = 10
)

func (dto SpawnSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.ProcTermRef, validation.Required),
		validation.Field(&dto.ProcCompRefs,
			validation.Length(MinValNr, MaxValNr),
			validation.Each(validation.Required),
		),
	)
}
