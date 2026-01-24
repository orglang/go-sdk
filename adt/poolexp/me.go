package poolexp

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
}

type expKind string

const (
	Hire    = expKind("hire")
	Fire    = expKind("fire")
	Apply   = expKind("apply")
	Quit    = expKind("quit")
	Acquire = expKind("acquire")
	Release = expKind("release")
	Accept  = expKind("accept")
	Detach  = expKind("detach")
)

type HireSpec struct {
	CommPH string   `json:"comm_ph"`
	BindPH string   `json:"bind_ph"`
	ProcQN string   `json:"proc_qn"`
	ValPHs []string `json:"val_phs"`
	ContES ExpSpec  `json:"cont_es"`
}

type FireSpec struct {
	ProcQN string
}

type ApplySpec struct {
	CommPH string  `json:"comm_ph"`
	ProcQN string  `json:"proc_qn"`
	ContES ExpSpec `json:"cont_es"`
}

type QuitSpec struct {
	ProcQN string `json:"proc_qn"`
}

type AcquireSpec struct {
	CommPH string  `json:"comm_ph"`
	ContES ExpSpec `json:"cont_es"`
}

type ReleaseSpec struct {
	CommPH string `json:"comm_ph"`
}

type AcceptSpec struct {
	CommPH string  `json:"comm_ph"`
	ContES ExpSpec `json:"cont_es"`
}

type DetachSpec struct {
	CommPH string `json:"comm_ph"`
}
