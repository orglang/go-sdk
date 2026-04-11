package procexec

import (
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/orglang/go-sdk/adt/procstep"
	"github.com/orglang/go-sdk/adt/semterm"
)

// Client-side secondary adapter
type RestySDK struct {
	Client *resty.Client
}

func (sdk *RestySDK) Take(spec procstep.StepSpec) error {
	var dto semterm.TermRef
	res, err := sdk.Client.R().
		SetResult(&dto).
		SetBody(&spec).
		SetPathParam("id", spec.ImplRef.TermID).
		Post("/procs/{id}/steps")
	if err != nil {
		return err
	}
	if res.IsError() {
		return fmt.Errorf("received: %v", string(res.Body()))
	}
	return nil
}

func (sdk *RestySDK) Retrieve(execRef semterm.TermRef) (ExecSnap, error) {
	var dto ExecSnap
	res, err := sdk.Client.R().
		SetPathParam("id", execRef.TermID).
		SetResult(&dto).
		Get("/procs/{id}")
	if err != nil {
		return ExecSnap{}, err
	}
	if res.IsError() {
		return ExecSnap{}, fmt.Errorf("received: %v", string(res.Body()))
	}
	return dto, nil
}
