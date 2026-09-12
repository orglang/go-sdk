package typeexp

import (
	"fmt"
	"slices"

	"github.com/alecthomas/participle/v2/lexer"
)

type ExpSpec struct {
	K      expKind    `parser:"@@" json:"kind"`
	Link   *LinkSpec  `parser:"( 'link' @@" json:"link,omitempty"`
	Plus   *LaborSpec `parser:"| 'plus' @@" json:"plus,omitempty"`
	With   *LaborSpec `parser:"| 'with' @@" json:"with,omitempty"`
	Tensor *ResSpec   `parser:"| 'tensor' @@" json:"tensor,omitempty"`
	Lolli  *ResSpec   `parser:"| 'lolli' @@" json:"lolli,omitempty"`
	Up     *ShiftSpec `parser:"| 'up' @@" json:"up,omitempty"`
	Down   *ShiftSpec `parser:"| 'down' @@ )" json:"down,omitempty"`
}

type LinkSpec struct {
	TypeQN string `parser:"@Ident" json:"type_qn"`
}

type LaborSpec struct {
	ProcQNs []string `parser:"'(' @Ident* ')'" json:"proc_qns"`
	ContExp ExpSpec  `parser:"'{' @@? '}'" json:"cont_exp"`
}

type ChoiceSpec struct {
	ProcQN  string  `parser:"" json:"proc_qn"`
	ContExp ExpSpec `parser:"'{' @@? '}'" json:"cont_exp"`
}

// Resource
type ResSpec struct {
	ValExp  ExpSpec `parser:"" json:"val_exp"`
	ContExp ExpSpec `parser:"'{' @@? '}'" json:"cont_exp"`
}

type ShiftSpec struct {
	ContExp ExpSpec `parser:"'{' @@? '}'" json:"cont_exp"`
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
	One    expKind = "one"
	Link   expKind = "link"
	Plus   expKind = "plus"
	With   expKind = "with"
	Tensor expKind = "tensor"
	Lolli  expKind = "lolli"
	Up     expKind = "up"
	Down   expKind = "down"
)

var (
	Kinds = []expKind{One, Link, Plus, With, Tensor, Lolli, Up, Down}
)

func ErrKindUnexpected(got expKind) error {
	return fmt.Errorf("exp kind unexpected: %v", got)
}
