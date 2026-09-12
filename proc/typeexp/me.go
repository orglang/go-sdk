package typeexp

import (
	"fmt"
	"slices"

	"github.com/alecthomas/participle/v2/lexer"
)

type ExpSpec struct {
	K      expKind    `parser:"@@" json:"kind"`
	Link   *LinkSpec  `parser:"( 'link' @@" json:"link,omitempty"`
	Tensor *ProdSpec  `parser:"| 'tensor' @@" json:"tensor,omitempty"`
	Lolli  *ProdSpec  `parser:"| 'lolli' @@" json:"lolli,omitempty"`
	Plus   *SumSpec   `parser:"| 'plus' @@" json:"plus,omitempty"`
	With   *SumSpec   `parser:"| 'with' @@" json:"with,omitempty"`
	Up     *ShiftSpec `parser:"| 'up' @@" json:"up,omitempty"`
	Down   *ShiftSpec `parser:"| 'down' @@ )" json:"down,omitempty"`
}

type LinkSpec struct {
	TypeQN string `parser:"@Ident" json:"type_qn"`
}

type ProdSpec struct {
	Val  ExpSpec `parser:"'{' @@? '}'" json:"val_es"`
	Cont ExpSpec `parser:"'{' @@? '}'" json:"cont_es"`
}

type SumSpec struct {
	Choices []ChoiceSpec `parser:"'{' @@* '}'" json:"choices"`
}

type ChoiceSpec struct {
	LabQN string  `parser:"'@'@Ident" json:"lab_qn"`
	Cont  ExpSpec `parser:"'{' @@? '}'" json:"cont_es"`
}

type ShiftSpec struct {
	Cont ExpSpec `parser:"'{' @@? '}'" json:"cont_es"`
}

type ExpRef struct {
	ExpVK int64   `json:"exp_vk" param:"vk"`
	K     expKind `json:"kind"`
}

type expKind string

func (k *expKind) Parse(lex *lexer.PeekingLexer) error {
	kind := expKind(lex.Peek().Value)
	if slices.Contains(Kinds, kind) {
		*k = kind
	}
	return nil
}

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

var (
	Kinds = []expKind{One, Link, Plus, With, Tensor, Lolli, Up, Down}
)

func ErrKindUnexpected(got expKind) error {
	return fmt.Errorf("exp kind unexpected: %v", got)
}
