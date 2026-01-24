package typeexp

import "fmt"

type ExpSpec struct {
	K      expKind    `json:"kind"`
	Link   *LinkSpec  `json:"link,omitempty"`
	Tensor *ProdSpec  `json:"tensor,omitempty"`
	Lolli  *ProdSpec  `json:"lolli,omitempty"`
	Plus   *SumSpec   `json:"plus,omitempty"`
	With   *SumSpec   `json:"with,omitempty"`
	Up     *ShiftSpec `json:"up,omitempty"`
	Down   *ShiftSpec `json:"down,omitempty"`
	Xact   *XactSpec  `json:"xact,omitempty"`
}

type LinkSpec struct {
	TypeQN string `json:"type_qn"`
}

type ProdSpec struct {
	ValES  ExpSpec `json:"val_es"`
	ContES ExpSpec `json:"cont_es"`
}

type SumSpec struct {
	Choices []ChoiceSpec `json:"choices"`
}

type ChoiceSpec struct {
	LabQN  string  `json:"lab_qn"`
	ContES ExpSpec `json:"cont_es"`
}

type ShiftSpec struct {
	ContES ExpSpec `json:"cont_es"`
}

type XactSpec struct {
	ContESs map[string]ExpSpec `json:"cont_ess"`
}

type ExpRef struct {
	ExpID string  `json:"exp_id" param:"id"`
	K     expKind `json:"kind"`
}

type expKind string

const (
	OneExp    = expKind("one")
	LinkExp   = expKind("link")
	TensorExp = expKind("tensor")
	LolliExp  = expKind("lolli")
	PlusExp   = expKind("plus")
	WithExp   = expKind("with")
	UpExp     = expKind("up")
	DownExp   = expKind("down")
	XactExp   = expKind("xact")
)

func ErrKindUnexpected(got expKind) error {
	return fmt.Errorf("exp kind unexpected: %v", got)
}
