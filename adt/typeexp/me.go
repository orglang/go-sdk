package typeexp

import "fmt"

type ExpSpecME struct {
	K      expKindME   `json:"kind"`
	Link   *LinkSpecME `json:"link,omitempty"`
	Tensor *ProdSpecME `json:"tensor,omitempty"`
	Lolli  *ProdSpecME `json:"lolli,omitempty"`
	Plus   *SumSpecME  `json:"plus,omitempty"`
	With   *SumSpecME  `json:"with,omitempty"`
	Up     *FooSpecME  `json:"up,omitempty"`
	Down   *FooSpecME  `json:"down,omitempty"`
	Xact   *XactSpecME `json:"xact,omitempty"`
}

type LinkSpecME struct {
	TypeQN string `json:"type_qn"`
}

type ProdSpecME struct {
	ValES  ExpSpecME `json:"val_es"`
	ContES ExpSpecME `json:"cont_es"`
}

type SumSpecME struct {
	Choices []ChoiceSpecME `json:"choices"`
}

type ChoiceSpecME struct {
	Label  string    `json:"label"`
	ContES ExpSpecME `json:"cont_es"`
}

type FooSpecME struct {
	ContES ExpSpecME `json:"cont_es"`
}

type XactSpecME struct {
	ContESs map[string]ExpSpecME `json:"cont_ess"`
}

type ExpRefME struct {
	ExpID string    `json:"exp_id" param:"id"`
	K     expKindME `json:"kind"`
}

type expKindME string

const (
	OneExp    = expKindME("one")
	LinkExp   = expKindME("link")
	TensorExp = expKindME("tensor")
	LolliExp  = expKindME("lolli")
	PlusExp   = expKindME("plus")
	WithExp   = expKindME("with")
	UpExp     = expKindME("up")
	DownExp   = expKindME("down")
	XactExp   = expKindME("xact")
)

func ErrKindUnexpected(got expKindME) error {
	return fmt.Errorf("kind unexpected: %v", got)
}
