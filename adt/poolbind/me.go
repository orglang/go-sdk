package poolbind

type BindSpec struct {
	ChnlPH string   `json:"chnl_ph"`
	ChnlBM bindMode `json:"chnl_bm"`
	XactQN string   `json:"xact_qn"`
}

type bindMode string

const (
	Internal bindMode = "internal"
	External bindMode = "external"
)
