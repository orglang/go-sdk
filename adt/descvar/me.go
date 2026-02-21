package descvar

type VarSpec struct {
	ChnlPH string `json:"chnl_ph"`
	// ChnlBM bindMode `json:"chnl_bm"`
	DescQN string `json:"desc_qn"`
}

type VarRec struct {
	ChnlPH string `json:"chnl_ph"`
	DescID string `json:"desc_id"`
}

type bindMode string

const (
	Internal bindMode = "internal"
	External bindMode = "external"
)
