package procexec

import (
	"github.com/go-resty/resty/v2"

	"github.com/orglang/go-sdk/adt/procstep"
)

// Client-side secondary adapter
type RestySDK struct {
	Client *resty.Client
}

func (sdk *RestySDK) Run(spec ExecSpec) error {
	var res ExecRef
	_, err := sdk.Client.R().
		SetPathParam("id", spec.ExecID).
		SetBody(&spec).
		SetResult(&res).
		Post("/procs/{id}/execs")
	if err != nil {
		return err
	}
	return nil
}

func (sdk *RestySDK) Take(spec procstep.StepSpec) error {
	var res ExecRef
	_, err := sdk.Client.R().
		SetResult(&res).
		SetBody(&spec).
		SetPathParam("poolID", spec.ExecID).
		SetPathParam("procID", spec.ProcID).
		Post("/pools/{poolID}/procs/{procID}/steps")
	if err != nil {
		return err
	}
	return nil
}

func (sdk *RestySDK) Retrieve(execID string) (ExecSnap, error) {
	var res ExecSnap
	_, err := sdk.Client.R().
		SetPathParam("id", execID).
		SetResult(&res).
		Get("/procs/{id}")
	if err != nil {
		return ExecSnap{}, err
	}
	return res, nil
}
