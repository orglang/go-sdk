package poolexp

import (
	"fmt"

	"github.com/orglang/go-sdk/adt/descsem"
	"github.com/orglang/go-sdk/adt/implsem"
)

type ExpSpec struct {
	K       expKind      `json:"kind"`
	Hire    *HireSpec    `json:"hire"`
	Fire    *FireSpec    `json:"fire"`
	Apply   *ApplySpec   `json:"apply"`
	Quit    *QuitSpec    `json:"quit"`
	Acquire *AcquireSpec `json:"acquire"`
	Release *ReleaseSpec `json:"release"`
	Accept  *AcceptSpec  `json:"accept"`
	Detach  *DetachSpec  `json:"detach"`
	Spawn   *SpawnSpec   `json:"spawn"`
}

type expKind string

const (
	Hire    expKind = "hire"
	Fire    expKind = "fire"
	Apply   expKind = "apply"
	Quit    expKind = "quit"
	Acquire expKind = "acquire"
	Release expKind = "release"
	Accept  expKind = "accept"
	Detach  expKind = "detach"
	Spawn   expKind = "spawn"
)

type HireSpec struct {
	CommChnlPH string  `json:"comm_ph"`
	ProcDescQN string  `json:"proc_qn"`
	ContES     ExpSpec `json:"cont_es"`
}

type FireSpec struct {
	CommChnlPH string
}

type ApplySpec struct {
	CommChnlPH string  `json:"comm_ph"`
	ProcDescQN string  `json:"proc_qn"`
	ContES     ExpSpec `json:"cont_es"`
}

type QuitSpec struct {
	CommChnlPH string `json:"comm_ph"`
	ProcDescQN string `json:"proc_qn"`
}

type AcquireSpec struct {
	CommChnlPH string  `json:"comm_ph"`
	ContExp    ExpSpec `json:"cont_exp"`
}

type ReleaseSpec struct {
	CommChnlPH string `json:"comm_ph"`
}

type AcceptSpec struct {
	CommChnlPH string  `json:"comm_ph"`
	ContExp    ExpSpec `json:"cont_exp"`
}

type DetachSpec struct {
	CommChnlPH string `json:"comm_ph"`
}

type SpawnSpec struct {
	// ссылка на описание порождаемого процесса
	ProcDescRef descsem.SemRef `json:"proc_dr"`
	// ссылки на воплощения потребляемых процессов
	ProcImplRefs []implsem.SemRef `json:"proc_irs"`
}

func ErrUnexpectedExpKind(k expKind) error {
	return fmt.Errorf("unexpected exp kind: %v", k)
}
