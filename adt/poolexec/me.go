package poolexec

import "github.com/orglang/go-sdk/adt/uniqref"

type ExecSpec struct {
	PoolQN  string   `json:"pool_qn"`
	ProcIDs []string `json:"proc_ids"`
	SupID   string   `json:"sup_id"`
}

type ExecRef = uniqref.Msg

type ExecSnap struct {
	ExecID string    `json:"exec_id"`
	Title  string    `json:"title"`
	Subs   []ExecRef `json:"subs"`
}

type PollSpec struct {
	ExecID string `json:"exec_id"`
}
