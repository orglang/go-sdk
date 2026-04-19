package typeexp

import (
	"fmt"
)

type ExpSpec struct {
	K      expKind    `json:"kind"`
	Link   *LinkSpec  `json:"link,omitempty"`
	Tensor *ProdSpec  `json:"tensor,omitempty"`
	Lolli  *ProdSpec  `json:"lolli,omitempty"`
	Plus   *SumSpec   `json:"plus,omitempty"`
	With   *SumSpec   `json:"with,omitempty"`
	Up     *ShiftSpec `json:"up,omitempty"`
	Down   *ShiftSpec `json:"down,omitempty"`
}

type LinkSpec struct {
	TypeQN string `json:"type_qn"`
}

type ProdSpec struct {
	Val  ExpSpec `json:"val_es"`
	Cont ExpSpec `json:"cont_es"`
}

type SumSpec struct {
	Choices []ChoiceSpec `json:"choices"`
}

type ChoiceSpec struct {
	LabQN string  `json:"lab_qn"`
	Cont  ExpSpec `json:"cont_es"`
}

type ShiftSpec struct {
	Cont ExpSpec `json:"cont_es"`
}

type ExpRef struct {
	ExpVK int64   `json:"exp_vk" param:"vk"`
	K     expKind `json:"kind"`
}

type expKind string

const (
	One    = expKind("one")
	Link   = expKind("link")
	Tensor = expKind("tensor")
	Lolli  = expKind("lolli")
	Plus   = expKind("plus")
	With   = expKind("with")
	Up     = expKind("up")
	Down   = expKind("down")
)

func ErrKindUnexpected(got expKind) error {
	return fmt.Errorf("exp kind unexpected: %v", got)
}
