package xactexp

import (
	"fmt"
)

type ExpSpec struct {
	K      expKind    `json:"kind"`
	Link   *LinkSpec  `json:"link,omitempty"`
	Plus   *LaborSpec `json:"plus,omitempty"`
	With   *LaborSpec `json:"with,omitempty"`
	Tensor *ResSpec   `json:"tensor,omitempty"`
	Lolli  *ResSpec   `json:"lolli,omitempty"`
	Up     *ShiftSpec `json:"up,omitempty"`
	Down   *ShiftSpec `json:"down,omitempty"`
}

type LinkSpec struct {
	XactQN string `json:"xact_qn"`
}

type LaborSpec struct {
	ProcQNs []string `json:"proc_qns"`
	ContExp ExpSpec  `json:"cont_exp"`
}

type ChoiceSpec struct {
	ProcQN  string  `json:"proc_qn"`
	ContExp ExpSpec `json:"cont_exp"`
}

// Resource
type ResSpec struct {
	ValExp  ExpSpec `json:"val_exp"`
	ContExp ExpSpec `json:"cont_exp"`
}

type ShiftSpec struct {
	ContExp ExpSpec `json:"cont_exp"`
}

type ExpRef struct {
	ExpVK int64   `json:"exp_vk" param:"vk"`
	K     expKind `json:"kind"`
}

type expKind string

const (
	One    expKind = "one"
	Link   expKind = "link"
	Plus   expKind = "plus"
	With   expKind = "with"
	Tensor expKind = "tensor"
	Lolli  expKind = "lolli"
	Up     expKind = "up"
	Down   expKind = "down"
)

func ErrKindUnexpected(got expKind) error {
	return fmt.Errorf("exp kind unexpected: %v", got)
}
