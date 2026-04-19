package compexec

import (
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/orglang/go-sdk/adt/compsem"
	"github.com/orglang/go-sdk/adt/termsem"
	"github.com/orglang/go-sdk/pool/compstep"
)

// Client-side secondary adapter
type RestySDK struct {
	Client *resty.Client
}

func (sdk *RestySDK) Create(spec ExecSpec) (compsem.SemRef, error) {
	var dto compsem.SemRef
	res, err := sdk.Client.R().
		SetResult(&dto).
		SetBody(&spec).
		Post("/pools/execs")
	if err != nil {
		return compsem.SemRef{}, err
	}
	if res.IsError() {
		return compsem.SemRef{}, errHTTPStatus(res.Body())
	}
	return dto, nil
}

func (sdk *RestySDK) Retrieve(ref compsem.SemRef) (ExecSnap, error) {
	var dto ExecSnap
	res, err := sdk.Client.R().
		SetResult(&dto).
		SetPathParam("id", ref.CompID).
		Get("/pools/{id}")
	if err != nil {
		return ExecSnap{}, err
	}
	if res.IsError() {
		return ExecSnap{}, errHTTPStatus(res.Body())
	}
	return dto, nil
}

func (sdk *RestySDK) RetreiveRefs() ([]termsem.SemRef, error) {
	refs := []termsem.SemRef{}
	return refs, nil
}

func (sdk *RestySDK) Take(spec compstep.StepSpec) error {
	res, err := sdk.Client.R().
		SetBody(&spec).
		Post("/pools/execs/steps")
	if err != nil {
		return err
	}
	if res.IsError() {
		return errHTTPStatus(res.Body())
	}
	return nil
}

func (sdk *RestySDK) Spawn(spec compstep.StepSpec) (compsem.SemRef, error) {
	var dto compsem.SemRef
	res, err := sdk.Client.R().
		SetResult(&dto).
		SetBody(&spec).
		Post("/pools/execs/spawns")
	if err != nil {
		return compsem.SemRef{}, err
	}
	if res.IsError() {
		return compsem.SemRef{}, errHTTPStatus(res.Body())
	}
	return dto, nil
}

func errHTTPStatus(body []byte) error {
	return fmt.Errorf("received: %v", string(body))
}
