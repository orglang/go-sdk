package procexp

import "fmt"

type expKind string

const (
	CloseExp = expKind("close")
	WaitExp  = expKind("wait")
	SendExp  = expKind("send")
	RecvExp  = expKind("recv")
	LabExp   = expKind("lab")
	CaseExp  = expKind("case")
	CallExp  = expKind("call")
	LinkExp  = expKind("link")
	SpawnExp = expKind("spawn")
	FwdExp   = expKind("fwd")
)

type ExpSpec struct {
	K     expKind    `json:"kind"`
	Close *CloseSpec `json:"close,omitempty"`
	Wait  *WaitSpec  `json:"wait,omitempty"`
	Send  *SendSpec  `json:"send,omitempty"`
	Recv  *RecvSpec  `json:"recv,omitempty"`
	Lab   *LabSpec   `json:"lab,omitempty"`
	Case  *CaseSpec  `json:"case,omitempty"`
	Spawn *SpawnSpec `json:"spawn,omitempty"`
	Fwd   *FwdSpec   `json:"fwd,omitempty"`
	Call  *CallSpec  `json:"call,omitempty"`
}

type CloseSpec struct {
	CommPH string `json:"comm_ph"`
}

type WaitSpec struct {
	CommPH string  `json:"comm_ph"`
	ContES ExpSpec `json:"cont_es"`
}

type SendSpec struct {
	CommPH string `json:"comm_ph"`
	ValPH  string `json:"val_ph"`
}

type RecvSpec struct {
	CommPH string  `json:"comm_ph"`
	BindPH string  `json:"bind_ph"`
	ContES ExpSpec `json:"cont_es"`
}

type LabSpec struct {
	CommPH string `json:"comm_ph"`
	LabQN  string `json:"lab_qn"`
}

type CaseSpec struct {
	CommPH  string       `json:"comm_ph"`
	ContBSs []BranchSpec `json:"cont_bss"`
}

type BranchSpec struct {
	LabQN  string  `json:"lab_qn"`
	ContES ExpSpec `json:"cont_es"`
}

type CallSpec struct {
	CommPH string   `json:"comm_ph"`
	BindPH string   `json:"bind_ph"`
	ProcQN string   `json:"proc_qn"` // раньше был SigPH
	ValPHs []string `json:"val_phs"`
	ContES ExpSpec  `json:"cont_es"`
}

type SpawnSpec struct {
	CommPH  string   `json:"comm_ph"`
	ProcQN  string   `json:"proc_qn"`
	BindPHs []string `json:"bind_phs"`
	ContES  ExpSpec  `json:"cont_es"`
}

type FwdSpec struct {
	CommPH string `json:"comm_ph"`
	ContPH string `json:"cont_ph"`
}

func ErrUnexpectedExpKind(k expKind) error {
	return fmt.Errorf("unexpected exp kind: %v", k)
}
