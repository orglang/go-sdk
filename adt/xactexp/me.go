package xactexp

import "fmt"

type ExpSpec struct {
	K      expKind    `json:"kind"`
	Link   *LinkSpec  `json:"link,omitempty"`
	Plus   *LaborSpec `json:"plus,omitempty"`
	With   *LaborSpec `json:"with,omitempty"`
	Tensor *ResSpec   `json:"tensor,omitempty"`
	Lolli  *ResSpec   `json:"lolli,omitempty"`
}

type LinkSpec struct {
	XactQN string `json:"xact_qn"`
}

type LaborSpec struct {
	Choices []ChoiceSpec `json:"choices"`
}

type ChoiceSpec struct {
	ProcQN string  `json:"proc_qn"`
	ContES ExpSpec `json:"cont_es"`
}

// Resource
type ResSpec struct {
	ValES  ExpSpec `json:"val_es"`
	ContES ExpSpec `json:"cont_es"`
}

type expKind string

const (
	One    = expKind("one")
	Link   = expKind("link")
	Plus   = expKind("plus")
	With   = expKind("with")
	Tensor = expKind("tensor")
	Lolli  = expKind("lolli")
)

func ErrKindUnexpected(got expKind) error {
	return fmt.Errorf("exp kind unexpected: %v", got)
}
