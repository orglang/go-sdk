package xactexp

import "fmt"

type ExpSpec struct {
	K    expKind   `json:"kind"`
	Link *LinkSpec `json:"link,omitempty"`
	Plus *SumSpec  `json:"plus,omitempty"`
	With *SumSpec  `json:"with,omitempty"`
}

type LinkSpec struct {
	XactQN string `json:"xact_qn"`
}

type SumSpec struct {
	Choices []ChoiceSpec `json:"choices"`
}

type ChoiceSpec struct {
	TermQN string  `json:"term_qn"`
	ContES ExpSpec `json:"cont_es"`
}

type expKind string

const (
	One  = expKind("one")
	Link = expKind("link")
	Plus = expKind("plus")
	With = expKind("with")
)

func ErrKindUnexpected(got expKind) error {
	return fmt.Errorf("exp kind unexpected: %v", got)
}
