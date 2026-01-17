package procstep

import (
	"github.com/orglang/go-sdk/adt/procexp"
)

type StepSpecME struct {
	ExecID string            `json:"exec_id"`
	ProcID string            `json:"proc_id"`
	ProcES procexp.ExpSpecME `json:"proc_es"`
}
