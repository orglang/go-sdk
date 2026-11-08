package prog

import (
	"github.com/alecthomas/participle/v2"

	pooltypedef "github.com/orglang/go-sdk/pool/typedef"
	proctypedef "github.com/orglang/go-sdk/proc/typedef"
)

func MsgFromText(text string) (Spec, error) {
	parser, buildErr := participle.Build[Spec](participle.Union[ProgExp](pooltypedef.DefSpec{}, proctypedef.DefSpec{}))
	if buildErr != nil {
		return Spec{}, buildErr
	}
	spec, parseErr := parser.ParseString("", text)
	if parseErr != nil {
		return Spec{}, parseErr
	}
	return *spec, nil
}
