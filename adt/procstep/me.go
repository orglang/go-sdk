package procstep

import (
	"github.com/orglang/go-sdk/adt/procexp"
)

type StepSpec struct {
	ExecID string          `json:"exec_id"`
	ProcID string          `json:"proc_id"`
	ProcES procexp.ExpSpec `json:"proc_es"`
}
