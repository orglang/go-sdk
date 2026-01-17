package procdec

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Client-side secondary adapter
type RestySDK struct {
	Client *resty.Client
}

func (sdk *RestySDK) Incept(decQN string) (DecRefME, error) {
	return DecRefME{}, nil
}

func (sdk *RestySDK) Create(spec DecSpecME) (DecSnapME, error) {
	var res DecSnapME
	resp, err := sdk.Client.R().
		SetResult(&res).
		SetBody(&spec).
		Post("/decs")
	if err != nil {
		return DecSnapME{}, err
	}
	if resp.IsError() {
		return DecSnapME{}, fmt.Errorf("received: %v", string(resp.Body()))
	}
	return res, nil
}

func (sdk *RestySDK) RetrieveSnap(decID string) (DecSnapME, error) {
	var res DecSnapME
	resp, err := sdk.Client.R().
		SetResult(&res).
		SetPathParam("id", decID).
		Get("/decs/{id}")
	if err != nil {
		return DecSnapME{}, err
	}
	if resp.IsError() {
		return DecSnapME{}, fmt.Errorf("received: %v", string(resp.Body()))
	}
	return res, nil
}

func (sdk *RestySDK) RetreiveRefs() ([]DecRefME, error) {
	refs := []DecRefME{}
	return refs, nil
}
