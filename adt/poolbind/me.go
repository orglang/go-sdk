package poolbind

type BindSpec struct {
	ChnlPH string   `json:"chnl_ph"`
	ChnlBM bindMode `json:"chnl_bm"`
	XactQN string   `json:"xact_qn"`
}

type bindMode string

const (
	Self     = bindMode("self")
	Required = bindMode("required")
	Optional = bindMode("optional")
)
