package procstep

import (
	"github.com/orglang/go-sdk/adt/procexp"
	"github.com/orglang/go-sdk/adt/uniqref"
)

type StepSpec struct {
	ExecID string          `json:"exec_id"`
	ProcID string          `json:"proc_id"`
	ProcES procexp.ExpSpec `json:"proc_es"`
}

type stepKind string

const (
	MsgStep = stepKind("msg")
	SvcStep = stepKind("svc")
)

type StepRec struct {
	ExecRef uniqref.Msg `json:"exec_ref"`
	Msg     *MsgRec     `json:"msg"`
	Svc     *SvcRec     `json:"svc"`
}

type MsgRec struct {
}

type SvcRec struct {
}
