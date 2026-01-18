package procdec

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Client-side secondary adapter
type RestySDK struct {
	Client *resty.Client
}

func (sdk *RestySDK) Incept(decQN string) (DecRef, error) {
	return DecRef{}, nil
}

func (sdk *RestySDK) Create(spec DecSpec) (DecSnap, error) {
	var res DecSnap
	resp, err := sdk.Client.R().
		SetResult(&res).
		SetBody(&spec).
		Post("/decs")
	if err != nil {
		return DecSnap{}, err
	}
	if resp.IsError() {
		return DecSnap{}, fmt.Errorf("received: %v", string(resp.Body()))
	}
	return res, nil
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

func (sdk *RestySDK) RetreiveRefs() ([]DecRef, error) {
	refs := []DecRef{}
	return refs, nil
}
