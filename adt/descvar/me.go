package descvar

type VarSpec struct {
	ChnlPH string `json:"chnl_ph"`
	// ChnlBM bindMode `json:"chnl_bm"`
	DescQN string `json:"xact_qn"`
}

type bindMode string

const (
	Internal bindMode = "internal"
	External bindMode = "external"
)
