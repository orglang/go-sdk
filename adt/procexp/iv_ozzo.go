package procexp

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/orglang/go-sdk/adt/identity"
	"github.com/orglang/go-sdk/adt/procbind"
	"github.com/orglang/go-sdk/adt/uniqsym"
)

var expKindRequired = []validation.Rule{
	validation.Required,
	validation.In(CloseExp, WaitExp, SendExp, RecvExp, LabExp, CaseExp, SpawnExp, FwdExp, CallExp),
}

func (dto ExpSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.K, expKindRequired...),
		validation.Field(&dto.Close, validation.Required.When(dto.K == CloseExp)),
		validation.Field(&dto.Wait, validation.Required.When(dto.K == WaitExp)),
		validation.Field(&dto.Send, validation.Required.When(dto.K == SendExp)),
		validation.Field(&dto.Recv, validation.Required.When(dto.K == RecvExp)),
		validation.Field(&dto.Lab, validation.Required.When(dto.K == LabExp)),
		validation.Field(&dto.Case, validation.Required.When(dto.K == CaseExp)),
		validation.Field(&dto.Spawn, validation.Required.When(dto.K == SpawnExp)),
		validation.Field(&dto.Fwd, validation.Required.When(dto.K == FwdExp)),
		validation.Field(&dto.Call, validation.Required.When(dto.K == CallExp)),
	)
}

func (dto CloseSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommPH, validation.Required),
	)
}

func (dto WaitSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommPH, validation.Required),
		validation.Field(&dto.ContES, validation.Required),
	)
}

func (dto SendSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommPH, validation.Required),
		validation.Field(&dto.ValPH, validation.Required),
	)
}

func (dto RecvSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommPH, validation.Required),
		validation.Field(&dto.BindPH, validation.Required),
		validation.Field(&dto.ContES, validation.Required),
	)
}

func (dto LabSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommPH, validation.Required),
		validation.Field(&dto.LabQN, uniqsym.Required...),
	)
}

func (dto CaseSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommPH, validation.Required),
		validation.Field(&dto.ContBSs,
			validation.Required,
			validation.Length(1, 10),
			validation.Each(validation.Required),
		),
	)
}

func (dto BranchSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.LabQN, uniqsym.Required...),
		validation.Field(&dto.ContES, validation.Required),
	)
}

func (dto CallSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommPH, validation.Required),
		validation.Field(&dto.ProcQN, identity.Required...),
		validation.Field(&dto.ValPHs, procbind.Optional...),
	)
}

func (dto SpawnSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommPH, validation.Required),
		validation.Field(&dto.DecID, identity.Required...),
		validation.Field(&dto.BindPHs, procbind.Optional...),
	)
}

func (dto FwdSpec) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.CommPH, validation.Required),
		validation.Field(&dto.ContPH, validation.Required),
	)
}
