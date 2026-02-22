package procdec

import (
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/orglang/go-sdk/adt/descsem"
)

// Client-side secondary adapter
type RestySDK struct {
	Client *resty.Client
}

func (sdk *RestySDK) Incept(decQN string) (descsem.SemRef, error) {
	return descsem.SemRef{}, nil
}

func (sdk *RestySDK) Create(spec DecSpec) (DecSnap, error) {
	var snap DecSnap
	resp, err := sdk.Client.R().
		SetResult(&snap).
		SetBody(&spec).
		Post("/procs/decs")
	if err != nil {
		return DecSnap{}, err
	}
	if resp.IsError() {
		return DecSnap{}, fmt.Errorf("received: %v", string(resp.Body()))
	}
	return snap, nil
}

func (sdk *RestySDK) RetrieveSnap(decID string) (DecSnap, error) {
	var res DecSnap
	resp, err := sdk.Client.R().
		SetResult(&res).
		SetPathParam("id", decID).
		Get("/decs/{id}")
	if err != nil {
		return DecSnap{}, err
	}
	if resp.IsError() {
		return DecSnap{}, fmt.Errorf("received: %v", string(resp.Body()))
	}
	return res, nil
}

func (sdk *RestySDK) RetreiveRefs() ([]descsem.SemRef, error) {
	refs := []descsem.SemRef{}
	return refs, nil
}
