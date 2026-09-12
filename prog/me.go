package prog

import (
	"github.com/orglang/go-sdk/internal/prog"
)

type Spec struct {
	Exps []prog.ProgExp `parser:"@@*"`
}
