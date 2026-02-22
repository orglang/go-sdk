package procexp

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/implvar"
	"github.com/orglang/go-sdk/adt/symbol"
	"github.com/orglang/go-sdk/adt/uniqsym"
)

var expKindRequired = []validation.Rule{
	validation.Required,
	validation.In(Close, Wait, Send, Recv, Lab, Case, Fwd, Call),
}

func (dto ExpSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.K, expKindRequired...),
		validation.Field(&dto.Close, validation.Required.When(dto.K == Close)),
		validation.Field(&dto.Wait, validation.Required.When(dto.K == Wait)),
		validation.Field(&dto.Send, validation.Required.When(dto.K == Send)),
		validation.Field(&dto.Recv, validation.Required.When(dto.K == Recv)),
		validation.Field(&dto.Lab, validation.Required.When(dto.K == Lab)),
		validation.Field(&dto.Case, validation.Required.When(dto.K == Case)),
		validation.Field(&dto.Fwd, validation.Required.When(dto.K == Fwd)),
		validation.Field(&dto.Call, validation.Required.When(dto.K == Call)),
	)
}

func (dto CloseSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommChnlPH, symbol.Required...),
	)
}

func (dto WaitSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommChnlPH, symbol.Required...),
		validation.Field(&dto.ContES, validation.Required),
	)
}

func (dto SendSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommChnlPH, symbol.Required...),
		validation.Field(&dto.ValChnlPH, symbol.Required...),
	)
}

func (dto RecvSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommChnlPH, symbol.Required...),
		validation.Field(&dto.BindChnlPH, symbol.Required...),
		validation.Field(&dto.ContES, validation.Required),
	)
}

func (dto LabSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommChnlPH, symbol.Required...),
		validation.Field(&dto.PatternQN, uniqsym.Required...),
	)
}

const (
	MinContNr = 1
	MaxContNr = 10
)

func (dto CaseSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommChnlPH, symbol.Required...),
		validation.Field(&dto.ContBSes,
			validation.Required,
			validation.Length(MinContNr, MaxContNr),
			validation.Each(validation.Required),
		),
	)
}

func (dto BranchSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.PatternQN, uniqsym.Required...),
		validation.Field(&dto.ContES, validation.Required),
	)
}

func (dto CallSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.BindChnlPH, symbol.Required...),
		validation.Field(&dto.ProcDescQN, uniqsym.Required...),
		validation.Field(&dto.ValChnlPHs, implvar.Required...),
	)
}

func (dto FwdSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommChnlPH, symbol.Required...),
		validation.Field(&dto.ContChnlPH, symbol.Required...),
	)
}
